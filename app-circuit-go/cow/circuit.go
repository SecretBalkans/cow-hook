package cow

import (
	"fmt"
	"github.com/brevis-network/brevis-sdk/sdk"
	"math/big"
)

type AppCircuit struct {
}

type Matching = sdk.Tuple2[
	sdk.Bytes32, // poolId
	sdk.Tuple2[
		sdk.Bytes32, // tickToSellAt
		sdk.Tuple2[
			sdk.Bytes32, // zeroForOne
			sdk.Bytes32, // matching amount
		],
	],
]

var _ sdk.AppCircuit = &AppCircuit{}

func (c *AppCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	// Our app is only ever going to use one storage data at a time so
	// we can simply limit the max number of data for storage to 1 and
	// 0 for all others
	return 2, 0, 1
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, input sdk.DataInput) error {
	// Fetch storage data from the input
	order0 := input.Receipts.Raw[0]
	order1 := input.Receipts.Raw[1]
	u248 := api.Uint248

	//tickSellAt0 := api.ToUint248(order0.Fields[0].Value)
	//tickSellAt1 := api.ToUint248(order1.Fields[0].Value)

	// TODO: assert it's the same tick sell at
	nQ96 := new(big.Int).Lsh(big.NewInt(1), 96)
	// TODO: price sell at instead of tickSellAt0
	Q96 := sdk.ConstUint248(nQ96) // 2^96

	multiplier := big.NewInt(10001)
	denominator := big.NewInt(10000)
	// Multiply 2^96 by 1.0001
	result := new(big.Int).Mul(nQ96, multiplier)
	result.Div(result, denominator)

	sqrtPrice0To1 := sdk.ConstUint248(result)
	sqrtPrice1To0, _ := u248.Div(u248.Mul(Q96, Q96), sqrtPrice0To1) // 1 unit ot Token0 is price0To1 of Token1

	fmt.Println("sqrtPrice0To1", sqrtPrice0To1.Val)
	fmt.Println("sqrtPrice1To0", sqrtPrice1To0.Val)

	zeroToOne0 := api.ToUint248(order0.Fields[1].Value)
	zeroToOne1 := sdk.ConstUint248(0) // TODO: api.ToUint248(order1.Fields[1].Value)

	fmt.Println("zeroToOne0", zeroToOne0.Val)
	fmt.Println("zeroToOne1", zeroToOne1.Val)

	inputAmount0 := api.ToUint248(order0.Fields[2].Value)
	inputAmount1 := api.ToUint248(order1.Fields[2].Value)

	fmt.Println("inputAmount0", inputAmount0.Val)
	fmt.Println("inputAmount1", inputAmount1.Val)

	maxBuyAmountToken0OfOrder1, _ := u248.Div(u248.Mul(inputAmount1, Q96), sqrtPrice0To1)
	//maxBuyAmount1Order0, _ := u248.Div(inputAmount0, tickSellAt0)
	fmt.Println("maxBuyAmountToken0OfOrder1", maxBuyAmountToken0OfOrder1.Val)

	matchAmount0 := u248.Select(
		u248.IsEqual(
			u248.Add(
				zeroToOne0,
				zeroToOne1,
			),
			sdk.ConstUint248(1),
		),
		u248.Select(u248.IsGreaterThan(maxBuyAmountToken0OfOrder1, inputAmount0), inputAmount0, maxBuyAmountToken0OfOrder1),
		sdk.ConstUint248(0),
	)

	matchAmount1, _ := u248.Div(u248.Mul(matchAmount0, Q96), sqrtPrice1To0)

	fmt.Println("matchAmount0", matchAmount0.Val)
	fmt.Println("matchAmount1", matchAmount1.Val)
	api.OutputUint(256, matchAmount0)
	api.OutputUint(256, matchAmount1)

	/*storageData := sdk.NewDataStream(api, input.StorageSlots)
	acc := Matching{}
	settlement := sdk.Reduce(storageData, acc, func(acc Matching, slot sdk.StorageSlot) (newAcc Matching) {
		// Decode storage slot values here
		poolId := slot.Slot.Values()[0:32]
		tickToSellAt := slot.Slot.Values()[32:36]
		zeroForOne := slot.Slot.Values()[36:37]
		inputAmount := slot.Value
		sdk.ZipMap3()
		currAmount := sdk.Filter(acc.Values())
		newAmount := u248.Add(inputAmount)
		if u248.IsGreaterThan(api.ToUint248(inputAmount), sdk.ConstUint248(0)) == sdk.ConstUint248(1) {
			// Update the accumulator with matching amount
			currentAmount := acc
			newAmount := api.Add(currentAmount, api.ToUint248(inputAmount))
			acc.SetUint(poolId, tickToSellAt, zeroForOne, newAmount)
		}

		return Matching{F0: api.ToBytes32(poolId), F1: api.ToBytes32(tickToSellAt), F2: api.ToBytes32(newAmount)}
	})*/
	/*txs := sdk.NewDataStream(api, in.Transactions)

	tx := sdk.GetUnderlying(txs, 0)
	// This is our main check logic
	api.Uint248.AssertIsEqual(tx.Nonce, sdk.ConstUint248(0))

	// Output variables can be later accessed in our app contract
	api.OutputAddress(tx.From)
	api.OutputUint(64, tx.BlockNum)*/
	return nil
}
