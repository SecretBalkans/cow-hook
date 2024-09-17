package cow

import (
	"fmt"
	"github.com/brevis-network/brevis-sdk/sdk"
	"math/big"
)

type AppCircuit struct {
}

type PriceAcc = sdk.Tuple3[
	sdk.Uint248, // zeroToOne = 0 inputAmount sum for position
	sdk.Uint248, // zeroToOne = 1 inputAmount sum for position
	sdk.Uint248, // sqrtPrice0To1 for position
]

var _ sdk.AppCircuit = &AppCircuit{}

func (c *AppCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	return 4, 0, 0
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, input sdk.DataInput) error {
	// Fetch storage data from the input
	u248 := api.Uint248
	/*	order0 := input.Receipts.Raw[0]
		order1 := input.Receipts.Raw[1]


		//tickSellAt0 := api.ToUint248(order0.Fields[0].Value)
		//tickSellAt1 := api.ToUint248(order1.Fields[0].Value)

		sqrtPrice0To1, zeroForOne0, sqrtPriceX96 := c.getZeroForOneSqrtPriceX96(api, order0.Fields[3].Value)
		fmt.Println("sqrtPriceX960", sqrtPriceX96.Val)

		sqrtPrice1To0, zeroForOne1, sqrtPriceX961 := c.getZeroForOneSqrtPriceX96(api, order1.Fields[3].Value)
		fmt.Println("sqrtPriceX961", sqrtPriceX961.Val)

		fmt.Println("zeroForOne0", zeroForOne0.Val)
		fmt.Println("zeroForOne1", zeroForOne1.Val)

		inputAmount0 := api.ToUint248(order0.Fields[1].Value)
		inputAmount1 := api.ToUint248(order1.Fields[1].Value)

		fmt.Println("inputAmount0", inputAmount0.Val)
		fmt.Println("inputAmount1", inputAmount1.Val)
		Q96, _ := c.getQ96()
		maxBuyAmountToken0OfOrder1, _ := u248.Div(u248.Mul(inputAmount1, Q96), sqrtPrice0To1)
		//maxBuyAmount1Order0, _ := u248.Div(inputAmount0, tickSellAt0)
		fmt.Println("maxBuyAmountToken0OfOrder1", maxBuyAmountToken0OfOrder1.Val)

		matchAmount0 := u248.Select(
			u248.IsEqual(
				u248.Add(
					zeroForOne0,
					zeroForOne1,
				),
				sdk.ConstUint248(1),
			),
			u248.Select(u248.IsGreaterThan(maxBuyAmountToken0OfOrder1, inputAmount0), inputAmount0, maxBuyAmountToken0OfOrder1),
			sdk.ConstUint248(0),
		)

		matchAmount1, _ := u248.Div(u248.Mul(matchAmount0, Q96), sqrtPrice1To0)

		fmt.Println("matchAmount0", matchAmount0.Val)
		fmt.Println("matchAmount1", matchAmount1.Val)

		//api.OutputUint(256, positionId0)
		api.OutputUint(256, matchAmount1)
		//api.OutputUint(256, positionId1)
		api.OutputUint(256, matchAmount0)
	*/
	priceAccumulatorDS, _ := sdk.GroupBy(sdk.NewDataStream(api, input.Receipts), func(accumulator PriceAcc, receipt sdk.Receipt) (newAccumulator PriceAcc) {
		_, zeroForOne, sqrtPriceVal := c.getZeroForOneSqrtPriceX96(api, receipt.Fields[3].Value)
		inputAmount := api.ToUint248(receipt.Fields[1].Value)
		fmt.Println(zeroForOne.Val)
		return PriceAcc{
			F0: u248.Add(accumulator.F0, u248.Select(zeroForOne, inputAmount, sdk.ConstUint248(0))), // zeroForOne == 0 total input amount
			F1: u248.Add(accumulator.F1, u248.Select(zeroForOne, sdk.ConstUint248(0), inputAmount)), // zeroForOne == 1 total input amount
			F2: sqrtPriceVal,                                                                        // price
		}
	}, PriceAcc{
		F0: sdk.ConstUint248(0),
		F1: sdk.ConstUint248(0),
		F2: sdk.ConstUint248(0),
	}, func(receiptToMap sdk.Receipt) sdk.Uint248 {
		//position := api.ToBytes32(receiptToMap.Fields[0].Value)
		//fmt.Println("positionId", position.Val)
		//fmt.Println("inputAmount", api.ToUint248(receiptToMap.Fields[1].Value).Val)
		//fmt.Println("expiry", api.ToUint248(receiptToMap.Fields[2].Value).Val)

		sqrtPrice0To1, zeroToOne, _ := c.getZeroForOneSqrtPriceX96(api, receiptToMap.Fields[3].Value)
		fmt.Println("sqrtPrice0To1", sqrtPrice0To1, "zeroToOne", zeroToOne)
		return sqrtPrice0To1
	})
	fmt.Println("priceAccumulatorDS")
	priceAccumulatorDS.Show()
	filtered := sdk.Filter(priceAccumulatorDS, func(current PriceAcc) sdk.Uint248 {
		return u248.IsGreaterThan(current.F2, sdk.ConstUint248(0)) // HACK: bugfix 0 price PriceAcc entering the map
	})
	fmt.Println("priceAccumulatorDS(>1)")
	filtered.Show()
	sdk.Map(filtered, func(priceAccToMatch PriceAcc) sdk.Uint248 {
		inputAmount0 := priceAccToMatch.F0
		inputAmount1 := priceAccToMatch.F1
		sqrtPrice0To1 := priceAccToMatch.F2
		fmt.Println("calculate match inputAmount0", inputAmount0.Val)
		fmt.Println("calculate match inputAmount1", inputAmount1.Val)
		fmt.Println("calculate match sqrtPrice0To1", sqrtPrice0To1.Val)
		sqrtPrice1To0 := c.reversePrice(api, sqrtPrice0To1)
		Q96, _ := c.getQ96()
		maxBuyAmountToken0OfOrder1, _ := u248.Div(u248.Mul(inputAmount1, Q96), sqrtPrice0To1)
		//maxBuyAmount1Order0, _ := u248.Div(inputAmount0, tickSellAt0)
		fmt.Println("maxBuyAmountToken0OfOrder1", maxBuyAmountToken0OfOrder1.Val)

		matchAmount0 := u248.Select(u248.IsGreaterThan(maxBuyAmountToken0OfOrder1, inputAmount0), inputAmount0, maxBuyAmountToken0OfOrder1)

		matchAmount1, _ := u248.Div(u248.Mul(matchAmount0, Q96), sqrtPrice1To0)

		ordersAtPrice := sdk.Filter(sdk.NewDataStream(api, input.Receipts), func(current sdk.Receipt) sdk.Uint248 {
			_, _, sqrtPriceVal := c.getZeroForOneSqrtPriceX96(api, current.Fields[3].Value)
			return u248.IsEqual(priceAccToMatch.F2, sqrtPriceVal)
		})
		// This Filter doesn't work either
		ordersAtPrice.Show()
		sdk.Map(ordersAtPrice,
			func(current sdk.Receipt) sdk.Uint248 {
				fmt.Println("Receipt found!")
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
		return sdk.ConstUint248(0)
	})

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

func (c *AppCircuit) reversePrice(api *sdk.CircuitAPI, sqrtPrice0To1 sdk.Uint248) sdk.Uint248 {
	u248 := api.Uint248
	Q96, _ := c.getQ96()
	fmt.Println("sqrtPrice0To1", sqrtPrice0To1.Val)
	quotient, _ := u248.Div(u248.Mul(Q96, Q96), sqrtPrice0To1)
	return quotient
}

func (c *AppCircuit) getQ96() (sdk.Uint248, *big.Int) {
	nQ96 := new(big.Int).Lsh(big.NewInt(1), 96)
	Q96 := sdk.ConstUint248(nQ96) // 2^96
	return Q96, nQ96
}

func (c *AppCircuit) getZeroForOneSqrtPriceX96(api *sdk.CircuitAPI, value sdk.Bytes32) (sdk.Uint248, sdk.Uint248, sdk.Uint248) {
	u248 := api.Uint248
	maxuint160 := sdk.ConstUint248(new(big.Int).Lsh(big.NewInt(1), 160))
	zeroForOneWithSqrtPriceX96 := api.ToUint248(value)
	zeroForOne := u248.Select(
		u248.IsGreaterThan(
			zeroForOneWithSqrtPriceX96,
			maxuint160,
		),
		sdk.ConstUint248(1),
		sdk.ConstUint248(0),
	)
	sqrtPriceX96 := u248.Sub(zeroForOneWithSqrtPriceX96, u248.Mul(zeroForOne, maxuint160))
	Q96, _ := c.getQ96()
	reversePrice, _ := u248.Div(u248.Mul(Q96, Q96), sqrtPriceX96)
	sqrtPriceEff := u248.Select(zeroForOne, reversePrice, sqrtPriceX96)
	return sqrtPriceEff, zeroForOne, sqrtPriceX96
}
