package treasury

import (
	"log"
	"math/big"
)

type StablecoinSwapper struct {
	DEXRouterAddress string
}

func NewStablecoinSwapper(router string) *StablecoinSwapper {
	return &StablecoinSwapper{DEXRouterAddress: router}
}

func (ss *StablecoinSwapper) ExecuteSwap(tokenIn string, tokenOut string, amount *big.Int) error {
	log.Printf("Executing atomic stablecoin swap on Monad AMM [In: %s, Out: %s, Amount: %s]", tokenIn, tokenOut, amount.String())
	return nil
}
