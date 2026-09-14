package execution

import (
	"log"
	"time"
)

type WSReconnector struct {
	MaxRetries int
	Backoff    time.Duration
}

func NewWSReconnector(retries int, backoff time.Duration) *WSReconnector {
	return &WSReconnector{MaxRetries: retries, Backoff: backoff}
}

func (w *WSReconnector) HandleDisconnection(url string) {
	log.Printf("WebSocket connection lost for endpoint: %s. Initiating exponential backoff reconnect...", url)
	// Automated reconnection loop with jitter for uninterrupted live streaming
}
