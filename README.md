# go-tradingview
> Note: This package is experimental and is not ready for production use.

### Usage
```go
package main

import (
	"context"
	"log"

	"github.com/kostaspt/go-tradingview"
)

func main() {
	c := tradingview.NewClient(nil)

	analysis, err := c.GetAnalysis(context.Background(), "crypto", []string{"BINANCE:BTCEUR"}, tradingview.DefaultInterval)
	if err != nil {
		panic(err)
	}

	log.Printf("%#v", analysis)
}
```

### Notes
- Based on [python-tradingview-ta](https://github.com/brian-the-dev/python-tradingview-ta)
