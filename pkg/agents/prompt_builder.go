package agents

import "fmt"

type PromptBuilder struct {
	BaseSystemRole string
}

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{
		BaseSystemRole: "You are the Gemini Pro Master Trading Orchestrator for On4Nem. Enforce strict risk limits and maximize alpha.",
	}
}

func (pb *PromptBuilder) BuildManagerPrompt(marketContext string, liveMetrics string) string {
	return fmt.Sprintf("%s\n\n[Context Data]: %s\n[Live In-Play Metrics]: %s\nGenerate the optimal 1-12 daily portfolio.", 
		pb.BaseSystemRole, marketContext, liveMetrics)
}
