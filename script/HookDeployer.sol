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

    function run() public {
        uint privateKey = vm.envUint("PRIVATE_KEY");

        uint160 flags = uint160(
            Hooks.BEFORE_SWAP_FLAG | Hooks.BEFORE_SWAP_RETURNS_DELTA_FLAG
        );
        console.log("Flags: ", flags);

        address CREATE2_DEPLOYER = 0x4e59b44847b379578588920cA78FbF26c0B4956C;
        (address hookAddress, bytes32 salt) = HookMiner.find(
            CREATE2_DEPLOYER,
            flags,
            type(CoWHook).creationCode,
            abi.encode(address(manager), "", 0x16206E4bc197A193755D35478e8F3BF6740C0088)
        );

        console.log("Hook pre mined address: ", hookAddress);

        vm.startBroadcast(privateKey);
        MockERC20 tokenA = new MockERC20("Token0", "TK0", 18);
        MockERC20 tokenB = new MockERC20("Token1", "TK1", 18);

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

        tokenA.approve(address(swapRouter), type(uint256).max);
        tokenB.approve(address(swapRouter), type(uint256).max);

        tokenA.mint(msg.sender, 100 * 10 ** 18);
        tokenB.mint(msg.sender, 100 * 10 ** 18);


        CoWHook hook = new CoWHook{salt: salt}(manager, "", 0x16206E4bc197A193755D35478e8F3BF6740C0088);
        console.log("CoWHook: ", address(hook));
        // require(address(hook) == hookAddress, "hook address mismatch");

        key = PoolKey({
            currency0: token0,
            currency1: token1,
            fee: 3000,
            tickSpacing: 120,
            hooks: hook
        });

        // the second argument here is SQRT_PRICE_1_1
        manager.initialize(key, 79228162514264337593543950336, new bytes(0));

        vm.stopBroadcast();
    }
}