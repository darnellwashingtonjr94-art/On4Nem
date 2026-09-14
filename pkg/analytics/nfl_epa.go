package analytics

type PlayMetrics struct {
	Down        int
	Distance    int
	YardLine    int
	Success     bool
}

func CalculateEPA(plays []PlayMetrics) float64 {
	totalEPA := 0.0
	for _, p := range plays {
		if p.Success {
			totalEPA += 0.35 // Standard positive play value contribution
		} else {
			totalEPA -= 0.42 // Turnover or failed conversion penalty
		}
	}
	return totalEPA / float64(len(plays))
}
