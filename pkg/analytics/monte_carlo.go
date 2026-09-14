package analytics

import (
	"math/rand"
)

type GameSimulator struct {
	Simulations int
}

func NewGameSimulator(sims int) *GameSimulator {
	return &GameSimulator{Simulations: sims}
}

func (gs *GameSimulator) RunSimulations(meanScoreDiff float64, stdDev float64) float64 {
	wins := 0
	for i := 0; i < gs.Simulations; i++ {
		simulatedDiff := rand.NormFloat64()*stdDev + meanScoreDiff
		if simulatedDiff > 0 {
			wins++
		}
	}
	return float64(wins) / float64(gs.Simulations)
}
