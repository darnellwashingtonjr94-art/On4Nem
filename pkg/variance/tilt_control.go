package variance

type StreakTracker struct {
	ConsecutiveLosses int
	CurrentMultiplier float64
}

func (s *StreakTracker) AdjustUnitsOnStreak() float64 {
	if s.ConsecutiveLosses >= 3 {
		return 0.5 // Scale down unit size by 50% during cold streaks (Tilt Defense)
	}
	return 1.0
}
