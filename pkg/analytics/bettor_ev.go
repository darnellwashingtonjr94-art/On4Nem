package analytics

import "math"

type BettorMetrics struct {
	TotalWagered float64
	TotalReturned float64
}

func (b *BettorMetrics) CalculateROI() float64 {
	if b.TotalWagered == 0 {
		return 0
	}
	return (b.TotalReturned - b.TotalWagered) / b.TotalWagered
}

func CalculateExpectedValue(trueProb float64, decimalOdds float64) float64 {
	return (trueProb * decimalOdds) - 1.0
}

func CalculateSharpeRatio(returns []float64, riskFreeRate float64) float64 {
	if len(returns) == 0 {
		return 0
	}
	sum, mean, variance := 0.0, 0.0, 0.0
	for _, r := range returns {
		sum += r
	}
	mean = sum / float64(len(returns))

	for _, r := range returns {
		variance += math.Pow(r-mean, 2)
	}
	stdDev := math.Sqrt(variance / float64(len(returns)))
	if stdDev == 0 {
		return 0
	}
	return (mean - riskFreeRate) / stdDev
}
