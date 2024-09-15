// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {PoolManager} from "v4-core/PoolManager.sol";
import {PoolSwapTest} from "v4-core/test/PoolSwapTest.sol";
import {CoWHook} from "../src/CowHook.sol";
import {Hooks} from "v4-core/libraries/Hooks.sol";
import {HookMiner} from "../test/HookMiner.sol";
import {MockERC20} from "solmate/src/test/utils/mocks/MockERC20.sol";
import {Currency} from "v4-core/types/Currency.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";

import "forge-std/console.sol";

contract HookDeployer is Script {
    PoolManager manager =
        PoolManager(0x7DD9eC4bE0e5459e1Be05a2C09c074379e61fD5e);
    PoolSwapTest swapRouter =
        PoolSwapTest(0x3Bc8AC81C4ecDcb9a9Dc2F76208C9f3a5aB6A6b1);

    Currency token0;
    Currency token1;

    PoolKey key;

    CoWHook hook = CoWHook(0x16206E4bc197A193755D35478e8F3BF6740C0088);

    function run() public {
        uint privateKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(privateKey);
        MockERC20 tokenA = MockERC20(0x0c3A3aB1addc53281f0F7cc338E668D3EEe9149F);
        MockERC20 tokenB = MockERC20(0x3a8511C697eedf6293c40BF7af46B6535fEF8384);

        tokenA.approve(address(hook), type(uint256).max);
        tokenB.approve(address(hook), type(uint256).max);

        if (address(tokenA) > address(tokenB)) {
            (token0, token1) = (
                Currency.wrap(address(tokenB)),
                Currency.wrap(address(tokenA))
            );
        } else {
            (token0, token1) = (
                Currency.wrap(address(tokenA)),
                Currency.wrap(address(tokenB))
            );
        }

        key = PoolKey({
            currency0: token0,
            currency1: token1,
            fee: 3000,
            tickSpacing: 120,
            hooks: hook
        });

        hook.placeOrder(key, 100, true, 0.001 ether, 10);

        vm.stopBroadcast();
    }
}