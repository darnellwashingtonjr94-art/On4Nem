package treasury

import (
	"log"
	"math/big"
)

type MultiSigVault struct {
	RequiredSignatures int
	SignerAddresses    []string
}

func (msv *MultiSigVault) ProposeHighValueTransfer(recipient string, amount *big.Int) {
	log.Printf("PROPOSAL: High-value treasury transfer of %s USDC to %s requires %d multi-sig approvals.",
		amount.String(), recipient, msv.RequiredSignatures)
}
