

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


#ignite scaffold module statether
 
ignite scaffold type ethaddress-storage-position-stored address block:uint nonce:uint storage-position:uint active:bool
ignite scaffold type ethaddress-storage-position block:uint nonce:uint storage-position:uint active:bool

# add address
ignite scaffold message add-address ethAddress --response id
ignite scaffold message remove-address ethAddress --response id
ignite scaffold message disable-address ethAddress --response id

# save the storage-position for an address
ignite scaffold message save-ethaddress-storage-position ethAddress data:EthaddressStoragePosition --response id

#get the storage-position for an address
ignite scaffold query get-ethaddress-storage-position ethAddress --response data:EthaddressStoragePosition

# display addresses information
ignite scaffold query get-all-ethaddresses-storage-positiona  --response data:EthaddressStoragePosition --paginated
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