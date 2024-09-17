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


