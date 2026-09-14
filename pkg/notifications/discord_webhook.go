package notifications

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type DiscordWebhook struct {
	WebhookURL string
}

func NewDiscordWebhook(url string) *DiscordWebhook {
	return &DiscordWebhook{WebhookURL: url}
}

func (dw *DiscordWebhook) SendEmbedAlert(title string, description string) error {
	payload := map[string]string{
		"content": fmt.Sprintf("**%s**\n%s", title, description),
	}
	body, _ := json.Marshal(payload)
	log.Printf("[Discord] Broadcasting alert webhook: %s", title)
	_, _ = http.Post(dw.WebhookURL, "application/json", bytes.NewBuffer(body))
	return nil
}
