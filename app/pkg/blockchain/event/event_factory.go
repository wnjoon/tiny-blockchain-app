package event

import (
	"tiny-blockchain-app/app/pkg/blockchain"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EventFactory struct {
	httpCli      *ethclient.Client
	websocketCli *ethclient.Client
}

func NewEventFactory(conf blockchain.Config) (*EventFactory, error) {

	httpCli, err := ethclient.Dial(conf.EndPoint)
	if err != nil {
		return nil, err
	}

	websocketCli, err := ethclient.Dial(conf.WebSocket)
	if err != nil {
		return nil, err
	}

	return &EventFactory{
		httpCli:      httpCli,
		websocketCli: websocketCli,
	}, nil
}
