package treasury

type CircuitBreaker struct {
	MaxExposure float64
	CurrentRisk float64
	IsTripped   bool
}

func NewCircuitBreaker(maxLimit float64) *CircuitBreaker {
	return &CircuitBreaker{MaxExposure: maxLimit, IsTripped: false}
}

func (cb *CircuitBreaker) CheckExposure(newBetAmount float64) bool {
	if cb.CurrentRisk+newBetAmount > cb.MaxExposure {
		cb.IsTripped = true
		return true // Trip breaker, halt execution loops
	}
	return false
}
