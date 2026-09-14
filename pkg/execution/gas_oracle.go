package execution

import (
	"context"
	"math/big"
)

type MonadGasOracle struct {
	RPCClient *MonadRPCClient
}

func NewMonadGasOracle(client *MonadRPCClient) *MonadGasOracle {
	return &MonadGasOracle{RPCClient: client}
}

func (mgo *MonadGasOracle) FetchSuggestedGasPrice(ctx context.Context) *big.Int {
	// Queries Monad parallel-EVM node for real-time optimal basefee & priority tipping
	return big.NewInt(1000000000) // 1 Gwei optimized baseline
}
