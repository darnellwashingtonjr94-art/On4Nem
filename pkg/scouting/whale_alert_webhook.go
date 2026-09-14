package scouting

import "log"

type WhaleAlertPayload struct {
	WalletAddress string
	AmountUSDC    float64
	TargetMarket  string
}

func ProcessWhaleWebhook(payload WhaleAlertPayload) {
	log.Printf("Whale movement detected: Wallet %s deployed $%.2f USDC into market %s", 
		payload.WalletAddress, payload.AmountUSDC, payload.TargetMarket)
}
