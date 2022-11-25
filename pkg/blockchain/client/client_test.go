package client

import (
	"context"
	"math/big"
	"testing"
	"tiny-blockchain-app/config"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	client, err := NewClient(conf)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, client)

	chainId, err := client.ChainID(context.Background())
	bigIntChainId, _ := new(big.Int).SetString("10", 10)

	assert.NoError(t, err)
	assert.Equal(t, bigIntChainId, chainId)
}
