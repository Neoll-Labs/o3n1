package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

func CmdListEthereumAddressState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ethereum-address-state",
		Short: "list all ethereum-address-state",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEthereumAddressStateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EthereumAddressStateAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEthereumAddressState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ethereum-address-state [index]",
		Short: "shows a ethereum-address-state",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetEthereumAddressStateRequest{
				Index: argIndex,
			}

			res, err := queryClient.EthereumAddressState(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
