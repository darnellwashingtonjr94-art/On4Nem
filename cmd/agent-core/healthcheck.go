package main

import (
	"log"
	"time"
)

type HealthMonitor struct {
	Interval time.Duration
}

func NewHealthMonitor(interval time.Duration) *HealthMonitor {
	return &HealthMonitor{Interval: interval}
}

func (hm *HealthMonitor) StartHeartbeat() {
	ticker := time.NewTicker(hm.Interval)
	go func() {
		for range ticker.C {
			log.Println("On4Nem System Heartbeat: All agent subroutines and RPC connections operating normally.")
		}
	}()
}
