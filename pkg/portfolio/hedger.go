package portfolio

type ActivePosition struct {
	MarketID    string
	CurrentOdds float64
	WagerAmount float64
}

func CalculateLiveHedge(pos ActivePosition, opposingOdds float64, targetProfit float64) float64 {
	// Calculates required hedge stake on opposing platform to lock in guaranteed returns
	hedgeStake := (pos.WagerAmount * pos.CurrentOdds) / opposingOdds
	return hedgeStake
}
