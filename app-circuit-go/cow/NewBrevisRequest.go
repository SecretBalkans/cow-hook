package cow

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
)

func NewBrevisProof(address common.Address, client *ethclient.Client) (*MainTransactor, error) {
	// Create a contract instance using NewBoundContract
	abiContent, err := os.ReadFile("BrevisRequest.abi")
	if err != nil {
		panic(err)
	}
	parsedABI, err := abi.JSON(strings.NewReader(string(abiContent)))
	if err != nil {
		return nil, err
	}

	// Bind contract
	boundContract := bind.NewBoundContract(address, parsedABI, client, client, client)

	// Return the custom MainTransactor struct
	return &MainTransactor{contract: boundContract}, nil
}
