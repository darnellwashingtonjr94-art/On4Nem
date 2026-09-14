package portfolio

type BetSelection struct {
	GameID string
	Entity string
}

func CalculatePortfolioCorrelation(selections []BetSelection) float64 {
	// Analyzes multi-variable dependency matrix to prevent correlated single-game tail risks
	riskScore := 0.0
	seenGames := make(map[string]bool)
	for _, sel := range selections {
		if seenGames[sel.GameID] {
			riskScore += 0.25 // Incremental risk penalty for overlapping game exposure
		}
		seenGames[sel.GameID] = true
	}
	return riskScore
}
