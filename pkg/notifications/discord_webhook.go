package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WebhookMessage struct {
	Username string `json:"username,omitempty"`
	Content  string `json:"content"`
}

type DiscordNotifier struct {
	WebhookURL string
	Client     *http.Client
}

func NewDiscordNotifier(webhookURL string) *DiscordNotifier {
	return &DiscordNotifier{
		WebhookURL: webhookURL,
		Client:     &http.Client{Timeout: 5 * time.Second},
	}
}

func (d *DiscordNotifier) SendAlert(message string) error {
	if d.WebhookURL == "" {
		return fmt.Errorf("discord webhook URL is not configured")
	}

	payload := WebhookMessage{
		Username: "On4Nem Alert Bot",
		Content:  message,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal discord payload: %w", err)
	}

	resp, err := d.Client.Post(d.WebhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send discord webhook request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook returned non-200 status: %d", resp.StatusCode)
	}

	return nil
}
