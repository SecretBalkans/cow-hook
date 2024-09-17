// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Foundry libraries
import {Test} from "forge-std/Test.sol";

import {Deployers} from "@uniswap/v4-core/test/utils/Deployers.sol";
import {PoolSwapTest} from "v4-core/test/PoolSwapTest.sol";
import {MockERC20} from "solmate/src/test/utils/mocks/MockERC20.sol";

import {PoolManager} from "v4-core/PoolManager.sol";
import {IPoolManager} from "v4-core/interfaces/IPoolManager.sol";

import {PoolId, PoolIdLibrary} from "v4-core/types/PoolId.sol";
import {Currency, CurrencyLibrary} from "v4-core/types/Currency.sol";
import {StateLibrary} from "v4-core/libraries/StateLibrary.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";

import {Hooks} from "v4-core/libraries/Hooks.sol";
import {TickMath} from "v4-core/libraries/TickMath.sol";

// Our contracts
import {CoWHook, SettleData} from "../src/CoWHook.sol";

import "forge-std/console.sol";

contract CoWHookTest is Test, Deployers {
    // Use the libraries
    using StateLibrary for IPoolManager;
    using PoolIdLibrary for PoolKey;
    using CurrencyLibrary for Currency;
    using TickMath for int24;

    // The two currencies (tokens) from the pool
    Currency token0;
    Currency token1;

    CoWHook hook;

	function setUp() public {
        // Deploy v4 core contracts
        deployFreshManagerAndRouters();

        // Deploy two test tokens
        (token0, token1) = deployMintAndApprove2Currencies();

        // Deploy our hook
        uint160 flags = uint160(
            Hooks.BEFORE_SWAP_FLAG | Hooks.BEFORE_SWAP_RETURNS_DELTA_FLAG
        );
        address hookAddress = address(flags);
        deployCodeTo(
            "CoWHook.sol",
            abi.encode(manager, "", 0x16206E4bc197A193755D35478e8F3BF6740C0088),
            hookAddress
        );
        hook = CoWHook(hookAddress);

        // Approve our hook address to spend these tokens as well
        MockERC20(Currency.unwrap(token0)).approve(
            address(hook),
            type(uint256).max
        );
        MockERC20(Currency.unwrap(token1)).approve(
            address(hook),
            type(uint256).max
        );

        // Initialize a pool with these two tokens
        (key, ) = initPool(
            token0,
            token1,
            hook,
            3000,
            SQRT_PRICE_1_1,
            ZERO_BYTES
        );

        // Add initial liquidity to the pool

        // Some liquidity from -60 to +60 tick range
        modifyLiquidityRouter.modifyLiquidity(
            key,
            IPoolManager.ModifyLiquidityParams({
                tickLower: -60,
                tickUpper: 60,
                liquidityDelta: 10 ether,
                salt: bytes32(0)
            }),
            ZERO_BYTES
        );
        // Some liquidity from -120 to +120 tick range
        modifyLiquidityRouter.modifyLiquidity(
            key,
            IPoolManager.ModifyLiquidityParams({
                tickLower: -120,
                tickUpper: 120,
                liquidityDelta: 10 ether,
                salt: bytes32(0)
            }),
            ZERO_BYTES
        );
        // some liquidity for full range
        modifyLiquidityRouter.modifyLiquidity(
            key,
            IPoolManager.ModifyLiquidityParams({
                tickLower: TickMath.minUsableTick(60),
                tickUpper: TickMath.maxUsableTick(60),
                liquidityDelta: 10 ether,
                salt: bytes32(0)
            }),
            ZERO_BYTES
        );
    }

    function test_placeOrder() public {
        // Place a zeroForOne take-profit order
        // for 10e18 token0 tokens
        // at tick 100

        // console.log("Token0 address: ", address(Currency.unwrap(token0)));
        // console.log("Token1 address: ", address(Currency.unwrap(token1)));
        // int24 tick = 100;
        // uint256 amount = 10e18;
        // bool zeroForOne = true;
        // uint256 blockLimit = 10;

        // console.log("Current block: ", block.number);

        // // Note the original balance of token0 we have
        // uint256 originalBalance = token0.balanceOfSelf();

        // // Place the order
        // int24 tickLower = hook.placeOrder(key, tick, zeroForOne, amount, blockLimit);
        // int24 tickLowerSame = hook.placeOrder(key, tick, zeroForOne, amount, blockLimit);
        // int24 tickLower2 = hook.placeOrder(key, tick, !zeroForOne, 2*amount, blockLimit);

        // console.log("Balance of token0: ", token0.balanceOf(address(hook)));
        // console.log("Balance of token1: ", token1.balanceOf(address(hook)));

        // // Note the new balance of token0 we have
        // uint256 newBalance = token0.balanceOfSelf();

        // // Since we deployed the pool contract with tick spacing = 60
        // // i.e. the tick can only be a multiple of 60
        // // the tickLower should be 60 since we placed an order at tick 100
        // assertEq(tickLower, 60);

        // // Ensure that our balance of token0 was reduced by `amount` tokens
        // assertEq(originalBalance - newBalance, amount*2);

        // // Check the balance of ERC-1155 tokens we received
        // uint256 positionId = hook.getPositionId(key, tickLower, zeroForOne);
        // uint256 tokenBalance = hook.balanceOf(address(this), positionId);

        // uint160 price = tickLower.getSqrtPriceAtTick();
        
        // //print all of the below line

        // uint256 matchedAmount = (amount * price) / (2**96);
        // console.log(matchedAmount);
        // console.log(amount);
        // console.log(price);

        // uint256 claimableOutputTokensBefore = hook.claimableOutputTokens(positionId);
        // console.log("Claimable output tokens before: ", claimableOutputTokensBefore);

        // hook.executeCoW(SettleData({
        //     positionId: positionId,
        //     matchedAmount: matchedAmount,
        //     expired: 0
        // }));

        // uint256 claimableOutputTokensAfter = hook.claimableOutputTokens(positionId);
        // console.log("Claimable output tokens after: ", claimableOutputTokensAfter);

        // assertEq(claimableOutputTokensAfter, claimableOutputTokensBefore + matchedAmount);
        // // Ensure that we were, in fact, given ERC-1155 tokens for the order

        // vm.roll(block.number + 1000);

        // console.log("Claim tokens supply: ", hook.claimTokensSupply(positionId));

        // hook.redeem(
        //     key,
        //     tickLower,
        //     zeroForOne,
        //     amount
        // );

        // uint256 claimableOutputTokensAfterRedeem = hook.claimableOutputTokens(positionId);
        // console.log("Claimable output tokens after redeem: ", claimableOutputTokensAfterRedeem);

        // assertEq(claimableOutputTokensAfterRedeem, matchedAmount/2);

        // // equal to the `amount` of token0 tokens we placed the order for
        // assertTrue(positionId != 0);
        // assertEq(tokenBalance, amount*2);
    }

    function test_executeCoW() public {
        // Place a zeroForOne take-profit order
        // for 10e18 token0 tokens
        // at tick 100
        int24 tick = 100;
        uint256 amount = 10e18;
        bool zeroForOne = true;
        uint256 blockLimit = 10;

        console.log("Current block: ", block.number);

        // Note the original balance of token0 we have
        uint256 originalBalance = token0.balanceOfSelf();

        // Place the order
        int24 tickLower = hook.placeOrder(key, tick, zeroForOne, amount, blockLimit);

        // Note the new balance of token0 we have
        uint256 newBalance = token0.balanceOfSelf();

        // Since we deployed the pool contract with tick spacing = 60
        // i.e. the tick can only be a multiple of 60
        // the tickLower should be 60 since we placed an order at tick 100
        assertEq(tickLower, 60);

        // Ensure that our balance of token0 was reduced by `amount` tokens
        assertEq(originalBalance - newBalance, amount);

        // Check the balance of ERC-1155 tokens we received
        uint256 positionId = hook.getPositionId(key, tickLower, zeroForOne);
        uint256 tokenBalance = hook.balanceOf(address(this), positionId);

        // Ensure that we were, in fact, given ERC-1155 tokens for the order
        // equal to the `amount` of token0 tokens we placed the order for
        assertTrue(positionId != 0);
        assertEq(tokenBalance, amount);
    }

    function test_beforeSwap() public {
        int24 tick = -100;
        uint256 amount = 10 ether;
        bool zeroForOne = false;
        uint256 blockLimit = 10;

        // Place our order at tick -100 for 10e18 token1 tokens
        // int24 tickLower = hook.placeOrder(key, tick, zeroForOne, amount, blockLimit);

        IPoolManager.SwapParams memory params = IPoolManager.SwapParams({
            zeroForOne: true,
            amountSpecified: -1 ether,
            sqrtPriceLimitX96: TickMath.MIN_SQRT_PRICE + 1
        });

        PoolSwapTest.TestSettings memory testSettings = PoolSwapTest
            .TestSettings({takeClaims: false, settleUsingBurn: false});

        swapRouter.swap(key, params, testSettings, abi.encode(blockLimit, tick));

    }

    function onERC1155Received(
        address,
        address,
        uint256,
        uint256,
        bytes calldata
    ) external pure returns (bytes4) {
        return this.onERC1155Received.selector;
    }

    function onERC1155BatchReceived(
        address,
        address,
        uint256[] calldata,
        uint256[] calldata,
        bytes calldata
    ) external pure returns (bytes4) {
        return this.onERC1155BatchReceived.selector;
    }
}