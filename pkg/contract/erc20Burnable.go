package contract

import (
	"context"
	"math/big"
	"tiny-blockchain-app/pkg/wallet"
	smartcontract "tiny-blockchain-app/smartcontract/golang"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type ERC20Constructor struct {
	Name     string
	Symbol   string
	Decimals uint8
}

func DeployERC20Burnable(client *ethclient.Client, keyPair wallet.KeyPair, c ERC20Constructor) (*ContractResponse, error) {
	auth, err := GetAuth(client, keyPair)
	if err != nil {
		return nil, err
	}

	_, tx, instance, err := smartcontract.DeployERC20Burnable(auth, client, c.Name, c.Symbol, c.Decimals)
	if err != nil {
		return nil, err
	}

	response := &ContractResponse{
		Tx:       tx,
		Instance: instance,
	}

	address, err := checkDeployed(client, response)
	if err != nil {
		return nil, err
	}

	response.Address = address
	return response, nil
}

func MintERC20Burnable(client *ethclient.Client, keyPair wallet.KeyPair, contractAddress_ string, toAddress_ string, amount_ int) (*types.Receipt, error) {
	auth, err := GetAuth(client, keyPair)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(contractAddress_)
	instance, err := smartcontract.NewERC20Burnable(contractAddress, client)
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(toAddress_)
	tx, err := instance.Mint(auth, toAddress, big.NewInt(int64(amount_)))

	if err != nil {
		return nil, err
	}

	response := &ContractResponse{
		Tx:       tx,
		Instance: instance,
	}

	receipt, err := checkMinted(client, response)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func ApproveERc20UsingABIGen(client *ethclient.Client, keyPair wallet.KeyPair, contractAddress_ string, spender common.Address, amount *big.Int) (*types.Receipt, error) {
	auth, err := GetAuth(client, keyPair)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(contractAddress_)
	instance, err := smartcontract.NewERC20Burnable(contractAddress, client)
	if err != nil {
		return nil, err
	}

	tx, err := instance.Approve(auth, spender, amount)
	if err != nil {
		return nil, err
	}

	response := &ContractResponse{
		Tx:       tx,
		Instance: instance,
	}

	receipt, err := checkMinted(client, response)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func TransferERC20UsingABIGen(client *ethclient.Client, keyPair wallet.KeyPair, contractAddress_ string, toAddress_ string, amount_ int) (*types.Receipt, error) {
	auth, err := GetAuth(client, keyPair)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(contractAddress_)
	instance, err := smartcontract.NewERC20Burnable(contractAddress, client)
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(toAddress_)
	tx, err := instance.Transfer(auth, toAddress, big.NewInt(int64(amount_)))

	if err != nil {
		return nil, err
	}

	response := &ContractResponse{
		Tx:       tx,
		Instance: instance,
	}

	receipt, err := checkMinted(client, response)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func BalanceERC20Burnable(client *ethclient.Client, contractAddress_ string, address_ string) (string, error) {
	instance, err := smartcontract.NewERC20Burnable(common.HexToAddress(contractAddress_), client)
	if err != nil {
		return "", err
	}
	balance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address_))
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

func TransferERC20Burnable(client *ethclient.Client, keyPair wallet.KeyPair, ca string, to string, amount_ int) (*types.Receipt, error) {

	nonce, err := client.PendingNonceAt(context.Background(), keyPair.PublicKey)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(to)
	contractAddress := common.HexToAddress(ca)

	transferFnSignature := []byte("transfer(address, uint256")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := big.NewInt(int64(amount_))
	// amount := new(big.Int)
	// amount.SetString("10000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &toAddress,
	// 	Data: data,
	// })
	// if err != nil {
	// 	return "", err
	// }
	gasLimit := uint64(12500000)

	tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), keyPair.PrivateKey)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	////
	tx, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
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
