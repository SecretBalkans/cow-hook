# APP_CIRCUIT GO using Brevis

### Requirements
1. go 1.2
2. environment varialb
   1. `PRIVATE_KEY` Sepolia ETH wallet private key for sending requests to Brevis Sepolia
   2. `ETH_RPC` RPC url for ETH network for constructing dummy data stream
   3. `SEPOLIA_RPC` RPC url for SEPOLIA network for making Brevis.sendRequest tx with proofId and callback info

## Running

`$ PRIVATE_KEY="xx" ETH_RPC="xx" SEPOLIA_RPC="xx" go test cow/circuit_test.go`
