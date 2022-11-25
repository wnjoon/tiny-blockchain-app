package transaction

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionManager struct {
	Client      *ethclient.Client
	Address     common.Address
	Transaction *types.Transaction
}
