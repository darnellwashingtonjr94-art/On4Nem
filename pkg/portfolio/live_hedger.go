package portfolio

import "log"

type LiveHedgingEngine struct {
	TriggerMargin float64
}

func NewLiveHedgingEngine() *LiveHedgingEngine {
	return &LiveHedgingEngine{TriggerMargin: 0.08}
}

func (lhe *LiveHedgingEngine) EvaluateInPlayHedging(impliedProbability float64, liveProbability float64) bool {
	delta := liveProbability - impliedProbability
	if delta >= lhe.TriggerMargin {
		log.Println("Live hedging threshold met: Executing counter-position to lock 100% outcome.")
		return true
	}
	return false
}
