package simulation

import (
	"math/rand"

	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/keeper"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEnableEthAddress(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEnableEthAddress{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the EnableEthAddress simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "EnableEthAddress simulation not implemented"), nil, nil
	}
}
