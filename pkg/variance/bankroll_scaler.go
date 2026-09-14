package variance

import "log"

type BankrollScaler struct {
	BaseMultiplier float64
}

func (bs *BankrollScaler) ScaleUnitBySize(currentBankroll float64, initialBankroll float64) float64 {
	ratio := currentBankroll / initialBankroll
	log.Printf("Scaling bankroll unit size based on current ratio: %.2f", ratio)
	return 100.0 * ratio // Dynamic base unit scaling
}
