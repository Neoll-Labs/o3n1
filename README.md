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

* Save Ethereum *address storage position* (state) with the block number and nonce

* Get Ethereum state from blockchain
  * query 

* List all Ethereum addresses enabled

  * relayer
    * listening the changes on the ethereum address 
      * on change subscrive event on ethereum 
        ```shell
          wscat -c wss://mainnet.infura.io/ws/v3/YOUR-API-KEY > {"jsonrpc":  "2.0",  "id":  1,  "method":  "eth_subscribe",  "params":  ["newHeads"]}
        ```


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

ignite chain serve -v
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


```

### command line to enable / disable the address
```shell
go build -o ether-state cmd/ether-stated/main.go

./ether-state query etherstate list-ethereum-address

./ether-state tx etherstate enable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8  --from cosmos10wtz2ckpzzgek0n4w4mpy4mrrnpwu3zx6nxm32 --gas auto  
./ether-state query etherstate show-ethereum-address  0xe8aCaaB95d1102D099F82F03f6106289ee19abA8

./ether-state tx etherstate disable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8  --from cosmos10wtz2ckpzzgek0n4w4mpy4mrrnpwu3zx6nxm32 --gas auto  
./ether-state query etherstate show-ethereum-address  0xe8aCaaB95d1102D099F82F03f6106289ee19abA8


./ether-state query etherstate list-ethereum-address-state

```

## save state

```shell
# eth address is the index
ignite scaffold map ethereum-address-state state:uint blockNumber:uint nonce:uint \
    --index index \
    --module ether-state \
    --no-message

ignite scaffold message save-ethereum-address-state address blockNumber:uint nonce:uint storage-position:uint \
    --response success:bool \
    --module ether-state

```

### command save state of ethereum address
```shell
# build
go build -o ether-state cmd/ether-stated/main.go

# enable an address
./ether-state tx etherstate enable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA8  --from cosmos1g5u6nm433jwpa2advpzrp80xpesnw3am3fxfek --gas auto  

# save state
./ether-state tx etherstate save-ethereum-address-state  0xe8aCaaB95d1102D099F82F03f6106289ee19abA8 10 10 10 --from cosmos1g5u6nm433jwpa2advpzrp80xpesnw3am3fxfek --gas auto
./ether-state query etherstate show-ethereum-address-state  0xe8aCaaB95d1102D099F82F03f6106289ee19abA8


# enable an address
./ether-state tx etherstate enable-eth-address 0xe8aCaaB95d1102D099F82F03f6106289ee19abA1  --from cosmos1g5u6nm433jwpa2advpzrp80xpesnw3am3fxfek --gas auto  

# save state
./ether-state tx etherstate save-ethereum-address-state  0xe8aCaaB95d1102D099F82F03f6106289ee19abA1 0 0 0 --from cosmos1g5u6nm433jwpa2advpzrp80xpesnw3am3fxfek --gas auto
./ether-state query etherstate show-ethereum-address-state  0xe8aCaaB95d1102D099F82F03f6106289ee19abA1

# list all 
./ether-state query etherstate list-ethereum-address-state

```

Connect to websocket server   
```shell
wscat -c ws://localhost:26657/websocket


	wscat -c ws://localhost:26657/websocket 
	
	{ "jsonrpc": "2.0", "method": "subscribe", "id": 0, "params": {"query": "tm.event = 'Tx'"}}
	{ "jsonrpc": "2.0", "method": "subscribe", "id": 1, "params": ["tm.event='NewBlock'"] }
	{ "jsonrpc": "2.0", "method": "subscribe", "id": 1, "params": ["tm.event='Tx'"] }
 
 
{"jsonrpc": "2.0","method": "subscribe", "id": 0,"params": {"query": "message.module = 'etherstate'"}}
 


{"jsonrpc": "2.0","method": "subscribe", "id": 0,"params": {"query": "message.action = '/etherstate.etherstate.MsgEnableEthAddress'"}}
{"jsonrpc": "2.0","method": "subscribe", "id": 0,"params": {"query": "message.action = '/etherstate.etherstate.MsgDisableEthAddress'"}}
 
```