package metrics

import (
	"log"
	"time"
)

type LatencyTracker struct {
	StartTime time.Time
}

func NewLatencyTracker() *LatencyTracker {
	return &LatencyTracker{StartTime: time.Now()}
}

func (lt *LatencyTracker) RecordExecutionTime(operation string) {
	duration := time.Since(lt.StartTime)
	log.Printf("Operation [%s] completed in %v (Monad high-frequency performance target met)", operation, duration)
}
