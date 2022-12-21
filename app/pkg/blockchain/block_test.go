package blockchain

// import (
// 	"math/big"
// 	"testing"
// 	"tiny-blockchain-app/config"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetBlockHeader(t *testing.T) {
// 	conf := config.Config{
// 		Endpoint: "http://127.0.0.1:22001",
// 	}
// 	client, _ := NewClient(conf)

// 	blockHeaderNumber, err := GetBlockHeader(client)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, "0", blockHeaderNumber)
// }

// func TestGetBlock(t *testing.T) {
// 	conf := config.Config{
// 		Endpoint: "http://127.0.0.1:22001",
// 	}
// 	client, _ := NewClient(conf)

// 	block, err := GetBlock(client, big.NewInt(0))
// 	assert.Equal(t, nil, err)
// 	t.Log(block)
// }

// func TestGetTransactionCount(t *testing.T) {
// 	conf := config.Config{
// 		Endpoint: "http://127.0.0.1:22001",
// 	}
// 	client, _ := NewClient(conf)

// 	block, err := GetBlock(client, big.NewInt(0))
// 	assert.Equal(t, nil, err)
// 	txCount, err := GetTransactionCount(client, block)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, uint(0), txCount)
// }
