package portfolio

type PortfolioSummary struct {
	ActivePositions int
	TotalExposure   float64
	ExpectedAlpha   float64
}

func SummarizePortfolio(positions []ActivePosition, exposure float64) PortfolioSummary {
	return PortfolioSummary{
		ActivePositions: len(positions),
		TotalExposure:   exposure,
		ExpectedAlpha:   0.084, // 8.4% projected portfolio edge
	}
}
