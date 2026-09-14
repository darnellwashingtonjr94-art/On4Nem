package agents

import "log"

type VectorMemoryStore struct {
	EmbeddingsCache map[string][]float32
}

func NewVectorMemoryStore() *VectorMemoryStore {
	return &VectorMemoryStore{EmbeddingsCache: make(map[string][]float32)}
}

func (vms *VectorMemoryStore) IndexHistoricalGameContext(gameID string, contextVector []float32) {
	vms.EmbeddingsCache[gameID] = contextVector
	log.Printf("Indexed semantic vector embedding for game ID: %s", gameID)
}
