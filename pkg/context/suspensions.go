package context

type SuspensionRecord struct {
	PlayerName       string
	IsSuspended      bool
	ChemistryPenalty float64 // Unit unit penalty for unit timing breakdown
}

func EvaluateSuspensionImpact(s SuspensionRecord) float64 {
	if !s.IsSuspended {
		return 0.0
	}
	return 0.15 + s.ChemistryPenalty // Heavy efficiency drop due to sudden roster voiding
}
