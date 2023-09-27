package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"storepositionether/x/storepositionether/keeper"
	"storepositionether/x/storepositionether/types"
)

func SimulateMsgSaveEthaddressStoragePosition(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSaveEthaddressStoragePosition{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SaveEthaddressStoragePosition simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SaveEthaddressStoragePosition simulation not implemented"), nil, nil
	}
}
