import { waffle, ethers, web3 } from "hardhat";
import { expect } from "chai";
import { v4 as uuid } from "uuid";

import SwapArtifact from "../artifacts/contracts/Swap.sol/Swap.json";
import {Swap} from "../typechain-types/Swap";

import MultisigArtifact from "../artifacts/contracts/MultiSig.sol/MultiSig.json";
import { MultiSig } from "../typechain-types/MultiSig";

import ERC20TokenArtifact from "../../solidity/artifacts/contracts/ERC20Burnable.sol/ERC20Burnable.json";
import { ERC20Burnable } from "../../solidity/typechain-types/ERC20Burnable";

import ERC1400TokenArtifact from "../artifacts/contracts/ERC1400Basic.sol/ERC1400Basic.json";
import { ERC1400Basic } from "../typechain-types/ERC1400Basic";
import { BigNumber, BigNumberish, providers, Wallet, utils } from "ethers";

import { constants } from "ethers";

const { deployContract } = waffle;
const provider = waffle.provider;

const zero_address = "0x0000000000000000000000000000000000000000";
const ZERO_BYTES32 = "0x0000000000000000000000000000000000000000000000000000000000000000";
const [master, user1, user2, user3, confirmer, recover, swapper] = waffle.provider.getWallets()

let swap: Swap;
let multisig_user1: MultiSig;
let multisig_user2: MultiSig;
let multisig_user3: MultiSig;
let erc20: ERC20Burnable;
let erc1400: ERC1400Basic;


describe("Test Swap Contract", function() {

    context('\n1. Deploy Contracts', async () => {

        it('- erc20', async() => {
            erc20 = await deployContract(
                master,
                ERC20TokenArtifact,
                [
                    "ERC20",
                    "FT",
                    0
                ]
            ) as ERC20Burnable;

            expect (await erc20.owner()).to.be.equal(master.address);
            expect (ethers.utils.hexValue(erc20.address)).to.be.not.equal(ZERO_BYTES32);
        })

        it('- erc1400 ', async() => {
            erc1400 = await deployContract(
                master,
                ERC1400TokenArtifact,
                [
                    "ERC1400",
                    "ST",
                    zero_address
                ]
            ) as ERC1400Basic;

            expect (await erc1400.owner()).to.be.equal(master.address);
            expect (ethers.utils.hexValue(erc1400.address)).to.be.not.equal(ZERO_BYTES32);
        })

        it('- swap ', async() => {
            swap = await deployContract(
                master,
                SwapArtifact,
                []
            ) as Swap;

            expect (await swap.owner()).to.be.equal(master.address);
            expect (ethers.utils.hexValue(swap.address)).to.be.not.equal(ZERO_BYTES32);
        })
    })

    context('\n2. Mint erc20 tokens', async () => {
        it('- user1 : 100 ', async() => {
            await erc20.mint(user1.address, 100); 
            expect (await erc20.balanceOf(user1.address)).to.be.equal(100)
        })
    })

    context('\n3. Issue erc1400 tokens', async () => {
        it('- check is issuable ', async() => {
            expect (await erc1400.isIssuable()).to.be.equal(true)
        })
        it('- user2 : 50 ', async() => {
            await erc1400.issue(user2.address, 50, ZERO_BYTES32); 
            expect (await erc1400.balanceOf(user2.address)).to.be.equal(50)
        })
    })
    
    context('\n3. Approve 30 erc20 from user1 to swap contract', async () => {

        it('- reverted when approve to the zero address', async() => {
            await expect (erc20.connect(user1).approve(zero_address, 30)).to.be.revertedWith("approve to the zero address.");
        })

        it('- reverted when approve to same address', async() => {
            await expect (erc20.connect(user1).approve(user1.address, 30)).to.be.revertedWith("approve to same address.");
        })

        it('- reverted when approve more than 0', async() => {
            await expect (erc20.connect(user1).approve(user1.address, 0)).to.be.revertedWith("approve amount is zero.");
        })

        it('- success', async() => {
            await erc20.connect(user1).approve(swap.address, 30);
            expect (await erc20.allowance(user1.address, swap.address)).to.be.equal(30)
        })
    })

    context('\n4. Approve 10 erc1400 from user2 to swap contract', async () => {

        it('- reverted when approve to the zero address', async() => {
            await expect (erc1400.connect(user2).approve(zero_address, 10)).to.be.revertedWith("56");
        })

        it('- success approve 10 to swap contract ', async() => {
            await erc1400.connect(user2).approve(swap.address, 10);
            expect (await erc1400.allowance(user2.address, swap.address)).to.be.equal(10)
        })
    })

    context('\n5. Check swap is available', async () => {

        it('- reverted when not enough erc20', async() => {
            await expect (swap.connect(master)
                .isSwapAvailable(erc20.address, user1.address, 1000, erc1400.address, user2.address, 10))
                .to.be.revertedWith("not enough erc20");
        })

        it('- reverted when not enough erc1400', async() => {
            await expect (swap.connect(master)
                .isSwapAvailable(erc20.address, user1.address, 30, erc1400.address, user2.address, 1000))
                .to.be.revertedWith("not enough erc1400");
        })
    })

    context('\n5. Swap', async () => {
        
        it('- reverted when not owner call swap', async() => {
            await expect (swap.connect(user3)
                .swapToken(erc20.address, user1.address, 30, erc1400.address, user2.address, 10))
                .to.be.revertedWith("only owner can call swap");
        })

        it('- reverted when send erc20 more than approved ', async() => {
            await expect (swap.connect(master)
            .swapToken(erc20.address, user1.address, 50, erc1400.address, user2.address, 10))
            .to.be.reverted;
        })

        it('- reverted when send erc1400 more than approved ', async() => {
            await expect (swap.connect(master)
            .swapToken(erc20.address, user1.address, 30, erc1400.address, user2.address, 20))
            .to.be.reverted;
        })

        it('- success', async() => {
            // prev erc20 amount
            expect (await erc20.balanceOf(user1.address)).to.be.equal(100);
            expect (await erc20.balanceOf(user2.address)).to.be.equal(0);

            // prev erc1400 amount
            expect (await erc1400.balanceOf(user1.address)).to.be.equal(0);
            expect (await erc1400.balanceOf(user2.address)).to.be.equal(50);

            // call swapToken function
            await expect (swap.connect(master)
            .swapToken(erc20.address, user1.address, 20, erc1400.address, user2.address, 10))
            .to.emit(swap, "SwapSuccess")
            .withArgs(user1.address, 20, user2.address, 10);

            // check allowance
            expect (await erc20.allowance(user1.address, swap.address)).to.be.equal(10);
            expect (await erc1400.allowance(user2.address, swap.address)).to.be.equal(0);

            // result erc20 amount
            expect (await erc20.balanceOf(user1.address)).to.be.equal(80);
            expect (await erc20.balanceOf(user2.address)).to.be.equal(20);

            // result erc1400 amount
            expect (await erc1400.balanceOf(user1.address)).to.be.equal(10);
            expect (await erc1400.balanceOf(user2.address)).to.be.equal(40);
        })
    })
})
