package contract

import (
	"context"
	"math/big"
	"strconv"
	"testing"
	"tiny-blockchain-app/app/config"
	"tiny-blockchain-app/app/pkg/blockchain/client"
	"tiny-blockchain-app/app/pkg/wallet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

type User struct {
	key *wallet.KeyPair
}

func newUser(privateKey string) User {
	keyPair, _ := wallet.GenerateKeyPair(privateKey)
	return User{
		key: keyPair,
	}
}

var contractAddress string = "0xb9D171F81716ee2Ce29b85Ba44B3966992512Ec9"

func TestERC20_Deploy(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}

	cli, err := client.NewClient(conf)
	assert.Equal(t, nil, err)

	prik := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	pubk := "0xb5ff8c7f64c1cfddb68edc1006dd5c58b07e1448"
	user := newUser(prik)
	assert.Equal(t, common.HexToAddress(pubk), user.key.PublicKey)

	c := ERC20Constructor{
		Name:     "ERC20Burnable",
		Symbol:   "E2B",
		Decimals: 10,
	}

	// 다음에 사용할 논스값
	nonce, err := cli.PendingNonceAt(context.Background(), user.key.PublicKey)
	assert.Equal(t, nil, err)

	// 컨트랙트 배포
	response, err := DeployERC20Burnable(cli, *user.key, c)
	assert.Equal(t, nil, err)

	// 다음에 사용할 논스값을 배포 과정에 정확히 사용하였는가?
	assert.Equal(t, nonce, response.Tx.Nonce())

	// Transaction이 정상적으로 만들어졌는지?
	receipt, err := cli.TransactionReceipt(context.Background(), response.Tx.Hash())
	assert.Equal(t, uint64(1), receipt.Status)
	assert.Equal(t, nil, err)

	// update to global variable contractAddress
	t.Log(receipt.ContractAddress)

}

func TestERC20_Mint(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}

	cli, err := client.NewClient(conf)
	assert.Equal(t, nil, err)

	prik := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	pubk := "0xb5ff8c7f64c1cfddb68edc1006dd5c58b07e1448"
	user := newUser(prik)
	assert.Equal(t, common.HexToAddress(pubk), user.key.PublicKey)

	// mint 하기 전의 잔액 확인
	prevBalance, err := BalanceERC20Burnable(cli, contractAddress, user.key.PublicKey.Hex())
	assert.Equal(t, nil, err)

	// amount 만큼의 erc20 토큰 mint
	amount := 1000000
	receipt, err := MintERC20Burnable(cli, *user.key, contractAddress, user.key.PublicKey.Hex(), amount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)

	// amount 만큼의 erc20 토큰이 정상적으로 mint 되었는지 확인
	currBalance, err := BalanceERC20Burnable(cli, contractAddress, user.key.PublicKey.Hex())
	assert.Equal(t, nil, err)
	prevBalanceInt, _ := strconv.Atoi(prevBalance)
	currBalanceInt, _ := strconv.Atoi(currBalance)
	assert.Equal(t, prevBalanceInt+amount, currBalanceInt)
}

func TestERC20_balance(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}

	cli, err := client.NewClient(conf)
	assert.Equal(t, nil, err)

	prik := "a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4"
	pubk := "0xb5ff8c7f64c1cfddb68edc1006dd5c58b07e1448"
	user := newUser(prik)
	assert.Equal(t, common.HexToAddress(pubk), user.key.PublicKey)

	balance, err := BalanceERC20Burnable(cli, contractAddress, user.key.PublicKey.Hex())
	assert.Equal(t, nil, err)
	t.Log(balance)
}

func TestERC20_Transfer(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}

	cli, err := client.NewClient(conf)
	assert.Equal(t, nil, err)

	user1 := newUser("a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4")
	// user2 := newUser("432024aaf30b6921f51f9c21ffad5485fa82550c7627fb2eb121ebe3f165648a")
	user3 := newUser("ef213cc132f97fa8b513b02bab6fa7f770b8a5269a763a64531b3c8a90c3a11c")

	// user1_prevBalance := balance(user1, cli)
	// user2_prevBalance := balance(user2, cli)

	amount := 10

	// Transfer amount of erc20 from user1 -> user2
	receipt, err := TransferERC20UsingABIGen(
		cli,
		*user1.key,
		contractAddress,
		user3.key.PublicKey.Hex(),
		amount,
	)

	// // Transfer amount of erc20 from user2 -> user1
	// receipt, err := TransferERC20UsingABIGen(
	// 	cli,
	// 	*user2.key,
	// 	contractAddress,
	// 	user1.key.PublicKey.Hex(),
	// 	amount,
	// )
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)

	// user1_currBalance := balance(user1, cli)
	// user2_currBalance := balance(user2, cli)

	// assert.Equal(t, user1_currBalance+amount, user1_prevBalance)
	// assert.Equal(t, user2_currBalance-amount, user2_prevBalance)
}

func TestERC20_Approval(t *testing.T) {
	conf := config.Config{
		Endpoint: "http://127.0.0.1:22001",
	}

	cli, err := client.NewClient(conf)
	assert.Equal(t, nil, err)

	user1 := newUser("a94b325ab4bea563fb94bffcf855fffb0bbac1a8e481d80f6816c6215c48bde4")
	user2 := newUser("432024aaf30b6921f51f9c21ffad5485fa82550c7627fb2eb121ebe3f165648a")

	amount := 10
	approveAmount := new(big.Int).SetUint64(uint64(amount))

	// Approval approveAmount of tokens of user1 to spender user2
	receipt, err := ApproveERc20UsingABIGen(cli, *user1.key, contractAddress, user2.key.PublicKey, approveAmount)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), receipt.Status)
	t.Log(receipt.TxHash.Hex())

}

func balance(user User, cli *ethclient.Client) int {
	balance, _ := BalanceERC20Burnable(cli, contractAddress, user.key.PublicKey.Hex())
	result, _ := strconv.Atoi(balance)
	return result
}
