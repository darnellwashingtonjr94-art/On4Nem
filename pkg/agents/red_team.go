package agents

import "log"

type RedTeamAgent struct {
	VetoThreshold float64
}

func (r *RedTeamAgent) ChallengeManagerPick(gameID string, modelEdge float64) bool {
	log.Printf("Red-Team Agent auditing pick for %s with edge %.2f...", gameID, modelEdge)
	if modelEdge < 0.03 {
		log.Println("VETO: Edge insufficient under adversarial stress-testing.")
		return true // Veto pick
	}
	return false
}
