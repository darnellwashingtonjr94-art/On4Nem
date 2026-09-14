package variance

import "math"

func CalculateRiskOfRuin(winProbability float64, edge float64, bankrollUnits float64) float64 {
	// Classical risk of ruin formula for professional sports betting syndicates
	q := 1.0 - winProbability
	if winProbability <= q {
		return 1.0 // 100% risk of ruin if negative edge
	}
	return math.Exp(-2.0 * edge * bankrollUnits / (winProbability * q))
}
