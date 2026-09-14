package portfolio

import "log"

type LiquidityAllocator struct {
	TotalCapital float64
}

func (la *LiquidityAllocator) AllocateCapital(confidenceScore float64) float64 {
	log.Printf("Allocating capital based on alpha confidence score: %.2f", confidenceScore)
	return la.TotalCapital * 0.05 * confidenceScore
}
