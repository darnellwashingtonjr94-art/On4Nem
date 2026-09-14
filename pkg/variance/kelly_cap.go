package variance

type RiskManager struct {
	MaxAllocationPct float64
}

func NewRiskManager(maxPct float64) *RiskManager {
	return &RiskManager{MaxAllocationPct: maxPct}
}

func (rm *RiskManager) ApplyKellyCap(rawKelly float64) float64 {
	if rawKelly > rm.MaxAllocationPct {
		return rm.MaxAllocationPct // Enforce hard cap to prevent over-leveraging single picks
	}
	return rawKelly
}
