package execution

import (
	"math/big"
)

type TransactionBuilder struct {
	ChainID *big.Int
	GasTip  *big.Int
}

func NewTransactionBuilder(chainID int64) *TransactionBuilder {
	return &TransactionBuilder{
		ChainID: big.NewInt(chainID),
		GasTip:  big.NewInt(1500000000), // Optimized for Monad parallel EVM gas efficiency
	}
}

func (tb *TransactionBuilder) BuildOptimizedTx(toAddress string, value *big.Int, data []byte) map[string]interface{} {
	return map[string]interface{}{
		"chainId": tb.ChainID,
		"to":      toAddress,
		"value":   value,
		"data":    data,
		"gasTip":  tb.GasTip,
	}
}
