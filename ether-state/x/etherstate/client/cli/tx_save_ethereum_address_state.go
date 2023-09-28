package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSaveEthereumAddressState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-ethereum-address-state [address] [blockNumber] [nonce] [storage-position]",
		Short: "Broadcast message save-ethereum-address-state",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argBlockNumber, err := cast.ToUint64E(args[1])
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

			msg := types.NewMsgSaveEthereumAddressState(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argBlockNumber,
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
