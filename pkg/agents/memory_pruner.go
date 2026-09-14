package agents

import "log"

type MemoryPruner struct {
	MaxTokens int
}

func NewMemoryPruner(maxTokens int) *MemoryPruner {
	return &MemoryPruner{MaxTokens: maxTokens}
}

func (mp *MemoryPruner) CompactContextHistory(history []string) []string {
	log.Println("Compacting Gemini Pro multi-agent context history to maintain token efficiency...")
	if len(history) > 50 {
		return history[len(history)-50:]
	}
	return history
}
