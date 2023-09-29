package playground

import (
	"context"
	"github.com/nelsonstr/o3n1/ether-state/relayer/cosmos"
	"github.com/stretchr/testify/assert"
	"testing"
)

// const ethAddress = "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"
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
