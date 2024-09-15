package cow

import (
	"github.com/brevis-network/brevis-sdk/sdk"
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
	return 0, 0, 1
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	// Fetch storage data from the input
	//storageData := sdk.NewDataStream(api, input.StorageSlots)
	//u248 := api.Uint248
	//acc := Matching{}
	/*settlement := sdk.Reduce(storageData, acc, func(acc Matching, slot sdk.StorageSlot) (newAcc Matching) {
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
	txs := sdk.NewDataStream(api, in.Transactions)

	tx := sdk.GetUnderlying(txs, 0)
	// This is our main check logic
	api.Uint248.AssertIsEqual(tx.Nonce, sdk.ConstUint248(0))

	// Output variables can be later accessed in our app contract
	api.OutputAddress(tx.From)
	api.OutputUint(64, tx.BlockNum)
	return nil
}
