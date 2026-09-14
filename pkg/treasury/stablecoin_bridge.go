package treasury

import "log"

type StablecoinBridge struct {
	TargetNetwork string
}

func (sb *StablecoinBridge) RebalanceLiquidity(amount float64) {
	log.Printf("Rebalancing $%.2f USDC across cross-chain bridges to %s", amount, sb.TargetNetwork)
}
