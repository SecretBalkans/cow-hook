package cow

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BrevisProofData is an auto generated low-level Go binding around an user-defined struct.
type BrevisProofData struct {
	CommitHash    [32]byte
	VkHash        [32]byte
	AppCommitHash [32]byte
	AppVkHash     [32]byte
	SmtRoot       [32]byte
}

// IAvsSigsVerifierBN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type IAvsSigsVerifierBN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IAvsSigsVerifierBN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type IAvsSigsVerifierBN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IAvsSigsVerifierNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IAvsSigsVerifierNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []IAvsSigsVerifierBN254G1Point
	QuorumApks                   []IAvsSigsVerifierBN254G1Point
	ApkG2                        IAvsSigsVerifierBN254G2Point
	Sigma                        IAvsSigsVerifierBN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IAvsSigsVerifierSigInfo is an auto generated low-level Go binding around an user-defined struct.
type IAvsSigsVerifierSigInfo struct {
	BlockNum uint64
	Params   IAvsSigsVerifierNonSignerStakesAndSignature
}

// IBrevisTypesCallback is an auto generated low-level Go binding around an user-defined struct.
type IBrevisTypesCallback struct {
	Target common.Address
	Gas    uint64
}

// IBvnSigsVerifierSigInfo is an auto generated low-level Go binding around an user-defined struct.
type IBvnSigsVerifierSigInfo struct {
	Sigs    [][]byte
	Signers []common.Address
	Powers  []*big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisProof\",\"name\":\"_brevisProof\",\"type\":\"address\"},{\"internalType\":\"contractIBvnSigsVerifier\",\"name\":\"_bvnSigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AvsSigsVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"}],\"name\":\"BaseDataUrlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BrevisDisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BrevisProofUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BvnSigsVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"nonces\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"appCommitHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"appVkHashes\",\"type\":\"bytes32[]\"}],\"name\":\"OpRequestsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ProverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ProverRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proofId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"RequestCallbackFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proofId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"RequestFeeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proofId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proofId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"RequestRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proofId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBrevisTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"RequestTimeoutUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"nonces\",\"type\":\"uint64[]\"}],\"name\":\"RequestsCallbackFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"nonces\",\"type\":\"uint64[]\"}],\"name\":\"RequestsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addProvers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBrevis.ProofData\",\"name\":\"_proofData\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"_nodeIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_appCircuitOutput\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_callbackTarget\",\"type\":\"address\"}],\"name\":\"applyBrevisAggProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBrevis.ProofData[]\",\"name\":\"_proofDataArray\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_appCircuitOutputs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_callbackTargets\",\"type\":\"address[]\"}],\"name\":\"applyBrevisAggProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_appCircuitOutput\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_callbackTarget\",\"type\":\"address\"}],\"name\":\"applyBrevisProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsSigsVerifier\",\"outputs\":[{\"internalType\":\"contractIAvsSigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseDataURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brevisDispute\",\"outputs\":[{\"internalType\":\"contractIBrevisDispute\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brevisProof\",\"outputs\":[{\"internalType\":\"contractIBrevisProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bvnSigsVerifier\",\"outputs\":[{\"internalType\":\"contractIBvnSigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"}],\"name\":\"dataURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proofIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_nonces\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_appCommitHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_appVkHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBvnSigsVerifier.SigInfo\",\"name\":\"_bvnSigInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsSigsVerifier.BN254_G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsSigsVerifier.BN254_G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIAvsSigsVerifier.BN254_G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsSigsVerifier.BN254_G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIAvsSigsVerifier.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"internalType\":\"structIAvsSigsVerifier.SigInfo\",\"name\":\"_avsSigInfo\",\"type\":\"tuple\"}],\"name\":\"fulfillOpRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_appCircuitOutput\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_callbackTarget\",\"type\":\"address\"}],\"name\":\"fulfillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proofIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_nonces\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"smtRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBrevis.ProofData[]\",\"name\":\"_proofDataArray\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_appCircuitOutputs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_callbackTargets\",\"type\":\"address[]\"}],\"name\":\"fulfillRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_addGas\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_currentFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundee\",\"type\":\"address\"}],\"name\":\"increaseGasFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisProof\",\"name\":\"_brevisProof\",\"type\":\"address\"},{\"internalType\":\"contractIBvnSigsVerifier\",\"name\":\"_bvnSigsVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestTimeout\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numPausers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numProvers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"onchainRequests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"feeHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"opdata\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pauserList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proverStates\",\"outputs\":[{\"internalType\":\"enumBrevisAccess.ProverState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"provers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"queryRequestStatus\",\"outputs\":[{\"internalType\":\"enumIBrevisTypes.RequestStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"}],\"name\":\"queryRequestStatus\",\"outputs\":[{\"internalType\":\"enumIBrevisTypes.RequestStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundee\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"removePausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeProvers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumIBrevisTypes.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_refundee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisTypes.Callback\",\"name\":\"_callback\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"_option\",\"type\":\"uint8\"}],\"name\":\"sendRequest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsSigsVerifier\",\"type\":\"address\"}],\"name\":\"setAvsSigsVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setBaseDataURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_brevisDispute\",\"type\":\"address\"}],\"name\":\"setBrevisDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_brevisProof\",\"type\":\"address\"}],\"name\":\"setBrevisProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bvnSigsVerifier\",\"type\":\"address\"}],\"name\":\"setBvnSigsVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestKey\",\"type\":\"bytes32\"},{\"internalType\":\"enumIBrevisTypes.RequestStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setRequestStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"setRequestTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_option\",\"type\":\"uint8\"}],\"name\":\"validateOpAppData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proofId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_appCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_appVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_option\",\"type\":\"uint8\"}],\"name\":\"validateOpAppData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proofIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_nonces\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_appCommitHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_appVkHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_option\",\"type\":\"uint8\"}],\"name\":\"validateOpAppData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// AvsSigsVerifier is a free data retrieval call binding the contract method 0x0b9c8c41.
//
// Solidity: function avsSigsVerifier() view returns(address)
func (_Main *MainCaller) AvsSigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "avsSigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsSigsVerifier is a free data retrieval call binding the contract method 0x0b9c8c41.
//
// Solidity: function avsSigsVerifier() view returns(address)
func (_Main *MainSession) AvsSigsVerifier() (common.Address, error) {
	return _Main.Contract.AvsSigsVerifier(&_Main.CallOpts)
}

// AvsSigsVerifier is a free data retrieval call binding the contract method 0x0b9c8c41.
//
// Solidity: function avsSigsVerifier() view returns(address)
func (_Main *MainCallerSession) AvsSigsVerifier() (common.Address, error) {
	return _Main.Contract.AvsSigsVerifier(&_Main.CallOpts)
}

// BaseDataURL is a free data retrieval call binding the contract method 0x25cda16d.
//
// Solidity: function baseDataURL() view returns(string)
func (_Main *MainCaller) BaseDataURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "baseDataURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseDataURL is a free data retrieval call binding the contract method 0x25cda16d.
//
// Solidity: function baseDataURL() view returns(string)
func (_Main *MainSession) BaseDataURL() (string, error) {
	return _Main.Contract.BaseDataURL(&_Main.CallOpts)
}

// BaseDataURL is a free data retrieval call binding the contract method 0x25cda16d.
//
// Solidity: function baseDataURL() view returns(string)
func (_Main *MainCallerSession) BaseDataURL() (string, error) {
	return _Main.Contract.BaseDataURL(&_Main.CallOpts)
}

// BrevisDispute is a free data retrieval call binding the contract method 0x7d91f017.
//
// Solidity: function brevisDispute() view returns(address)
func (_Main *MainCaller) BrevisDispute(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "brevisDispute")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrevisDispute is a free data retrieval call binding the contract method 0x7d91f017.
//
// Solidity: function brevisDispute() view returns(address)
func (_Main *MainSession) BrevisDispute() (common.Address, error) {
	return _Main.Contract.BrevisDispute(&_Main.CallOpts)
}

// BrevisDispute is a free data retrieval call binding the contract method 0x7d91f017.
//
// Solidity: function brevisDispute() view returns(address)
func (_Main *MainCallerSession) BrevisDispute() (common.Address, error) {
	return _Main.Contract.BrevisDispute(&_Main.CallOpts)
}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_Main *MainCaller) BrevisProof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "brevisProof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_Main *MainSession) BrevisProof() (common.Address, error) {
	return _Main.Contract.BrevisProof(&_Main.CallOpts)
}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_Main *MainCallerSession) BrevisProof() (common.Address, error) {
	return _Main.Contract.BrevisProof(&_Main.CallOpts)
}

// BvnSigsVerifier is a free data retrieval call binding the contract method 0x3f2b17b1.
//
// Solidity: function bvnSigsVerifier() view returns(address)
func (_Main *MainCaller) BvnSigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "bvnSigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BvnSigsVerifier is a free data retrieval call binding the contract method 0x3f2b17b1.
//
// Solidity: function bvnSigsVerifier() view returns(address)
func (_Main *MainSession) BvnSigsVerifier() (common.Address, error) {
	return _Main.Contract.BvnSigsVerifier(&_Main.CallOpts)
}

// BvnSigsVerifier is a free data retrieval call binding the contract method 0x3f2b17b1.
//
// Solidity: function bvnSigsVerifier() view returns(address)
func (_Main *MainCallerSession) BvnSigsVerifier() (common.Address, error) {
	return _Main.Contract.BvnSigsVerifier(&_Main.CallOpts)
}

// DataURL is a free data retrieval call binding the contract method 0x666d1651.
//
// Solidity: function dataURL(bytes32 _proofId) view returns(string)
func (_Main *MainCaller) DataURL(opts *bind.CallOpts, _proofId [32]byte) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "dataURL", _proofId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURL is a free data retrieval call binding the contract method 0x666d1651.
//
// Solidity: function dataURL(bytes32 _proofId) view returns(string)
func (_Main *MainSession) DataURL(_proofId [32]byte) (string, error) {
	return _Main.Contract.DataURL(&_Main.CallOpts, _proofId)
}

// DataURL is a free data retrieval call binding the contract method 0x666d1651.
//
// Solidity: function dataURL(bytes32 _proofId) view returns(string)
func (_Main *MainCallerSession) DataURL(_proofId [32]byte) (string, error) {
	return _Main.Contract.DataURL(&_Main.CallOpts, _proofId)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Main *MainCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Main *MainSession) FeeCollector() (common.Address, error) {
	return _Main.Contract.FeeCollector(&_Main.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Main *MainCallerSession) FeeCollector() (common.Address, error) {
	return _Main.Contract.FeeCollector(&_Main.CallOpts)
}

// IsActiveProver is a free data retrieval call binding the contract method 0xec64842e.
//
// Solidity: function isActiveProver(address _account) view returns(bool)
func (_Main *MainCaller) IsActiveProver(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isActiveProver", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveProver is a free data retrieval call binding the contract method 0xec64842e.
//
// Solidity: function isActiveProver(address _account) view returns(bool)
func (_Main *MainSession) IsActiveProver(_account common.Address) (bool, error) {
	return _Main.Contract.IsActiveProver(&_Main.CallOpts, _account)
}

// IsActiveProver is a free data retrieval call binding the contract method 0xec64842e.
//
// Solidity: function isActiveProver(address _account) view returns(bool)
func (_Main *MainCallerSession) IsActiveProver(_account common.Address) (bool, error) {
	return _Main.Contract.IsActiveProver(&_Main.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Main *MainCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Main *MainSession) IsPauser(account common.Address) (bool, error) {
	return _Main.Contract.IsPauser(&_Main.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Main *MainCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Main.Contract.IsPauser(&_Main.CallOpts, account)
}

// NumPausers is a free data retrieval call binding the contract method 0x58a16b44.
//
// Solidity: function numPausers() view returns(uint256)
func (_Main *MainCaller) NumPausers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "numPausers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPausers is a free data retrieval call binding the contract method 0x58a16b44.
//
// Solidity: function numPausers() view returns(uint256)
func (_Main *MainSession) NumPausers() (*big.Int, error) {
	return _Main.Contract.NumPausers(&_Main.CallOpts)
}

// NumPausers is a free data retrieval call binding the contract method 0x58a16b44.
//
// Solidity: function numPausers() view returns(uint256)
func (_Main *MainCallerSession) NumPausers() (*big.Int, error) {
	return _Main.Contract.NumPausers(&_Main.CallOpts)
}

// NumProvers is a free data retrieval call binding the contract method 0x4f4fef18.
//
// Solidity: function numProvers() view returns(uint256)
func (_Main *MainCaller) NumProvers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "numProvers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumProvers is a free data retrieval call binding the contract method 0x4f4fef18.
//
// Solidity: function numProvers() view returns(uint256)
func (_Main *MainSession) NumProvers() (*big.Int, error) {
	return _Main.Contract.NumProvers(&_Main.CallOpts)
}

// NumProvers is a free data retrieval call binding the contract method 0x4f4fef18.
//
// Solidity: function numProvers() view returns(uint256)
func (_Main *MainCallerSession) NumProvers() (*big.Int, error) {
	return _Main.Contract.NumProvers(&_Main.CallOpts)
}

// OnchainRequests is a free data retrieval call binding the contract method 0xc49af0fa.
//
// Solidity: function onchainRequests(bytes32 ) view returns(bytes32 feeHash, (address,uint64) callback)
func (_Main *MainCaller) OnchainRequests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	FeeHash  [32]byte
	Callback IBrevisTypesCallback
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "onchainRequests", arg0)

	outstruct := new(struct {
		FeeHash  [32]byte
		Callback IBrevisTypesCallback
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Callback = *abi.ConvertType(out[1], new(IBrevisTypesCallback)).(*IBrevisTypesCallback)

	return *outstruct, err

}

// OnchainRequests is a free data retrieval call binding the contract method 0xc49af0fa.
//
// Solidity: function onchainRequests(bytes32 ) view returns(bytes32 feeHash, (address,uint64) callback)
func (_Main *MainSession) OnchainRequests(arg0 [32]byte) (struct {
	FeeHash  [32]byte
	Callback IBrevisTypesCallback
}, error) {
	return _Main.Contract.OnchainRequests(&_Main.CallOpts, arg0)
}

// OnchainRequests is a free data retrieval call binding the contract method 0xc49af0fa.
//
// Solidity: function onchainRequests(bytes32 ) view returns(bytes32 feeHash, (address,uint64) callback)
func (_Main *MainCallerSession) OnchainRequests(arg0 [32]byte) (struct {
	FeeHash  [32]byte
	Callback IBrevisTypesCallback
}, error) {
	return _Main.Contract.OnchainRequests(&_Main.CallOpts, arg0)
}

// Opdata is a free data retrieval call binding the contract method 0xc2eaa931.
//
// Solidity: function opdata(bytes32 ) view returns(bytes32)
func (_Main *MainCaller) Opdata(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "opdata", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Opdata is a free data retrieval call binding the contract method 0xc2eaa931.
//
// Solidity: function opdata(bytes32 ) view returns(bytes32)
func (_Main *MainSession) Opdata(arg0 [32]byte) ([32]byte, error) {
	return _Main.Contract.Opdata(&_Main.CallOpts, arg0)
}

// Opdata is a free data retrieval call binding the contract method 0xc2eaa931.
//
// Solidity: function opdata(bytes32 ) view returns(bytes32)
func (_Main *MainCallerSession) Opdata(arg0 [32]byte) ([32]byte, error) {
	return _Main.Contract.Opdata(&_Main.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainSession) Paused() (bool, error) {
	return _Main.Contract.Paused(&_Main.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Main *MainCallerSession) Paused() (bool, error) {
	return _Main.Contract.Paused(&_Main.CallOpts)
}

// PauserList is a free data retrieval call binding the contract method 0x158535ff.
//
// Solidity: function pauserList(uint256 ) view returns(address)
func (_Main *MainCaller) PauserList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "pauserList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserList is a free data retrieval call binding the contract method 0x158535ff.
//
// Solidity: function pauserList(uint256 ) view returns(address)
func (_Main *MainSession) PauserList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.PauserList(&_Main.CallOpts, arg0)
}

// PauserList is a free data retrieval call binding the contract method 0x158535ff.
//
// Solidity: function pauserList(uint256 ) view returns(address)
func (_Main *MainCallerSession) PauserList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.PauserList(&_Main.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Main *MainCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Main *MainSession) Pausers(arg0 common.Address) (bool, error) {
	return _Main.Contract.Pausers(&_Main.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Main *MainCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Main.Contract.Pausers(&_Main.CallOpts, arg0)
}

// ProverStates is a free data retrieval call binding the contract method 0xfabc74f5.
//
// Solidity: function proverStates(address ) view returns(uint8)
func (_Main *MainCaller) ProverStates(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "proverStates", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProverStates is a free data retrieval call binding the contract method 0xfabc74f5.
//
// Solidity: function proverStates(address ) view returns(uint8)
func (_Main *MainSession) ProverStates(arg0 common.Address) (uint8, error) {
	return _Main.Contract.ProverStates(&_Main.CallOpts, arg0)
}

// ProverStates is a free data retrieval call binding the contract method 0xfabc74f5.
//
// Solidity: function proverStates(address ) view returns(uint8)
func (_Main *MainCallerSession) ProverStates(arg0 common.Address) (uint8, error) {
	return _Main.Contract.ProverStates(&_Main.CallOpts, arg0)
}

// Provers is a free data retrieval call binding the contract method 0xfd1190ea.
//
// Solidity: function provers(uint256 ) view returns(address)
func (_Main *MainCaller) Provers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "provers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Provers is a free data retrieval call binding the contract method 0xfd1190ea.
//
// Solidity: function provers(uint256 ) view returns(address)
func (_Main *MainSession) Provers(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.Provers(&_Main.CallOpts, arg0)
}

// Provers is a free data retrieval call binding the contract method 0xfd1190ea.
//
// Solidity: function provers(uint256 ) view returns(address)
func (_Main *MainCallerSession) Provers(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.Provers(&_Main.CallOpts, arg0)
}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xa65b3c06.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce) view returns(uint8, uint8)
func (_Main *MainCaller) QueryRequestStatus(opts *bind.CallOpts, _proofId [32]byte, _nonce uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "queryRequestStatus", _proofId, _nonce)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xa65b3c06.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce) view returns(uint8, uint8)
func (_Main *MainSession) QueryRequestStatus(_proofId [32]byte, _nonce uint64) (uint8, uint8, error) {
	return _Main.Contract.QueryRequestStatus(&_Main.CallOpts, _proofId, _nonce)
}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xa65b3c06.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce) view returns(uint8, uint8)
func (_Main *MainCallerSession) QueryRequestStatus(_proofId [32]byte, _nonce uint64) (uint8, uint8, error) {
	return _Main.Contract.QueryRequestStatus(&_Main.CallOpts, _proofId, _nonce)
}

// QueryRequestStatus0 is a free data retrieval call binding the contract method 0xb33e1a39.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce, uint256 _challengeWindow) view returns(uint8, uint8)
func (_Main *MainCaller) QueryRequestStatus0(opts *bind.CallOpts, _proofId [32]byte, _nonce uint64, _challengeWindow *big.Int) (uint8, uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "queryRequestStatus0", _proofId, _nonce, _challengeWindow)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// QueryRequestStatus0 is a free data retrieval call binding the contract method 0xb33e1a39.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce, uint256 _challengeWindow) view returns(uint8, uint8)
func (_Main *MainSession) QueryRequestStatus0(_proofId [32]byte, _nonce uint64, _challengeWindow *big.Int) (uint8, uint8, error) {
	return _Main.Contract.QueryRequestStatus0(&_Main.CallOpts, _proofId, _nonce, _challengeWindow)
}

// QueryRequestStatus0 is a free data retrieval call binding the contract method 0xb33e1a39.
//
// Solidity: function queryRequestStatus(bytes32 _proofId, uint64 _nonce, uint256 _challengeWindow) view returns(uint8, uint8)
func (_Main *MainCallerSession) QueryRequestStatus0(_proofId [32]byte, _nonce uint64, _challengeWindow *big.Int) (uint8, uint8, error) {
	return _Main.Contract.QueryRequestStatus0(&_Main.CallOpts, _proofId, _nonce, _challengeWindow)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_Main *MainCaller) RequestTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "requestTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_Main *MainSession) RequestTimeout() (*big.Int, error) {
	return _Main.Contract.RequestTimeout(&_Main.CallOpts)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_Main *MainCallerSession) RequestTimeout() (*big.Int, error) {
	return _Main.Contract.RequestTimeout(&_Main.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, uint8 option)
func (_Main *MainCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status    uint8
	Timestamp uint64
	Option    uint8
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Status    uint8
		Timestamp uint64
		Option    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Option = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, uint8 option)
func (_Main *MainSession) Requests(arg0 [32]byte) (struct {
	Status    uint8
	Timestamp uint64
	Option    uint8
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, uint8 option)
func (_Main *MainCallerSession) Requests(arg0 [32]byte) (struct {
	Status    uint8
	Timestamp uint64
	Option    uint8
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// ValidateOpAppData is a free data retrieval call binding the contract method 0x54eee2f0.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainCaller) ValidateOpAppData(opts *bind.CallOpts, _proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "validateOpAppData", _proofId, _nonce, _appCommitHash, _appVkHash, _challengeWindow, _option)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOpAppData is a free data retrieval call binding the contract method 0x54eee2f0.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainSession) ValidateOpAppData(_proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData(&_Main.CallOpts, _proofId, _nonce, _appCommitHash, _appVkHash, _challengeWindow, _option)
}

// ValidateOpAppData is a free data retrieval call binding the contract method 0x54eee2f0.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainCallerSession) ValidateOpAppData(_proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData(&_Main.CallOpts, _proofId, _nonce, _appCommitHash, _appVkHash, _challengeWindow, _option)
}

// ValidateOpAppData0 is a free data retrieval call binding the contract method 0x84c2d279.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint8 _option) view returns(bool)
func (_Main *MainCaller) ValidateOpAppData0(opts *bind.CallOpts, _proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _option uint8) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "validateOpAppData0", _proofId, _nonce, _appCommitHash, _appVkHash, _option)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOpAppData0 is a free data retrieval call binding the contract method 0x84c2d279.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint8 _option) view returns(bool)
func (_Main *MainSession) ValidateOpAppData0(_proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData0(&_Main.CallOpts, _proofId, _nonce, _appCommitHash, _appVkHash, _option)
}

// ValidateOpAppData0 is a free data retrieval call binding the contract method 0x84c2d279.
//
// Solidity: function validateOpAppData(bytes32 _proofId, uint64 _nonce, bytes32 _appCommitHash, bytes32 _appVkHash, uint8 _option) view returns(bool)
func (_Main *MainCallerSession) ValidateOpAppData0(_proofId [32]byte, _nonce uint64, _appCommitHash [32]byte, _appVkHash [32]byte, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData0(&_Main.CallOpts, _proofId, _nonce, _appCommitHash, _appVkHash, _option)
}

// ValidateOpAppData1 is a free data retrieval call binding the contract method 0xa90f4bd7.
//
// Solidity: function validateOpAppData(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainCaller) ValidateOpAppData1(opts *bind.CallOpts, _proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "validateOpAppData1", _proofIds, _nonces, _appCommitHashes, _appVkHashes, _challengeWindow, _option)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOpAppData1 is a free data retrieval call binding the contract method 0xa90f4bd7.
//
// Solidity: function validateOpAppData(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainSession) ValidateOpAppData1(_proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData1(&_Main.CallOpts, _proofIds, _nonces, _appCommitHashes, _appVkHashes, _challengeWindow, _option)
}

// ValidateOpAppData1 is a free data retrieval call binding the contract method 0xa90f4bd7.
//
// Solidity: function validateOpAppData(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, uint256 _challengeWindow, uint8 _option) view returns(bool)
func (_Main *MainCallerSession) ValidateOpAppData1(_proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _challengeWindow *big.Int, _option uint8) (bool, error) {
	return _Main.Contract.ValidateOpAppData1(&_Main.CallOpts, _proofIds, _nonces, _appCommitHashes, _appVkHashes, _challengeWindow, _option)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Main *MainTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Main *MainSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddPauser(&_Main.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Main *MainTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddPauser(&_Main.TransactOpts, account)
}

// AddPausers is a paid mutator transaction binding the contract method 0xe79b7a51.
//
// Solidity: function addPausers(address[] accounts) returns()
func (_Main *MainTransactor) AddPausers(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addPausers", accounts)
}

// AddPausers is a paid mutator transaction binding the contract method 0xe79b7a51.
//
// Solidity: function addPausers(address[] accounts) returns()
func (_Main *MainSession) AddPausers(accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddPausers(&_Main.TransactOpts, accounts)
}

// AddPausers is a paid mutator transaction binding the contract method 0xe79b7a51.
//
// Solidity: function addPausers(address[] accounts) returns()
func (_Main *MainTransactorSession) AddPausers(accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddPausers(&_Main.TransactOpts, accounts)
}

// AddProvers is a paid mutator transaction binding the contract method 0x677625f2.
//
// Solidity: function addProvers(address[] _accounts) returns()
func (_Main *MainTransactor) AddProvers(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addProvers", _accounts)
}

// AddProvers is a paid mutator transaction binding the contract method 0x677625f2.
//
// Solidity: function addProvers(address[] _accounts) returns()
func (_Main *MainSession) AddProvers(_accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddProvers(&_Main.TransactOpts, _accounts)
}

// AddProvers is a paid mutator transaction binding the contract method 0x677625f2.
//
// Solidity: function addProvers(address[] _accounts) returns()
func (_Main *MainTransactorSession) AddProvers(_accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddProvers(&_Main.TransactOpts, _accounts)
}

// ApplyBrevisAggProof is a paid mutator transaction binding the contract method 0x8cd2ed66.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32) _proofData, bytes32 _merkleRoot, bytes32[] _merkleProof, uint8 _nodeIndex, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactor) ApplyBrevisAggProof(opts *bind.TransactOpts, _chainId uint64, _proofData BrevisProofData, _merkleRoot [32]byte, _merkleProof [][32]byte, _nodeIndex uint8, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "applyBrevisAggProof", _chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex, _appCircuitOutput, _callbackTarget)
}

// ApplyBrevisAggProof is a paid mutator transaction binding the contract method 0x8cd2ed66.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32) _proofData, bytes32 _merkleRoot, bytes32[] _merkleProof, uint8 _nodeIndex, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainSession) ApplyBrevisAggProof(_chainId uint64, _proofData BrevisProofData, _merkleRoot [32]byte, _merkleProof [][32]byte, _nodeIndex uint8, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisAggProof(&_Main.TransactOpts, _chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex, _appCircuitOutput, _callbackTarget)
}

// ApplyBrevisAggProof is a paid mutator transaction binding the contract method 0x8cd2ed66.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32) _proofData, bytes32 _merkleRoot, bytes32[] _merkleProof, uint8 _nodeIndex, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactorSession) ApplyBrevisAggProof(_chainId uint64, _proofData BrevisProofData, _merkleRoot [32]byte, _merkleProof [][32]byte, _nodeIndex uint8, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisAggProof(&_Main.TransactOpts, _chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex, _appCircuitOutput, _callbackTarget)
}

// ApplyBrevisAggProof0 is a paid mutator transaction binding the contract method 0xbc7c2050.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainTransactor) ApplyBrevisAggProof0(opts *bind.TransactOpts, _chainId uint64, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "applyBrevisAggProof0", _chainId, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// ApplyBrevisAggProof0 is a paid mutator transaction binding the contract method 0xbc7c2050.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainSession) ApplyBrevisAggProof0(_chainId uint64, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisAggProof0(&_Main.TransactOpts, _chainId, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// ApplyBrevisAggProof0 is a paid mutator transaction binding the contract method 0xbc7c2050.
//
// Solidity: function applyBrevisAggProof(uint64 _chainId, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainTransactorSession) ApplyBrevisAggProof0(_chainId uint64, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisAggProof0(&_Main.TransactOpts, _chainId, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// ApplyBrevisProof is a paid mutator transaction binding the contract method 0x0dbeefaf.
//
// Solidity: function applyBrevisProof(bytes32 _proofId, bytes32 _appVkHash, bytes32 _appCommitHash, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactor) ApplyBrevisProof(opts *bind.TransactOpts, _proofId [32]byte, _appVkHash [32]byte, _appCommitHash [32]byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "applyBrevisProof", _proofId, _appVkHash, _appCommitHash, _appCircuitOutput, _callbackTarget)
}

// ApplyBrevisProof is a paid mutator transaction binding the contract method 0x0dbeefaf.
//
// Solidity: function applyBrevisProof(bytes32 _proofId, bytes32 _appVkHash, bytes32 _appCommitHash, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainSession) ApplyBrevisProof(_proofId [32]byte, _appVkHash [32]byte, _appCommitHash [32]byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisProof(&_Main.TransactOpts, _proofId, _appVkHash, _appCommitHash, _appCircuitOutput, _callbackTarget)
}

// ApplyBrevisProof is a paid mutator transaction binding the contract method 0x0dbeefaf.
//
// Solidity: function applyBrevisProof(bytes32 _proofId, bytes32 _appVkHash, bytes32 _appCommitHash, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactorSession) ApplyBrevisProof(_proofId [32]byte, _appVkHash [32]byte, _appCommitHash [32]byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.ApplyBrevisProof(&_Main.TransactOpts, _proofId, _appVkHash, _appCommitHash, _appCircuitOutput, _callbackTarget)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_Main *MainTransactor) CollectFee(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "collectFee", _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_Main *MainSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.Contract.CollectFee(&_Main.TransactOpts, _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_Main *MainTransactorSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.Contract.CollectFee(&_Main.TransactOpts, _amount, _to)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x35d43733.
//
// Solidity: function fulfillOpRequests(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, (bytes[],address[],uint256[]) _bvnSigInfo, (uint64,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])) _avsSigInfo) returns()
func (_Main *MainTransactor) FulfillOpRequests(opts *bind.TransactOpts, _proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _bvnSigInfo IBvnSigsVerifierSigInfo, _avsSigInfo IAvsSigsVerifierSigInfo) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "fulfillOpRequests", _proofIds, _nonces, _appCommitHashes, _appVkHashes, _bvnSigInfo, _avsSigInfo)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x35d43733.
//
// Solidity: function fulfillOpRequests(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, (bytes[],address[],uint256[]) _bvnSigInfo, (uint64,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])) _avsSigInfo) returns()
func (_Main *MainSession) FulfillOpRequests(_proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _bvnSigInfo IBvnSigsVerifierSigInfo, _avsSigInfo IAvsSigsVerifierSigInfo) (*types.Transaction, error) {
	return _Main.Contract.FulfillOpRequests(&_Main.TransactOpts, _proofIds, _nonces, _appCommitHashes, _appVkHashes, _bvnSigInfo, _avsSigInfo)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x35d43733.
//
// Solidity: function fulfillOpRequests(bytes32[] _proofIds, uint64[] _nonces, bytes32[] _appCommitHashes, bytes32[] _appVkHashes, (bytes[],address[],uint256[]) _bvnSigInfo, (uint64,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][])) _avsSigInfo) returns()
func (_Main *MainTransactorSession) FulfillOpRequests(_proofIds [][32]byte, _nonces []uint64, _appCommitHashes [][32]byte, _appVkHashes [][32]byte, _bvnSigInfo IBvnSigsVerifierSigInfo, _avsSigInfo IAvsSigsVerifierSigInfo) (*types.Transaction, error) {
	return _Main.Contract.FulfillOpRequests(&_Main.TransactOpts, _proofIds, _nonces, _appCommitHashes, _appVkHashes, _bvnSigInfo, _avsSigInfo)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xcd978249.
//
// Solidity: function fulfillRequest(bytes32 _proofId, uint64 _nonce, uint64 _chainId, bytes _proof, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactor) FulfillRequest(opts *bind.TransactOpts, _proofId [32]byte, _nonce uint64, _chainId uint64, _proof []byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "fulfillRequest", _proofId, _nonce, _chainId, _proof, _appCircuitOutput, _callbackTarget)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xcd978249.
//
// Solidity: function fulfillRequest(bytes32 _proofId, uint64 _nonce, uint64 _chainId, bytes _proof, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainSession) FulfillRequest(_proofId [32]byte, _nonce uint64, _chainId uint64, _proof []byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.FulfillRequest(&_Main.TransactOpts, _proofId, _nonce, _chainId, _proof, _appCircuitOutput, _callbackTarget)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xcd978249.
//
// Solidity: function fulfillRequest(bytes32 _proofId, uint64 _nonce, uint64 _chainId, bytes _proof, bytes _appCircuitOutput, address _callbackTarget) returns()
func (_Main *MainTransactorSession) FulfillRequest(_proofId [32]byte, _nonce uint64, _chainId uint64, _proof []byte, _appCircuitOutput []byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _Main.Contract.FulfillRequest(&_Main.TransactOpts, _proofId, _nonce, _chainId, _proof, _appCircuitOutput, _callbackTarget)
}

// FulfillRequests is a paid mutator transaction binding the contract method 0xfb22b7a8.
//
// Solidity: function fulfillRequests(bytes32[] _proofIds, uint64[] _nonces, uint64 _chainId, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainTransactor) FulfillRequests(opts *bind.TransactOpts, _proofIds [][32]byte, _nonces []uint64, _chainId uint64, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "fulfillRequests", _proofIds, _nonces, _chainId, _proof, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// FulfillRequests is a paid mutator transaction binding the contract method 0xfb22b7a8.
//
// Solidity: function fulfillRequests(bytes32[] _proofIds, uint64[] _nonces, uint64 _chainId, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainSession) FulfillRequests(_proofIds [][32]byte, _nonces []uint64, _chainId uint64, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.Contract.FulfillRequests(&_Main.TransactOpts, _proofIds, _nonces, _chainId, _proof, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// FulfillRequests is a paid mutator transaction binding the contract method 0xfb22b7a8.
//
// Solidity: function fulfillRequests(bytes32[] _proofIds, uint64[] _nonces, uint64 _chainId, bytes _proof, (bytes32,bytes32,bytes32,bytes32,bytes32)[] _proofDataArray, bytes[] _appCircuitOutputs, address[] _callbackTargets) returns()
func (_Main *MainTransactorSession) FulfillRequests(_proofIds [][32]byte, _nonces []uint64, _chainId uint64, _proof []byte, _proofDataArray []BrevisProofData, _appCircuitOutputs [][]byte, _callbackTargets []common.Address) (*types.Transaction, error) {
	return _Main.Contract.FulfillRequests(&_Main.TransactOpts, _proofIds, _nonces, _chainId, _proof, _proofDataArray, _appCircuitOutputs, _callbackTargets)
}

// IncreaseGasFee is a paid mutator transaction binding the contract method 0xc33529ae.
//
// Solidity: function increaseGasFee(bytes32 _proofId, uint64 _nonce, uint64 _addGas, uint256 _currentFee, address _refundee) payable returns()
func (_Main *MainTransactor) IncreaseGasFee(opts *bind.TransactOpts, _proofId [32]byte, _nonce uint64, _addGas uint64, _currentFee *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "increaseGasFee", _proofId, _nonce, _addGas, _currentFee, _refundee)
}

// IncreaseGasFee is a paid mutator transaction binding the contract method 0xc33529ae.
//
// Solidity: function increaseGasFee(bytes32 _proofId, uint64 _nonce, uint64 _addGas, uint256 _currentFee, address _refundee) payable returns()
func (_Main *MainSession) IncreaseGasFee(_proofId [32]byte, _nonce uint64, _addGas uint64, _currentFee *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.Contract.IncreaseGasFee(&_Main.TransactOpts, _proofId, _nonce, _addGas, _currentFee, _refundee)
}

// IncreaseGasFee is a paid mutator transaction binding the contract method 0xc33529ae.
//
// Solidity: function increaseGasFee(bytes32 _proofId, uint64 _nonce, uint64 _addGas, uint256 _currentFee, address _refundee) payable returns()
func (_Main *MainTransactorSession) IncreaseGasFee(_proofId [32]byte, _nonce uint64, _addGas uint64, _currentFee *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.Contract.IncreaseGasFee(&_Main.TransactOpts, _proofId, _nonce, _addGas, _currentFee, _refundee)
}

// Init is a paid mutator transaction binding the contract method 0x46639dba.
//
// Solidity: function init(address _feeCollector, address _brevisProof, address _bvnSigsVerifier, uint256 _requestTimeout) returns()
func (_Main *MainTransactor) Init(opts *bind.TransactOpts, _feeCollector common.Address, _brevisProof common.Address, _bvnSigsVerifier common.Address, _requestTimeout *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "init", _feeCollector, _brevisProof, _bvnSigsVerifier, _requestTimeout)
}

// Init is a paid mutator transaction binding the contract method 0x46639dba.
//
// Solidity: function init(address _feeCollector, address _brevisProof, address _bvnSigsVerifier, uint256 _requestTimeout) returns()
func (_Main *MainSession) Init(_feeCollector common.Address, _brevisProof common.Address, _bvnSigsVerifier common.Address, _requestTimeout *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Init(&_Main.TransactOpts, _feeCollector, _brevisProof, _bvnSigsVerifier, _requestTimeout)
}

// Init is a paid mutator transaction binding the contract method 0x46639dba.
//
// Solidity: function init(address _feeCollector, address _brevisProof, address _bvnSigsVerifier, uint256 _requestTimeout) returns()
func (_Main *MainTransactorSession) Init(_feeCollector common.Address, _brevisProof common.Address, _bvnSigsVerifier common.Address, _requestTimeout *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Init(&_Main.TransactOpts, _feeCollector, _brevisProof, _bvnSigsVerifier, _requestTimeout)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Main *MainTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Main *MainSession) Pause() (*types.Transaction, error) {
	return _Main.Contract.Pause(&_Main.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Main *MainTransactorSession) Pause() (*types.Transaction, error) {
	return _Main.Contract.Pause(&_Main.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfc0cfccc.
//
// Solidity: function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) returns()
func (_Main *MainTransactor) Refund(opts *bind.TransactOpts, _proofId [32]byte, _nonce uint64, _amount *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "refund", _proofId, _nonce, _amount, _refundee)
}

// Refund is a paid mutator transaction binding the contract method 0xfc0cfccc.
//
// Solidity: function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) returns()
func (_Main *MainSession) Refund(_proofId [32]byte, _nonce uint64, _amount *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.Contract.Refund(&_Main.TransactOpts, _proofId, _nonce, _amount, _refundee)
}

// Refund is a paid mutator transaction binding the contract method 0xfc0cfccc.
//
// Solidity: function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) returns()
func (_Main *MainTransactorSession) Refund(_proofId [32]byte, _nonce uint64, _amount *big.Int, _refundee common.Address) (*types.Transaction, error) {
	return _Main.Contract.Refund(&_Main.TransactOpts, _proofId, _nonce, _amount, _refundee)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Main *MainTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Main *MainSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemovePauser(&_Main.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Main *MainTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemovePauser(&_Main.TransactOpts, account)
}

// RemovePausers is a paid mutator transaction binding the contract method 0xa036e799.
//
// Solidity: function removePausers(address[] accounts) returns()
func (_Main *MainTransactor) RemovePausers(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removePausers", accounts)
}

// RemovePausers is a paid mutator transaction binding the contract method 0xa036e799.
//
// Solidity: function removePausers(address[] accounts) returns()
func (_Main *MainSession) RemovePausers(accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemovePausers(&_Main.TransactOpts, accounts)
}

// RemovePausers is a paid mutator transaction binding the contract method 0xa036e799.
//
// Solidity: function removePausers(address[] accounts) returns()
func (_Main *MainTransactorSession) RemovePausers(accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemovePausers(&_Main.TransactOpts, accounts)
}

// RemoveProvers is a paid mutator transaction binding the contract method 0xe6c6fcec.
//
// Solidity: function removeProvers(address[] _accounts) returns()
func (_Main *MainTransactor) RemoveProvers(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeProvers", _accounts)
}

// RemoveProvers is a paid mutator transaction binding the contract method 0xe6c6fcec.
//
// Solidity: function removeProvers(address[] _accounts) returns()
func (_Main *MainSession) RemoveProvers(_accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveProvers(&_Main.TransactOpts, _accounts)
}

// RemoveProvers is a paid mutator transaction binding the contract method 0xe6c6fcec.
//
// Solidity: function removeProvers(address[] _accounts) returns()
func (_Main *MainTransactorSession) RemoveProvers(_accounts []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveProvers(&_Main.TransactOpts, _accounts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Main *MainTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Main *MainSession) RenouncePauser() (*types.Transaction, error) {
	return _Main.Contract.RenouncePauser(&_Main.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Main *MainTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Main.Contract.RenouncePauser(&_Main.TransactOpts)
}

// SendRequest is a paid mutator transaction binding the contract method 0x191fa9b6.
//
// Solidity: function sendRequest(bytes32 _proofId, uint64 _nonce, address _refundee, (address,uint64) _callback, uint8 _option) payable returns()
func (_Main *MainTransactor) SendRequest(opts *bind.TransactOpts, _proofId [32]byte, _nonce uint64, _refundee common.Address, _callback IBrevisTypesCallback, _option uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "sendRequest", _proofId, _nonce, _refundee, _callback, _option)
}

// SendRequest is a paid mutator transaction binding the contract method 0x191fa9b6.
//
// Solidity: function sendRequest(bytes32 _proofId, uint64 _nonce, address _refundee, (address,uint64) _callback, uint8 _option) payable returns()
func (_Main *MainSession) SendRequest(_proofId [32]byte, _nonce uint64, _refundee common.Address, _callback IBrevisTypesCallback, _option uint8) (*types.Transaction, error) {
	return _Main.Contract.SendRequest(&_Main.TransactOpts, _proofId, _nonce, _refundee, _callback, _option)
}

// SendRequest is a paid mutator transaction binding the contract method 0x191fa9b6.
//
// Solidity: function sendRequest(bytes32 _proofId, uint64 _nonce, address _refundee, (address,uint64) _callback, uint8 _option) payable returns()
func (_Main *MainTransactorSession) SendRequest(_proofId [32]byte, _nonce uint64, _refundee common.Address, _callback IBrevisTypesCallback, _option uint8) (*types.Transaction, error) {
	return _Main.Contract.SendRequest(&_Main.TransactOpts, _proofId, _nonce, _refundee, _callback, _option)
}

// SetAvsSigsVerifier is a paid mutator transaction binding the contract method 0xa8135374.
//
// Solidity: function setAvsSigsVerifier(address _avsSigsVerifier) returns()
func (_Main *MainTransactor) SetAvsSigsVerifier(opts *bind.TransactOpts, _avsSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setAvsSigsVerifier", _avsSigsVerifier)
}

// SetAvsSigsVerifier is a paid mutator transaction binding the contract method 0xa8135374.
//
// Solidity: function setAvsSigsVerifier(address _avsSigsVerifier) returns()
func (_Main *MainSession) SetAvsSigsVerifier(_avsSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetAvsSigsVerifier(&_Main.TransactOpts, _avsSigsVerifier)
}

// SetAvsSigsVerifier is a paid mutator transaction binding the contract method 0xa8135374.
//
// Solidity: function setAvsSigsVerifier(address _avsSigsVerifier) returns()
func (_Main *MainTransactorSession) SetAvsSigsVerifier(_avsSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetAvsSigsVerifier(&_Main.TransactOpts, _avsSigsVerifier)
}

// SetBaseDataURL is a paid mutator transaction binding the contract method 0x7a784a1c.
//
// Solidity: function setBaseDataURL(string _url) returns()
func (_Main *MainTransactor) SetBaseDataURL(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setBaseDataURL", _url)
}

// SetBaseDataURL is a paid mutator transaction binding the contract method 0x7a784a1c.
//
// Solidity: function setBaseDataURL(string _url) returns()
func (_Main *MainSession) SetBaseDataURL(_url string) (*types.Transaction, error) {
	return _Main.Contract.SetBaseDataURL(&_Main.TransactOpts, _url)
}

// SetBaseDataURL is a paid mutator transaction binding the contract method 0x7a784a1c.
//
// Solidity: function setBaseDataURL(string _url) returns()
func (_Main *MainTransactorSession) SetBaseDataURL(_url string) (*types.Transaction, error) {
	return _Main.Contract.SetBaseDataURL(&_Main.TransactOpts, _url)
}

// SetBrevisDispute is a paid mutator transaction binding the contract method 0xb5c06c33.
//
// Solidity: function setBrevisDispute(address _brevisDispute) returns()
func (_Main *MainTransactor) SetBrevisDispute(opts *bind.TransactOpts, _brevisDispute common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setBrevisDispute", _brevisDispute)
}

// SetBrevisDispute is a paid mutator transaction binding the contract method 0xb5c06c33.
//
// Solidity: function setBrevisDispute(address _brevisDispute) returns()
func (_Main *MainSession) SetBrevisDispute(_brevisDispute common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBrevisDispute(&_Main.TransactOpts, _brevisDispute)
}

// SetBrevisDispute is a paid mutator transaction binding the contract method 0xb5c06c33.
//
// Solidity: function setBrevisDispute(address _brevisDispute) returns()
func (_Main *MainTransactorSession) SetBrevisDispute(_brevisDispute common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBrevisDispute(&_Main.TransactOpts, _brevisDispute)
}

// SetBrevisProof is a paid mutator transaction binding the contract method 0xc772c87f.
//
// Solidity: function setBrevisProof(address _brevisProof) returns()
func (_Main *MainTransactor) SetBrevisProof(opts *bind.TransactOpts, _brevisProof common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setBrevisProof", _brevisProof)
}

// SetBrevisProof is a paid mutator transaction binding the contract method 0xc772c87f.
//
// Solidity: function setBrevisProof(address _brevisProof) returns()
func (_Main *MainSession) SetBrevisProof(_brevisProof common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBrevisProof(&_Main.TransactOpts, _brevisProof)
}

// SetBrevisProof is a paid mutator transaction binding the contract method 0xc772c87f.
//
// Solidity: function setBrevisProof(address _brevisProof) returns()
func (_Main *MainTransactorSession) SetBrevisProof(_brevisProof common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBrevisProof(&_Main.TransactOpts, _brevisProof)
}

// SetBvnSigsVerifier is a paid mutator transaction binding the contract method 0x967a6581.
//
// Solidity: function setBvnSigsVerifier(address _bvnSigsVerifier) returns()
func (_Main *MainTransactor) SetBvnSigsVerifier(opts *bind.TransactOpts, _bvnSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setBvnSigsVerifier", _bvnSigsVerifier)
}

// SetBvnSigsVerifier is a paid mutator transaction binding the contract method 0x967a6581.
//
// Solidity: function setBvnSigsVerifier(address _bvnSigsVerifier) returns()
func (_Main *MainSession) SetBvnSigsVerifier(_bvnSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBvnSigsVerifier(&_Main.TransactOpts, _bvnSigsVerifier)
}

// SetBvnSigsVerifier is a paid mutator transaction binding the contract method 0x967a6581.
//
// Solidity: function setBvnSigsVerifier(address _bvnSigsVerifier) returns()
func (_Main *MainTransactorSession) SetBvnSigsVerifier(_bvnSigsVerifier common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetBvnSigsVerifier(&_Main.TransactOpts, _bvnSigsVerifier)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_Main *MainTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_Main *MainSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFeeCollector(&_Main.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_Main *MainTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFeeCollector(&_Main.TransactOpts, _feeCollector)
}

// SetRequestStatus is a paid mutator transaction binding the contract method 0x19d43eb5.
//
// Solidity: function setRequestStatus(bytes32 requestKey, uint8 _status) returns()
func (_Main *MainTransactor) SetRequestStatus(opts *bind.TransactOpts, requestKey [32]byte, _status uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRequestStatus", requestKey, _status)
}

// SetRequestStatus is a paid mutator transaction binding the contract method 0x19d43eb5.
//
// Solidity: function setRequestStatus(bytes32 requestKey, uint8 _status) returns()
func (_Main *MainSession) SetRequestStatus(requestKey [32]byte, _status uint8) (*types.Transaction, error) {
	return _Main.Contract.SetRequestStatus(&_Main.TransactOpts, requestKey, _status)
}

// SetRequestStatus is a paid mutator transaction binding the contract method 0x19d43eb5.
//
// Solidity: function setRequestStatus(bytes32 requestKey, uint8 _status) returns()
func (_Main *MainTransactorSession) SetRequestStatus(requestKey [32]byte, _status uint8) (*types.Transaction, error) {
	return _Main.Contract.SetRequestStatus(&_Main.TransactOpts, requestKey, _status)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_Main *MainTransactor) SetRequestTimeout(opts *bind.TransactOpts, _timeout *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRequestTimeout", _timeout)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_Main *MainSession) SetRequestTimeout(_timeout *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetRequestTimeout(&_Main.TransactOpts, _timeout)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_Main *MainTransactorSession) SetRequestTimeout(_timeout *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetRequestTimeout(&_Main.TransactOpts, _timeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Main *MainTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Main *MainSession) Unpause() (*types.Transaction, error) {
	return _Main.Contract.Unpause(&_Main.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Main *MainTransactorSession) Unpause() (*types.Transaction, error) {
	return _Main.Contract.Unpause(&_Main.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactorSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}

// MainAvsSigsVerifierUpdatedIterator is returned from FilterAvsSigsVerifierUpdated and is used to iterate over the raw logs and unpacked data for AvsSigsVerifierUpdated events raised by the Main contract.
type MainAvsSigsVerifierUpdatedIterator struct {
	Event *MainAvsSigsVerifierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainAvsSigsVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAvsSigsVerifierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainAvsSigsVerifierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainAvsSigsVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAvsSigsVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAvsSigsVerifierUpdated represents a AvsSigsVerifierUpdated event raised by the Main contract.
type MainAvsSigsVerifierUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAvsSigsVerifierUpdated is a free log retrieval operation binding the contract event 0x06297769d0d4011879e4b190001e695e641c69d3dbcaf41b33a926214ad82d8c.
//
// Solidity: event AvsSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) FilterAvsSigsVerifierUpdated(opts *bind.FilterOpts) (*MainAvsSigsVerifierUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AvsSigsVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &MainAvsSigsVerifierUpdatedIterator{contract: _Main.contract, event: "AvsSigsVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchAvsSigsVerifierUpdated is a free log subscription operation binding the contract event 0x06297769d0d4011879e4b190001e695e641c69d3dbcaf41b33a926214ad82d8c.
//
// Solidity: event AvsSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) WatchAvsSigsVerifierUpdated(opts *bind.WatchOpts, sink chan<- *MainAvsSigsVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AvsSigsVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAvsSigsVerifierUpdated)
				if err := _Main.contract.UnpackLog(event, "AvsSigsVerifierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAvsSigsVerifierUpdated is a log parse operation binding the contract event 0x06297769d0d4011879e4b190001e695e641c69d3dbcaf41b33a926214ad82d8c.
//
// Solidity: event AvsSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) ParseAvsSigsVerifierUpdated(log types.Log) (*MainAvsSigsVerifierUpdated, error) {
	event := new(MainAvsSigsVerifierUpdated)
	if err := _Main.contract.UnpackLog(event, "AvsSigsVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBaseDataUrlUpdatedIterator is returned from FilterBaseDataUrlUpdated and is used to iterate over the raw logs and unpacked data for BaseDataUrlUpdated events raised by the Main contract.
type MainBaseDataUrlUpdatedIterator struct {
	Event *MainBaseDataUrlUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainBaseDataUrlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBaseDataUrlUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainBaseDataUrlUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainBaseDataUrlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBaseDataUrlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBaseDataUrlUpdated represents a BaseDataUrlUpdated event raised by the Main contract.
type MainBaseDataUrlUpdated struct {
	From string
	To   string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBaseDataUrlUpdated is a free log retrieval operation binding the contract event 0xe7fa7d4cb7253455bd011caadc607a0db44090fd2ea468dd50c5613ac9e9820f.
//
// Solidity: event BaseDataUrlUpdated(string from, string to)
func (_Main *MainFilterer) FilterBaseDataUrlUpdated(opts *bind.FilterOpts) (*MainBaseDataUrlUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "BaseDataUrlUpdated")
	if err != nil {
		return nil, err
	}
	return &MainBaseDataUrlUpdatedIterator{contract: _Main.contract, event: "BaseDataUrlUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseDataUrlUpdated is a free log subscription operation binding the contract event 0xe7fa7d4cb7253455bd011caadc607a0db44090fd2ea468dd50c5613ac9e9820f.
//
// Solidity: event BaseDataUrlUpdated(string from, string to)
func (_Main *MainFilterer) WatchBaseDataUrlUpdated(opts *bind.WatchOpts, sink chan<- *MainBaseDataUrlUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "BaseDataUrlUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBaseDataUrlUpdated)
				if err := _Main.contract.UnpackLog(event, "BaseDataUrlUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaseDataUrlUpdated is a log parse operation binding the contract event 0xe7fa7d4cb7253455bd011caadc607a0db44090fd2ea468dd50c5613ac9e9820f.
//
// Solidity: event BaseDataUrlUpdated(string from, string to)
func (_Main *MainFilterer) ParseBaseDataUrlUpdated(log types.Log) (*MainBaseDataUrlUpdated, error) {
	event := new(MainBaseDataUrlUpdated)
	if err := _Main.contract.UnpackLog(event, "BaseDataUrlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBrevisDisputeUpdatedIterator is returned from FilterBrevisDisputeUpdated and is used to iterate over the raw logs and unpacked data for BrevisDisputeUpdated events raised by the Main contract.
type MainBrevisDisputeUpdatedIterator struct {
	Event *MainBrevisDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainBrevisDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBrevisDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainBrevisDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainBrevisDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBrevisDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBrevisDisputeUpdated represents a BrevisDisputeUpdated event raised by the Main contract.
type MainBrevisDisputeUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBrevisDisputeUpdated is a free log retrieval operation binding the contract event 0xe04e2434309a3c5548fe58cd48e2054b9f2c3a0581f0df4277aa52d191a8b794.
//
// Solidity: event BrevisDisputeUpdated(address from, address to)
func (_Main *MainFilterer) FilterBrevisDisputeUpdated(opts *bind.FilterOpts) (*MainBrevisDisputeUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "BrevisDisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &MainBrevisDisputeUpdatedIterator{contract: _Main.contract, event: "BrevisDisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchBrevisDisputeUpdated is a free log subscription operation binding the contract event 0xe04e2434309a3c5548fe58cd48e2054b9f2c3a0581f0df4277aa52d191a8b794.
//
// Solidity: event BrevisDisputeUpdated(address from, address to)
func (_Main *MainFilterer) WatchBrevisDisputeUpdated(opts *bind.WatchOpts, sink chan<- *MainBrevisDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "BrevisDisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBrevisDisputeUpdated)
				if err := _Main.contract.UnpackLog(event, "BrevisDisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBrevisDisputeUpdated is a log parse operation binding the contract event 0xe04e2434309a3c5548fe58cd48e2054b9f2c3a0581f0df4277aa52d191a8b794.
//
// Solidity: event BrevisDisputeUpdated(address from, address to)
func (_Main *MainFilterer) ParseBrevisDisputeUpdated(log types.Log) (*MainBrevisDisputeUpdated, error) {
	event := new(MainBrevisDisputeUpdated)
	if err := _Main.contract.UnpackLog(event, "BrevisDisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBrevisProofUpdatedIterator is returned from FilterBrevisProofUpdated and is used to iterate over the raw logs and unpacked data for BrevisProofUpdated events raised by the Main contract.
type MainBrevisProofUpdatedIterator struct {
	Event *MainBrevisProofUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainBrevisProofUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBrevisProofUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainBrevisProofUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainBrevisProofUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBrevisProofUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBrevisProofUpdated represents a BrevisProofUpdated event raised by the Main contract.
type MainBrevisProofUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBrevisProofUpdated is a free log retrieval operation binding the contract event 0xddb7d4b45d521a6718ed9ccac62f1faa18b869772bca7e77ab6f392912a4ec18.
//
// Solidity: event BrevisProofUpdated(address from, address to)
func (_Main *MainFilterer) FilterBrevisProofUpdated(opts *bind.FilterOpts) (*MainBrevisProofUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "BrevisProofUpdated")
	if err != nil {
		return nil, err
	}
	return &MainBrevisProofUpdatedIterator{contract: _Main.contract, event: "BrevisProofUpdated", logs: logs, sub: sub}, nil
}

// WatchBrevisProofUpdated is a free log subscription operation binding the contract event 0xddb7d4b45d521a6718ed9ccac62f1faa18b869772bca7e77ab6f392912a4ec18.
//
// Solidity: event BrevisProofUpdated(address from, address to)
func (_Main *MainFilterer) WatchBrevisProofUpdated(opts *bind.WatchOpts, sink chan<- *MainBrevisProofUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "BrevisProofUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBrevisProofUpdated)
				if err := _Main.contract.UnpackLog(event, "BrevisProofUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBrevisProofUpdated is a log parse operation binding the contract event 0xddb7d4b45d521a6718ed9ccac62f1faa18b869772bca7e77ab6f392912a4ec18.
//
// Solidity: event BrevisProofUpdated(address from, address to)
func (_Main *MainFilterer) ParseBrevisProofUpdated(log types.Log) (*MainBrevisProofUpdated, error) {
	event := new(MainBrevisProofUpdated)
	if err := _Main.contract.UnpackLog(event, "BrevisProofUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBvnSigsVerifierUpdatedIterator is returned from FilterBvnSigsVerifierUpdated and is used to iterate over the raw logs and unpacked data for BvnSigsVerifierUpdated events raised by the Main contract.
type MainBvnSigsVerifierUpdatedIterator struct {
	Event *MainBvnSigsVerifierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainBvnSigsVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBvnSigsVerifierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainBvnSigsVerifierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainBvnSigsVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBvnSigsVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBvnSigsVerifierUpdated represents a BvnSigsVerifierUpdated event raised by the Main contract.
type MainBvnSigsVerifierUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBvnSigsVerifierUpdated is a free log retrieval operation binding the contract event 0x032d76868790fdb5fb04c6f54308ddd47ecf8a80e5995610816aa0b29773b520.
//
// Solidity: event BvnSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) FilterBvnSigsVerifierUpdated(opts *bind.FilterOpts) (*MainBvnSigsVerifierUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "BvnSigsVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &MainBvnSigsVerifierUpdatedIterator{contract: _Main.contract, event: "BvnSigsVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchBvnSigsVerifierUpdated is a free log subscription operation binding the contract event 0x032d76868790fdb5fb04c6f54308ddd47ecf8a80e5995610816aa0b29773b520.
//
// Solidity: event BvnSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) WatchBvnSigsVerifierUpdated(opts *bind.WatchOpts, sink chan<- *MainBvnSigsVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "BvnSigsVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBvnSigsVerifierUpdated)
				if err := _Main.contract.UnpackLog(event, "BvnSigsVerifierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBvnSigsVerifierUpdated is a log parse operation binding the contract event 0x032d76868790fdb5fb04c6f54308ddd47ecf8a80e5995610816aa0b29773b520.
//
// Solidity: event BvnSigsVerifierUpdated(address from, address to)
func (_Main *MainFilterer) ParseBvnSigsVerifierUpdated(log types.Log) (*MainBvnSigsVerifierUpdated, error) {
	event := new(MainBvnSigsVerifierUpdated)
	if err := _Main.contract.UnpackLog(event, "BvnSigsVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the Main contract.
type MainFeeCollectedIterator struct {
	Event *MainFeeCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFeeCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainFeeCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFeeCollected represents a FeeCollected event raised by the Main contract.
type MainFeeCollected struct {
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0xf10cda68996dfb656d49ab0db3c62cc5f0849710633671a337171c3ad9255186.
//
// Solidity: event FeeCollected(uint256 amount, address receiver)
func (_Main *MainFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*MainFeeCollectedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &MainFeeCollectedIterator{contract: _Main.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xf10cda68996dfb656d49ab0db3c62cc5f0849710633671a337171c3ad9255186.
//
// Solidity: event FeeCollected(uint256 amount, address receiver)
func (_Main *MainFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *MainFeeCollected) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFeeCollected)
				if err := _Main.contract.UnpackLog(event, "FeeCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeCollected is a log parse operation binding the contract event 0xf10cda68996dfb656d49ab0db3c62cc5f0849710633671a337171c3ad9255186.
//
// Solidity: event FeeCollected(uint256 amount, address receiver)
func (_Main *MainFilterer) ParseFeeCollected(log types.Log) (*MainFeeCollected, error) {
	event := new(MainFeeCollected)
	if err := _Main.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the Main contract.
type MainFeeCollectorUpdatedIterator struct {
	Event *MainFeeCollectorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFeeCollectorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainFeeCollectorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the Main contract.
type MainFeeCollectorUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_Main *MainFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts) (*MainFeeCollectorUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return &MainFeeCollectorUpdatedIterator{contract: _Main.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_Main *MainFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *MainFeeCollectorUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFeeCollectorUpdated)
				if err := _Main.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_Main *MainFilterer) ParseFeeCollectorUpdated(log types.Log) (*MainFeeCollectorUpdated, error) {
	event := new(MainFeeCollectorUpdated)
	if err := _Main.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOpRequestsFulfilledIterator is returned from FilterOpRequestsFulfilled and is used to iterate over the raw logs and unpacked data for OpRequestsFulfilled events raised by the Main contract.
type MainOpRequestsFulfilledIterator struct {
	Event *MainOpRequestsFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainOpRequestsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOpRequestsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainOpRequestsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainOpRequestsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOpRequestsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOpRequestsFulfilled represents a OpRequestsFulfilled event raised by the Main contract.
type MainOpRequestsFulfilled struct {
	ProofIds        [][32]byte
	Nonces          []uint64
	AppCommitHashes [][32]byte
	AppVkHashes     [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOpRequestsFulfilled is a free log retrieval operation binding the contract event 0x599dc3a262ef12090a85966f98eb3576ad3412b27513e677d6eaaee82a82bea6.
//
// Solidity: event OpRequestsFulfilled(bytes32[] proofIds, uint64[] nonces, bytes32[] appCommitHashes, bytes32[] appVkHashes)
func (_Main *MainFilterer) FilterOpRequestsFulfilled(opts *bind.FilterOpts) (*MainOpRequestsFulfilledIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "OpRequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return &MainOpRequestsFulfilledIterator{contract: _Main.contract, event: "OpRequestsFulfilled", logs: logs, sub: sub}, nil
}

// WatchOpRequestsFulfilled is a free log subscription operation binding the contract event 0x599dc3a262ef12090a85966f98eb3576ad3412b27513e677d6eaaee82a82bea6.
//
// Solidity: event OpRequestsFulfilled(bytes32[] proofIds, uint64[] nonces, bytes32[] appCommitHashes, bytes32[] appVkHashes)
func (_Main *MainFilterer) WatchOpRequestsFulfilled(opts *bind.WatchOpts, sink chan<- *MainOpRequestsFulfilled) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "OpRequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOpRequestsFulfilled)
				if err := _Main.contract.UnpackLog(event, "OpRequestsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpRequestsFulfilled is a log parse operation binding the contract event 0x599dc3a262ef12090a85966f98eb3576ad3412b27513e677d6eaaee82a82bea6.
//
// Solidity: event OpRequestsFulfilled(bytes32[] proofIds, uint64[] nonces, bytes32[] appCommitHashes, bytes32[] appVkHashes)
func (_Main *MainFilterer) ParseOpRequestsFulfilled(log types.Log) (*MainOpRequestsFulfilled, error) {
	event := new(MainOpRequestsFulfilled)
	if err := _Main.contract.UnpackLog(event, "OpRequestsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Main contract.
type MainOwnershipTransferredIterator struct {
	Event *MainOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnershipTransferred represents a OwnershipTransferred event raised by the Main contract.
type MainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnershipTransferredIterator{contract: _Main.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnershipTransferred)
				if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) ParseOwnershipTransferred(log types.Log) (*MainOwnershipTransferred, error) {
	event := new(MainOwnershipTransferred)
	if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Main contract.
type MainPausedIterator struct {
	Event *MainPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPaused represents a Paused event raised by the Main contract.
type MainPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Main *MainFilterer) FilterPaused(opts *bind.FilterOpts) (*MainPausedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MainPausedIterator{contract: _Main.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Main *MainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MainPaused) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPaused)
				if err := _Main.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Main *MainFilterer) ParsePaused(log types.Log) (*MainPaused, error) {
	event := new(MainPaused)
	if err := _Main.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Main contract.
type MainPauserAddedIterator struct {
	Event *MainPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPauserAdded represents a PauserAdded event raised by the Main contract.
type MainPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Main *MainFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*MainPauserAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &MainPauserAddedIterator{contract: _Main.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Main *MainFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *MainPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPauserAdded)
				if err := _Main.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Main *MainFilterer) ParsePauserAdded(log types.Log) (*MainPauserAdded, error) {
	event := new(MainPauserAdded)
	if err := _Main.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Main contract.
type MainPauserRemovedIterator struct {
	Event *MainPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPauserRemoved represents a PauserRemoved event raised by the Main contract.
type MainPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Main *MainFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*MainPauserRemovedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &MainPauserRemovedIterator{contract: _Main.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Main *MainFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *MainPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPauserRemoved)
				if err := _Main.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Main *MainFilterer) ParsePauserRemoved(log types.Log) (*MainPauserRemoved, error) {
	event := new(MainPauserRemoved)
	if err := _Main.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainProverAddedIterator is returned from FilterProverAdded and is used to iterate over the raw logs and unpacked data for ProverAdded events raised by the Main contract.
type MainProverAddedIterator struct {
	Event *MainProverAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainProverAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainProverAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainProverAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainProverAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainProverAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainProverAdded represents a ProverAdded event raised by the Main contract.
type MainProverAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProverAdded is a free log retrieval operation binding the contract event 0xef1fa0a4d797341645c201a742cf59be633da0589e0e3cda511cfc90cd039684.
//
// Solidity: event ProverAdded(address account)
func (_Main *MainFilterer) FilterProverAdded(opts *bind.FilterOpts) (*MainProverAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ProverAdded")
	if err != nil {
		return nil, err
	}
	return &MainProverAddedIterator{contract: _Main.contract, event: "ProverAdded", logs: logs, sub: sub}, nil
}

// WatchProverAdded is a free log subscription operation binding the contract event 0xef1fa0a4d797341645c201a742cf59be633da0589e0e3cda511cfc90cd039684.
//
// Solidity: event ProverAdded(address account)
func (_Main *MainFilterer) WatchProverAdded(opts *bind.WatchOpts, sink chan<- *MainProverAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ProverAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainProverAdded)
				if err := _Main.contract.UnpackLog(event, "ProverAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProverAdded is a log parse operation binding the contract event 0xef1fa0a4d797341645c201a742cf59be633da0589e0e3cda511cfc90cd039684.
//
// Solidity: event ProverAdded(address account)
func (_Main *MainFilterer) ParseProverAdded(log types.Log) (*MainProverAdded, error) {
	event := new(MainProverAdded)
	if err := _Main.contract.UnpackLog(event, "ProverAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainProverRemovedIterator is returned from FilterProverRemoved and is used to iterate over the raw logs and unpacked data for ProverRemoved events raised by the Main contract.
type MainProverRemovedIterator struct {
	Event *MainProverRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainProverRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainProverRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainProverRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainProverRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainProverRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainProverRemoved represents a ProverRemoved event raised by the Main contract.
type MainProverRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProverRemoved is a free log retrieval operation binding the contract event 0xd64d11086d859b73cf85a91ca06cbc484398acabe3d9a1b26d4366dff377d985.
//
// Solidity: event ProverRemoved(address account)
func (_Main *MainFilterer) FilterProverRemoved(opts *bind.FilterOpts) (*MainProverRemovedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ProverRemoved")
	if err != nil {
		return nil, err
	}
	return &MainProverRemovedIterator{contract: _Main.contract, event: "ProverRemoved", logs: logs, sub: sub}, nil
}

// WatchProverRemoved is a free log subscription operation binding the contract event 0xd64d11086d859b73cf85a91ca06cbc484398acabe3d9a1b26d4366dff377d985.
//
// Solidity: event ProverRemoved(address account)
func (_Main *MainFilterer) WatchProverRemoved(opts *bind.WatchOpts, sink chan<- *MainProverRemoved) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ProverRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainProverRemoved)
				if err := _Main.contract.UnpackLog(event, "ProverRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProverRemoved is a log parse operation binding the contract event 0xd64d11086d859b73cf85a91ca06cbc484398acabe3d9a1b26d4366dff377d985.
//
// Solidity: event ProverRemoved(address account)
func (_Main *MainFilterer) ParseProverRemoved(log types.Log) (*MainProverRemoved, error) {
	event := new(MainProverRemoved)
	if err := _Main.contract.UnpackLog(event, "ProverRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestCallbackFailedIterator is returned from FilterRequestCallbackFailed and is used to iterate over the raw logs and unpacked data for RequestCallbackFailed events raised by the Main contract.
type MainRequestCallbackFailedIterator struct {
	Event *MainRequestCallbackFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestCallbackFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestCallbackFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestCallbackFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestCallbackFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestCallbackFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestCallbackFailed represents a RequestCallbackFailed event raised by the Main contract.
type MainRequestCallbackFailed struct {
	ProofId [32]byte
	Nonce   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestCallbackFailed is a free log retrieval operation binding the contract event 0x65ec418e0b7b50180dd133ac2495d373d2df1b616648284452f7b57ac532ce4a.
//
// Solidity: event RequestCallbackFailed(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) FilterRequestCallbackFailed(opts *bind.FilterOpts) (*MainRequestCallbackFailedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestCallbackFailed")
	if err != nil {
		return nil, err
	}
	return &MainRequestCallbackFailedIterator{contract: _Main.contract, event: "RequestCallbackFailed", logs: logs, sub: sub}, nil
}

// WatchRequestCallbackFailed is a free log subscription operation binding the contract event 0x65ec418e0b7b50180dd133ac2495d373d2df1b616648284452f7b57ac532ce4a.
//
// Solidity: event RequestCallbackFailed(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) WatchRequestCallbackFailed(opts *bind.WatchOpts, sink chan<- *MainRequestCallbackFailed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestCallbackFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestCallbackFailed)
				if err := _Main.contract.UnpackLog(event, "RequestCallbackFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestCallbackFailed is a log parse operation binding the contract event 0x65ec418e0b7b50180dd133ac2495d373d2df1b616648284452f7b57ac532ce4a.
//
// Solidity: event RequestCallbackFailed(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) ParseRequestCallbackFailed(log types.Log) (*MainRequestCallbackFailed, error) {
	event := new(MainRequestCallbackFailed)
	if err := _Main.contract.UnpackLog(event, "RequestCallbackFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestFeeIncreasedIterator is returned from FilterRequestFeeIncreased and is used to iterate over the raw logs and unpacked data for RequestFeeIncreased events raised by the Main contract.
type MainRequestFeeIncreasedIterator struct {
	Event *MainRequestFeeIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestFeeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestFeeIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestFeeIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestFeeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestFeeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestFeeIncreased represents a RequestFeeIncreased event raised by the Main contract.
type MainRequestFeeIncreased struct {
	ProofId [32]byte
	Nonce   uint64
	Gas     *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestFeeIncreased is a free log retrieval operation binding the contract event 0x26313df9aa6577df866c6588c888b4c89d59707d6df9fd0befdb29cb976945b6.
//
// Solidity: event RequestFeeIncreased(bytes32 proofId, uint64 nonce, uint256 gas, uint256 fee)
func (_Main *MainFilterer) FilterRequestFeeIncreased(opts *bind.FilterOpts) (*MainRequestFeeIncreasedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestFeeIncreased")
	if err != nil {
		return nil, err
	}
	return &MainRequestFeeIncreasedIterator{contract: _Main.contract, event: "RequestFeeIncreased", logs: logs, sub: sub}, nil
}

// WatchRequestFeeIncreased is a free log subscription operation binding the contract event 0x26313df9aa6577df866c6588c888b4c89d59707d6df9fd0befdb29cb976945b6.
//
// Solidity: event RequestFeeIncreased(bytes32 proofId, uint64 nonce, uint256 gas, uint256 fee)
func (_Main *MainFilterer) WatchRequestFeeIncreased(opts *bind.WatchOpts, sink chan<- *MainRequestFeeIncreased) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestFeeIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestFeeIncreased)
				if err := _Main.contract.UnpackLog(event, "RequestFeeIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFeeIncreased is a log parse operation binding the contract event 0x26313df9aa6577df866c6588c888b4c89d59707d6df9fd0befdb29cb976945b6.
//
// Solidity: event RequestFeeIncreased(bytes32 proofId, uint64 nonce, uint256 gas, uint256 fee)
func (_Main *MainFilterer) ParseRequestFeeIncreased(log types.Log) (*MainRequestFeeIncreased, error) {
	event := new(MainRequestFeeIncreased)
	if err := _Main.contract.UnpackLog(event, "RequestFeeIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the Main contract.
type MainRequestFulfilledIterator struct {
	Event *MainRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestFulfilled represents a RequestFulfilled event raised by the Main contract.
type MainRequestFulfilled struct {
	ProofId [32]byte
	Nonce   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0xc2b28def014c5f6ccb93ba212bc842430ed5b3d5a6e1a28d27e98a568e5fc02f.
//
// Solidity: event RequestFulfilled(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) FilterRequestFulfilled(opts *bind.FilterOpts) (*MainRequestFulfilledIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &MainRequestFulfilledIterator{contract: _Main.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0xc2b28def014c5f6ccb93ba212bc842430ed5b3d5a6e1a28d27e98a568e5fc02f.
//
// Solidity: event RequestFulfilled(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *MainRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestFulfilled)
				if err := _Main.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFulfilled is a log parse operation binding the contract event 0xc2b28def014c5f6ccb93ba212bc842430ed5b3d5a6e1a28d27e98a568e5fc02f.
//
// Solidity: event RequestFulfilled(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) ParseRequestFulfilled(log types.Log) (*MainRequestFulfilled, error) {
	event := new(MainRequestFulfilled)
	if err := _Main.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestRefundedIterator is returned from FilterRequestRefunded and is used to iterate over the raw logs and unpacked data for RequestRefunded events raised by the Main contract.
type MainRequestRefundedIterator struct {
	Event *MainRequestRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestRefunded represents a RequestRefunded event raised by the Main contract.
type MainRequestRefunded struct {
	ProofId [32]byte
	Nonce   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestRefunded is a free log retrieval operation binding the contract event 0x6c005c1eef4990013ad08bd5c941274dcd8d109eb0ca4ac049587173d21be21f.
//
// Solidity: event RequestRefunded(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) FilterRequestRefunded(opts *bind.FilterOpts) (*MainRequestRefundedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestRefunded")
	if err != nil {
		return nil, err
	}
	return &MainRequestRefundedIterator{contract: _Main.contract, event: "RequestRefunded", logs: logs, sub: sub}, nil
}

// WatchRequestRefunded is a free log subscription operation binding the contract event 0x6c005c1eef4990013ad08bd5c941274dcd8d109eb0ca4ac049587173d21be21f.
//
// Solidity: event RequestRefunded(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) WatchRequestRefunded(opts *bind.WatchOpts, sink chan<- *MainRequestRefunded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestRefunded)
				if err := _Main.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestRefunded is a log parse operation binding the contract event 0x6c005c1eef4990013ad08bd5c941274dcd8d109eb0ca4ac049587173d21be21f.
//
// Solidity: event RequestRefunded(bytes32 proofId, uint64 nonce)
func (_Main *MainFilterer) ParseRequestRefunded(log types.Log) (*MainRequestRefunded, error) {
	event := new(MainRequestRefunded)
	if err := _Main.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the Main contract.
type MainRequestSentIterator struct {
	Event *MainRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestSent represents a RequestSent event raised by the Main contract.
type MainRequestSent struct {
	ProofId  [32]byte
	Nonce    uint64
	Refundee common.Address
	Fee      *big.Int
	Callback IBrevisTypesCallback
	Option   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xa067fba43ae94917e16e55db2317937adb10c6b47b9c9e3bb33843fef25874ad.
//
// Solidity: event RequestSent(bytes32 proofId, uint64 nonce, address refundee, uint256 fee, (address,uint64) callback, uint8 option)
func (_Main *MainFilterer) FilterRequestSent(opts *bind.FilterOpts) (*MainRequestSentIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &MainRequestSentIterator{contract: _Main.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xa067fba43ae94917e16e55db2317937adb10c6b47b9c9e3bb33843fef25874ad.
//
// Solidity: event RequestSent(bytes32 proofId, uint64 nonce, address refundee, uint256 fee, (address,uint64) callback, uint8 option)
func (_Main *MainFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *MainRequestSent) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestSent)
				if err := _Main.contract.UnpackLog(event, "RequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestSent is a log parse operation binding the contract event 0xa067fba43ae94917e16e55db2317937adb10c6b47b9c9e3bb33843fef25874ad.
//
// Solidity: event RequestSent(bytes32 proofId, uint64 nonce, address refundee, uint256 fee, (address,uint64) callback, uint8 option)
func (_Main *MainFilterer) ParseRequestSent(log types.Log) (*MainRequestSent, error) {
	event := new(MainRequestSent)
	if err := _Main.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestTimeoutUpdatedIterator is returned from FilterRequestTimeoutUpdated and is used to iterate over the raw logs and unpacked data for RequestTimeoutUpdated events raised by the Main contract.
type MainRequestTimeoutUpdatedIterator struct {
	Event *MainRequestTimeoutUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestTimeoutUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestTimeoutUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestTimeoutUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestTimeoutUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestTimeoutUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestTimeoutUpdated represents a RequestTimeoutUpdated event raised by the Main contract.
type MainRequestTimeoutUpdated struct {
	From *big.Int
	To   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestTimeoutUpdated is a free log retrieval operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_Main *MainFilterer) FilterRequestTimeoutUpdated(opts *bind.FilterOpts) (*MainRequestTimeoutUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return &MainRequestTimeoutUpdatedIterator{contract: _Main.contract, event: "RequestTimeoutUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestTimeoutUpdated is a free log subscription operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_Main *MainFilterer) WatchRequestTimeoutUpdated(opts *bind.WatchOpts, sink chan<- *MainRequestTimeoutUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestTimeoutUpdated)
				if err := _Main.contract.UnpackLog(event, "RequestTimeoutUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestTimeoutUpdated is a log parse operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_Main *MainFilterer) ParseRequestTimeoutUpdated(log types.Log) (*MainRequestTimeoutUpdated, error) {
	event := new(MainRequestTimeoutUpdated)
	if err := _Main.contract.UnpackLog(event, "RequestTimeoutUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestsCallbackFailedIterator is returned from FilterRequestsCallbackFailed and is used to iterate over the raw logs and unpacked data for RequestsCallbackFailed events raised by the Main contract.
type MainRequestsCallbackFailedIterator struct {
	Event *MainRequestsCallbackFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestsCallbackFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestsCallbackFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestsCallbackFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestsCallbackFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestsCallbackFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestsCallbackFailed represents a RequestsCallbackFailed event raised by the Main contract.
type MainRequestsCallbackFailed struct {
	ProofIds [][32]byte
	Nonces   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestsCallbackFailed is a free log retrieval operation binding the contract event 0xf97038d0ee914b31475cdb033a5264b8f2c4d5a89fc292214bea2ac641b49e76.
//
// Solidity: event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) FilterRequestsCallbackFailed(opts *bind.FilterOpts) (*MainRequestsCallbackFailedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestsCallbackFailed")
	if err != nil {
		return nil, err
	}
	return &MainRequestsCallbackFailedIterator{contract: _Main.contract, event: "RequestsCallbackFailed", logs: logs, sub: sub}, nil
}

// WatchRequestsCallbackFailed is a free log subscription operation binding the contract event 0xf97038d0ee914b31475cdb033a5264b8f2c4d5a89fc292214bea2ac641b49e76.
//
// Solidity: event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) WatchRequestsCallbackFailed(opts *bind.WatchOpts, sink chan<- *MainRequestsCallbackFailed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestsCallbackFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestsCallbackFailed)
				if err := _Main.contract.UnpackLog(event, "RequestsCallbackFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestsCallbackFailed is a log parse operation binding the contract event 0xf97038d0ee914b31475cdb033a5264b8f2c4d5a89fc292214bea2ac641b49e76.
//
// Solidity: event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) ParseRequestsCallbackFailed(log types.Log) (*MainRequestsCallbackFailed, error) {
	event := new(MainRequestsCallbackFailed)
	if err := _Main.contract.UnpackLog(event, "RequestsCallbackFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestsFulfilledIterator is returned from FilterRequestsFulfilled and is used to iterate over the raw logs and unpacked data for RequestsFulfilled events raised by the Main contract.
type MainRequestsFulfilledIterator struct {
	Event *MainRequestsFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainRequestsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainRequestsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainRequestsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestsFulfilled represents a RequestsFulfilled event raised by the Main contract.
type MainRequestsFulfilled struct {
	ProofIds [][32]byte
	Nonces   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestsFulfilled is a free log retrieval operation binding the contract event 0x5222b9309369092c135ac01f588e14730b87b05890b5c016dcb1f9b9a55d14f0.
//
// Solidity: event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) FilterRequestsFulfilled(opts *bind.FilterOpts) (*MainRequestsFulfilledIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return &MainRequestsFulfilledIterator{contract: _Main.contract, event: "RequestsFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestsFulfilled is a free log subscription operation binding the contract event 0x5222b9309369092c135ac01f588e14730b87b05890b5c016dcb1f9b9a55d14f0.
//
// Solidity: event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) WatchRequestsFulfilled(opts *bind.WatchOpts, sink chan<- *MainRequestsFulfilled) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestsFulfilled)
				if err := _Main.contract.UnpackLog(event, "RequestsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestsFulfilled is a log parse operation binding the contract event 0x5222b9309369092c135ac01f588e14730b87b05890b5c016dcb1f9b9a55d14f0.
//
// Solidity: event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces)
func (_Main *MainFilterer) ParseRequestsFulfilled(log types.Log) (*MainRequestsFulfilled, error) {
	event := new(MainRequestsFulfilled)
	if err := _Main.contract.UnpackLog(event, "RequestsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Main contract.
type MainUnpausedIterator struct {
	Event *MainUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUnpaused represents a Unpaused event raised by the Main contract.
type MainUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Main *MainFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MainUnpausedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MainUnpausedIterator{contract: _Main.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Main *MainFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MainUnpaused) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUnpaused)
				if err := _Main.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Main *MainFilterer) ParseUnpaused(log types.Log) (*MainUnpaused, error) {
	event := new(MainUnpaused)
	if err := _Main.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
