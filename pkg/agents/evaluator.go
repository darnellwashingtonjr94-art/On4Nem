package agents

import "log"

type AgentEvaluator struct {
	AccuracyWeight float64
}

func (ae *AgentEvaluator) EvaluateSubAgentPerformance(agentID string, successfulPicks int, totalPicks int) float64 {
	if totalPicks == 0 {
		return 0.0
	}
	score := float64(successfulPicks) / float64(totalPicks)
	log.Printf("Agent [%s] performance score evaluated: %.2f", agentID, score)
	return score
}
