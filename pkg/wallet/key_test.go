package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetKeyPairFromPrivateKey(t *testing.T) {
	samplePrivateKey := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	expectedAddress := "b5ff8c7f64c1cfddb68edc1006dd5c58b07e1448"
	keyPair, err := GetKeyPairFromPrivateKey(samplePrivateKey)

	assert.Equal(t, nil, err)
	assert.Equal(t, common.HexToAddress(expectedAddress), keyPair.PublicKey)

}
