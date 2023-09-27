package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"statether/x/statether/types"
)

var _ = strconv.Itoa(0)

func CmdSaveEthaddressStoragePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-ethaddress-storage-position [data]",
		Short: "Broadcast message save-ethaddress-storage-position",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argData := new(types.EthaddressStoragePosition)
			err = json.Unmarshal([]byte(args[0]), argData)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveEthaddressStoragePosition(
				clientCtx.GetFromAddress().String(),
				argData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
