package variance

type TiltGuard struct {
	MaxAllowedLossStreak int
	IsCoolingOff         bool
}

func NewTiltGuard(maxLossStreak int) *TiltGuard {
	return &TiltGuard{
		MaxAllowedLossStreak: maxLossStreak,
	}
}

// EvaluateStatus checks the loss streak using StreakTracker without redefining the type.
func (tg *TiltGuard) EvaluateStatus(tracker *StreakTracker) bool {
	if tracker.ConsecutiveLosses >= tg.MaxAllowedLossStreak {
		tg.IsCoolingOff = true
		return true // Tilt triggered
	}
	tg.IsCoolingOff = false
	return false
}
