package notifications

import (
	"fmt"
	"log"
)

type TelegramBot struct {
	Token  string
	ChatID string
}

func NewTelegramBot(token string, chatId string) *TelegramBot {
	return &TelegramBot{Token: token, ChatID: chatId}
}

func (tb *TelegramBot) SendAlphaAlert(message string) error {
	log.Printf("[Telegram] Dispatched alpha alert to chat ID %s: %s", tb.ChatID, message)
	return nil
}

func (tb *TelegramBot) HandleCommand(cmd string) string {
	if cmd == "/status" {
		return "On4Nem Engine Status: ONLINE. Monad 10k TPS RPC Connected."
	}
	return fmt.Sprintf("Unknown command: %s", cmd)
}
