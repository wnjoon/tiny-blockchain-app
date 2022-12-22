// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity >=0.7.0 <0.9.0;

import "./ERC1400Burnable.sol";
import "./ERC20Burnable.sol";

contract Swap {

    address _owner;

    // 거래를 위한 TrasnferFrom 호출
    bytes4 private constant TRANSFER_FROM_FUNC =
        bytes4(keccak256("transferFrom(address,address,uint256)"));

    // 잔고 조회를 위한 BalanceOf 호출
    bytes4 private constant BALANCE_OF_FUNC =
        bytes4(keccak256("balanceOf(address)"));


    constructor() {
        _owner = msg.sender;
    }

    function owner() public view virtual returns (address) {
        return _owner;
    }

    event SwapSuccess(address indexed erc20Owner, uint256 erc20Amount, address indexed erc1400Owner, uint256 erc1400Amount);
    // event SwapTest(address erc20Token, address erc1400Token, address indexed erc1400Owner, address indexed erc20Owner, uint256 erc20Amount, uint256 erc1400Amount);
    // event calledFunction(bool isSuccess, bytes outputOfFunction);

    /*
     * Swap 가능 여부 확인
     */
    function isSwapAvailable(
        address erc20Token, address erc20Owner, uint256 erc20Amount, 
        address erc1400Token, address erc1400Owner, uint256 erc1400Amount
    ) public {

        // check erc20 owner's balance
        (bool successCallBalanceOfErc20, bytes memory dataCallBalanceOfErc20) = erc20Token.call(
            abi.encodeWithSelector(
                BALANCE_OF_FUNC, 
                erc20Owner
            )
        );
        require(successCallBalanceOfErc20, "failed to get balance of erc20");
        uint256 availableToSendErc20 = bytesToUint256(dataCallBalanceOfErc20);
        require(availableToSendErc20 >= erc20Amount, "not enough erc20");

        // check erc1400 owner's balance
        (bool successCallBalanceOfErc1400, bytes memory dataCallBalanceOfErc1400) = erc1400Token.call(
            abi.encodeWithSelector(
                BALANCE_OF_FUNC, 
                erc1400Owner
            )
        );
        require(successCallBalanceOfErc1400, "failed to get balance of erc1400");
        uint256 availableToSendErc1400 = bytesToUint256(dataCallBalanceOfErc1400);
        require(availableToSendErc1400 >= erc1400Amount, "not enough erc1400");
    }

    /*
     * ERC1400과 ERC20 교환
     */
    function swapToken(
        address erc20Token, address erc20Owner, uint256 erc20Amount, 
        address erc1400Token, address erc1400Owner, uint256 erc1400Amount
    ) public {

        require(msg.sender == _owner, "only owner can call swap");

        // send erc20 from erc20owner to erc1400owner
        (bool successCallTransferErc20, ) = erc20Token.call(
            abi.encodeWithSelector(
                TRANSFER_FROM_FUNC, 
                erc20Owner,
                erc1400Owner,
                erc20Amount
            )
        );
        require(successCallTransferErc20, "failed to transfer erc20");
        
        // send erc1400 from erc1400owner to erc20owner
        (bool successCallTransferErc1400, ) = erc1400Token.call(
            abi.encodeWithSelector(
                TRANSFER_FROM_FUNC, 
                erc1400Owner,
                erc20Owner,
                erc1400Amount
            )
        );
        require(successCallTransferErc1400, "failed to transfer erc1400");
        
        // Emit event
        emit SwapSuccess(erc20Owner, erc20Amount, erc1400Owner, erc1400Amount);
    }

    function bytesToAddress(bytes memory bys)
        private
        pure
        returns (address addr)
    {
        assembly {
            addr := mload(add(bys, 32))
        }
    }

    function bytesToUint256(bytes memory bys)
        private
        pure
        returns (uint256 amount)
    {
        return uint256(bytes32(bys));
    }
}