package main

import (
	"context"
	"log"
	"time"
)

type DaemonOrchestrator struct {
	PollInterval time.Duration
}

func NewDaemonOrchestrator(interval time.Duration) *DaemonOrchestrator {
	return &DaemonOrchestrator{PollInterval: interval}
}

func (do *DaemonOrchestrator) StartBackgroundDaemon(ctx context.Context) {
	ticker := time.NewTicker(do.PollInterval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Background daemon orchestrator stopped.")
				return
			case <-ticker.C:
				log.Println("Daemon heartbeat: Polling feeds, refreshing odds matrix, and evaluating +EV alpha picks.")
			}
		}
	}()
}
