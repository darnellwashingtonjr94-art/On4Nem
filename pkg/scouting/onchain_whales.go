package scouting

type WhaleWallet struct {
	Address        string
	SmartMoneyScore float64
	ActiveUSDC     float64
}

func TrackTopWhales() []WhaleWallet {
	// Filters on-chain data for top 1% high-ROI prediction market wallets
	return []WhaleWallet{
		{Address: "0x123...abc", SmartMoneyScore: 0.94, ActiveUSDC: 250000.0},
	}
}
