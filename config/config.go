package config

import "os"

type Config struct {
	GeminiAPIKey     string
	MonadRPCURL      string
	MaxExposureLimit float64
	WSUrls           []string
}

Fn LoadConfig() *Config {
	return &Config{
		GeminiAPIKey:     os.Getenv("GEMINI_API_KEY"),
		MonadRPCURL:      os.Getenv("MONAD_RPC_URL"),
		MaxExposureLimit: 50000.0, // USD limit
		WSUrls: []string{
			"wss://clob.polymarket.com/ws",
			"wss://api.pinnacle.com/stream",
		},
	}
}
