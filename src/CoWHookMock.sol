pragma solidity ^0.8.26;

import {BaseHook} from "v4-periphery/src/base/hooks/BaseHook.sol";
import {Hooks} from "v4-core/libraries/Hooks.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {PoolId, PoolIdLibrary} from "v4-core/types/PoolId.sol";
import {BalanceDelta} from "v4-core/types/BalanceDelta.sol";
import {BeforeSwapDelta, toBeforeSwapDelta} from "v4-core/types/BeforeSwapDelta.sol";
import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IPoolManager} from "v4-core/interfaces/IPoolManager.sol";
import {Currency, CurrencyLibrary} from "v4-core/types/Currency.sol";
import {FixedPointMathLib} from "solmate/src/utils/FixedPointMathLib.sol";
import {TickMath} from "v4-core/libraries/TickMath.sol";

contract CoWHookMock {
    using PoolIdLibrary for PoolKey;
    using FixedPointMathLib for uint256;

    //events
    event OrderPlaced(
        PoolKey key,
        int24 tickToSellAt, 
        bool zeroForOne, 
        uint256 inputAmount, 
        uint256 blockLimit
    );

    //errors
    error InvalidOrder();

    mapping(PoolId poolId => mapping(int24 tickToSellAt => mapping(bool zeroForOne => mapping(uint256 blockLimit => uint256 inputAmount)))) pendingOrders;

    constructor(){}

    function placeOrder(PoolKey calldata key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) public returns (int24){
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        pendingOrders[key.toId()][tick][zeroForOne][blockLimit] += inputAmount;

        emit OrderPlaced(key, tick, zeroForOne, inputAmount, blockLimit);

        return tick;
    }

    function getLowerUsableTick(
        int24 tick,
        int24 tickSpacing
    ) private pure returns (int24) {
        int24 intervals = tick / tickSpacing;
        if (tick < 0 && tick % tickSpacing != 0) intervals--;
        return intervals * tickSpacing;
    }

    function brevisCallback(bytes calldata callbackData) external {
    }
}