package variance

import "math"

// RiskOfRuin calculates probability of hitting zero bankroll based on win rate and unit risk.
func RiskOfRuin(winRate, winLossRatio, riskPerTrade float64) float64 {
	if winRate <= 0 || winRate >= 1 {
		return 1.0
	}
	
	// Formula: ((1 - (w - l)) / (1 + (w - l))) ^ units
	edge := (winRate * winLossRatio) - (1.0 - winRate)
	if edge <= 0 {
		return 1.0
	}

	units := 1.0 / riskPerTrade
	prob := math.Pow((1.0 - edge) / (1.0 + edge), units)
	
	if prob > 1.0 {
		return 1.0
	}
	return prob
}
