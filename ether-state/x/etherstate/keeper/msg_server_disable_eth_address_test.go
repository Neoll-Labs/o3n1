package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/testutil/sample"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_msgServer_DisableEthAddress(t *testing.T) {
	alice := sample.AccAddress()
	address := "0xe8aCaaB95d1102D099F82F03f6106289ee19abA1"
	msgSrvr, context := setupMsgServer(t)
	msgSrvr.DisableEthAddress(context, &types.MsgDisableEthAddress{
		Creator: alice,
		Address: address,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "etherstate.etherstate.MsgDisableEthAddress",
		Attributes: []sdk.Attribute{
			{Key: "address", Value: "\"" + address + "\""},
			{Key: "creator", Value: "\"" + alice + "\""},
		},
	}, event)
}
