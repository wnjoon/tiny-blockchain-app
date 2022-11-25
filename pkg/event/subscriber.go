package event

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventListenHandler struct {
	cli *ethclient.Client
}

type EventInfo struct {
	BlockNumber uint64
	TxHash      string
	Index       uint
	Event       map[string]interface{}
}

type ABIEventInput struct {
	Indexed bool
	Type    string
	Name    string
}

func NewEventListenHandler(client *ethclient.Client) EventListenHandler {
	return EventListenHandler{
		cli: client,
	}
}

func (h *EventListenHandler) SubscribeContractEvent(contractAddress string, contractABI abi.ABI) error {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	logs := make(chan types.Log)
	sub, err := h.cli.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

// func (h *EventListenHandler) SubscribeContractEvent(contractAddress string) error {

// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{common.HexToAddress(contractAddress)},
// 	}

// 	logs := make(chan types.Log)
// 	sub, err := h.cli.SubscribeFilterLogs(context.Background(), query, logs)

// 	if err != nil {
// 		return err
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			return err
// 		case vLog := <-logs:
// 			fmt.Println(vLog)
// 		}
// 	}
// }

// Get Input Information from ABI
func getEventInputFromABI(contractABI abi.ABI, eventName string) []ABIEventInput {
	inputArgs := contractABI.Events[eventName].Inputs
	eventInputs := make([]ABIEventInput, 0)
	for i := 0; i < len(inputArgs); i++ {
		eventInput := ABIEventInput{
			Indexed: inputArgs[i].Indexed,
			Type:    inputArgs[i].Type.String(),
			Name:    inputArgs[i].Name,
		}
		eventInputs = append(eventInputs, eventInput)
	}
	return eventInputs
}

//
func (h *EventListenHandler) EventListByBlockNumber(contractAddress, abiStr, eventDescription string, _fromBlock, _toBlock uint64) ([]EventInfo, error) {

	// 이벤트 명 확인
	eventName := strings.Split(eventDescription, "(")[0]

	// from, to 값 확인
	if _toBlock == 0 {
		recentBlockNumber, err := h.cli.BlockNumber(context.Background())
		if err != nil {
			return nil, err
		}
		_toBlock = recentBlockNumber
		// _toBlock = strconv.FormatUint(recentBlockNumber, 10)
	}
	fromBlock := new(big.Int).SetUint64(_fromBlock)
	toBlock := new(big.Int).SetUint64(_toBlock)

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(contractAddress),
		},
	}

	// filter logs by query
	logs, err := h.cli.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	// get contract ABI
	contractABI, err := abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		return nil, err
	}

	// get event input information
	eventInputs := getEventInputFromABI(contractABI, eventName)

	// make event array
	events := make([]EventInfo, 0)

	// iterate logs
	for _, vLog := range logs {

		eventSignature := []byte(eventDescription)
		eventSignatureHash := crypto.Keccak256Hash(eventSignature)

		// Only append correspond event
		if eventSignatureHash.Hex() == vLog.Topics[0].Hex() {
			eventLog := map[string]interface{}{}

			topicIndex := 1
			for i := 0; i < len(eventInputs); i++ {
				// indexed should be from Topics
				if eventInputs[i].Indexed {
					eventLog[eventInputs[i].Name] = vLog.Topics[topicIndex]
					topicIndex++
				}
			}

			err = contractABI.UnpackIntoMap(eventLog, "Transfer", vLog.Data)
			if err != nil {
				return nil, err
			}

			event := EventInfo{
				BlockNumber: vLog.BlockNumber,
				TxHash:      vLog.TxHash.Hex(),
				Index:       vLog.Index,
				Event:       eventLog,
			}
			events = append(events, event)
		}
	}
	return events, nil
}
