package config

import "os"

type Config struct {
	Environment string
	Port        string
	MaxSlippage float64
	DiscordURL  string
}

// Fix: Replaced 'Fn' typo with 'func'
func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	discordURL := os.Getenv("DISCORD_WEBHOOK_URL")

	return &Config{
		Environment: env,
		Port:        port,
		MaxSlippage: 0.02,
		DiscordURL:  discordURL,
	}
}
