package event

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventSubscriber struct {
	cli            *ethclient.Client
	contractAddr   common.Address
	contractABI    abi.ABI
	eventSignature []byte
	outch          chan EventInfo
	errch          chan error
}

type EventInfo struct {
	BlockNumber uint64
	TxHash      string
	Index       uint
	Name        string
	Event       map[string]interface{}
}

type abiEventInput struct {
	Indexed bool
	Type    string
	Name    string
}

func NewEventSubscriber(cli *ethclient.Client, contractAddr common.Address, contractABI abi.ABI, eventSignature []byte) *EventSubscriber {
	return &EventSubscriber{
		cli:            cli,
		contractAddr:   contractAddr,
		contractABI:    contractABI,
		eventSignature: eventSignature,
		outch:          make(chan EventInfo),
		errch:          make(chan error),
	}
}

func (es *EventSubscriber) NewSubscription() chan EventInfo {
	query := filterQuery(nil, nil, es.contractAddr, es.eventSignature)
	go subscribe(es, query)
	return es.outch
}

func subscribe(es *EventSubscriber, query ethereum.FilterQuery) {

	logs := make(chan types.Log)
	sub, err := es.cli.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event, err := getEvent(es.contractABI, vLog)
			if err != nil {
				log.Fatal(err)
			}
			es.outch <- event
		}
	}
}

// func (es *EventSubscriber) SubscribeFunctionEvent() {}

// func subscribe(es *EventSubscriber, query ethereum.FilterQuery) {

// 	logs := make(chan types.Log)
// 	sub, _ := es.cli.SubscribeFilterLogs(context.Background(), query, logs)

// 	// sub, err := es.cli.SubscribeFilterLogs(context.Background(), query, logs)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			es.errch <- err
// 		case vLog := <-logs:
// 			event, err := getEvent(es.contractABI, vLog)
// 			if err != nil {
// 				es.errch <- err
// 			}
// 			es.outch <- event
// 		}
// 	}
// }

func filterQuery(from, to *big.Int, contractAddress common.Address, eventSignature []byte) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{
			contractAddress,
		},
	}
	if eventSignature != nil {
		query.Topics = [][]common.Hash{
			{crypto.Keccak256Hash(eventSignature)},
		}
	}
	return query
}

func eventInfoList(cli *ethclient.Client, contractABI abi.ABI, query ethereum.FilterQuery) ([]EventInfo, error) {

	// filter logs by query
	logs, err := cli.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	// make event array
	events := make([]EventInfo, len(logs))

	// iterate logs
	for idx, vLog := range logs {
		event, err := getEvent(contractABI, vLog)
		if err != nil {
			return nil, err
		}
		events[idx] = event
	}
	return events, nil
}

// Get event from filtered log
func getEvent(contractABI abi.ABI, vLog types.Log) (EventInfo, error) {

	// Get event from log
	eventSignature := vLog.Topics[0]
	abiEvent, err := contractABI.EventByID(eventSignature)
	if err != nil {
		return EventInfo{}, err
	}
	eventInputs := eventInputFromABI(contractABI, abiEvent.Name)

	// Only append correspond event
	eventLog := map[string]interface{}{}

	topicIndex := 1
	for i := 0; i < len(eventInputs); i++ {
		// indexed should be from Topics
		if eventInputs[i].Indexed {
			eventLog[eventInputs[i].Name] = vLog.Topics[topicIndex]
			topicIndex++
		}
	}

	err = contractABI.UnpackIntoMap(eventLog, abiEvent.Name, vLog.Data)
	if err != nil {
		return EventInfo{}, err
	}

	event := EventInfo{
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Hex(),
		Index:       vLog.Index,
		Name:        abiEvent.Name,
		Event:       eventLog,
	}
	return event, nil
}

// Get Input Information from ABI
func eventInputFromABI(contractABI abi.ABI, eventName string) []abiEventInput {
	inputArgs := contractABI.Events[eventName].Inputs
	eventInputs := make([]abiEventInput, len(inputArgs))
	for i := 0; i < len(inputArgs); i++ {
		eventInput := abiEventInput{
			Indexed: inputArgs[i].Indexed,
			Type:    inputArgs[i].Type.String(),
			Name:    inputArgs[i].Name,
		}
		eventInputs[i] = eventInput
	}
	return eventInputs
}

func BlockchainEventPrinter(ch <-chan EventInfo, index int) {
	for {
		select {
		case event := <-ch:
			fmt.Println("[", index, "] : ", event)
		}
	}
}
