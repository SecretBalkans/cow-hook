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
import {Currency} from "v4-core/types/Currency.sol";
import {CurrencySettler} from "@uniswap/v4-core/test/utils/CurrencySettler.sol";
import {StateLibrary} from "v4-core/libraries/StateLibrary.sol";
import {FixedPointMathLib} from "solmate/src/utils/FixedPointMathLib.sol";
import {TickMath} from "v4-core/libraries/TickMath.sol";
import {BrevisApp} from "./BrevisApp.sol";

contract CoWHook is BaseHook, ERC1155, BrevisApp {
    using StateLibrary for IPoolManager;
    using PoolIdLibrary for PoolKey;
    using FixedPointMathLib for uint256;
    using CurrencySettler for Currency;
    using TickMath for int24;
    //events
    event OrderPlaced(
        uint256 positionId,
        PoolKey key,
        int24 tickToSellAt, 
        bool zeroForOne, 
        uint256 inputAmount,
        uint160 sqrtPriceX96,
        uint256 expiryBlock
    );

    //data for the CoW execution
    //we exchange the claim tokens between the two positions, respecting matched amounts
    struct SettleData {
        uint256 positionId;
        uint256 matchedAmount;
        uint256 expired;
    }

    struct PositionData {
        PoolKey key;
        int24 tick;
        bool zeroForOne;
        uint256 inputAmount;
    }

    //errors
    error InvalidOrder();
    error NothingToClaim();
    error NotEnoughToClaim();
    error BadCallbackData();
    error NotOwner();
    error PositionNotExpired();

    //modifiers
    modifier onlyOwner() {
        if(msg.sender != owner) revert NotOwner();
        _;
    }

    mapping(PoolId poolId => mapping(int24 tickToSellAt => mapping(bool zeroForOne => uint256 inputAmount))) public pendingOrders;

    mapping(uint256 positionId => uint256 claimsSupply) public claimTokensSupply;

    mapping(uint256 positionId => uint256 outputClaimable) public claimableOutputTokens;

    mapping(uint256 positionId => uint256 expiryBlock) public positionExpiryBlock;

    mapping(uint256 positionId => uint256 startingBlock) public positionStartingBlock;

    mapping(uint256 positionId => PositionData positionData) public positionData;

    address public owner;

    bytes32 public vkHash;

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
        string memory _uri,
        address _brevisRequest
    ) BaseHook(_manager) ERC1155(_uri) BrevisApp(_brevisRequest) {
        owner = msg.sender;
    }

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
        pendingOrders[key.toId()][tick][zeroForOne] += inputAmount;
        uint256 positionId = getPositionId(key, tick, zeroForOne);
        positionStartingBlock[positionId] = block.number;
        positionExpiryBlock[positionId] = block.number + blockLimit;
        positionData[positionId] = PositionData({
            key: key,
            tick: tick,
            zeroForOne: zeroForOne,
            inputAmount: inputAmount
        });

        //issue claim tokens to the users and account for total claimed tokens
        claimTokensSupply[positionId] += inputAmount;
        _mint(msg.sender, positionId, inputAmount, "");

        address sellToken = zeroForOne
            ? Currency.unwrap(key.currency0)
            : Currency.unwrap(key.currency1);
        IERC20(sellToken).transferFrom(msg.sender, address(this), inputAmount);

        uint160 sqrtPriceAtTick = tick.getSqrtPriceAtTick();

        emit OrderPlaced(positionId, key, tick, zeroForOne, inputAmount, sqrtPriceAtTick, positionExpiryBlock[positionId]);

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
        positionExpiryBlock[positionId] = 0;
        positionStartingBlock[positionId] = 0;

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


    /**
    * @param params - struct SwapParams {
                bool zeroForOne;
                int256 amountSpecified;
                uint160 sqrtPriceLimitX96;
            }
     */
    //here we will check for CoW and noOp if possible to match orders
    function beforeSwap(address sender, PoolKey calldata poolKey, IPoolManager.SwapParams calldata params, bytes calldata data) override external returns (bytes4, BeforeSwapDelta, uint24){
        //preventing reentrancy. Is this needed?
        // if (sender == address(this)) return (this.beforeSwap.selector, BeforeSwapDelta.unwrap(0), 0);

        (uint256 blockLimit, int24 tickToSellAt) = abi.decode(data, (uint256, int24));
        BeforeSwapDelta beforeSwapDelta;
        
        if(blockLimit > 0) {
            
            //this will skip the core logic of the swap, by setting beforeSwapDelta to amountSpecified, which will make swapAmount 0 when added to params.specifiedAmount
            beforeSwapDelta = toBeforeSwapDelta(int128(-params.amountSpecified),int128(params.amountSpecified));

            // //take the custody of the input tokens
            // if (params.zeroForOne) {
            //     // If user is selling Token 0 and buying Token 1

            //     // They will be sending Token 0 to the PM, creating a debit of Token 0 in the PM
            //     // We will take claim tokens for that Token 0 from the PM and keep it in the hook to create an equivalent credit for ourselves
            //     poolKey.currency0.take(
            //         poolManager,
            //         address(this),
            //         amountInOutPositive,
            //         true
            //     );

            //     // They will be receiving Token 1 from the PM, creating a credit of Token 1 in the PM
            //     // We will burn claim tokens for Token 1 from the hook so PM can pay the user
            //     poolKey.currency1.settle(
            //         poolManager,
            //         address(this),
            //         amountInOutPositive,
            //         true
            //     );
            // } else {
            //     poolKey.currency0.settle(
            //         poolManager,
            //         address(this),
            //         amountInOutPositive,
            //         true
            //     );
            //     poolKey.currency1.take(
            //         poolManager,
            //         address(this),
            //         amountInOutPositive,
            //         true
            //     );
            // }
            
            placeOrder(poolKey, tickToSellAt, params.zeroForOne, uint256(params.amountSpecified), blockLimit);
        }

        return (this.beforeSwap.selector, beforeSwapDelta, 0);
    }

    // function afterSwap(address, PoolKey calldata, IPoolManager.SwapParams calldata, BalanceDelta, bytes calldata) override external returns (bytes4, int128) {
    //     //update tick value

    //     //call secret orderbook and check if some of the pending orders are executable

    //     //if the proof is valid execute pending orders
    //     return(this.afterSwap.selector, 0);
    // }


    function executeCoW(SettleData memory settleData) internal {
        uint256 positionId = settleData.positionId;
        uint256 matchedAmount = settleData.matchedAmount;

        //matched amount should be in the output currency
        claimableOutputTokens[positionId] += matchedAmount;
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
        uint256 inputAmount
    ) internal {
        // Do the actual swap and settle all balances
        BalanceDelta delta = poolManager.swap(
            key,
            IPoolManager.SwapParams({
                zeroForOne: zeroForOne,
                // We provide a negative value here to signify an "exact input for output" swap
                amountSpecified: -int256(inputAmount),
                // No slippage limits (maximum slippage possible)
                sqrtPriceLimitX96: zeroForOne
                    ? TickMath.MIN_SQRT_PRICE + 1
                    : TickMath.MAX_SQRT_PRICE - 1
            }),
            ""
        );

        // `inputAmount` has been deducted from this position
        pendingOrders[key.toId()][tick][zeroForOne] -= inputAmount;
        uint256 positionId = getPositionId(key, tick, zeroForOne);
        positionExpiryBlock[positionId] = 0;
        positionStartingBlock[positionId] = 0;
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
        uint256 inputAmountToClaimFor
    ) external {
        int24 tick = getLowerUsableTick(tickToSellAt, key.tickSpacing);
        uint256 positionId = getPositionId(key, tick, zeroForOne);

        //position can be redeemed only after expiry
        if(positionExpiryBlock[positionId] > block.number) revert PositionNotExpired();

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

    function parseBytesToUint256Array(bytes memory data) public pure returns (uint256[] memory) {
        require(data.length % 32 == 0, "Invalid input length");

        uint256[] memory result = new uint256[](data.length / 32);

        for (uint256 i = 0; i < result.length; i++) {
            uint256 value;
            assembly {
                value := mload(add(data, add(32, mul(i, 32))))
            }
            result[i] = value;
        }

        return result;
    }

    function handleProofResult(
        bytes32 _vkHash,
        bytes calldata _circuitOutput
    ) internal override {
        //should return data about all the positions
        //we check the order conditions (each order should have specified block number until which CoW can be executed)
        //for the orders that didn't reach block limit, check if there is a matching order
        //for the orders that reached block limit, execute them through AMM

        require(vkHash == _vkHash, "invalid vk");

        //return just the matching orders, and read the rest from the storage

        (bytes memory settleData) = abi.decode(_circuitOutput, (bytes));

        uint256[] memory result = parseBytesToUint256Array(settleData);

        SettleData[50] memory matches;

        uint256 j = 0;
        // Use assembly to parse bytes into uint8 array
        for (uint256 i = 0; i < result.length;) {
            matches[j] = SettleData({
                positionId: result[i],
                matchedAmount: result[i+1],
                expired: result[i+2]
            });
            i += 3;
            j += 1;
        }

        //TODO: Go over each position in the proof and execute CoW, which will remove it from positionExpiryBlock

        //try to execute each position
        for(uint256 i = 0; i < matches.length; i++){
            executeCoW(matches[i]);
        }
    }

    function settleBalances(
        PoolKey calldata key,
        BalanceDelta delta,
        bool zeroForOne
    ) internal returns (BalanceDelta) {
        if (zeroForOne) {
            if (delta.amount0() < 0) {
                _settle(key.currency0, uint128(-delta.amount0()));
            }
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

    function decodeOutput(
        bytes calldata output
    ) internal pure returns (uint64, uint248) {
        uint64 minBlockNum = uint64(bytes8(output[0:8])); 
        uint248 sum = uint248(bytes31(output[8:8+31])); 
        return (minBlockNum, sum);
    }

    function setVkHash(bytes32 _vkHash) external onlyOwner {
        vkHash = _vkHash;
    }

}