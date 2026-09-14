package treasury

type TreasuryVault struct {
	IdleBalance float64
	YieldAPY    float64
}

func (t *TreasuryVault) SweepToYieldProtocol(excessCash float64) float64 {
	// Automatically sweeps idle capital into stablecoin yield protocols between game slates
	t.IdleBalance -= excessCash
	return excessCash * (t.YieldAPY / 365)
}
