package ml

import (
	"log"
)

type FeatureStore struct {
	Features map[string]map[string]float64
}

func NewFeatureStore() *FeatureStore {
	return &FeatureStore{Features: make(map[string]map[string]float64)}
}

func (fs *FeatureStore) SaveGameFeature(gameID string, featureName string, value float64) {
	if _, exists := fs.Features[gameID]; !exists {
		fs.Features[gameID] = make(map[string]float64)
	}
	fs.Features[gameID][featureName] = value
	log.Printf("[FeatureStore] Saved feature [%s] = %.4f for game %s", featureName, value, gameID)
}
