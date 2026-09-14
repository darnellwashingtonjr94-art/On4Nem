package analytics

import "log"

type MicrostructureAnalyzer struct {
	OrderBookDepth float64
}

func (ma *MicrostructureAnalyzer) AnalyzeBookImbalance(bidsSum, asksSum float64) float64 {
	if (bidsSum + asksSum) == 0 {
		return 0.0
	}
	imbalance := (bidsSum - asksSum) / (bidsSum + asksSum)
	log.Printf("Order book microstructure imbalance calculated: %.4f", imbalance)
	return imbalance
}
