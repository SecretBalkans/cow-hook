pragma solidity ^0.8.26;

import {BaseHook} from "v4-periphery/src/base/hooks/BaseHook.sol";
import {Hooks} from "v4-core/libraries/Hooks.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {PoolId, PoolIdLibrary} from "v4-core/types/PoolId.sol";
import {BalanceDelta} from "v4-core/types/BalanceDelta.sol";
import {BeforeSwapDelta} from "v4-core/types/BeforeSwapDelta.sol";
import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IPoolManager} from "v4-core/interfaces/IPoolManager.sol";
import {Currency, CurrencyLibrary} from "v4-core/types/Currency.sol";

contract IcebergHook is BaseHook, ERC1155 {
    using PoolIdLibrary for PoolKey;

    //errors
    error InvalidOrder();

    struct OrderData {
        address user;
        PoolKey poolKey;
        int24 tick;
        bool zeroForOne;
        uint256 inputAmount;
    }

    mapping(PoolId poolId => mapping(int24 tickToSellAt => mapping(bool zeroForOne => uint256 inputAmount))) pendingOrders;

    mapping(uint256 positionId => uint256 claimsSupply) public claimTokensSupply;

    function getPositionId(
        PoolKey calldata key,
        int24 tick,
        bool zeroForOne
    ) public pure returns (uint256) {
        return
            uint256(keccak256(abi.encode(key.toId(), tick, zeroForOne)));
    }

    constructor(
        IPoolManager _manager,
        string memory _uri
    ) BaseHook(_manager) ERC1155(_uri) {}

    function getHookPermissions()
        public
        pure
        override
        returns (Hooks.Permissions memory)
    {
        return
            Hooks.Permissions({
                beforeInitialize: false,
                afterInitialize: false,
                beforeAddLiquidity: false,
                afterAddLiquidity: false,
                beforeRemoveLiquidity: false,
                afterRemoveLiquidity: false,
                beforeSwap: true,
                afterSwap: true,
                beforeDonate: false,
                afterDonate: false,
                beforeSwapReturnDelta: false,
                afterSwapReturnDelta: false,
                afterAddLiquidityReturnDelta: false,
                afterRemoveLiquidityReturnDelta: false
            });
    }

    /**
        * @param key - Key of the pool
        * @param tickToSellAt - Tick at which we want to execute order
        * @param zeroForOne - direction of the swap
        * @param inputAmount - swap amount
     */
    function placeOrder(PoolKey calldata key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount) public returns (int24){
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        pendingOrders[key.toId()][tick][zeroForOne] += inputAmount;

        //issue claim tokens to the users and account for total claimed tokens
        uint256 positionId = getPositionId(key, tick, zeroForOne);
        claimTokensSupply[positionId] += inputAmount;
        _mint(msg.sender, positionId, inputAmount, "");

        address sellToken = zeroForOne
            ? Currency.unwrap(key.currency0)
            : Currency.unwrap(key.currency1);
        IERC20(sellToken).transferFrom(msg.sender, address(this), inputAmount);

        return tick;
    }

    function cancelOrder(PoolKey calldata key, int24 tickToSellAt, bool zeroForOne) public returns (int24){
        //1. remove pending order from the mapping
        //2. decline total claim amount for the position
        //3. burn claim tokens
        //4. send tokens back to the user
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        uint256 positionId = getPositionId(key, tick, zeroForOne);

        uint256 positionTokens = balanceOf(msg.sender, positionId);
        if (positionTokens == 0) revert InvalidOrder();

        pendingOrders[key.toId()][tick][zeroForOne] -= positionTokens;

        claimTokensSupply[positionId] -= positionTokens;
        _burn(msg.sender, positionId, positionTokens);

        address sellToken = zeroForOne
            ? Currency.unwrap(key.currency0)
            : Currency.unwrap(key.currency1);
        IERC20(sellToken).transfer(msg.sender, positionTokens);

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

    //here we will check for CoW and noOp if possible to match orders
    function beforeSwap(address, PoolKey calldata, IPoolManager.SwapParams calldata, bytes calldata) override external returns (bytes4, BeforeSwapDelta, uint24){
        revert HookNotImplemented();
    }

    function afterSwap(address, PoolKey calldata, IPoolManager.SwapParams calldata, BalanceDelta, bytes calldata) override external returns (bytes4, int128) {
        //update tick value

        //call secret orderbook and check if some of the pending orders are executable

        //if the proof is valid execute pending orders
        return(this.afterSwap.selector, 0);
    }

}