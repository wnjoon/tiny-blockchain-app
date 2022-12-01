package contract

import (
	"context"
	"math/big"
	"strconv"
	"testing"
	"tiny-blockchain-app/config"
	"tiny-blockchain-app/pkg/blockchain/client"
	"tiny-blockchain-app/pkg/wallet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTransferERC20Burnable(t *testing.T) {

	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	// client.NewClient(conf)
	client, _ := client.NewClient(conf)

	samplePrivateKey1 := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	expectedAddress1 := "0xb5ff8c7f64c1cfddb68edc1006dd5c58b07e1448"
	keyPair1, err := wallet.GetKeyPairFromPrivateKey(samplePrivateKey1)
	assert.Equal(t, nil, err)
	assert.Equal(t, common.HexToAddress(expectedAddress1), keyPair1.PublicKey)

	samplePrivateKey2 := "432024aaf30b6921f51f9c21ffad5485fa82550c7627fb2eb121ebe3f165648a"
	expectedAddress2 := "0xa11c095c7262e0f1c001c397dfb288cc2c1c516b"
	keyPair2, err := wallet.GetKeyPairFromPrivateKey(samplePrivateKey2)
	assert.Equal(t, nil, err)
	assert.Equal(t, common.HexToAddress(expectedAddress2), keyPair2.PublicKey)

	contractAddress := "0x5B0CFA20e02145ea529983c7954702Ab5E9e7949"

	// prevBalance1, err := BalanceERC20Burnable(client, contractAddress, keyPair1.PublicKey.Hex())
	// assert.Equal(t, nil, err)
	// prevBalance2, err := BalanceERC20Burnable(client, contractAddress, keyPair2.PublicKey.Hex())
	// assert.Equal(t, nil, err)
	// assert.Equal(t, "30000", prevBalance1)
	// assert.Equal(t, "0", prevBalance2)

	amount := 100

	toAddress := keyPair2.PublicKey.Hex()
	// assert.Equal(t, expectedAddress1, toAddress)

	receipt, err := TransferERC20Burnable(client, *keyPair2, contractAddress, toAddress, amount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)

	// currBalance1, err := BalanceERC20Burnable(client, contractAddress, keyPair1.PublicKey.Hex())
	// assert.Equal(t, nil, err)
	// prevBalanceInt1, _ := strconv.Atoi(prevBalance1)
	// currBalanceInt1, _ := strconv.Atoi(currBalance1)
	// assert.Equal(t, prevBalanceInt1-amount, currBalanceInt1)

	// currBalance2, err := BalanceERC20Burnable(client, contractAddress, keyPair2.PublicKey.Hex())
	// assert.Equal(t, nil, err)
	// prevBalanceInt2, _ := strconv.Atoi(prevBalance2)
	// currBalanceInt2, _ := strconv.Atoi(currBalance2)
	// assert.Equal(t, prevBalanceInt2+amount, currBalanceInt2)

}

func TestDeployERC20Burnable(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	client, _ := client.NewClient(conf)

	samplePrivateKey := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	keyPair, _ := wallet.GetKeyPairFromPrivateKey(samplePrivateKey)

	constructor := ERC20Constructor{
		Name:     "ERC20Burnable",
		Symbol:   "E2B",
		Decimals: 10,
	}

	// 다음에 사용할 논스값
	nonce, err := client.PendingNonceAt(context.Background(), keyPair.PublicKey)
	assert.Equal(t, nil, err)

	// 컨트랙트 배포
	response, err := DeployERC20Burnable(client, *keyPair, constructor)
	assert.Equal(t, nil, err)

	// 다음에 사용할 논스값을 배포 과정에 정확히 사용하였는가?
	assert.Equal(t, nonce, response.Tx.Nonce())

	// Transaction이 정상적으로 만들어졌는지?
	receipt, err := client.TransactionReceipt(context.Background(), response.Tx.Hash())
	assert.Equal(t, uint64(1), receipt.Status)
	assert.Equal(t, nil, err)

	t.Log(receipt.ContractAddress)
}

func TestMintERC20Burnable(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	client, _ := client.NewClient(conf)

	samplePrivateKey := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	keyPair, _ := wallet.GetKeyPairFromPrivateKey(samplePrivateKey)

	contractAddress := "0x5B0CFA20e02145ea529983c7954702Ab5E9e7949"

	prevBalance, err := BalanceERC20Burnable(client, contractAddress, keyPair.PublicKey.Hex())
	assert.Equal(t, nil, err)

	amount := 1000000
	receipt, err := MintERC20Burnable(client, *keyPair, contractAddress, keyPair.PublicKey.Hex(), amount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)

	currBalance, err := BalanceERC20Burnable(client, contractAddress, keyPair.PublicKey.Hex())
	assert.Equal(t, nil, err)
	prevBalanceInt, _ := strconv.Atoi(prevBalance)
	currBalanceInt, _ := strconv.Atoi(currBalance)
	assert.Equal(t, prevBalanceInt+amount, currBalanceInt)
}

func TestTransferERC20BurnableUsingABIGen(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	client, _ := client.NewClient(conf)

	fromPrivateKey := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	fromKeyPair, _ := wallet.GetKeyPairFromPrivateKey(fromPrivateKey)

	toPrivateKey := "432024aaf30b6921f51f9c21ffad5485fa82550c7627fb2eb121ebe3f165648a"
	toKeyPair, _ := wallet.GetKeyPairFromPrivateKey(toPrivateKey)

	contractAddress := "0x5B0CFA20e02145ea529983c7954702Ab5E9e7949"

	fromPrevBalance, _ := BalanceERC20Burnable(client, contractAddress, fromKeyPair.PublicKey.Hex())
	toPrevBalance, _ := BalanceERC20Burnable(client, contractAddress, toKeyPair.PublicKey.Hex())
	fromPrevBalanceInt, _ := strconv.Atoi(fromPrevBalance) // from prev balance of Int
	toPrevBalanceInt, _ := strconv.Atoi(toPrevBalance)     // to prev balance of Int

	amount := 40
	receipt, err := TransferERC20UsingABIGen(client, *fromKeyPair, contractAddress, toKeyPair.PublicKey.Hex(), amount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)
	t.Log(receipt.TxHash.Hex())

	fromCurrBalance, _ := BalanceERC20Burnable(client, contractAddress, fromKeyPair.PublicKey.Hex())
	toCurrBalance, _ := BalanceERC20Burnable(client, contractAddress, toKeyPair.PublicKey.Hex())
	fromCurrBalanceInt, _ := strconv.Atoi(fromCurrBalance)
	toCurrBalanceInt, _ := strconv.Atoi(toCurrBalance)

	assert.Equal(t, fromPrevBalanceInt-amount, fromCurrBalanceInt)
	assert.Equal(t, toPrevBalanceInt+amount, toCurrBalanceInt)

}

func TestApproveERC20BurnableUsingABIGen(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}
	client, _ := client.NewClient(conf)

	fromPrivateKey := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	fromKeyPair, _ := wallet.GetKeyPairFromPrivateKey(fromPrivateKey)

	toPrivateKey := "432024aaf30b6921f51f9c21ffad5485fa82550c7627fb2eb121ebe3f165648a"
	toKeyPair, _ := wallet.GetKeyPairFromPrivateKey(toPrivateKey)

	contractAddress := "0x5B0CFA20e02145ea529983c7954702Ab5E9e7949"

	amount := new(big.Int).SetUint64(uint64(50))

	receipt, err := ApproveERc20UsingABIGen(client, *fromKeyPair, contractAddress, toKeyPair.PublicKey, amount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)
	t.Log(receipt.TxHash.Hex())
}
