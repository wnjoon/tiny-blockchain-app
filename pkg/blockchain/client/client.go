package client

import (
	"errors"

	config "tiny-blockchain-app/config"
	"tiny-blockchain-app/pkg/wallet"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumController struct {
	Client  *ethclient.Client
	KeyPair wallet.KeyPair
}

func NewClient(conf config.Config) (*ethclient.Client, error) {
	if conf.Endpoint == "" {
		return nil, errors.New("no endpoint info")
	}

	client, err := ethclient.Dial(conf.Endpoint)
	if err != nil {
		return nil, err
	}

	return client, nil
}
