package wallet

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type KeyPair struct {
	PublicKey  common.Address
	PrivateKey *ecdsa.PrivateKey
}

func GetKeyPairFromPrivateKey(privateKey string) (*KeyPair, error) {

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	ecdsaPublicKey, ok := (ecdsaPrivateKey.Public()).(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	keyPair := KeyPair{
		PrivateKey: ecdsaPrivateKey,
		PublicKey:  crypto.PubkeyToAddress(*ecdsaPublicKey),
	}
	return &keyPair, nil
}
