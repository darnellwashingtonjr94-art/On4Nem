package portfolio

type ParlayLeg struct {
	Odds  float64
	Prob  float64
}

func CalculateParlayEV(legs []ParlayLeg, stake float64) float64 {
	combinedProb := 1.0
	combinedDecimalOdds := 1.0
	for _, leg := range legs {
		combinedProb *= leg.Prob
		combinedDecimalOdds *= leg.Odds
	}
	return (combinedProb * combinedDecimalOdds * stake) - stake
}
