package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSaveEthaddressState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-ethaddress-state [eth-address] [block] [nonce] [storage-position]",
		Short: "Broadcast message save-ethaddress-state",
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

			msg := types.NewMsgSaveEthaddressState(
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
