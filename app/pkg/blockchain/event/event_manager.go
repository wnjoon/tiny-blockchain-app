package event

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Structs
type EventSubscriber struct {
	cli     *ethclient.Client
	request EventRequest
	outch   chan EventResponse
	errch   chan error
	subch   chan bool
}

type EventHistoryFinder struct {
	cli     *ethclient.Client
	request EventRequest
	from    *big.Int
	to      *big.Int
}

type EventDescription struct {
	Name  string
	Rules map[string][]interface{}
}

type EventRequest struct {
	ABI       abi.ABI
	Addresses []common.Address
	// [TODO] 우선 단건 이벤트에 대해서만 제대로 조회 -> 향후 다건 이벤트도 적용할 것 (2022.12.20)
	Events EventDescription // Transfer : from : {xxx, yyy}
}

type EventResponse struct {
	BlockNumber uint64
	TxHash      string
	Index       uint
	Name        string                 // Event Name
	Event       map[string]interface{} // Event rules
}

type abiInput struct {
	Indexed bool
	Type    string
	Name    string
}

//// Main Functions
// NewEventSubscriber
func (e *EventFactory) NewEventSubscriber(request EventRequest) *EventSubscriber {
	return &EventSubscriber{
		cli:     e.websocketCli,
		request: request,
		outch:   make(chan EventResponse),
		errch:   make(chan error),
		subch:   make(chan bool),
	}
}

// NewEventHistoryFinder
func (e *EventFactory) NewEventHistoryFinder(request EventRequest, from, to *big.Int) *EventHistoryFinder {
	return &EventHistoryFinder{
		cli:     e.httpCli,
		request: request,
		from:    from,
		to:      to,
	}
}

// subscribe
func (e *EventSubscriber) Subscribe() (chan EventResponse, chan error, error) {
	query, err := queryRequest(e.request, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		logs := make(chan types.Log)
		sub, err := e.cli.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			return
		}
		for {
			select {
			case err := <-sub.Err():
				e.errch <- err
			case vLog := <-logs:
				event, err := getEvent(e.request, vLog)
				if err != nil {
					e.errch <- err
				}
				e.outch <- event
			case <-e.subch:
				sub.Unsubscribe() // Unsubscribe cancels the sending of events to the data channel and closes the error channel.
				return
			}
		}
	}()
	return e.outch, e.errch, nil
}

// History
func (e *EventHistoryFinder) History() ([]EventResponse, error) {

	query, err := queryRequest(e.request, e.from, e.to)
	if err != nil {
		return nil, err
	}

	logs, err := e.cli.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	response := make([]EventResponse, 0)

	for _, vLog := range logs {
		event, err := getEvent(e.request, vLog)
		if err != nil {
			return nil, err
		}
		response = append(response, event)
	}
	return response, nil
}

// filterTopics
func filterTopics(contractAbi abi.ABI, d EventDescription) ([][]common.Hash, error) {

	totalRules := make([][]interface{}, 0)
	abiInputs := getEventAbiInput(contractAbi, d.Name)

	for _, input := range abiInputs {

		if d.Rules != nil {
			if input.Indexed {
				rules := make([]interface{}, 0)
				item, exist := d.Rules[input.Name]
				if exist {
					typedRule := typeConverter(input.Type, item)
					rules = append(rules, typedRule...)
				}
				totalRules = append(totalRules, rules)
			}
		}
	}
	result := append([][]interface{}{{contractAbi.Events[d.Name].ID}}, totalRules...)

	topics, err := abi.MakeTopics(result...)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

// getEventConstruct : ABI 문서로부터 이벤트 입력 순서에 맞게 검색 조건 구조체 생성
func getEventAbiInput(abi abi.ABI, event string) []abiInput {
	inputs := abi.Events[event].Inputs
	eventConstruct := make([]abiInput, len(inputs))

	for i, input := range inputs {
		eventDescription := abiInput{
			Indexed: input.Indexed,
			Type:    input.Type.String(),
			Name:    input.Name,
		}
		eventConstruct[i] = eventDescription
	}
	return eventConstruct
}

// Get event from filtered log
func getEvent(r EventRequest, vLog types.Log) (EventResponse, error) {

	eventLog := map[string]interface{}{}
	abiEvent, err := r.ABI.EventByID(vLog.Topics[0])
	if err != nil {
		return EventResponse{}, err
	}

	abiInputs := getEventAbiInput(r.ABI, abiEvent.Name)

	// vLog.Topics[0] => event signature이며, 실제 topic 값은 1부터 시작
	tIdx := 1
	for _, input := range abiInputs {
		if input.Indexed {
			eventLog[input.Name] = vLog.Topics[tIdx]
			tIdx++
		}
	}

	err = r.ABI.UnpackIntoMap(eventLog, abiEvent.Name, vLog.Data)
	if err != nil {
		return EventResponse{}, err
	}

	event := EventResponse{
		BlockNumber: vLog.BlockNumber,
		TxHash:      vLog.TxHash.Hex(),
		Index:       vLog.Index,
		Name:        abiEvent.Name,
		Event:       eventLog,
	}
	return event, nil
}

// getQuery : 쿼리문 생성
func queryRequest(r EventRequest, from, to *big.Int) (ethereum.FilterQuery, error) {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, r.Addresses...)

	query := ethereum.FilterQuery{
		Addresses: addresses,
		FromBlock: from,
		ToBlock:   to,
	}

	// Filter events when request has conditions
	if !reflect.ValueOf(r.Events).IsZero() {
		topics, err := filterTopics(r.ABI, r.Events)
		if err != nil {
			return ethereum.FilterQuery{}, err
		}
		query.Topics = topics
	}

	return query, nil
}

// TypeConverter : All parameters should be inserted with string
func typeConverter(ruleType string, rules []interface{}) []interface{} {
	convertedRules := make([]interface{}, 0)
	for _, rule := range rules {
		str := fmt.Sprintf("%v", rule)
		switch ruleType {
		case "string":
			rule = str
		case "address":
			rule = common.HexToAddress(str)
		case "hash":
			rule = common.HexToHash(str)
		case "bool":
			rule, _ = strconv.ParseBool(str)
		}
		convertedRules = append(convertedRules, rule)
	}
	return convertedRules
}
