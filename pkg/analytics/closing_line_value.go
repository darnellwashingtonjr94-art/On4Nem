package analytics

import "log"

type CLVRecord struct {
	GameID       string
	PlacedOdds   float64
	ClosingOdds  float64
}

func CalculateCLV(rec CLVRecord) float64 {
	// Closing Line Value is the ultimate metric of a sharp sports betting edge
	clvPercentage := (rec.PlacedOdds / rec.ClosingOdds) - 1.0
	log.Printf("CLV Analyzed for Game %s: %.2f%% beat-the-close efficiency", rec.GameID, clvPercentage*100)
	return clvPercentage
}
