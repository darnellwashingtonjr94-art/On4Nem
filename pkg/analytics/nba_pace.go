package analytics

type TeamStats struct {
	PointsFor     float64
	PointsAgainst float64
	Possessions   float64
}

func CalculateNetRating(t TeamStats) float64 {
	if t.Possessions == 0 {
		return 0
	}
	offensiveRating := (t.PointsFor / t.Possessions) * 100
	defensiveRating := (t.PointsAgainst / t.Possessions) * 100
	return offensiveRating - defensiveRating
}
