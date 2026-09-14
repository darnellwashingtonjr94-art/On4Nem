package variance

type StreakTracker struct {
	History []bool // true = win, false = loss
}

func (s *StreakTracker) GetCurrentStreak() (string, int) {
	if len(s.History) == 0 {
		return "None", 0
	}
	lastResult := s.History[len(s.History)-1]
	count := 0
	for i := len(s.History) - 1; i >= 0; i-- {
		if s.History[i] == lastResult {
			count++
		} else {
			break
		}
	}
	if lastResult {
		return "Winning", count
	}
	return "Losing", count
}
