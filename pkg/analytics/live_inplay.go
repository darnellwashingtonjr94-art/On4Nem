package analytics

type GameState struct {
	PossessionTeam string
	FieldPosition  int
	TimeRemaining  int
	ScoreDiff      int
}

type LiveInPlayEngine struct {
	Threshold float64
}

func NewLiveInPlayEngine() *LiveInPlayEngine {
	return &LiveInPlayEngine{Threshold: 0.04} // 4% edge threshold for live arbs
}

func (e *LiveInPlayEngine) EvaluateLiveEPV(state GameState) float64 {
	// High-frequency Expected Possession Value calculation
	baseValue := float64(state.FieldPosition) * 0.05
	if state.TimeRemaining < 300 && state.ScoreDiff < 0 {
		baseValue *= 1.25 // Urgency multiplier
	}
	return baseValue
}
