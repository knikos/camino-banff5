# Release Notes

## [v1.9.5](https://github.com/ava-labs/avalanchego/releases/tag/v1.9.5)

This version is backwards compatible to [v1.9.0](https://github.com/ava-labs/avalanchego/releases/tag/v1.9.0). It is optional, but encouraged. The supported plugin version is `21`.

### Subnet Messaging

- Added subnet message serialization format
- Added subnet message signing
- Replaced `bls.SecretKey` with a `teleporter.Signer` in the `snow.Context`
- Moved `SNLookup` into the `validators.State` interface to support non-whitelisted chainID to subnetID lookups
- Added support for non-whitelisted subnetIDs for fetching the validator set at a given height
- Added subnet message verification
- Added `teleporter.AnycastID` to denote a subnet message not intended for a specific chain

### Fixes

- Added re-gossip of updated validator IPs
- Fixed `rpcchainvm.BatchedParseBlock` to correctly wrap returned blocks
- Removed incorrect `uintptr` handling in the generic codec
- Removed message latency tracking on messages being sent to itself

### Coreth

- Added support for eth_call over VM2VM messaging
- Added config flags for tx pool behavior

### Miscellaneous

- Added networking package README.md
- Removed pagination of large db messages over gRPC
- Added `Size` to the generic codec to reduce allocations
- Added `UnpackLimitedBytes` and `UnpackLimitedStr` to the manual packer
- Added SECURITY.md
- Exposed proposer list from the `proposervm`'s `Windower` interface
- Added health and bootstrapping client helpers that block until the node is healthy
- Moved bit sets from the `ids` package to the `set` package
- Added more wallet examples

## [v1.9.4](https://github.com/ava-labs/avalanchego/releases/tag/v1.9.4)

This version is the first one running on our testnet. The supported plugin version is `20`.

### AvalancheGo version
- Based on avalanchego [v1.9.4](https://github.com/ava-labs/avalanchego/releases/tag/v1.9.4)

### Camino features
- Added AddressStates to track KYC and other address related states
- NodeSignature verification (instead BLS): NodeID / privateKey SECP256K1 pair
- Deposit and Bonding on P-Chain instead Staking and Delegation
- DepositOffer (IncentivePools)
- P-Chain GetConfiguration API to initialize clients
