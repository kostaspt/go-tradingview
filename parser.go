package tradingview

import "github.com/shopspring/decimal"

func parse(r Response) []Analysis {
	cKeys := ColumnKeys()

	s := make([]Analysis, len(r.Data))
	for i, d := range r.Data {
		// Summary
		s[i].Recommendations.RecommendAll = computeRecommend(d.D[cKeys["Recommend.All"]])
		s[i].Recommendations.RecommendOscillators = computeRecommend(d.D[cKeys["Recommend.Other"]])
		s[i].Recommendations.RecommendMA = computeRecommend(d.D[cKeys["Recommend.MA"]])

		// Oscillators
		s[i].Recommendations.RSI = computeRSI(d.D[cKeys["RSI"]], d.D[cKeys["RSI[1]"]])
		s[i].Recommendations.StochK = computeStoch(
			d.D[cKeys["Stoch.K"]], d.D[cKeys["Stoch.D"]],
			d.D[cKeys["Stoch.D[1]"]], d.D[cKeys["Stoch.D[1]"]],
		)
		s[i].Recommendations.CCI = computeCCI(d.D[cKeys["CCI20"]], d.D[cKeys["CCI20[1]"]])
		s[i].Recommendations.ADX = computeADX(
			d.D[cKeys["ADX"]],
			d.D[cKeys["ADX+DI"]], d.D[cKeys["ADX-DI"]],
			d.D[cKeys["ADX+DI[1]"]], d.D[cKeys["ADX-DI[1]"]],
		)
		s[i].Recommendations.AO = computeAO(d.D[cKeys["AO"]], d.D[cKeys["AO[1]"]], d.D[cKeys["AO[1]"]])
		s[i].Recommendations.Mom = computeMom(d.D[cKeys["Mom"]], d.D[cKeys["Mom[1]"]])
		s[i].Recommendations.MACD = computeMACD(d.D[cKeys["MACD.macd"]], d.D[cKeys["MACD.signal"]])
		s[i].Recommendations.StochRSI = computeSimple(d.D[cKeys["Rec.Stoch.RSI"]])
		s[i].Recommendations.WPR = computeSimple(d.D[cKeys["Rec.WR"]])
		s[i].Recommendations.BBPower = computeSimple(d.D[cKeys["Rec.BBPower"]])
		s[i].Recommendations.UO = computeSimple(d.D[cKeys["Rec.UO"]])

		// Moving Averages
		s[i].Recommendations.EMA10 = computeMA(d.D[cKeys["EMA10"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA10 = computeMA(d.D[cKeys["SMA10"]], d.D[cKeys["close"]])
		s[i].Recommendations.EMA20 = computeMA(d.D[cKeys["EMA20"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA20 = computeMA(d.D[cKeys["SMA20"]], d.D[cKeys["close"]])
		s[i].Recommendations.EMA30 = computeMA(d.D[cKeys["EMA30"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA30 = computeMA(d.D[cKeys["SMA30"]], d.D[cKeys["close"]])
		s[i].Recommendations.EMA50 = computeMA(d.D[cKeys["EMA50"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA50 = computeMA(d.D[cKeys["SMA50"]], d.D[cKeys["close"]])
		s[i].Recommendations.EMA100 = computeMA(d.D[cKeys["EMA100"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA100 = computeMA(d.D[cKeys["SMA100"]], d.D[cKeys["close"]])
		s[i].Recommendations.EMA200 = computeMA(d.D[cKeys["EMA200"]], d.D[cKeys["close"]])
		s[i].Recommendations.SMA200 = computeMA(d.D[cKeys["SMA200"]], d.D[cKeys["close"]])
		s[i].Recommendations.Ichimoku = computeSimple(d.D[cKeys["Rec.Ichimoku"]])
		s[i].Recommendations.VWMA = computeSimple(d.D[cKeys["Rec.VWMA"]])
		s[i].Recommendations.HullMA = computeSimple(d.D[cKeys["Rec.HullMA9"]])

		s[i].Pivots.ClassicS3 = d.D[cKeys["Pivot.M.Classic.S3"]]
		s[i].Pivots.ClassicS2 = d.D[cKeys["Pivot.M.Classic.S2"]]
		s[i].Pivots.ClassicS1 = d.D[cKeys["Pivot.M.Classic.S1"]]
		s[i].Pivots.ClassicMid = d.D[cKeys["Pivot.M.Classic.Middle"]]
		s[i].Pivots.ClassicR1 = d.D[cKeys["Pivot.M.Classic.R1"]]
		s[i].Pivots.ClassicR2 = d.D[cKeys["Pivot.M.Classic.R2"]]
		s[i].Pivots.ClassicR3 = d.D[cKeys["Pivot.M.Classic.R3"]]

		s[i].Pivots.FibonacciS3 = d.D[cKeys["Pivot.M.Fibonacci.S3"]]
		s[i].Pivots.FibonacciS2 = d.D[cKeys["Pivot.M.Fibonacci.S2"]]
		s[i].Pivots.FibonacciS1 = d.D[cKeys["Pivot.M.Fibonacci.S1"]]
		s[i].Pivots.FibonacciMid = d.D[cKeys["Pivot.M.Fibonacci.Middle"]]
		s[i].Pivots.FibonacciR1 = d.D[cKeys["Pivot.M.Fibonacci.R1"]]
		s[i].Pivots.FibonacciR2 = d.D[cKeys["Pivot.M.Fibonacci.R2"]]
		s[i].Pivots.FibonacciR3 = d.D[cKeys["Pivot.M.Fibonacci.R3"]]

		s[i].Pivots.CamarillaS3 = d.D[cKeys["Pivot.M.Camarilla.S3"]]
		s[i].Pivots.CamarillaS2 = d.D[cKeys["Pivot.M.Camarilla.S2"]]
		s[i].Pivots.CamarillaS1 = d.D[cKeys["Pivot.M.Camarilla.S1"]]
		s[i].Pivots.CamarillaMid = d.D[cKeys["Pivot.M.Camarilla.Middle"]]
		s[i].Pivots.CamarillaR1 = d.D[cKeys["Pivot.M.Camarilla.R1"]]
		s[i].Pivots.CamarillaR2 = d.D[cKeys["Pivot.M.Camarilla.R2"]]
		s[i].Pivots.CamarillaR3 = d.D[cKeys["Pivot.M.Camarilla.R3"]]

		s[i].Pivots.WoodieS3 = d.D[cKeys["Pivot.M.Woodie.S3"]]
		s[i].Pivots.WoodieS2 = d.D[cKeys["Pivot.M.Woodie.S2"]]
		s[i].Pivots.WoodieS1 = d.D[cKeys["Pivot.M.Woodie.S1"]]
		s[i].Pivots.WoodieMid = d.D[cKeys["Pivot.M.Woodie.Middle"]]
		s[i].Pivots.WoodieR1 = d.D[cKeys["Pivot.M.Woodie.R1"]]
		s[i].Pivots.WoodieR2 = d.D[cKeys["Pivot.M.Woodie.R2"]]
		s[i].Pivots.WoodieR3 = d.D[cKeys["Pivot.M.Woodie.R3"]]

		s[i].Pivots.DemarkS1 = d.D[cKeys["Pivot.M.Demark.S1"]]
		s[i].Pivots.DemarkMid = d.D[cKeys["Pivot.M.Demark.Middle"]]
		s[i].Pivots.DemarkR1 = d.D[cKeys["Pivot.M.Demark.R1"]]

		s[i].RawValues = make(map[string]decimal.Decimal)
		for k := range cKeys {
			s[i].RawValues[k] = d.D[cKeys[k]]
		}
	}

	return s
}

func computeSimple(val decimal.Decimal) Recommendation {
	if val.Equal(decimal.NewFromInt(1)) {
		return RecommendationBuy
	} else if val.Equal(decimal.NewFromInt(-1)) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeRecommend(val decimal.Decimal) Recommendation {
	if val.GreaterThan(decimal.NewFromFloat(0.5)) {
		return RecommendationStrongBuy
	} else if val.GreaterThan(decimal.NewFromFloat(0.1)) && val.LessThanOrEqual(decimal.NewFromFloat(0.5)) {
		return RecommendationBuy
	} else if val.GreaterThanOrEqual(decimal.NewFromFloat(-0.5)) && val.LessThan(decimal.NewFromFloat(-0.1)) {
		return RecommendationSell
	} else if val.LessThan(decimal.NewFromFloat(-0.5)) {
		return RecommendationStrongSell
	}
	return RecommendationNeutral
}

func computeRSI(rsi, rsi1 decimal.Decimal) Recommendation {
	if rsi.LessThan(decimal.NewFromInt(30)) && rsi1.LessThan(rsi) {
		return RecommendationBuy
	} else if rsi.GreaterThan(decimal.NewFromInt(70)) && rsi1.GreaterThan(rsi) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeStoch(k, d, k1, d1 decimal.Decimal) Recommendation {
	if k.LessThan(decimal.NewFromInt(20)) &&
		d.LessThan(decimal.NewFromInt(20)) &&
		k.GreaterThan(d) &&
		k1.LessThan(d1) {
		return RecommendationBuy
	} else if k.GreaterThan(decimal.NewFromInt(80)) &&
		d.GreaterThan(decimal.NewFromInt(80)) &&
		k.LessThan(d) &&
		k1.GreaterThan(d1) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeCCI(cci, cci1 decimal.Decimal) Recommendation {
	if cci.LessThan(decimal.NewFromInt(-100)) && cci.GreaterThan(cci1) {
		return RecommendationBuy
	} else if cci.GreaterThan(decimal.NewFromInt(100)) && cci.LessThan(cci1) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeADX(adx, adxPosDI, adxNegDI, adxPosDI1, adxNegDI1 decimal.Decimal) Recommendation {
	if adx.GreaterThan(decimal.NewFromInt(20)) && adxPosDI1.LessThan(adxNegDI1) && adxPosDI.GreaterThan(adxNegDI) {
		return RecommendationBuy
	} else if adx.GreaterThan(decimal.NewFromInt(20)) && adxPosDI1.GreaterThan(adxNegDI1) && adxPosDI.LessThan(adxNegDI) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeAO(ao, ao1, ao2 decimal.Decimal) Recommendation {
	if ao.GreaterThan(decimal.NewFromInt(0)) &&
		(ao1.LessThan(decimal.NewFromInt(0)) || (ao1.GreaterThan(decimal.NewFromInt(0))) && ao.GreaterThan(ao1) && ao2.GreaterThan(ao1)) {
		return RecommendationBuy
	} else if ao.LessThan(decimal.NewFromInt(0)) &&
		(ao1.GreaterThan(decimal.NewFromInt(0)) || (ao1.LessThan(decimal.NewFromInt(0))) && ao.LessThan(ao1) && ao2.LessThan(ao1)) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeMom(mom, mom1 decimal.Decimal) Recommendation {
	if mom.GreaterThan(mom1) {
		return RecommendationBuy
	} else if mom.LessThan(mom1) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeMACD(macd, signal decimal.Decimal) Recommendation {
	if macd.GreaterThan(signal) {
		return RecommendationBuy
	} else if macd.LessThan(signal) {
		return RecommendationSell
	}
	return RecommendationNeutral
}

func computeMA(ma, close decimal.Decimal) Recommendation {
	if ma.LessThan(close) {
		return RecommendationBuy
	} else if ma.GreaterThan(close) {
		return RecommendationSell
	}
	return RecommendationNeutral
}
