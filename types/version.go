package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
)

// mainnet
const (
	MainnetChainId       = "fxcore"
	mainnetEvmChainID    = 530
	MainnetGenesisHash   = "56629F685970FEC1E35521FC943ACE9AEB2C53448544A0560E4DD5799E1A5593"
	MainnetBlockHeightV2 = 5_713_000
	MainnetBlockHeightV3 = 8_756_000
)

// testnet
const (
	TestnetChainId       = "dhobyghaut"
	testnetEvmChainID    = 90001
	TestnetGenesisHash   = "06D0A9659E1EC5B0E57E8E2E5F1B1266094808BC9B4081E1A55011FEF4586ACE"
	TestnetBlockHeightV2 = 3418880
	TestnetBlockHeightV3 = 6578000
)

var (
	chainId = MainnetChainId
	once    sync.Once
)

func SetChainId(id string) {
	if id != MainnetChainId && id != TestnetChainId {
		panic("invalid chainId: " + id)
	}
	once.Do(func() {
		chainId = id
	})
}

func ChainId() string {
	return chainId
}

func EIP155ChainID() *big.Int {
	if TestnetChainId == ChainId() {
		return big.NewInt(testnetEvmChainID)
	}
	return big.NewInt(mainnetEvmChainID)
}

func ChainIdWithEIP155() string {
	if TestnetChainId == ChainId() {
		return fmt.Sprintf("%s_%d-1", TestnetChainId, testnetEvmChainID)
	}
	return fmt.Sprintf("%s_%d-1", MainnetChainId, mainnetEvmChainID)
}

// Sha256Hex calculate SHA-256 hash
func Sha256Hex(b []byte) string {
	sha := sha256.New()
	sha.Write(b)
	return strings.ToUpper(hex.EncodeToString(sha.Sum(nil)))
}
