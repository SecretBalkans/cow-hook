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

contract CoWHook is BaseHook, ERC1155 {
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
    error NothingToClaim();
    error NotEnoughToClaim();
    error BadCallbackData();

    struct OrderData {
        address user;
        PoolKey poolKey;
        int24 tick;
        bool zeroForOne;
        uint256 inputAmount;
    }

    mapping(PoolId poolId => mapping(int24 tickToSellAt => mapping(bool zeroForOne => mapping(uint256 blockLimit => uint256 inputAmount)))) pendingOrders;

    mapping(uint256 positionId => uint256 claimsSupply) public claimTokensSupply;

    mapping(uint256 positionId => uint256 outputClaimable) public claimableOutputTokens;

    function getPositionId(
        PoolKey calldata key,
        int24 tick,
        bool zeroForOne,
        uint256 blockLimit //users specify this value
    ) public pure returns (uint256) {
        return
            uint256(keccak256(abi.encode(key.toId(), tick, zeroForOne, blockLimit)));
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
                afterSwap: false,
                beforeDonate: false,
                afterDonate: false,
                beforeSwapReturnDelta: true,
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
    function placeOrder(PoolKey calldata key, int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) public returns (int24){
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        pendingOrders[key.toId()][tick][zeroForOne][blockLimit] += inputAmount;

        //issue claim tokens to the users and account for total claimed tokens
        uint256 positionId = getPositionId(key, tick, zeroForOne, blockLimit);
        claimTokensSupply[positionId] += inputAmount;
        _mint(msg.sender, positionId, inputAmount, "");

        address sellToken = zeroForOne
            ? Currency.unwrap(key.currency0)
            : Currency.unwrap(key.currency1);
        IERC20(sellToken).transferFrom(msg.sender, address(this), inputAmount);

        emit OrderPlaced(key, tick, zeroForOne, inputAmount, blockLimit);

        return tick;
    }

    function cancelOrder(PoolKey calldata key, int24 tickToSellAt, bool zeroForOne, uint256 blockLimit) public returns (int24){
        //1. remove pending order from the mapping
        //2. decline total claim amount for the position
        //3. burn claim tokens
        //4. send tokens back to the user
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        uint256 positionId = getPositionId(key, tick, zeroForOne, blockLimit);

        uint256 positionTokens = balanceOf(msg.sender, positionId);
        if (positionTokens == 0) revert InvalidOrder();

        pendingOrders[key.toId()][tick][zeroForOne][blockLimit] -= positionTokens;

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
    function beforeSwap(address sender, PoolKey calldata poolKey, IPoolManager.SwapParams calldata params, bytes calldata data) override external returns (bytes4, BeforeSwapDelta, uint24){
        //this will skip the core logic of the swap, by setting beforeSwapDelta to amountSpecified, which will make swapAmount 0 when added to params.specifiedAmount
        BeforeSwapDelta beforeSwapDelta = toBeforeSwapDelta(int128(-params.amountSpecified),int128(params.amountSpecified));

        //preventing reentrancy. Is this needed?
        if (sender == address(this)) return (this.beforeSwap.selector, beforeSwapDelta, 0);
        
        //take the custody of the input tokens

        (int24 tickToSellAt, bool zeroForOne, uint256 inputAmount, uint256 blockLimit) = abi.decode(data, (int24, bool, uint256, uint256));
        placeOrder(poolKey, tickToSellAt, zeroForOne, inputAmount, blockLimit);

        //our server will call the brevis

        return (this.beforeSwap.selector, beforeSwapDelta, 0);
    }

    // function afterSwap(address, PoolKey calldata, IPoolManager.SwapParams calldata, BalanceDelta, bytes calldata) override external returns (bytes4, int128) {
    //     //update tick value

    //     //call secret orderbook and check if some of the pending orders are executable

    //     //if the proof is valid execute pending orders
    //     return(this.afterSwap.selector, 0);
    // }

    function executeCoW(PoolKey calldata poolKey, bool zeroForOne, int24 tick, uint256 blockLimit, uint256 matchedAmount) internal {
        //remember that the hook has the custody over both tokens
        
        address sellToken = zeroForOne ? Currency.unwrap(poolKey.currency0) : Currency.unwrap(poolKey.currency1);
        address buyToken = zeroForOne ? Currency.unwrap(poolKey.currency1) : Currency.unwrap(poolKey.currency0);

        //send sellToken to another position -> the hook should do internal accounting. Then when users want to claim using their 1155, the accounting logic will be applied
        
        //we should se. the outputTokens for both positions. Need to calculate appropriate amounts using the tick
        uint256 position1Id = getPositionId(poolKey, tick, zeroForOne, blockLimit);
        uint256 position2Id = getPositionId(poolKey, tick, !zeroForOne, blockLimit);

        //how is the tick converted to price?
        claimableOutputTokens[position1Id] += 0; //buyToken
        claimableOutputTokens[position2Id] += 0; //sellToken

        //take token0 from position1 and send them to position2
        //take token1
    }

    function brevisCallback(bytes calldata callbackData) external {
        //should return data about all the positions
        //we check the order conditions (each order should have specified block number until which CoW can be executed)
        //for the orders that didn't reach block limit, check if there is a matching order
        //for the orders that reached block limit, execute them through AMM

        //verify proof, by calling brevis handleProof function

        //return just the matching orders, and read the rest from the storage

        // (PoolKey[] memory poolKey, 
        // int24[] memory tick, 
        // bool[] memory zeroForOne, 
        // uint256[] memory blockLimit, 
        // uint256[] memory amount, 
        // uint256[] memory matchedAmount, 
        // bytes memory proof) = abi.decode(callbackData, (PoolKey[], int24[], bool[], uint256[], uint256[], uint256[], bytes));
        
        // if(poolKey.length != 0 && (poolKey.length != tick.length || poolKey.length == zeroForOne.length || poolKey.length == blockLimit.length || poolKey.length == amount.length)) {
        //     revert BadCallbackData();
        // }

        // //try to execute each position
        // for(uint256 i = 0; i < poolKey.length; i++){
        //     if(block.timestamp > blockLimit[i]){

        //         // PoolKey calldata newPoolKey = PoolKey({
        //         //     currency0: poolKey[i].currency0,
        //         //     currency1: poolKey[i].currency1,
        //         //     fee: poolKey[i].fee,
        //         //     tickSpacing: poolKey[i].tickSpacing,
        //         //     hooks: Hooks
        //         // });
        //         //execute regular AMM swap
        //         executeAMMSwap(poolKey[i], tick[i], zeroForOne[i], amount[i], blockLimit[i]);
        //     }else{
        //         //check if there is a CoW matching order
        //         if(matchedAmount[i] != 0){
        //             executeCoW(poolKey[i], zeroForOne[i], tick[i], blockLimit[i], matchedAmount[i]);
        //         }
        //     }
        // }
    }

    function swapAndSettleBalances(
        PoolKey calldata key,
        IPoolManager.SwapParams memory params
    ) internal returns (BalanceDelta) {
        // Conduct the swap inside the Pool Manager
        BalanceDelta delta = poolManager.swap(key, params, "");

        // If we just did a zeroForOne swap
        // We need to send Token 0 to PM, and receive Token 1 from PM
        if (params.zeroForOne) {
            // Negative Value => Money leaving user's wallet
            // Settle with PoolManager
            if (delta.amount0() < 0) {
                _settle(key.currency0, uint128(-delta.amount0()));
            }

            // Positive Value => Money coming into user's wallet
            // Take from PM
            if (delta.amount1() > 0) {
                _take(key.currency1, uint128(delta.amount1()));
            }
        } else {
            if (delta.amount1() < 0) {
                _settle(key.currency1, uint128(-delta.amount1()));
            }

            if (delta.amount0() > 0) {
                _take(key.currency0, uint128(delta.amount0()));
            }
        }

        return delta;
    }

    function _settle(Currency currency, uint128 amount) internal {
        // Transfer tokens to PM and let it know
        poolManager.sync(currency);
        currency.transfer(address(poolManager), amount);
        poolManager.settle();
    }

    function _take(Currency currency, uint128 amount) internal {
        // Take tokens out of PM to our hook contract
        poolManager.take(currency, address(this), amount);
    }

    function executeAMMSwap(
        PoolKey calldata key,
        int24 tick,
        bool zeroForOne,
        uint256 inputAmount,
        uint256 blockLimit
    ) internal {
        // Do the actual swap and settle all balances
        BalanceDelta delta = swapAndSettleBalances(
            key,
            IPoolManager.SwapParams({
                zeroForOne: zeroForOne,
                // We provide a negative value here to signify an "exact input for output" swap
                amountSpecified: -int256(inputAmount),
                // No slippage limits (maximum slippage possible)
                sqrtPriceLimitX96: zeroForOne
                    ? TickMath.MIN_SQRT_PRICE + 1
                    : TickMath.MAX_SQRT_PRICE - 1
            })
        );

        // `inputAmount` has been deducted from this position
        pendingOrders[key.toId()][tick][zeroForOne][blockLimit] -= inputAmount;
        uint256 positionId = getPositionId(key, tick, zeroForOne, blockLimit);
        uint256 outputAmount = zeroForOne
            ? uint256(int256(delta.amount1()))
            : uint256(int256(delta.amount0()));

        // `outputAmount` worth of tokens now can be claimed/redeemed by position holders
        claimableOutputTokens[positionId] += outputAmount;
    }

    function redeem(
        PoolKey calldata key,
        int24 tickToSellAt,
        bool zeroForOne,
        uint256 inputAmountToClaimFor,
        uint256 blockLimit
    ) external {
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        uint256 positionId = getPositionId(key, tick, zeroForOne, blockLimit);

        if(claimableOutputTokens[positionId] == 0) revert NothingToClaim();
    
        uint256 positionTokens = balanceOf(msg.sender, positionId);
        if(positionTokens < inputAmountToClaimFor) revert NotEnoughToClaim();

        uint256 totalInputAmountForPosition = claimTokensSupply[positionId];
        uint256 totalClaimableForPosition = claimableOutputTokens[positionId];
        
        uint256 outputAmount = inputAmountToClaimFor.mulDivDown(totalClaimableForPosition, totalInputAmountForPosition);

        //reduce the claimable output tokens amount
        //reduce claim token total supply for position
        //burn claim tokens
        claimableOutputTokens[positionId] -= outputAmount;
        claimTokensSupply[positionId] -= inputAmountToClaimFor;
        _burn(msg.sender, positionId, inputAmountToClaimFor);

        //transfer output tokens to the user
        Currency token = zeroForOne ? key.currency0 : key.currency1;
        token.transfer(msg.sender, outputAmount);
    }

}