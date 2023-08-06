package crypto

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Public, Secret string
}

func GenerateWallet() Wallet {
	key, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	return Wallet{strings.ToLower(address), privateKey}
}
