# GOALS
* Your goal is to utilize the Cosmos SDK to build a prototype blockchain that can read state from Ethereum, 
verify its validity, and store it on your chain.


## features

* register Ethereum addresses for monitoring

* Save Ethereum *address storage position* (state) with the block number and nonce

* suspend the Ethereum address monitoring

* remove the Ethereum address monitoring


* List all Ethereum addresses enabled



# required tools
https://docs.ignite.com/guide/getting-started
git
golang

# init demo chain

init `demoChain`
```shell
ignite scaffold chain statether 
```
expect output
```shell
⭐️ Successfully created a new blockchain 'statether'.
👉 Get started with the following commands:

```

start the blockchain with
```shell
cd statether
ignite chain serve
```

expected output
```shell
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500
```


# cms ignite boilerplate

```shell

ignite scaffold chain storepositionether 

cd storepositionether


ignite scaffold map ethAddress block:uint nonce:uint storage-position:uint active:bool \
    --index index \
    --module storepositionether \
    --no-message    


ignite scaffold module statether
 
ignite scaffold type ethaddress-storage-position ethAddress block:uint nonce:uint storage-position:uint active:bool \
    --module storepositionether
 

# add address
ignite scaffold message add-address creator ethAddress --response ethAddress \
    --module storepositionether
ignite scaffold message remove-address creator ethAddress --response ethAddress \
    --module storepositionether
ignite scaffold message disable-address creator ethAddress --response ethAddress \
    --module storepositionether

# save the storage-position for an address
ignite scaffold message save-ethaddress-storage-position ethAddress block:uint nonce:uint storage-position:uint --response ethAddress \
    --module storepositionether

#get the storage-position for an address
ignite scaffold query get-ethaddress-storage-position ethAddress --response data:EthaddressStoragePosition\
    --module storepositionether

# display addresses information
ignite scaffold query get-all-ethaddresses-storage-positiona  --response data:EthaddressStoragePosition --paginated\
    --module storepositionether
```


# command line

## add ethereum address
```shell
 ./main tx  storepositionether add-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8 --from cosmos1kcr3cx8lxc0w6uxdh93c3mef5v0wkgfcs6vvzq --gas auto
```

## save storage position for ethereum address(also store the nonce and block number)

```shell
./main tx  storepositionether save-ethaddress-storage-position 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8 0 0 0 --from cosmos1kcr3cx8lxc0w6uxdh93c3mef5v0wkgfcs6vvzq --gas auto
```


```go
package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/statether module sentinel errors
var (
	ErrInvalidEthereumAddress = errors.Register(ModuleName, 1101, "invalid ethereum address")
	ErrInvalidNonce           = errors.Register(ModuleName, 1102, "invalid nonce")
	ErrInvalidBlockNumber     = errors.Register(ModuleName, 1103, "invalid block number")
	ErrInvalidStoragePosition = errors.Register(ModuleName, 1104, "invalid storage position")
)

```