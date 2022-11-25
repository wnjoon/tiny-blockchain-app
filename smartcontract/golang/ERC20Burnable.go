// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20BurnableMetaData contains all meta data concerning the ERC20Burnable contract.
var ERC20BurnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200252e3803806200252e833981810160405281019062000037919062000371565b82828282828282828282828282600190805190602001906200005b929190620000e6565b50816002908051906020019062000074929190620000e6565b5080600360006101000a81548160ff021916908360ff16021790555050505050505033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050506200046f565b828054620000f4906200043a565b90600052602060002090601f01602090048101928262000118576000855562000164565b82601f106200013357805160ff191683800117855562000164565b8280016001018555821562000164579182015b828111156200016357825182559160200191906001019062000146565b5b50905062000173919062000177565b5090565b5b808211156200019257600081600090555060010162000178565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001ff82620001b4565b810181811067ffffffffffffffff82111715620002215762000220620001c5565b5b80604052505050565b60006200023662000196565b9050620002448282620001f4565b919050565b600067ffffffffffffffff821115620002675762000266620001c5565b5b6200027282620001b4565b9050602081019050919050565b60005b838110156200029f57808201518184015260208101905062000282565b83811115620002af576000848401525b50505050565b6000620002cc620002c68462000249565b6200022a565b905082815260208101848484011115620002eb57620002ea620001af565b5b620002f88482856200027f565b509392505050565b600082601f830112620003185762000317620001aa565b5b81516200032a848260208601620002b5565b91505092915050565b600060ff82169050919050565b6200034b8162000333565b81146200035757600080fd5b50565b6000815190506200036b8162000340565b92915050565b6000806000606084860312156200038d576200038c620001a0565b5b600084015167ffffffffffffffff811115620003ae57620003ad620001a5565b5b620003bc8682870162000300565b935050602084015167ffffffffffffffff811115620003e057620003df620001a5565b5b620003ee8682870162000300565b925050604062000401868287016200035a565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200045357607f821691505b6020821081036200046957620004686200040b565b5b50919050565b6120af806200047f6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806370a0823111610097578063a9059cbb11610066578063a9059cbb1461028b578063dd62ed3e146102bb578063f2fde38b146102eb578063f7b188a51461030757610100565b806370a08231146102155780638456cb59146102455780638da5cb5b1461024f57806395d89b411461026d57610100565b8063313ce567116100d3578063313ce567146101a157806340c10f19146101bf57806342966c68146101db5780635c975abb146101f757610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d610311565b60405161011a9190611617565b60405180910390f35b61013d600480360381019061013891906116d2565b6103a3565b60405161014a919061172d565b60405180910390f35b61015b6105b4565b6040516101689190611757565b60405180910390f35b61018b60048036038101906101869190611772565b6105bd565b604051610198919061172d565b60405180910390f35b6101a9610727565b6040516101b691906117e1565b60405180910390f35b6101d960048036038101906101d491906116d2565b61073e565b005b6101f560048036038101906101f091906117fc565b610841565b005b6101ff610943565b60405161020c919061172d565b60405180910390f35b61022f600480360381019061022a9190611829565b61095a565b60405161023c9190611757565b60405180910390f35b61024d6109a3565b005b610257610a93565b6040516102649190611865565b60405180910390f35b610275610abd565b6040516102829190611617565b60405180910390f35b6102a560048036038101906102a091906116d2565b610b4f565b6040516102b2919061172d565b60405180910390f35b6102d560048036038101906102d09190611880565b610b66565b6040516102e29190611757565b60405180910390f35b61030560048036038101906103009190611829565b610bed565b005b61030f610db2565b005b606060018054610320906118ef565b80601f016020809104026020016040519081016040528092919081815260200182805461034c906118ef565b80156103995780601f1061036e57610100808354040283529160200191610399565b820191906000526020600020905b81548152906001019060200180831161037c57829003601f168201915b5050505050905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040a9061196c565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610481576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610478906119d8565b60405180910390fd5b600082116104c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bb90611a44565b60405180910390fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105a29190611757565b60405180910390a36001905092915050565b60008054905090565b600081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561067e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067590611ab0565b60405180910390fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461070a9190611aff565b9250508190555061071c848484610ea2565b600190509392505050565b6000600360009054906101000a900460ff16905090565b3373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c590611ba5565b60405180910390fd5b6107d88282610ec8565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8836040516108359190611757565b60405180910390a35050565b3373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c890611ba5565b60405180910390fd5b6108db3382610eee565b3373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453836040516109389190611757565b60405180910390a350565b6000600660149054906101000a900460ff16905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2a90611ba5565b60405180910390fd5b6001600660146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25860405160405180910390a2565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060028054610acc906118ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610af8906118ef565b8015610b455780601f10610b1a57610100808354040283529160200191610b45565b820191906000526020600020905b815481529060010190602001808311610b2857829003601f168201915b5050505050905090565b6000610b5c338484610ea2565b6001905092915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b3373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7490611ba5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce390611c37565b60405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b3373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3990611ba5565b60405180910390fd5b6000600660146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167fff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e60405160405180910390a2565b610ead838383610f14565b610eb8838383610f61565b610ec3838383611252565b505050565b610ed460008383610f14565b610ede8282611257565b610eea60008383611252565b5050565b610efa82600083610f14565b610f0482826113e1565b610f1082600083611252565b5050565b610f1c610943565b15610f5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5390611ca3565b60405180910390fd5b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610fd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc790611d0f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361103f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103690611d7b565b60405180910390fd5b60008111611082576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107990611de7565b60405180910390fd5b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611104576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110fb90611e53565b60405180910390fd5b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461114f9190611aff565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546111e19190611e73565b925050819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516112459190611757565b60405180910390a3505050565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112bd90611f15565b60405180910390fd5b60008111611309576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130090611f81565b60405180910390fd5b8060008082825461131a9190611e73565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113709190611e73565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516113d59190611757565b60405180910390a35050565b60008111611424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141b90611fed565b60405180910390fd5b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156114a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149d90612059565b60405180910390fd5b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546114f59190611aff565b925050819055508060008082825461150d9190611aff565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115729190611757565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115b857808201518184015260208101905061159d565b838111156115c7576000848401525b50505050565b6000601f19601f8301169050919050565b60006115e98261157e565b6115f38185611589565b935061160381856020860161159a565b61160c816115cd565b840191505092915050565b6000602082019050818103600083015261163181846115de565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006116698261163e565b9050919050565b6116798161165e565b811461168457600080fd5b50565b60008135905061169681611670565b92915050565b6000819050919050565b6116af8161169c565b81146116ba57600080fd5b50565b6000813590506116cc816116a6565b92915050565b600080604083850312156116e9576116e8611639565b5b60006116f785828601611687565b9250506020611708858286016116bd565b9150509250929050565b60008115159050919050565b61172781611712565b82525050565b6000602082019050611742600083018461171e565b92915050565b6117518161169c565b82525050565b600060208201905061176c6000830184611748565b92915050565b60008060006060848603121561178b5761178a611639565b5b600061179986828701611687565b93505060206117aa86828701611687565b92505060406117bb868287016116bd565b9150509250925092565b600060ff82169050919050565b6117db816117c5565b82525050565b60006020820190506117f660008301846117d2565b92915050565b60006020828403121561181257611811611639565b5b6000611820848285016116bd565b91505092915050565b60006020828403121561183f5761183e611639565b5b600061184d84828501611687565b91505092915050565b61185f8161165e565b82525050565b600060208201905061187a6000830184611856565b92915050565b6000806040838503121561189757611896611639565b5b60006118a585828601611687565b92505060206118b685828601611687565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061190757607f821691505b60208210810361191a576119196118c0565b5b50919050565b7f617070726f766520746f20746865207a65726f20616464726573732e00000000600082015250565b6000611956601c83611589565b915061196182611920565b602082019050919050565b6000602082019050818103600083015261198581611949565b9050919050565b7f617070726f766520746f2073616d6520616464726573732e0000000000000000600082015250565b60006119c2601883611589565b91506119cd8261198c565b602082019050919050565b600060208201905081810360008301526119f1816119b5565b9050919050565b7f617070726f766520616d6f756e74206973207a65726f2e000000000000000000600082015250565b6000611a2e601783611589565b9150611a39826119f8565b602082019050919050565b60006020820190508181036000830152611a5d81611a21565b9050919050565b7f696e73756666696369656e7420616c6c6f77616e63652e000000000000000000600082015250565b6000611a9a601783611589565b9150611aa582611a64565b602082019050919050565b60006020820190508181036000830152611ac981611a8d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b0a8261169c565b9150611b158361169c565b925082821015611b2857611b27611ad0565b5b828203905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611b8f602183611589565b9150611b9a82611b33565b604082019050919050565b60006020820190508181036000830152611bbe81611b82565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611c21602683611589565b9150611c2c82611bc5565b604082019050919050565b60006020820190508181036000830152611c5081611c14565b9050919050565b7f636f6e7472616374206973207061757365642e00000000000000000000000000600082015250565b6000611c8d601383611589565b9150611c9882611c57565b602082019050919050565b60006020820190508181036000830152611cbc81611c80565b9050919050565b7f7472616e736665722066726f6d20746865207a65726f20616464726573732e00600082015250565b6000611cf9601f83611589565b9150611d0482611cc3565b602082019050919050565b60006020820190508181036000830152611d2881611cec565b9050919050565b7f7472616e7366657220746f20746865207a65726f20616464726573732e000000600082015250565b6000611d65601d83611589565b9150611d7082611d2f565b602082019050919050565b60006020820190508181036000830152611d9481611d58565b9050919050565b7f7472616e7366657220616d6f756e74206973207a65726f2e0000000000000000600082015250565b6000611dd1601883611589565b9150611ddc82611d9b565b602082019050919050565b60006020820190508181036000830152611e0081611dc4565b9050919050565b7f7472616e7366657220616d6f756e7420657863656564732062616c616e63652e600082015250565b6000611e3d602083611589565b9150611e4882611e07565b602082019050919050565b60006020820190508181036000830152611e6c81611e30565b9050919050565b6000611e7e8261169c565b9150611e898361169c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611ebe57611ebd611ad0565b5b828201905092915050565b7f6d696e7420746f20746865207a65726f20616464726573732e00000000000000600082015250565b6000611eff601983611589565b9150611f0a82611ec9565b602082019050919050565b60006020820190508181036000830152611f2e81611ef2565b9050919050565b7f6d696e7420616d6f756e74206973207a65726f2e000000000000000000000000600082015250565b6000611f6b601483611589565b9150611f7682611f35565b602082019050919050565b60006020820190508181036000830152611f9a81611f5e565b9050919050565b7f6275726e20616d6f756e74206973207a65726f2e000000000000000000000000600082015250565b6000611fd7601483611589565b9150611fe282611fa1565b602082019050919050565b6000602082019050818103600083015261200681611fca565b9050919050565b7f6275726e20616d6f756e7420657863656564732062616c616e63652e00000000600082015250565b6000612043601c83611589565b915061204e8261200d565b602082019050919050565b6000602082019050818103600083015261207281612036565b905091905056fea26469706673582212208939e1bc13364f5117feb4d7f442557f778e9196b2ec27819af721dafb7b938964736f6c634300080d0033",
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableMetaData.ABI instead.
var ERC20BurnableABI = ERC20BurnableMetaData.ABI

// ERC20BurnableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20BurnableMetaData.Bin instead.
var ERC20BurnableBin = ERC20BurnableMetaData.Bin

// DeployERC20Burnable deploys a new Ethereum contract, binding an instance of ERC20Burnable to it.
func DeployERC20Burnable(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *ERC20Burnable, error) {
	parsed, err := ERC20BurnableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20BurnableBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Burnable *ERC20BurnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Burnable *ERC20BurnableSession) Owner() (common.Address, error) {
	return _ERC20Burnable.Contract.Owner(&_ERC20Burnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Burnable *ERC20BurnableCallerSession) Owner() (common.Address, error) {
	return _ERC20Burnable.Contract.Owner(&_ERC20Burnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Burnable *ERC20BurnableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Paused() (bool, error) {
	return _ERC20Burnable.Contract.Paused(&_ERC20Burnable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Burnable *ERC20BurnableCallerSession) Paused() (bool, error) {
	return _ERC20Burnable.Contract.Paused(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Mint(&_ERC20Burnable.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Mint(&_ERC20Burnable.TransactOpts, account, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Burnable *ERC20BurnableSession) Pause() (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Pause(&_ERC20Burnable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Pause(&_ERC20Burnable.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Burnable *ERC20BurnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferOwnership(&_ERC20Burnable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferOwnership(&_ERC20Burnable.TransactOpts, newOwner)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC20Burnable *ERC20BurnableTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC20Burnable *ERC20BurnableSession) UnPause() (*types.Transaction, error) {
	return _ERC20Burnable.Contract.UnPause(&_ERC20Burnable.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) UnPause() (*types.Transaction, error) {
	return _ERC20Burnable.Contract.UnPause(&_ERC20Burnable.TransactOpts)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ERC20Burnable contract.
type ERC20BurnableBurnIterator struct {
	Event *ERC20BurnableBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableBurn represents a Burn event raised by the ERC20Burnable contract.
type ERC20BurnableBurn struct {
	Account common.Address
	From    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed account, address indexed from, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address, from []common.Address) (*ERC20BurnableBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Burn", accountRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableBurnIterator{contract: _ERC20Burnable.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed account, address indexed from, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ERC20BurnableBurn, account []common.Address, from []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Burn", accountRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableBurn)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed account, address indexed from, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseBurn(log types.Log) (*ERC20BurnableBurn, error) {
	event := new(ERC20BurnableBurn)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ERC20Burnable contract.
type ERC20BurnableMintIterator struct {
	Event *ERC20BurnableMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableMint represents a Mint event raised by the ERC20Burnable contract.
type ERC20BurnableMint struct {
	Account  common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed account, address indexed receiver, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address, receiver []common.Address) (*ERC20BurnableMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Mint", accountRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableMintIterator{contract: _ERC20Burnable.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed account, address indexed receiver, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ERC20BurnableMint, account []common.Address, receiver []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Mint", accountRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableMint)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed account, address indexed receiver, uint256 amount)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseMint(log types.Log) (*ERC20BurnableMint, error) {
	event := new(ERC20BurnableMint)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20Burnable contract.
type ERC20BurnableOwnershipTransferredIterator struct {
	Event *ERC20BurnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Burnable contract.
type ERC20BurnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20BurnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableOwnershipTransferredIterator{contract: _ERC20Burnable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20BurnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableOwnershipTransferred)
				if err := _ERC20Burnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20BurnableOwnershipTransferred, error) {
	event := new(ERC20BurnableOwnershipTransferred)
	if err := _ERC20Burnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Burnable contract.
type ERC20BurnablePausedIterator struct {
	Event *ERC20BurnablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnablePaused represents a Paused event raised by the ERC20Burnable contract.
type ERC20BurnablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ERC20BurnablePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnablePausedIterator{contract: _ERC20Burnable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20BurnablePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnablePaused)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) ParsePaused(log types.Log) (*ERC20BurnablePaused, error) {
	event := new(ERC20BurnablePaused)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the ERC20Burnable contract.
type ERC20BurnableUnPausedIterator struct {
	Event *ERC20BurnableUnPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableUnPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableUnPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableUnPaused represents a UnPaused event raised by the ERC20Burnable contract.
type ERC20BurnableUnPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterUnPaused(opts *bind.FilterOpts, account []common.Address) (*ERC20BurnableUnPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "UnPaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUnPausedIterator{contract: _ERC20Burnable.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *ERC20BurnableUnPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "UnPaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableUnPaused)
				if err := _ERC20Burnable.contract.UnpackLog(event, "UnPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnPaused is a log parse operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address indexed account)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseUnPaused(log types.Log) (*ERC20BurnableUnPaused, error) {
	event := new(ERC20BurnableUnPaused)
	if err := _ERC20Burnable.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
