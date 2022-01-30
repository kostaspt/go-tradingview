package tradingview

import (
	"github.com/shopspring/decimal"
)

type Analysis struct {
	Recommendations struct {
		// Summary
		RecommendAll         Recommendation `json:"recommend_all"`
		RecommendOscillators Recommendation `json:"recommend_oscillators"`
		RecommendMA          Recommendation `json:"recommend_ma"`

		// Oscillators
		RSI      Recommendation `json:"rsi"`
		StochK   Recommendation `json:"stoch_k"`
		CCI      Recommendation `json:"cci"`
		ADX      Recommendation `json:"adx"`
		AO       Recommendation `json:"ao"`
		Mom      Recommendation `json:"mom"`
		MACD     Recommendation `json:"macd"`
		StochRSI Recommendation `json:"stoch_rsi_fast"`
		WPR      Recommendation `json:"wpr"`
		BBPower  Recommendation `json:"bb_power"`
		UO       Recommendation `json:"uo"`

		// Moving Averages
		EMA10    Recommendation `json:"ema10"`
		SMA10    Recommendation `json:"sma10"`
		EMA20    Recommendation `json:"ema20"`
		SMA20    Recommendation `json:"sma20"`
		EMA30    Recommendation `json:"ema30"`
		SMA30    Recommendation `json:"sma30"`
		EMA50    Recommendation `json:"ema50"`
		SMA50    Recommendation `json:"sma50"`
		EMA100   Recommendation `json:"ema100"`
		SMA100   Recommendation `json:"sma100"`
		EMA200   Recommendation `json:"ema200"`
		SMA200   Recommendation `json:"sma200"`
		Ichimoku Recommendation `json:"ichimoku"`
		VWMA     Recommendation `json:"vwma"`
		HullMA   Recommendation `json:"hull_ma"`
	} `json:"recommendations"`
	Pivots struct {
		ClassicS3  decimal.Decimal `json:"classic_s3"`
		ClassicS2  decimal.Decimal `json:"classic_s2"`
		ClassicS1  decimal.Decimal `json:"classic_s1"`
		ClassicMid decimal.Decimal `json:"classic_middle"`
		ClassicR1  decimal.Decimal `json:"classic_r1"`
		ClassicR2  decimal.Decimal `json:"classic_r2"`
		ClassicR3  decimal.Decimal `json:"classic_r3"`

		FibonacciS3  decimal.Decimal `json:"fibonacci_s3"`
		FibonacciS2  decimal.Decimal `json:"fibonacci_s2"`
		FibonacciS1  decimal.Decimal `json:"fibonacci_s1"`
		FibonacciMid decimal.Decimal `json:"fibonacci_middle"`
		FibonacciR1  decimal.Decimal `json:"fibonacci_r1"`
		FibonacciR2  decimal.Decimal `json:"fibonacci_r2"`
		FibonacciR3  decimal.Decimal `json:"fibonacci_r3"`

		CamarillaS3  decimal.Decimal `json:"camarilla_s3"`
		CamarillaS2  decimal.Decimal `json:"camarilla_s2"`
		CamarillaS1  decimal.Decimal `json:"camarilla_s1"`
		CamarillaMid decimal.Decimal `json:"camarilla_middle"`
		CamarillaR1  decimal.Decimal `json:"camarilla_r1"`
		CamarillaR2  decimal.Decimal `json:"camarilla_r2"`
		CamarillaR3  decimal.Decimal `json:"camarilla_r3"`

		WoodieS3  decimal.Decimal `json:"woodie_s3"`
		WoodieS2  decimal.Decimal `json:"woodie_s2"`
		WoodieS1  decimal.Decimal `json:"woodie_s1"`
		WoodieMid decimal.Decimal `json:"woodie_middle"`
		WoodieR1  decimal.Decimal `json:"woodie_r1"`
		WoodieR2  decimal.Decimal `json:"woodie_r2"`
		WoodieR3  decimal.Decimal `json:"woodie_r3"`

		DemarkS1  decimal.Decimal `json:"demark_s1"`
		DemarkMid decimal.Decimal `json:"demark_middle"`
		DemarkR1  decimal.Decimal `json:"demark_r1"`
	} `json:"pivots"`
	RawValues map[string]decimal.Decimal `json:"raw_values"`
}

type Recommendation string

const (
	RecommendationStrongBuy  = Recommendation("strong-buy")
	RecommendationBuy        = Recommendation("buy")
	RecommendationNeutral    = Recommendation("neutral")
	RecommendationSell       = Recommendation("sell")
	RecommendationStrongSell = Recommendation("strong-sell")
)

type Interval string

const DefaultInterval = Interval1D

const (
	Interval1Min  Interval = "1m"
	Interval5Min  Interval = "5m"
	Interval15Min Interval = "15m"
	Interval30Min Interval = "30m"
	Interval60Min Interval = "1h"
	Interval1H    Interval = "1h"
	Interval2H    Interval = "2h"
	Interval4H    Interval = "4h"
	Interval1D    Interval = "1d"
	Interval1W    Interval = "1w"
	Interval1M    Interval = "1M"
)

func (i Interval) ForColumn() string {
	switch i {
	case Interval1Min:
		return "|1"
	case Interval5Min:
		return "|5"
	case Interval15Min:
		return "|15"
	case Interval30Min:
		return "|30"
	case Interval1H:
		return "|60"
	case Interval2H:
		return "|120"
	case Interval4H:
		return "|240"
	case Interval1W:
		return "|1W"
	case Interval1M:
		return "|1M"
	}
	return ""
}

func ColumnKeys() map[string]int {
	m := make(map[string]int)
	for key, val := range Columns() {
		m[val] = key
	}
	return m
}

func Columns() []string {
	return []string{
		"Recommend.Other",
		"Recommend.All",
		"Recommend.MA",
		"RSI",
		"RSI[1]",
		"Stoch.K",
		"Stoch.D",
		"Stoch.K[1]",
		"Stoch.D[1]",
		"CCI20",
		"CCI20[1]",
		"ADX",
		"ADX+DI",
		"ADX-DI",
		"ADX+DI[1]",
		"ADX-DI[1]",
		"AO",
		"AO[1]",
		"AO[2]",
		"Mom",
		"Mom[1]",
		"MACD.macd",
		"MACD.signal",
		"Rec.Stoch.RSI",
		"Stoch.RSI.K",
		"Rec.WR",
		"W.R",
		"Rec.BBPower",
		"BBPower",
		"BB.lower",
		"BB.upper",
		"Rec.UO",
		"UO",
		"P.SAR",
		"EMA5",
		"SMA5",
		"EMA10",
		"SMA10",
		"EMA20",
		"SMA20",
		"EMA30",
		"SMA30",
		"EMA50",
		"SMA50",
		"EMA100",
		"SMA100",
		"EMA200",
		"SMA200",
		"Rec.Ichimoku",
		"Ichimoku.BLine",
		"Rec.VWMA",
		"VWMA",
		"Rec.HullMA9",
		"HullMA9",
		"Pivot.M.Classic.S3",
		"Pivot.M.Classic.S2",
		"Pivot.M.Classic.S1",
		"Pivot.M.Classic.Middle",
		"Pivot.M.Classic.R1",
		"Pivot.M.Classic.R2",
		"Pivot.M.Classic.R3",
		"Pivot.M.Fibonacci.S3",
		"Pivot.M.Fibonacci.S2",
		"Pivot.M.Fibonacci.S1",
		"Pivot.M.Fibonacci.Middle",
		"Pivot.M.Fibonacci.R1",
		"Pivot.M.Fibonacci.R2",
		"Pivot.M.Fibonacci.R3",
		"Pivot.M.Camarilla.S3",
		"Pivot.M.Camarilla.S2",
		"Pivot.M.Camarilla.S1",
		"Pivot.M.Camarilla.Middle",
		"Pivot.M.Camarilla.R1",
		"Pivot.M.Camarilla.R2",
		"Pivot.M.Camarilla.R3",
		"Pivot.M.Woodie.S3",
		"Pivot.M.Woodie.S2",
		"Pivot.M.Woodie.S1",
		"Pivot.M.Woodie.Middle",
		"Pivot.M.Woodie.R1",
		"Pivot.M.Woodie.R2",
		"Pivot.M.Woodie.R3",
		"Pivot.M.Demark.S1",
		"Pivot.M.Demark.Middle",
		"Pivot.M.Demark.R1",
		"open",
		"low",
		"high",
		"close",
		"volume",
		"change",
	}
}
