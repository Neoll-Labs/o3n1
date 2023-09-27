package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"storepositionether/x/storepositionether/types"
)

var _ = strconv.Itoa(0)

func CmdSaveEthaddressStoragePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-ethaddress-storage-position [eth-address] [block] [nonce] [storage-position]",
		Short: "Broadcast message save-ethaddress-storage-position",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEthAddress := args[0]
			argBlock, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argNonce, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argStoragePosition, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveEthaddressStoragePosition(
				clientCtx.GetFromAddress().String(),
				argEthAddress,
				argBlock,
				argNonce,
				argStoragePosition,
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
