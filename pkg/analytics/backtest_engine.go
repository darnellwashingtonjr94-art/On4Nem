package analytics

import "log"

type HistoricalFixture struct {
	FixtureID string
	ClosingOdds float64
	ActualOutcome bool
}

func RunBacktestSimulation(fixtures []HistoricalFixture) float64 {
	log.Printf("Running historical backtest simulation across %d fixtures...", len(fixtures))
	wins := 0
	for _, f := range fixtures {
		if f.ActualOutcome {
			wins++
		}
	}
	return float64(wins) / float64(len(fixtures))
}
