package playground

import (
	"context"
	"github.com/nelsonstr/o3n1/ether-state/relayer/cosmos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnableAddress(t *testing.T) {
	//const ethAddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA1"
	const ethAddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8"

	ctx := context.Background()
	c := cosmos.NewEtherStateClient(ctx)

	addressEnable, err := c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)

	if addressEnable.Active {
		err = c.DisableEthereumAddress(ctx, cosmos.Alice, ethAddress)
		assert.NoError(t, err)
	}

	addressEnable, err = c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)
	assert.False(t, addressEnable.Active)

	err = c.EnableEthereumAddress(ctx, cosmos.Alice, ethAddress)
	assert.NoError(t, err)

	addressEnable, err = c.QueryEthereumAddress(ctx, ethAddress)
	assert.NoError(t, err)
	assert.True(t, addressEnable.Active)
}
