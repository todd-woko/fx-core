package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdkcfg "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/evmos/ethermint/crypto/hd"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/functionx/fx-core/v4/app"
	"github.com/functionx/fx-core/v4/client/cli"
	fxserver "github.com/functionx/fx-core/v4/server"
	fxcfg "github.com/functionx/fx-core/v4/server/config"
	fxtypes "github.com/functionx/fx-core/v4/types"
	avalanchecli "github.com/functionx/fx-core/v4/x/avalanche/client/cli"
	bsccli "github.com/functionx/fx-core/v4/x/bsc/client/cli"
	crosschaincli "github.com/functionx/fx-core/v4/x/crosschain/client/cli"
	ethcli "github.com/functionx/fx-core/v4/x/eth/client/cli"
	polygoncli "github.com/functionx/fx-core/v4/x/polygon/client/cli"
	troncli "github.com/functionx/fx-core/v4/x/tron/client/cli"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() *cobra.Command {
	fxtypes.SetConfig(false)

	encodingConfig := app.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithOutput(os.Stdout).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(fxtypes.GetDefaultNodeHome()).
		WithViper(fxtypes.EnvPrefix).
		WithKeyringOptions(hd.EthSecp256k1Option())

	rootCmd := &cobra.Command{
		Use:   fxtypes.Name + "d",
		Short: "FunctionX Core BlockChain App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			// read flag
			initClientCtx, err = client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// read client.toml
			initClientCtx, err = sdkcfg.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			// set clientCtx
			if err = client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := fxcfg.AppConfig(fmt.Sprintf("4000000000000%s", fxtypes.DefaultDenom))
			if err = server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig()); err != nil {
				return err
			}
			return nil
		},
	}

	initRootCmd(rootCmd, encodingConfig, fxtypes.GetDefaultNodeHome())
	return rootCmd
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig app.EncodingConfig, defaultNodeHome string) {
	rootCmd.AddCommand(
		cli.Debug(),
		cli.InitCmd(defaultNodeHome, app.NewDefAppGenesisByDenom(fxtypes.DefaultDenom, encodingConfig.Codec), app.CustomGenesisConsensusParams()),
		cli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		cli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, defaultNodeHome),
		cli.AddGenesisAccountCmd(defaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(),
		configCmd(),
	)

	myAppCreator := appCreator{encodingConfig}

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		cli.StatusCommand(),
		keyCommands(defaultNodeHome),
		queryCommand(),
		txCommand(),
		version.NewVersionCommand(),
		server.NewRollbackCmd(myAppCreator.newApp, defaultNodeHome),
		fxserver.DataCmd(),
		fxserver.ExportSateCmd(myAppCreator.appExport, defaultNodeHome),
		fxserver.StartCmd(myAppCreator.newApp, defaultNodeHome),
		fxserver.TendermintCommand(),
		app.GetUpgrades().GetLatest().PreUpgradeCmd,
		doctorCmd(),
	)

	// add rosetta
	rootCmd.AddCommand(server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Codec))
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		cli.BlockCommand(),
		cli.QueryTxsByEventsCmd(),
		cli.QueryTxCmd(),
		cli.QueryStoreCmd(),
		cli.QueryValidatorByConsAddr(),
		cli.QueryBlockResultsCmd(),
		cli.QueryGasPricesCmd(),
		crosschaincli.GetQueryCmd(
			avalanchecli.GetQueryCmd(),
			bsccli.GetQueryCmd(),
			ethcli.GetQueryCmd(),
			polygoncli.GetQueryCmd(),
			troncli.GetQueryCmd(),
		),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		crosschaincli.GetTxCmd(
			avalanchecli.GetTxCmd(),
			bsccli.GetTxCmd(),
			ethcli.GetTxCmd(),
			polygoncli.GetTxCmd(),
			troncli.GetTxCmd(),
		),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg app.EncodingConfig
}

// newApp is an AppCreator
func (a appCreator) newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := dbm.NewDB("metadata", sdkserver.GetAppDBBackend(appOpts), snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	gasPrice := cast.ToString(appOpts.Get(server.FlagMinGasPrices))
	if strings.Contains(gasPrice, ".") {
		panic("Invalid gas price, cannot contain decimals")
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(
		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
	)
	return app.New(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		a.encCfg,
		appOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(gasPrice),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		baseapp.SetIAVLCacheSize(cast.ToInt(appOpts.Get(server.FlagIAVLCacheSize))),
		baseapp.SetIAVLDisableFastNode(cast.ToBool(appOpts.Get(server.FlagDisableIAVLFastNode))),
	)
}

// appExport creates a new simapp (optionally at a given height)
func (a appCreator) appExport(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	var anApp *app.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	if height != -1 {
		anApp = app.New(logger, db, traceStore, false, map[int64]bool{},
			homePath, cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), a.encCfg, appOpts,
		)

		if err := anApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		anApp = app.New(logger, db, traceStore, true, map[int64]bool{},
			homePath, cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), a.encCfg, appOpts,
		)
	}

	return anApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
