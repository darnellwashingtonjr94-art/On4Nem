package config

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

type NetworkConfig struct {
	NetworkName      string
	RPCURL           string
	ChainID          *big.Int
	OperatorWallet   string
	ExplorerURL      string
	GasMultiplier    float64
}

func LoadNetworkConfig() map[string]*NetworkConfig {
	operatorAddress := "0x95481E7BE683720C68B47AC3BCa8e8d0c62F78DB"

	return map[string]*NetworkConfig{
		"mainnet": {
			NetworkName:    "Monad Mainnet",
			RPCURL:         getEnvOrDefault("MONAD_MAINNET_RPC", "https://rpc.monad.xyz"),
			ChainID:        big.NewInt(143), // Official Monad Mainnet Chain ID
			OperatorWallet: operatorAddress,
			ExplorerURL:    "https://monadvision.com",
			GasMultiplier:  1.20, // 20% aggressive priority tip for high-frequency sandwich defense
		},
		"testnet": {
			NetworkName:    "Monad Testnet",
			RPCURL:         getEnvOrDefault("MONAD_TESTNET_RPC", "https://testnet-rpc.monad.xyz"),
			ChainID:        big.NewInt(10143), // Monad Testnet Chain ID
			OperatorWallet: operatorAddress,
			ExplorerURL:    "https://testnet.monadexplorer.com",
			GasMultiplier:  1.05,
		},
	}
}

func ValidateOperatorKey(privateKeyHex string, expectedAddress string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	derivedAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if derivedAddress != expectedAddress {
		log.Printf("[SECURITY WARNING] Derived address %s does not match expected operator %s!", derivedAddress, expectedAddress)
	} else {
		log.Printf("[OK] Operator wallet successfully verified: %s", derivedAddress)
	}

	return privateKey, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
