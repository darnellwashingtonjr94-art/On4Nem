package analytics

import "math"

type ExtendedEVCalculator struct {
	VigMultiplier float64
}

func (e *ExtendedEVCalculator) CalculateTrueNoVigProbability(oddsA, oddsB float64) (float64, float64) {
	impliedA := 1.0 / oddsA
	impliedB := 1.0 / oddsB
	totalImplied := impliedA + impliedB
	
	trueProbA := impliedA / totalImplied
	trueProbB := impliedB / totalImplied
	return trueProbA, trueProbB
}

func CalculateKellyCriterionStake(bankroll float64, trueProb float64, decimalOdds float64) float64 {
	b := decimalOdds - 1.0
	q := 1.0 - trueProb
	kelly := (b*trueProb - q) / b
	if kelly <= 0 {
		return 0
	}
	return math.Round(bankroll * kelly * 0.25 * 100.0) / 100.0 // Quarter Kelly sizing
}
