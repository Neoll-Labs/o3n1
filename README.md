# GOALS
* Your goal is to utilize the Cosmos SDK to build a prototype blockchain that can read state from Ethereum, 
verify its validity, and store it on your chain.

* Your chain should be able to agree upon some Ethereum state.

* The mechanism that facilitates this agreement is up to you.

* Your chain should be able to store the state value for an Ethereum address  + storage slot.

* The address/storage slot pair should be parameterized such that any address/slot 
combination can be stored on your chain.

* Once some storage value is agreed upon,
it should be stored on your blockchain along with any necessary metadata that can be used to identify it.

* Users should be able to read data this from your blockchain.




## features

* register Ethereum addresses for monitoring
  * create a map of addresses and boolean
  * create update read eth address

* suspend the Ethereum address monitoring

* remove the Ethereum address monitoring

* Save Ethereum *address storage position* (state) with the block number and nonce

* Get Ethereum state from blockchain
  * query 


* List all Ethereum addresses enabled



# required tools
https://docs.ignite.com/guide/getting-started
git
golang

# init demo chain

init `statether`
```shell
ignite scaffold chain ether-state
```
expect output
```shell
⭐️ Successfully created a new blockchain 'ether-state'.
👉 Get started with the following commands:

```

start the blockchain with
```shell
cd ether-state
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

ignite scaffold chain github.com/nelsonstr/o3n1/ether-state 

ignite generate ts-client

cd ether-state

ignite chain serve 
ignite chain debug --server --server-address 127.0.0.1:30500

```

http://localhost:1317/

```shell

# create map for ethereum addresses and 

# eth address is the index
ignite scaffold map ethereum-address active:bool \
    --index index \
    --module ether-state \
    --no-message

# add address
ignite scaffold message enable-eth-address  address --response success:bool --module ether-state
ignite scaffold message disable-eth-address address --response success:bool --module ether-state



# eth address is the index
ignite scaffold map statether state:uint block:uint nonce:uint \
    --index index \
    --module ether-state \
    --no-message

ignite scaffold message save-ethaddress-state ethAddress block:uint nonce:uint storage-position:uint \
    --response ethAddress \
    --module statether


ignite scaffold message get-ethaddress-state ethAddress \
    --response data:Statether \
    --module statether

--------------------------

ignite scaffold type ethaddress-storage-position ethAddress block:uint nonce:uint storage-position:uint active:bool \
    --module statether
 
 # save the storage-position for an address
ignite scaffold message save-ethaddress-state ethAddress block:uint nonce:uint storage-position:uint --response ethAddress \
    --module statether

#get the state for an address
ignite scaffold query get-ethaddress-state ethAddress --response data:EthaddressStoragePosition\
    --module statether



# save the storage-position for an address
ignite scaffold message save-ethaddress-state ethAddress block:uint nonce:uint storage-position:uint --response ethAddress \
    --module storepositionether



# display addresses information
ignite scaffold query get-all-ethaddresses-storage-positiona  --response data:EthaddressStoragePosition --paginated\
    --module storepositionether
```


# command line


##  ethereum address enable

```shell

./main tx etherstate disable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8  --from cosmos1ke87p8tf0a28v4muvezhrxz9r93ewddg3q4ag6 --gas auto  
./main tx etherstate enable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8  --from cosmos1ke87p8tf0a28v4muvezhrxz9r93ewddg3q4ag6 --gas auto  


./main query etherstate show-statether 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8



```



## add ethereum address
```shell
 ./main tx statether add-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8 --from cosmos17tzhfv8zpjx2weplrgzwetsg5ah53xu9hza0sr --gas auto
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


```shell

curl -X GET "http://localhost:1317/cosmos/auth/v1beta1/accounts/cosmos17tzhfv8zpjx2weplrgzwetsg5ah53xu9hza0sr" -H  "accept: application/json"

```