package analytics

import "fmt"

func ConvertAmericanToDecimal(americanOdds int) (float64, error) {
	if americanOdds == 0 {
		return 0, fmt.Errorf("invalid american odds format")
	}
	if americanOdds > 0 {
		return (float64(americanOdds) / 100.0) + 1.0, nil
	}
	return (100.0 / float64(-americanOdds)) + 1.0, nil
}

func CalculateImpliedProbability(decimalOdds float64) float64 {
	if decimalOdds <= 0 {
		return 0
	}
	return 1.0 / decimalOdds
}
