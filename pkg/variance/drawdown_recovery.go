package variance

type RecoveryCalculator struct {
	CurrentBankroll  float64
	BaselineBankroll float64
}

func (rc *RecoveryCalculator) RequiredGainToRecover() float64 {
	if rc.CurrentBankroll >= rc.BaselineBankroll {
		return 0.0
	}
	deficit := rc.BaselineBankroll - rc.CurrentBankroll
	return deficit / rc.CurrentBankroll // Percentage gain required to return to high-water mark
}
