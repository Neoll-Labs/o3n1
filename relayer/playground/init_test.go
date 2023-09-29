package playground

import (
	"context"
	"fmt"
	"github.com/cometbft/cometbft/types"
	"github.com/nelsonstr/o3n1/ether-state/relayer/cosmos"
	"github.com/stretchr/testify/assert"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"testing"
	"time"
)

// const ethAddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA1"
const ethAddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8"

func TestDisableAddress(t *testing.T) {
	ctx := context.Background()
	c := cosmos.NewEtherStateClient(ctx)

	err := c.DisableEthereumAddress(ctx, cosmos.Alice, ethAddress)
	assert.NoError(t, err)
}

func TestEnableAddress(t *testing.T) {
	ctx := context.Background()
	c := cosmos.NewEtherStateClient(ctx)

	err := c.EnableEthereumAddress(ctx, cosmos.Alice, ethAddress)
	assert.NoError(t, err)
}

func TestFlow_Disable_Query_Enable_Query(t *testing.T) {
	ctx := context.Background()
	c := cosmos.NewEtherStateClient(ctx)
	addressEnable, err := c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)

	err = c.DisableEthereumAddress(ctx, cosmos.Alice, ethAddress)
	assert.NoError(t, err)

	addressEnable, err = c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)
	assert.False(t, addressEnable.Active)

	err = c.EnableEthereumAddress(ctx, cosmos.Alice, ethAddress)
	assert.NoError(t, err)

	addressEnable, err = c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)
	assert.True(t, addressEnable.Active)
}

func TestSubscribe(t *testing.T) {
	client, err := rpchttp.New("tcp://localhost:26657")
	assert.NoError(t, err)
	err = client.Start()
	assert.NoError(t, err)

	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_ = ctx
	//
	err = client.UnsubscribeAll(ctx, "test-client")
	assert.NoError(t, err)
	//
	query := "message.action = '/etherstate.etherstate.MsgEnableEthAddress'"
	query = "tm.event='Tx'"
	txs, err := client.Subscribe(ctx, "test-client", query)
	assert.NoError(t, err)

	go func() {
		for e := range txs {
			fmt.Println("got ", e.Data.(types.EventDataTx))
		}
	}()
	assert.NoError(t, err)
}
