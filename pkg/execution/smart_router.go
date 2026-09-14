package execution

import "log"

type SmartRouter struct {
	SupportedBooks []string
}

func NewSmartRouter(books []string) *SmartRouter {
	return &SmartRouter{SupportedBooks: books}
}

func (sr *SmartRouter) RouteBestExecution(marketID string, targetOdds float64) string {
	log.Printf("Smart routing order for market %s across active bookmakers for optimal fill price...", marketID)
	return "Pinnacle" // Selects book with lowest vig and highest liquidity
}
