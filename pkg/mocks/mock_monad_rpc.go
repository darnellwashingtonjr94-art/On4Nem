package mocks

import (
	"log"
	"math/big"
)

type MockMonadNode struct {
	ChainID int64
}

func NewMockMonadNode(chainID int64) *MockMonadNode {
	return &MockMonadNode{ChainID: chainID}
}

func (m *MockMonadNode) SimulateParallelTxExecution(txPayload []byte) (*big.Int, error) {
	log.Printf("[Mock Monad RPC] Simulating parallel-EVM execution for payload size %d bytes...", len(txPayload))
	return big.NewInt(21000), nil // Simulated gas used
}
