package cosmos

import (
	"context"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	types2 "github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"log"
	"os"
	"path/filepath"
)

const addressPrefix = "cosmos"
const Alice = "alice"
const Bob = "bob"

type etherStateClient struct {
	c cosmosclient.Client
}

type Client interface {
	CosmosQueryClient() (types2.QueryClient, error)
	EnableEthereumAddress(ctx context.Context, accountName, address string) error
	DisableEthereumAddress(ctx context.Context, accountName, address string) error
	SaveEthereumAddressState(ctx context.Context, accountName, address string, blockNumber, nonce, storagePosition uint64) error
	QueryAllEthereumAddress(ctx context.Context) ([]types2.EthereumAddress, error)
	QueryEthereumAddress(ctx context.Context, address string) (*types2.EthereumAddress, error)
}

var _ Client = (*etherStateClient)(nil)

func NewEtherStateClient(ctx context.Context) etherStateClient {

	return etherStateClient{
		c: createClient(ctx),
	}
}

func createClient(ctx context.Context) cosmosclient.Client {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	homePath := filepath.Join(home, ".ether-state")

	cosmosOptions := []cosmosclient.Option{
		cosmosclient.WithHome(homePath),
		cosmosclient.WithAddressPrefix(addressPrefix),
	}

	client, err := cosmosclient.New(context.Background(), cosmosOptions...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (c etherStateClient) CosmosQueryClient() (types2.QueryClient, error) {
	queryClient := types2.NewQueryClient(c.c.Context())

	return queryClient, nil
}

func (c etherStateClient) QueryAllEthereumAddress(ctx context.Context) ([]types2.EthereumAddress, error) {
	queryClient := types2.NewQueryClient(c.c.Context())
	all, err := queryClient.EthereumAddressAll(ctx, &types2.QueryAllEthereumAddressRequest{})
	if err != nil {
		return nil, err
	}
	return all.EthereumAddress, err
}

func (c etherStateClient) QueryEthereumAddress(ctx context.Context, address string) (*types2.EthereumAddress, error) {
	queryClient := types2.NewQueryClient(c.c.Context())

	all, err := queryClient.EthereumAddress(ctx, &types2.QueryGetEthereumAddressRequest{Index: address})
	if err != nil {
		return nil, err
	}

	return &all.EthereumAddress, nil
}

func (c etherStateClient) SaveEthereumAddressState(ctx context.Context, accountName, address string, blocknumber, nonce, storagePosition uint64) error {
	account, addr, err := c.address(accountName)
	if err != nil {
		return err
	}

	// Define a message to create a post
	msg := &types2.MsgSaveEthereumAddressState{
		Creator:         addr,
		Address:         address,
		BlockNumber:     blocknumber,
		Nonce:           nonce,
		StoragePosition: storagePosition,
	}

	txResp, err := c.c.BroadcastTx(ctx, *account, msg)
	if err != nil {
		return err
	}

	_ = txResp

	return nil
}

func (c etherStateClient) EnableEthereumAddress(ctx context.Context, accountName, address string) error {
	account, addr, err := c.address(accountName)
	if err != nil {
		return err
	}

	// Define a message to create a post
	msg := &types2.MsgEnableEthAddress{
		Creator: addr,
		Address: address,
	}

	txResp, err := c.c.BroadcastTx(ctx, *account, msg)
	if err != nil {
		return err
	}

	_ = txResp

	return nil
}

func (c etherStateClient) DisableEthereumAddress(ctx context.Context, accountName, address string) error {
	account, addr, err := c.address(accountName)
	if err != nil {
		return err
	}

	// Define a message to create a post
	msg := &types2.MsgDisableEthAddress{
		Creator: addr,
		Address: address,
	}

	txResp, err := c.c.BroadcastTx(ctx, *account, msg)
	if err != nil {
		return err
	}

	_ = txResp

	return nil
}

func (c etherStateClient) address(accountName string) (*cosmosaccount.Account, string, error) {
	account, err := c.c.Account(accountName)
	if err != nil {
		return nil, "", err
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		return nil, "", err
	}
	return &account, addr, nil
}
