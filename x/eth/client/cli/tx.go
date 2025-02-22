package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/functionx/fx-core/v4/x/crosschain/client/cli"
	"github.com/functionx/fx-core/v4/x/eth/types"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "ETH transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(cli.GetTxSubCmds(types.ModuleName)...)
	return cmd
}
