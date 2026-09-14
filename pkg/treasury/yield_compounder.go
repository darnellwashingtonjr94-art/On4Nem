package treasury

import "log"

type YieldCompounder struct {
	ProtocolAPY float64
}

func NewYieldCompounder(apy float64) *YieldCompounder {
	return &YieldCompounder{ProtocolAPY: apy}
}

func (yc *YieldCompounder) CompoundIdleUSDC(balance float64) float64 {
	dailyYield := balance * (yc.ProtocolAPY / 365.0)
	log.Printf("Compounding idle USDC: Generated daily yield of $%.2f", dailyYield)
	return dailyYield
}
