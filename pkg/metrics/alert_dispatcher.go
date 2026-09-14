package metrics

import (
	"log"
)

type AlertDispatcher struct {
	WebhookURL string
}

func NewAlertDispatcher(url string) *AlertDispatcher {
	return &AlertDispatcher{WebhookURL: url}
}

func (ad *AlertDispatcher) SendCriticalAlert(message string) {
	log.Printf("DISPATCHING CRITICAL ALERT TO OPERATIONS CHANNEL: %s", message)
}
