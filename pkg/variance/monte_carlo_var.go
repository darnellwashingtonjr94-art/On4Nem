package variance

import (
	"math"
	"math/rand"
)

type ValueAtRiskEngine struct {
	ConfidenceLevel float64
}

func (va *ValueAtRiskEngine) CalculateVaR(portfolioValue float64, volatility float64, days int) float64 {
	// Parametric Value-at-Risk calculation for sports portfolio tail-risk management
	zScore := 1.645 // 95% confidence interval
	return portfolioValue * volatility * zScore * math.Sqrt(float64(days)/365.0)
}
