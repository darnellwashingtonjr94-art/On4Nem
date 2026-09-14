package variance

import "log"

type DrawdownProtector struct {
	PeakBankroll   float64
	MaxDrawdownPct float64
}

func NewDrawdownProtector(peak float64, maxPct float64) *DrawdownProtector {
	return &DrawdownProtector{PeakBankroll: peak, MaxDrawdownPct: maxPct}
}

func (dp *DrawdownProtector) CheckMaxLossLock(currentBankroll float64) bool {
	drawdown := (dp.PeakBankroll - currentBankroll) / dp.PeakBankroll
	if drawdown >= dp.MaxDrawdownPct {
		log.Println("MAX DRAWDOWN REACHED: Halting all automated execution loops immediately.")
		return true
	}
	return false
}
