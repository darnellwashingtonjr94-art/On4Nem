import "log"

type SlippageGuard struct {
	MaxSlippagePct float64
}

func NewSlippageGuard(maxSlippage float64) *SlippageGuard {
	return &SlippageGuard{MaxSlippagePct: maxSlippage}
}

func (sg *SlippageGuard) ValidatePriceImpact(expectedPrice, executionPrice float64) bool {
	slippage := (executionPrice - expectedPrice) / expectedPrice
	if slippage > sg.MaxSlippagePct {
		log.Printf("SLIPPAGE ALERT: Deviation of %.2f%% exceeds max limit. Aborting trade.", slippage*100)
		return false
	}
	return true
}
