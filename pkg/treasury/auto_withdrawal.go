package treasury

import "log"

type AutoWithdrawalEngine struct {
	ThresholdAmount float64
}

func (awe *AutoWithdrawalEngine) CheckProfitSweep(currentBalance float64, baselineBankroll float64) {
	profit := currentBalance - baselineBankroll
	if profit >= awe.ThresholdAmount {
		log.Printf("Profit threshold reached: Sweeping $%.2f into cold storage vault.", profit)
	}
}
