// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {PoolManager} from "v4-core/PoolManager.sol";
import {PoolSwapTest} from "v4-core/test/PoolSwapTest.sol";
import {PoolModifyLiquidityTest} from "v4-core/test/PoolModifyLiquidityTest.sol";
import {PoolDonateTest} from "v4-core/test/PoolDonateTest.sol";
import {PoolTakeTest} from "v4-core/test/PoolTakeTest.sol";
import {PoolClaimsTest} from "v4-core/test/PoolClaimsTest.sol";

import "forge-std/console.sol";

contract V4Deployer is Script {
    function run() public {
        uint privateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(privateKey);

        PoolManager manager = new PoolManager(500000);
        console.log("Deployed PoolManager at", address(manager));
        PoolSwapTest swapRouter = new PoolSwapTest(manager);
        console.log("Deployed PoolSwapTest at", address(swapRouter));

        vm.stopBroadcast();
    }
}