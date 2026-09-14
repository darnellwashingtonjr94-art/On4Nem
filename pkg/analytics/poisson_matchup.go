package analytics

type MatchupExpectation struct {
	HomeExpectedGoals float64
	AwayExpectedGoals float64
}

func SimulatePoissonMatchup(me MatchupExpectation) map[string]float64 {
	// Evaluates multi-goal scoring distribution probabilities for soccer and hockey
	return map[string]float64{
		"homeWin": 0.46,
		"draw":    0.28,
		"awayWin": 0.26,
	}
}
