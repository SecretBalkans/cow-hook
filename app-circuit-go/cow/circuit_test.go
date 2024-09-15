package cow

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/gwproto"
	"github.com/brevis-network/brevis-sdk/test"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestCircuit(t *testing.T) {
	app, err := sdk.NewBrevisApp()
	check(err)

	ec, err := ethclient.Dial("https://sepolia.infura.io/v3/6f3271b349f74f9487b77b5a51174e18")
	check(err)

	orderReceipt, err := ec.TransactionReceipt(context.Background(), common.HexToHash("0x5f01a517902cdaa49dfb980e7b619474106a455bc7aedc91246d5ee6e3d5c21e"))
	check(err)
	orderLog := orderReceipt.Logs[2]
	orderData := orderReceipt.Logs[2].Data

	/*fromToken := hexutils.BytesToHex(orderData[0:32])
	toToken := hexutils.BytesToHex(orderData[32:64])
	//feeRaw := orderData[64:96]
	//tickSpacingRaw := orderData[96:128]
	//hookAddressRaw := orderData[128:160]
	tickSellAt := new(big.Int).SetBytes(orderData[160:192]).Int64()
	zeroToOne := new(big.Int).SetBytes(orderData[192:224]).Int64()
	check(err)
	inputAmount := new(big.Int).SetBytes(orderData[224:256]).Int64()
	check(err)
	blockLimit := new(big.Int).SetBytes(orderData[256:288]).Int64()
	check(err)*/

	app.AddReceipt(sdk.ReceiptData{
		BlockNum: orderReceipt.BlockNumber,
		TxHash:   orderReceipt.TxHash,
		Fields: [4]sdk.LogFieldData{
			//{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 0, Value: common.BytesToHash(orderData[0:32])},
			//{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 1, Value: common.BytesToHash(orderData[0:32])},
			//tickSellAt
			{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 5, Value: common.BytesToHash(orderData[160:192])},
			//zeroToOne
			{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 6, Value: common.BytesToHash(orderData[192:224])},
			//inputAmount
			{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 7, Value: common.BytesToHash(orderData[224:256])},
		},
	})

	orderReceipt1, err := ec.TransactionReceipt(context.Background(), common.HexToHash("0x5f01a517902cdaa49dfb980e7b619474106a455bc7aedc91246d5ee6e3d5c21e"))
	check(err)
	orderLog1 := orderReceipt1.Logs[2]
	orderData1 := orderReceipt1.Logs[2].Data

	/*fromToken := hexutils.BytesToHex(orderData[0:32])
	toToken := hexutils.BytesToHex(orderData[32:64])
	//feeRaw := orderData[64:96]
	//tickSpacingRaw := orderData[96:128]
	//hookAddressRaw := orderData[128:160]
	tickSellAt := new(big.Int).SetBytes(orderData[160:192]).Int64()
	zeroToOne := new(big.Int).SetBytes(orderData[192:224]).Int64()
	check(err)
	inputAmount := new(big.Int).SetBytes(orderData[224:256]).Int64()
	check(err)
	blockLimit := new(big.Int).SetBytes(orderData[256:288]).Int64()
	check(err)*/

	app.AddReceipt(sdk.ReceiptData{
		BlockNum: orderReceipt.BlockNumber,
		TxHash:   orderReceipt.TxHash,
		Fields: [4]sdk.LogFieldData{
			//{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 0, Value: common.BytesToHash(orderData[0:32])},
			//{Contract: orderLog.Address, LogIndex: orderLog.Index, EventID: orderLog.Topics[0], IsTopic: false, FieldIndex: 1, Value: common.BytesToHash(orderData[0:32])},
			//tickSellAt
			{Contract: orderLog1.Address, LogIndex: orderLog1.Index, EventID: orderLog1.Topics[0], IsTopic: false, FieldIndex: 5, Value: common.BytesToHash(orderData1[160:192])},
			//zeroToOne
			{Contract: orderLog1.Address, LogIndex: orderLog1.Index, EventID: orderLog1.Topics[0], IsTopic: false, FieldIndex: 6, Value: common.BytesToHash(orderData1[192:224])},
			//inputAmount
			{Contract: orderLog1.Address, LogIndex: orderLog1.Index, EventID: orderLog1.Topics[0], IsTopic: false, FieldIndex: 7, Value: common.BytesToHash(orderData1[224:256])},
		},
	})

	appCircuit := &AppCircuit{}
	appCircuitAssignment := &AppCircuit{}

	circuitInput, err := app.BuildCircuitInput(appCircuit)
	check(err)

	///////////////////////////////////////////////////////////////////////////////
	// Testing
	///////////////////////////////////////////////////////////////////////////////

	//test.IsSolved(t, appCircuit, appCircuitAssignment, circuitInput)
	test.ProverSucceeded(t, appCircuit, appCircuitAssignment, circuitInput)
}

func TestE2E(t *testing.T) {
	app, err := sdk.NewBrevisApp()
	check(err)

	txHash := common.HexToHash(
		"0x6dc75e61220cc775aafa17796c20e49ac08030020fce710e3e546aa4e003454c")

	ethRPCUrl := os.Getenv("ETH_RPC")
	fmt.Printf("env.ETH_RPC: %s\n", ethRPCUrl)
	ethDial, err := ethclient.Dial(ethRPCUrl)
	sepoliaRPCUrl := os.Getenv("SEPOLIA_RPC")
	fmt.Printf("env.SEPOLIA_RPC: %s\n", sepoliaRPCUrl)
	ec, err := ethclient.Dial(sepoliaRPCUrl)
	check(err)
	tx, _, err := ethDial.TransactionByHash(context.Background(), txHash)
	check(err)
	receipt, err := ethDial.TransactionReceipt(context.Background(), txHash)
	check(err)
	from, err := types.Sender(types.NewLondonSigner(tx.ChainId()), tx)
	check(err)

	app.AddTransaction(sdk.TransactionData{
		Hash:                txHash,
		ChainId:             tx.ChainId(),
		BlockNum:            receipt.BlockNumber,
		Nonce:               tx.Nonce(),
		GasTipCapOrGasPrice: tx.GasTipCap(),
		GasFeeCap:           tx.GasFeeCap(),
		GasLimit:            tx.Gas(),
		From:                from,
		To:                  *tx.To(),
		Value:               tx.Value(),
	})

	appCircuit := &AppCircuit{}
	appCircuitAssignment := &AppCircuit{}

	circuitInput, err := app.BuildCircuitInput(appCircuitAssignment)
	check(err)

	///////////////////////////////////////////////////////////////////////////////
	// Testing
	///////////////////////////////////////////////////////////////////////////////

	//test.IsSolved(t, appCircuit, appCircuitAssignment, circuitInput)
	test.ProverSucceeded(t, appCircuit, appCircuitAssignment, circuitInput)

	///////////////////////////////////////////////////////////////////////////////
	// Compiling and Setup
	///////////////////////////////////////////////////////////////////////////////

	outDir := "./circuit-data/circuit-out/age"
	srsDir := "./circuit-data/kzgsrs"
	// The compiled circuit, proving key, and verifying key are saved to outDir, and
	// the downloaded SRS in the process is saved to srsDir
	compiledCircuit, pk, vk, err := sdk.Compile(appCircuit, outDir, srsDir)
	check(err)

	fmt.Println("compilation/setup complete")

	///////////////////////////////////////////////////////////////////////////////
	// Proving
	///////////////////////////////////////////////////////////////////////////////

	// Once you saved your ccs, pk, and vk files, you can read them back into memory
	// for use with the provided utils
	compiledCircuit, pk, vk, err = sdk.ReadSetupFrom(outDir)
	check(err)

	witness, publicWitness, err := sdk.NewFullWitness(appCircuitAssignment, circuitInput)
	check(err)
	proof, err := sdk.Prove(compiledCircuit, pk, witness)
	check(err)
	err = sdk.Verify(vk, publicWitness, proof) // optional
	check(err)

	///////////////////////////////////////////////////////////////////////////////
	// Initiating Brevis Request
	///////////////////////////////////////////////////////////////////////////////

	fmt.Println(">> Initiate Brevis Request")
	appContract := common.HexToAddress("0x73090023b8D731c4e87B3Ce9Ac4A9F4837b4C1bd")
	refundee := common.HexToAddress("0x164Ef8f77e1C88Fb2C724D3755488bE4a3ba4342")

	calldata, requestId, nonceBrevis, feeValue, err := app.PrepareRequest(vk, 11155111, 11155111, refundee, appContract, 400000, gwproto.QueryOption_ZK_MODE.Enum(), "")
	check(err)
	fmt.Printf("calldata 0x%x\nrequestId %x\nnonce %d\nfeeValue %d Wei\n", calldata, requestId, nonceBrevis, feeValue)

	///////////////////////////////////////////////////////////////////////////////
	// Submit Proof to Brevis
	///////////////////////////////////////////////////////////////////////////////
	fmt.Println(">> Submit Proof to Brevis")
	err = app.SubmitProof(proof)
	check(err)
	fmt.Println("Submitted proof. Instantiating wallet for sendRequest...")
	// [Call BrevisProof.sendRequest() with the above calldata] ////////////////////

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		log.Fatalf("Failed to get the nonce: %v", err)
	}

	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	chainID, err := ec.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	// Create the transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = feeValue          // Amount of ETH to send along with the request
	auth.GasLimit = uint64(400000) // Gas limit
	auth.GasPrice = gasPrice
	fmt.Printf("Public address: %s\nNonce: %d\nGasPrice: %d\nchainID: %d\n", fromAddress, nonce, gasPrice, chainID)

	// Contract address (replace with your contract's address)
	brevisContractAddress := common.HexToAddress("0x841ce48F9446C8E281D3F1444cB859b4A6D0738C")

	contractInstance, err := NewBrevisProof(brevisContractAddress, ec)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	var proofBuf bytes.Buffer
	n, err := proof.WriteTo(&proofBuf)
	if err != nil {
		log.Fatalf("Failed to write proof: %v", err)
	}
	fmt.Printf("Write proof to buffer: %d\n", n)
	proofData := proofBuf.Bytes()
	var proofArray32 [32]byte
	copy(proofArray32[:], proofData[:32])

	refundeeAddress := common.HexToAddress("0xae4976143Cd0886ce6d5347BC53C7890Cbd401B6") // Replace with refundee address
	callbackAddress := common.HexToAddress("0xD9A4fE38CC031Cfac4eE7C4319dFd3901AdD14BD") // CoWHookMock Sepolia address

	// Constructing Callback struct
	callback := IBrevisTypesCallback{
		Target: callbackAddress,
		Gas:    uint64(300000),
	}
	option := uint8(0)

	// Call the sendRequest method
	txSr, err := contractInstance.SendRequest(auth, requestId, nonceBrevis, refundeeAddress, callback, option)
	if err != nil {
		log.Fatalf("Failed to send request transaction: %v", err)
	}

	fmt.Printf("SendRequest Transaction submitted: %s\n", txSr.Hash().Hex())

	////////////////////////////////////////////////////////////////////////////////
	// Poll Brevis gateway for query status till the final proof is submitted
	// on-chain by Brevis and your contract is called

	submitTx, err := app.WaitFinalProofSubmitted(context.Background())
	check(err)
	fmt.Printf("tx hash %s\n", submitTx)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
