

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


# ignite boilerplate
```shell

#ignite scaffold module statether

ignite scaffold type ethaddress-storage-position address:string block:uint nonce:uint storage-position:uint 

# add address
ignite scaffold message add-address address:string --response id

# save the storage-position for an address
ignite scaffold message save-ethaddress-storage-position data:EthaddressStoragePosition --response id

#get the storage-position for an address
ignite scaffold query get-ethaddress-storage-position address:string --response data:EthaddressStoragePosition

# display addresses information
ignite scaffold query fetch-ethaddresses-storage-positiona  --response data:EthaddressStoragePosition --paginated
```
