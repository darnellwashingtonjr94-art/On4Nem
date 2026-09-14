package execution

import (
	"log"
	"math/big"
)

type AtomicBundler struct {
	BuilderEndpoint string
}

func NewAtomicBundler(endpoint string) *AtomicBundler {
	return &AtomicBundler{BuilderEndpoint: endpoint}
}

func (ab *AtomicBundler) BundleTransactions(txData [][]byte, maxTip *big.Int) error {
	log.Printf("Bundling %d transactions for atomic execution via Monad parallel sequencer...", len(txData))
	return nil
}
