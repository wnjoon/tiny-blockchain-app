package blockchain

// import (
// 	"context"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func GetBlockHeader(client *ethclient.Client) (string, error) {
// 	header, err := client.HeaderByNumber(context.Background(), nil)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return header.Number.String(), nil
// }

// func GetBlock(client *ethclient.Client, blockNumber *big.Int) (*types.Block, error) {
// 	block, err := client.BlockByNumber(context.Background(), blockNumber)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return block, nil
// }

// func GetTransactionCount(client *ethclient.Client, block *types.Block) (uint, error) {
// 	txCount, err := client.TransactionCount(context.Background(), block.Hash())
// 	if err != nil {
// 		return 0, err
// 	}
// 	return txCount, nil
// }
