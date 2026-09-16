package variance

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

type MonteCarloConfig struct {
	InitialBankroll float64
	WinRate         float64
	AvgWin          float64
	AvgLoss         float64
	Simulations     int
	Trades          int
	Confidence      float64
}

func MonteCarloVaR(cfg MonteCarloConfig) float64 {
	if cfg.Simulations <= 0 || cfg.Trades <= 0 {
		return 0.0
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	results := make([]float64, cfg.Simulations)

	for i := 0; i < cfg.Simulations; i++ {
		bankroll := cfg.InitialBankroll
		for t := 0; t < cfg.Trades; t++ {
			if rng.Float64() < cfg.WinRate {
				bankroll += cfg.AvgWin
			} else {
				bankroll -= cfg.AvgLoss
			}
		}
		results[i] = bankroll
	}

	sort.Float64s(results)

	cutoff := int(math.Floor((1.0 - cfg.Confidence) * float64(cfg.Simulations)))
	if cutoff < 0 {
		cutoff = 0
	}

	varAtRisk := cfg.InitialBankroll - results[cutoff]
	if varAtRisk < 0 {
		return 0.0
	}

	return varAtRisk
}
