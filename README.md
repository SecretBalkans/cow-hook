# Introduction
CoW hook uses smart indexing engine, which is able to match existing orders if there is a Coincidence of Wants. In order to achieve end-to-end trustless flow, it uses Brevis as a verifiability layer. If there is no CoW after certain blocks (specified by user), trade will be executed through AMM regularly.

# Test
To test the hook locally run: `forge test --via-ir`

# Architecture
CoWHook solution consists of hook smart contract and Brevis Circuit for order matching.

Hook implements before swap function to intersect the swap and place CoW orders.

If the orders are matched trades are executed p2p, without slippage.

# Progress
We implemented Brevis Cirtuit, which verifies the order matching algorithm.

We imagined to place order from beforeSwap hook, by using noop, but faced issues with claim token issuance (sender needs to be manager).

Need to think of alternative solutions.

### Update: 00:15am Found a critical bug in Brevis sdk.Filter:
#### Example 1
sdk.Map after filter called with values that shouldn't be included
```go
fmt.Println("priceAccumulatorDS")
priceAccumulatorDS.Show()

filtered := sdk.Filter(priceAccumulatorDS, func(current PriceAcc) sdk.Uint248 {
    return u248.IsGreaterThan(current.F2, sdk.ConstUint248(0)) // HACK: bugfix 0 price PriceAcc entering the map
})
fmt.Println("priceAccumulatorDS(>1)")
filtered.Show()
// Bug: Filtered should not include values equal to 1. It works the same with any other predicate
sdk.Map(filtered, func(priceAccToMatch PriceAcc) sdk.Uint248 {
    // Bug: This map is entered with filtered values
		inputAmount0 := priceAccToMatch.F0
		....
 
```
```text
priceAccumulatorDS
┌───┬────────────────────────────────────────────────────────────────────┬────────┐
│ # │ DATA                                                               │ TOGGLE │
├───┼────────────────────────────────────────────────────────────────────┼────────┤
│ 0 │ (300000000000001, 1, 79704936542881920863903188246)                │ 1      │
│ 1 │ (1, 200000000000001, 79704936542881920863903188246)                │ 1      │
│ 2 │ (100000000000001, 1000000000000001, 79228162514264337593543950336) │ 1      │
│ 3 │ (0, 0, 0)                                                          │ 0      │
└───┴────────────────────────────────────────────────────────────────────┴────────┘
priceAccumulatorDS(F2>1)
┌───┬────────────────────────────────────────────────────────────────────┬────────┐
│ # │ DATA                                                               │ TOGGLE │
├───┼────────────────────────────────────────────────────────────────────┼────────┤
│ 0 │ (300000000000001, 1, 79704936542881920863903188246)                │ 1      │
│ 1 │ (1, 200000000000001, 79704936542881920863903188246)                │ 1      │
│ 2 │ (100000000000001, 1000000000000001, 79228162514264337593543950336) │ 1      │
│ 3 │ (0, 0, 0)                                                          │ 0      │
└───┴────────────────────────────────────────────────────────────────────┴────────┘
```
#### Example 2
sdk.Map after filter called again not allowing to filter Receipts
The Toggle field seems to be correct but sdk.Map doesn't look at it perhaps.
```go
ordersAtPrice.Show()
sdk.Map(ordersAtPrice,
    func(current sdk.Receipt) sdk.Uint248 {
        fmt.Println("Receipt found!")
         // Bug: This map is entered with filtered values
        _, zeroForOne, _ := c.getZeroForOneSqrtPriceX96(api, current.Fields[3].Value)
        matchAmount := u248.Select(zeroForOne, matchAmount0, matchAmount1)
        expired := sdk.ConstUint248(0)
        //fmt.Println("positionId", current.Fields[0].Value)
        api.OutputBytes32(current.Fields[0].Value)
        api.OutputUint(248, matchAmount)
        //fmt.Println("matchAmount", matchAmount)
        api.OutputUint(248, expired)
        //fmt.Println("expired", expired)
        return sdk.ConstUint248(0)
})
```

```text
┌───┬─────────┬────────┐
│ # │ DATA    │ TOGGLE │
├───┼─────────┼────────┤
│ 0 │ Receipt │ 1      │
│ 1 │ Receipt │ 1      │
│ 2 │ Receipt │ 0      │
│ 3 │ Receipt │ 0      │
└───┴─────────┴────────┘
Receipt found!
added bytes32 output: aee45fa367f44258022cf9b7d8cacff4d6cd4d10f8bef958e23d9cc642c71ac4
added uint248 output: 0
added uint248 output: 0
Receipt found!
added bytes32 output: dad39c468f4b9dc058a102604d8eebdf6bc151eaf0d994096dc38dab6160c360
added uint248 output: 0
added uint248 output: 0
Receipt found!
added bytes32 output: 7847f7785fd0b98945ffe0c8dba44295de7107382f024fc0b8f09fee4fb58f0c
added uint248 output: 0
added uint248 output: 0
Receipt found!
added bytes32 output: be1007e9117ff8b445586944dabe07c9abefbd2c92b0927c99e15ff2c6ca0d7a
added uint248 output: 0
added uint248 output: 0
```
