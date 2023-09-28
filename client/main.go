package main

import (
	"context"
	"fmt"
	types2 "github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"log"
	"os"
	"path/filepath"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func main() {
	ctx := context.Background()
	addressPrefix := "cosmos"

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	homePath := filepath.Join(home, ".ether-state")

	cosmosOptions := []cosmosclient.Option{
		cosmosclient.WithHome(homePath),
		cosmosclient.WithAddressPrefix(addressPrefix),
	}

	// create an instance of cosmosclient
	client, err := cosmosclient.New(context.Background(), cosmosOptions...)
	if err != nil {
		log.Fatal(err)
	}

	//Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Define a message to create a post
	msg := &types2.MsgEnableEthAddress{
		Creator: addr,
		Address: "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8",
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	// Instantiate a query client for your `blog` blockchain
	queryClient := types2.NewQueryClient(client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.EthereumAddressAll(ctx, &types2.QueryAllEthereumAddressRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
