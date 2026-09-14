package portfolio

func CalculateKellyFraction(winProb float64, decimalOdds float64, bankroll float64) float64 {
	// Fractional Kelly Criterion sizing (e.g., quarter-Kelly for variance control)
	b := decimalOdds - 1.0
	q := 1.0 - winProb
	kelly := (b*winProb - q) / b
	if kelly < 0 {
		return 0
	}
	return bankroll * kelly * 0.25 // Quarter Kelly allocation
}
