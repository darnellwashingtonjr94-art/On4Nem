package execution

import (
	"math/big"
)

type GasOptimizer struct {
	BaseMultiplier float64
}

func NewGasOptimizer() *GasOptimizer {
	return &GasOptimizer{BaseMultiplier: 1.15}
}

func (go_opt *GasOptimizer) CalculatePriorityFee(currentNetworkGas *big.Int) *big.Int {
	// Optimizes priority gas bidding for Monad's parallel-EVM blockspace during high-volatility sports slates
	multiplierBig := big.NewInt(int64(go_opt.BaseMultiplier * 100))
	optimizedFee := new(big.Int).Mul(currentNetworkGas, multiplierBig)
	return optimizedFee.Div(optimizedFee, big.NewInt(100))
}
