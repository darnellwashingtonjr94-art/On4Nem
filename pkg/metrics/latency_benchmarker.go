package metrics

import (
	"log"
	"time"
)

type LatencyBenchmarker struct {
	TargetBooks []string
}

func (lb *LatencyBenchmarker) MeasureBookPing(book string) time.Duration {
	start := time.Now()
	// Simulates lightweight TCP ping to sportsbook gateway
	duration := time.Since(start)
	log.Printf("Book [%s] network latency: %v", book, duration)
	return duration
}
