package agents

import "log"

type OrchestratorHook struct {
	ActiveAgents int
}

func (oh *OrchestratorHook) BroadcastStateUpdate(stateID string) {
	log.Printf("Broadcasting state update [%s] across multi-agent Gemini Pro cluster", stateID)
}
