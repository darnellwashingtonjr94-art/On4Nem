package agents

import (
	"log"
	"time"
)

type AgentHeartbeat struct {
	AgentID string
	LastPing time.Time
}

func (ah *AgentHeartbeat) RecordHeartbeat() {
	ah.LastPing = time.Now()
	log.Printf("Agent [%s] heartbeat confirmed active at %v", ah.AgentID, ah.LastPing)
}
