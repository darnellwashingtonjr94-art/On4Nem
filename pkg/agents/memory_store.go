package agents

type AgentMemory struct {
	ContextHistory []string
}

func NewAgentMemory() *AgentMemory {
	return &AgentMemory{ContextHistory: make([]string, 0)}
}

func (am *AgentMemory) AppendMemory(thought string) {
	if len(am.ContextHistory) >= 100 {
		am.ContextHistory = am.ContextHistory[1:] // Keep rolling 100-thought window for Gemini Pro
	}
	am.ContextHistory = append(am.ContextHistory, thought)
}
