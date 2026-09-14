package treasury

import (
	"time"
)

type TaxRecord struct {
	Timestamp   time.Time
	Asset       string
	CostBasis   float64
	Proceeds    float64
	PlatformFee float64
}

type TaxLedger struct {
	Records []TaxRecord
}

func (tl *TaxLedger) LogTransaction(asset string, basis float64, proceeds float64, fee float64) {
	tl.Records = append(tl.Records, TaxRecord{
		Timestamp:   time.Now(),
		Asset:       asset,
		CostBasis:   basis,
		Proceeds:    proceeds,
		PlatformFee: fee,
	})
}
