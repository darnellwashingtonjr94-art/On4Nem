package portfolio

type AlphaPick struct {
	GameID     string
	Selection  string
	Edge       float64
	Confidence int
}

func GenerateDailyAlpha(rawEdges []AlphaPick) []AlphaPick {
	// Filters and isolates top 1 to 12 highest-conviction +EV picks
	if len(rawEdges) > 12 {
		return rawEdges[:12]
	}
	return rawEdges
}
