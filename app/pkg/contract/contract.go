package contract

import (
	"context"
	"math/big"
	"tiny-blockchain-app/app/pkg/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractResponse struct {
	Address  common.Address
	Tx       *types.Transaction
	Instance interface{}
}

func GetAuth(client *ethclient.Client, keyPair wallet.KeyPair) (*bind.TransactOpts, error) {

	nonce, err := client.PendingNonceAt(context.Background(), keyPair.PublicKey)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(keyPair.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(12500000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	// gasPrice, _ := big.NewInt(0).SetString("0", 10)
	auth.GasPrice = gasPrice

	return auth, nil
}

func checkMinted(client *ethclient.Client, response *ContractResponse) (*types.Receipt, error) {
	tx, isPending, err := client.TransactionByHash(context.Background(), response.Tx.Hash())
	if err != nil {
		return nil, err
	}

	if !isPending {
		return nil, nil
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func checkDeployed(client *ethclient.Client, response *ContractResponse) (common.Address, error) {
	tx, isPending, err := client.TransactionByHash(context.Background(), response.Tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	if !isPending {
		return response.Address, nil
	}

	address, err := bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}
