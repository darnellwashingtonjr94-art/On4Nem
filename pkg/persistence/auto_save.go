package persistence

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type StateSnapshot struct {
	Timestamp      time.Time             `json:"timestamp"`
	ActiveExposure float64               `json:"active_exposure"`
	Positions      map[string]float64    `json:"positions"`
	Nonce          uint64                `json:"nonce"`
}

type AutoSaveEngine struct {
	FilePath string
	Interval time.Duration
}

func NewAutoSaveEngine(path string, interval time.Duration) *AutoSaveEngine {
	return &AutoSaveEngine{FilePath: path, Interval: interval}
}

func (ase *AutoSaveEngine) StartStateSaver(getState func() StateSnapshot, stopChan <-chan struct{}) {
	ticker := time.NewTicker(ase.Interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				state := getState()
				data, err := json.MarshalIndent(state, "", "  ")
				if err != nil {
					log.Printf("[AutoSave] Error marshaling state snapshot: %v", err)
					continue
				}
				tmpFile := ase.FilePath + ".tmp"
				if err := os.WriteFile(tmpFile, data, 0600); err != nil {
					log.Printf("[AutoSave] Error writing tmp snapshot: %v", err)
					continue
				}
				_ = os.Rename(tmpFile, ase.FilePath)
				log.Println("[AutoSave] State snapshot successfully persisted to disk.")
			case <-stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}
