package context

type ScheduleContext struct {
	TeamID          string
	DaysOfRest      int
	IsAway          bool
	CrossedTimeZones int
}

func CalculateRestPenalty(s ScheduleContext) float64 {
	// Non-linear rest penalty curve (Back-to-back vs 5+ days rest)
	if s.DaysOfRest == 0 {
		return 0.12 // 12% efficiency downgrade for zero rest
	} else if s.DaysOfRest == 1 {
		return 0.05
	}
	return 0.0
}
