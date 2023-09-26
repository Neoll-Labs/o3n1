

# required tools
https://docs.ignite.com/guide/getting-started
git
golang

# init demo chain

init `demoChain`
```shell
ignite scaffold chain example
```
expect output
```shell
⭐️ Successfully created a new blockchain 'demo'.
👉 Get started with the following commands:

 % cd demo
 % ignite chain serve
```

start the blockchain with
```shell
 % cd demo
 % ignite chain serve
```

expected output
```shell
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500
```


# ignite boilerplate
```shell

ignite scaffold module store

ignite scaffold message storeStatus retriveStatus to:string from:string \
--module store \
--response capturedX:int

```
