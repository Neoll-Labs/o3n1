package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"statether/x/statether/types"
)

var _ = strconv.Itoa(0)

func CmdGetEthaddressStoragePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-ethaddress-storage-position [address]",
		Short: "Query get-ethaddress-storage-position",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEthaddressStoragePositionRequest{

				Address: reqAddress,
			}

			res, err := queryClient.GetEthaddressStoragePosition(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
