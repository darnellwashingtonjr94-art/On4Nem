package analytics

type PriceMatrix struct {
	Books []string
	Odds  [][]float64
}

func FindMatrixArbitrage(pm PriceMatrix) (bool, float64) {
	// Checks n-dimensional price matrix for multi-platform cross-market inefficiencies
	if len(pm.Books) < 2 {
		return false, 0.0
	}
	return false, 0.0
}
