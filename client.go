package tradingview

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"
)

var ErrInvalidScreener = errors.New("invalid screener")
var ErrNoSymbols = errors.New("no symbols given")

type Client struct {
	client *http.Client
}

type Request struct {
	Symbols struct {
		Tickers []string `json:"tickers"`
		Query   struct {
			Types []interface{} `json:"types"`
		} `json:"query"`
	} `json:"symbols"`
	Columns []string `json:"columns"`
}

type Response struct {
	Data []struct {
		D []decimal.Decimal `json:"d"`
	} `json:"data"`
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	return &Client{
		client: client,
	}
}

func (c *Client) GetAnalysis(ctx context.Context, screener string, symbols []string, interval Interval) ([]Analysis, error) {
	resp, err := c.do(ctx, screener, symbols, interval)
	if err != nil {
		return nil, err
	}

	return parse(resp), nil
}

func (c Client) do(ctx context.Context, screener string, symbols []string, interval Interval) (Response, error) {
	var r Response

	if screener == "" {
		return r, ErrInvalidScreener
	}

	if len(symbols) == 0 {
		return r, ErrNoSymbols
	}

	url := fmt.Sprintf("https://scanner.tradingview.com/%s/scan", strings.ToLower(screener))

	payload, err := c.payload(symbols, interval)
	if err != nil {
		return r, err
	}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return r, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go-tradingview")

	resp, err := c.client.Do(req)
	if err != nil {
		return r, err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return r, fmt.Errorf("HTTP error %d: returned %s", resp.StatusCode, raw)
	}

	err = json.Unmarshal(raw, &r)
	return r, err
}

func (c Client) payload(symbols []string, interval Interval) (*bytes.Reader, error) {
	data := Request{}
	data.Symbols.Tickers = symbols
	data.Symbols.Query.Types = []interface{}{}
	data.Columns = c.columns(interval)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(payload), nil
}

func (c Client) columns(period Interval) []string {
	cls := Columns()

	suffix := period.ForColumn()
	if len(suffix) > 0 {
		for i := range cls {
			cls[i] += suffix
		}
	}

	return cls
}
