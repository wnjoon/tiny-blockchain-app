import { waffle } from "hardhat";
import { expect } from "chai";


import ERC20TokenArtifact from "../../artifacts/contracts/ERC20Burnable.sol/ERC20Burnable.json";
import { ERC20Burnable } from "../../typechain-types/ERC20Burnable";

const { deployContract } = waffle;
const provider = waffle.provider;


describe("ERC20BurnableToken", function() {
    let erc20Token: ERC20Burnable;
    // const provider = waffle.provider;
    const [user1, user2, user3] = waffle.provider.getWallets()
    
    // contract information
    const contractName = "ERC20BurnableToken";
    const contractSymbol = "E2B";
    const contractDecimal = 0;

    /*
     * beforeEach : 모든 테스트(it) 마다 해당 동작을 수행
     * before : 본 테스트파일을 실행하는 초기에 한번만 해당 동작으로 수행
     */
    before(async () => {
      erc20Token = await deployContract(
          user1,
          ERC20TokenArtifact,
          [
              contractName,
              contractSymbol,
              contractDecimal
          ]
      ) as ERC20Burnable;
        // user1에게 10000개의 토큰 발행
        // await erc20Token.mint(user1.address, 10000)
    })

    context('Deploy', async () => {
        it('- Check contract information ', async() => {
            expect(await erc20Token.name()).to.be.equal(contractName)
            expect(await erc20Token.symbol()).to.be.equal(contractSymbol)
            expect(await erc20Token.decimals()).to.be.equal(contractDecimal)
        })
    })
    context('Mint', async () => {
        it('- [ERROR] account != address(0) --> mint to the zero address.', async() => {
            await expect(erc20Token.mint("0x0000000000000000000000000000000000000000", 10000))
                .to.be.revertedWith("mint to the zero address.")
        })
        it('- [ERROR] amount > 0 --> mint amount is zero.', async() => {
            await expect(erc20Token.mint(user1.address, 0))
                .to.be.revertedWith("mint amount is zero.")
        })
        it('- Mint 10000 tokens to user1', async() => {
            await erc20Token.mint(user1.address, 10000)
            expect(await erc20Token.owner()).to.be.equal(user1.address)
            expect(await erc20Token.totalSupply()).to.be.equal(10000)
            expect(await erc20Token.balanceOf(user1.address)).to.be.equal(10000)
        })
    })   
    context('Transfer', async () => {
        it('- [ERROR] from != address(0) --> transfer from the zero address.', async() => {
            await expect(erc20Token.connect("0x0000000000000000000000000000000000000000")
                .transfer(user2.address, 10))
                .to.be.revertedWith("transfer from the zero address.")
        })
        it('- [ERROR] to != address(0) --> transfer to the zero address.', async() => {
            await expect(erc20Token.transfer("0x0000000000000000000000000000000000000000", 10))
                .to.be.revertedWith("transfer to the zero address.")
        })
        it('- [ERROR] amount > 0 --> transfer amount is zero.', async() => {
            await expect(erc20Token.transfer(user2.address, 0))
                .to.be.revertedWith("transfer amount is zero.")
        })
        it('- [ERROR] _balances[from] >= amount --> transfer amount exceeds balance.', async() => {
            await expect(erc20Token.transfer(user2.address, 10001))
                .to.be.revertedWith("transfer amount exceeds balance.")
        })
        it('- Transfer 3000 tokens to user2', async() => {
            const tx = await erc20Token.transfer(user2.address, 3000)
            const receipt = await tx.wait();
            

            expect(await erc20Token.balanceOf(user1.address)).to.be.equal(7000)
            expect(await erc20Token.balanceOf(user2.address)).to.be.equal(3000)

            console.log("Transaction Receipt : ", receipt)
        })        
    })
    context('Approve', async () => {
        it('- [ERROR] spender != address(0) --> approve to the zero address.', async() => {
            await expect(erc20Token.approve("0x0000000000000000000000000000000000000000", 1000))
                .to.be.revertedWith("approve to the zero address.")
        })
        it('- [ERROR] spender != msg.sender --> approve to same address.', async() => {
            await expect(erc20Token.approve(user1.address, 1000))
                .to.be.revertedWith("approve to same address.")
        })
        it('- [ERROR] amount > 0 --> approve amount is zero.', async() => {
            await expect(erc20Token.approve(user2.address, 0))
                .to.be.revertedWith("approve amount is zero.")
        })
        it('- Approve 1000 tokens of user2 to user1', async() => {
            await erc20Token.connect(user2).approve(user1.address, 1000)
            expect(await erc20Token.allowance(user2.address, user1.address)).to.be.equal(1000)
        })
    })
    context('TransferFrom', async () => {
        it('- [ERROR] _allowances[from][msg.sender] >= amount --> insufficient allowance.', async() => {
            await expect(erc20Token.transferFrom(user2.address, user3.address, 2000))
                .to.be.revertedWith("insufficient allowance.")
        })
        it('- Transfer 1000 tokens from user2 to user3 by user1', async() => {
            expect(await erc20Token.allowance(user2.address, user1.address)).to.be.equal(1000)
            await erc20Token.transferFrom(user2.address, user3.address, 1000);
            expect(await erc20Token.balanceOf(user2.address)).to.be.equal(2000)
            expect(await erc20Token.balanceOf(user3.address)).to.be.equal(1000)
        })
    })
    context('Pause', async () => {
        it('- Pause contract', async() => {
            await erc20Token.pause()
            expect(await erc20Token.paused()).to.be.equal(true)
        })
        it('- [ERROR] !paused() --> contract is paused.', async() => {
            await expect(erc20Token.transfer(user2.address, 1000)).to.be.revertedWith("contract is paused.")
        })
        
        it('- Unpause contract', async() => {
            await erc20Token.unPause()
            expect(await erc20Token.paused()).to.be.equal(false)
        })
    })
    context('Burn', async () => {
        it('- [ERROR] account != owner --> Replaced to Ownable: caller is not the owner.', async() => {
            await expect(erc20Token.connect("0x0000000000000000000000000000000000000000")
                .burn(1000))
                .to.be.revertedWith("Ownable: caller is not the owner.")
        })
        it('- [ERROR] amount > 0 --> burn amount is zero.', async() => {
            await expect(erc20Token.burn(0)).to.be.revertedWith('burn amount is zero.')
        })
        it('- [ERROR] _balances[account] >= amount --> burn amount exceeds balance.', async() => {
            await expect(erc20Token.burn(10000)).to.be.revertedWith('burn amount exceeds balance.')
        })
        it('- Burn 1000 tokens of user1', async() => {
            await erc20Token.burn(1000)
            expect(await erc20Token.balanceOf(user1.address)).to.be.equal(6000)
        })
    })
    context('TransferOwnership', async () => {
        it('- [ERROR] newOwner != address(0) --> Ownable: new owner is the zero address', async() => {
            await expect(erc20Token.transferOwnership("0x0000000000000000000000000000000000000000")).to.be.revertedWith('Ownable: new owner is the zero address')
        })
        it('- Transfer Ownership of contract to user2', async() => {
            await erc20Token.transferOwnership(user2.address)
            expect(await erc20Token.owner()).to.be.equal(user2.address)
        })
        it('- Burn 500 tokens of user2', async() => {
            await erc20Token.connect(user2).burn(500)
            expect(await erc20Token.balanceOf(user2.address)).to.be.equal(1500)
        })
        it('- Transfer Ownership of contract to user1', async() => {
            await erc20Token.connect(user2).transferOwnership(user1.address)
            expect(await erc20Token.owner()).to.be.equal(user1.address)
        })
    })

})