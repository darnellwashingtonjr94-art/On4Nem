package variance

// StreakTracker measures consecutive wins and losses for variance monitoring.
type StreakTracker struct {
	ConsecutiveWins   int
	ConsecutiveLosses int
	MaxLossStreak     int
}

func NewStreakTracker() *StreakTracker {
	return &StreakTracker{}
}

func (s *StreakTracker) RecordOutcome(isWin bool) {
	if isWin {
		s.ConsecutiveWins++
		s.ConsecutiveLosses = 0
	} else {
		s.ConsecutiveLosses++
		s.ConsecutiveWins = 0
		if s.ConsecutiveLosses > s.MaxLossStreak {
			s.MaxLossStreak = s.ConsecutiveLosses
		}
	}
}
