package portfolio

type CompoundingEngine struct {
	CompoundRatio float64 // 1.0 = 100% reinvestment
}

func NewCompoundingEngine() *CompoundingEngine {
	return &CompoundingEngine{CompoundRatio: 1.0}
}

func (c *CompoundingEngine) ReinvestNetWinnings(previousDayNetProfit float64, currentBankroll float64) float64 {
	reinvestedAmount := previousDayNetProfit * c.CompoundRatio
	return currentBankroll + reinvestedAmount
}
