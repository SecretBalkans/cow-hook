// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {PoolManager} from "v4-core/PoolManager.sol";
import {PoolSwapTest} from "v4-core/test/PoolSwapTest.sol";
import {CoWHookMock} from "../src/CoWHookMock.sol";
import {Hooks} from "v4-core/libraries/Hooks.sol";
import {MockERC20} from "solmate/src/test/utils/mocks/MockERC20.sol";
import {Currency} from "v4-core/types/Currency.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";

import "forge-std/console.sol";

contract HookDeployer is Script {
    function run() public {
        vm.startBroadcast();
        CoWHookMock hook = new CoWHookMock();
        console.log("Hook address: ", address(hook));
        vm.stopBroadcast();
    }
}