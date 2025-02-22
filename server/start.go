package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	pruningtypes "github.com/cosmos/cosmos-sdk/pruning/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servergrpc "github.com/cosmos/cosmos-sdk/server/grpc"
	"github.com/cosmos/cosmos-sdk/server/rosetta"
	crgserver "github.com/cosmos/cosmos-sdk/server/rosetta/lib/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	ethmetricsexp "github.com/ethereum/go-ethereum/metrics/exp"
	"github.com/evmos/ethermint/indexer"
	ethermintserver "github.com/evmos/ethermint/server"
	ethermintconfig "github.com/evmos/ethermint/server/config"
	srvflags "github.com/evmos/ethermint/server/flags"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	abciserver "github.com/tendermint/tendermint/abci/server"
	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	tmcfg "github.com/tendermint/tendermint/config"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tendermint/rpc/client/local"
	"github.com/tendermint/tendermint/store"
	tmtypes "github.com/tendermint/tendermint/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	fxcfg "github.com/functionx/fx-core/v4/server/config"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
//
//gocyclo:ignore
func StartCmd(appCreator types.AppCreator, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Run the full node",
		Long: `Run the full node application with Tendermint in or out of process. By
default, the application will run with Tendermint in process.

Pruning options can be provided via the '--pruning' flag or alternatively with '--pruning-keep-recent', and 'pruning-interval' together.

For '--pruning' the options are as follows:

default: the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals
nothing: all historic states will be saved, nothing will be deleted (i.e. archiving node)
everything: all saved states will be deleted, storing only the current state; pruning at 10 block intervals
custom: allow pruning options to be manually specified through 'pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval'

Node halting configurations exist in the form of two flags: '--halt-height' and '--halt-time'. During
the ABCI Commit phase, the node will check if the current block height is greater than or equal to
the halt-height or if the current block time is greater than or equal to the halt-time. If so, the
node will attempt to gracefully shutdown and the block will not be committed. In addition, the node
will not be able to commit subsequent blocks.

For profiling and benchmarking purposes, CPU profiling can be enabled via the '--cpu-profile' flag
which accepts a path for the resulting pprof file.
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)

			if zeroLog, ok := serverCtx.Logger.(server.ZeroLogWrapper); ok {
				if strings.ToLower(serverCtx.Viper.GetString(flags.FlagLogFormat)) == tmcfg.LogFormatPlain {
					zeroLog.Logger = zeroLog.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "03:04:05PM"})
					serverCtx.Logger = zeroLog
				}
				filterLogTypes, _ := cmd.Flags().GetStringSlice(FlagLogFilter)
				if len(filterLogTypes) > 0 {
					serverCtx.Logger = NewFxZeroLogWrapper(zeroLog.Logger, filterLogTypes)
				}
			}

			// Bind flags to the Context's Viper so the app construction can set
			// options accordingly.
			if err := serverCtx.Viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			if _, err := server.GetPruningOptionsFromFlags(serverCtx.Viper); err != nil {
				return err
			}

			genDocFile := serverCtx.Config.GenesisFile()
			genesisBytes, err := os.ReadFile(genDocFile)
			if err != nil {
				return fmt.Errorf("couldn't read GenesisDoc file: %w", err)
			}
			expectGenesisHash := serverCtx.Viper.GetString("genesis_hash")
			actualGenesisHash := fxtypes.Sha256Hex(genesisBytes)
			if len(expectGenesisHash) != 0 && fxtypes.Sha256Hex(genesisBytes) != expectGenesisHash {
				return fmt.Errorf("--genesis_hash=%s does not match %s hash: %s", expectGenesisHash, genDocFile, actualGenesisHash)
			}
			genesisDoc, err := tmtypes.GenesisDocFromJSON(genesisBytes)
			if err != nil {
				return fmt.Errorf("error reading GenesisDoc at %s: %w", genDocFile, err)
			}
			if err = checkMainnetAndBlock(genesisDoc, serverCtx.Config); err != nil {
				return err
			}
			fxtypes.SetChainId(genesisDoc.ChainID)
			return err
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithChainID(fxtypes.ChainId())

			withTM, _ := cmd.Flags().GetBool(srvflags.WithTendermint)
			if !withTM {
				serverCtx.Logger.Info("starting ABCI without Tendermint")
				return wrapCPUProfile(serverCtx, func() error {
					return startStandAlone(serverCtx, appCreator)
				})
			}

			// amino is needed here for backwards compatibility of REST routes
			err = wrapCPUProfile(serverCtx, func() error {
				return startInProcess(serverCtx, clientCtx, appCreator)
			})
			errCode, ok := err.(server.ErrorCode)
			if !ok {
				return err
			}

			serverCtx.Logger.Debug(fmt.Sprintf("received quit signal: %d", errCode.Code))
			return nil
		},
	}

	cmd.Flags().StringSlice(FlagLogFilter, nil, `The logging filter can discard custom log type (ABCIQuery)`)
	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().Bool(srvflags.WithTendermint, true, "Run abci app embedded in-process with tendermint")
	cmd.Flags().String(srvflags.Address, "tcp://0.0.0.0:26658", "Listen address")
	cmd.Flags().String(srvflags.Transport, "socket", "Transport protocol: socket, grpc")
	cmd.Flags().String(srvflags.TraceStore, "", "Enable KVStore tracing to an output file")
	cmd.Flags().String(server.FlagMinGasPrices, "", "Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photon;0.0001stake)")
	cmd.Flags().IntSlice(server.FlagUnsafeSkipUpgrades, []int{}, "Skip a set of upgrade heights to continue the old binary")
	cmd.Flags().Uint64(server.FlagHaltHeight, 0, "Block height at which to gracefully halt the chain and shutdown the node")
	cmd.Flags().Uint64(server.FlagHaltTime, 0, "Minimum block time (in Unix seconds) at which to gracefully halt the chain and shutdown the node")
	cmd.Flags().Bool(server.FlagInterBlockCache, true, "Enable inter-block caching")
	cmd.Flags().String(srvflags.CPUProfile, "", "Enable CPU profiling and write to the provided file")
	cmd.Flags().Bool(server.FlagTrace, false, "Provide full stack traces for errors in ABCI Log")
	cmd.Flags().String(server.FlagPruning, pruningtypes.PruningOptionDefault, "Pruning strategy (default|nothing|everything|custom)")
	cmd.Flags().Uint64(server.FlagPruningKeepRecent, 0, "Number of recent heights to keep on disk (ignored if pruning is not 'custom')")
	cmd.Flags().Uint64(server.FlagPruningInterval, 0, "Height interval at which pruned heights are removed from disk (ignored if pruning is not 'custom')")
	cmd.Flags().Uint(server.FlagInvCheckPeriod, 0, "Assert registered invariants every N blocks")
	cmd.Flags().Uint64(server.FlagMinRetainBlocks, 0, "Minimum block height offset during ABCI commit to prune Tendermint blocks")
	cmd.Flags().String(srvflags.AppDBBackend, "", "The type of database for application and snapshots databases")

	cmd.Flags().Bool(srvflags.GRPCOnly, false, "Start the node in gRPC query only mode without Tendermint process")
	cmd.Flags().Bool(srvflags.GRPCEnable, true, "Define if the gRPC server should be enabled")
	cmd.Flags().String(srvflags.GRPCAddress, serverconfig.DefaultGRPCAddress, "the gRPC server address to listen on")
	cmd.Flags().Bool(srvflags.GRPCWebEnable, true, "Define if the gRPC-Web server should be enabled. (Note: gRPC must also be enabled.)")
	cmd.Flags().String(srvflags.GRPCWebAddress, serverconfig.DefaultGRPCWebAddress, "The gRPC-Web server address to listen on")

	cmd.Flags().Bool(srvflags.RPCEnable, false, "Defines if Cosmos-sdk REST server should be enabled")
	cmd.Flags().Bool(srvflags.EnabledUnsafeCors, false, "Defines if CORS should be enabled (unsafe - use it at your own risk)")

	cmd.Flags().Bool(srvflags.JSONRPCEnable, true, "Define if the JSON-RPC server should be enabled")
	cmd.Flags().StringSlice(srvflags.JSONRPCAPI, ethermintconfig.GetDefaultAPINamespaces(), "Defines a list of JSON-RPC namespaces that should be enabled")
	cmd.Flags().String(srvflags.JSONRPCAddress, ethermintconfig.DefaultJSONRPCAddress, "the JSON-RPC server address to listen on")
	cmd.Flags().String(srvflags.JSONWsAddress, ethermintconfig.DefaultJSONRPCWsAddress, "the JSON-RPC WS server address to listen on")
	cmd.Flags().Uint64(srvflags.JSONRPCGasCap, fxcfg.DefaultGasCap, "Sets a cap on gas that can be used in eth_call/estimateGas unit is aphoton (0=infinite)")
	cmd.Flags().Float64(srvflags.JSONRPCTxFeeCap, ethermintconfig.DefaultTxFeeCap, "Sets a cap on transaction fee that can be sent via the RPC APIs (1 = default 1 photon)")
	cmd.Flags().Int32(srvflags.JSONRPCFilterCap, ethermintconfig.DefaultFilterCap, "Sets the global cap for total number of filters that can be created")
	cmd.Flags().Duration(srvflags.JSONRPCEVMTimeout, ethermintconfig.DefaultEVMTimeout, "Sets a timeout used for eth_call (0=infinite)")
	cmd.Flags().Duration(srvflags.JSONRPCHTTPTimeout, ethermintconfig.DefaultHTTPTimeout, "Sets a read/write timeout for json-rpc http server (0=infinite)")
	cmd.Flags().Duration(srvflags.JSONRPCHTTPIdleTimeout, ethermintconfig.DefaultHTTPIdleTimeout, "Sets a idle timeout for json-rpc http server (0=infinite)")
	cmd.Flags().Bool(srvflags.JSONRPCAllowUnprotectedTxs, ethermintconfig.DefaultAllowUnprotectedTxs, "Allow for unprotected (non EIP155 signed) transactions to be submitted via the node's RPC when the global parameter is disabled")
	cmd.Flags().Int32(srvflags.JSONRPCLogsCap, ethermintconfig.DefaultLogsCap, "Sets the max number of results can be returned from single `eth_getLogs` query")
	cmd.Flags().Int32(srvflags.JSONRPCBlockRangeCap, ethermintconfig.DefaultBlockRangeCap, "Sets the max block range allowed for `eth_getLogs` query")
	cmd.Flags().Int(srvflags.JSONRPCMaxOpenConnections, ethermintconfig.DefaultMaxOpenConnections, "Sets the maximum number of simultaneous connections for the server listener")
	cmd.Flags().Bool(srvflags.JSONRPCEnableIndexer, false, "Enable the custom tx indexer for json-rpc")
	cmd.Flags().Bool(srvflags.JSONRPCEnableMetrics, false, "Define if EVM rpc metrics server should be enabled")

	cmd.Flags().String(srvflags.EVMTracer, ethermintconfig.DefaultEVMTracer, "the EVM tracer type to collect execution traces from the EVM transaction execution (json|struct|access_list|markdown)")
	cmd.Flags().Uint64(srvflags.EVMMaxTxGasWanted, ethermintconfig.DefaultMaxTxGasWanted, "the gas wanted for each eth tx returned in ante handler in check tx mode")

	cmd.Flags().String(srvflags.TLSCertPath, "", "the cert.pem file path for the server TLS configuration")
	cmd.Flags().String(srvflags.TLSKeyPath, "", "the key.pem file path for the server TLS configuration")

	cmd.Flags().Uint64(server.FlagStateSyncSnapshotInterval, 0, "State sync snapshot interval")
	cmd.Flags().Uint32(server.FlagStateSyncSnapshotKeepRecent, 2, "State sync snapshot to keep")

	cmd.Flags().Bool(server.FlagDisableIAVLFastNode, true, "Disable fast node for IAVL tree")

	// add support for all Tendermint-specific command line options
	tcmd.AddNodeFlags(cmd)
	crisis.AddModuleInitFlags(cmd)
	return cmd
}

func startStandAlone(ctx *server.Context, appCreator types.AppCreator) error {
	addr := ctx.Viper.GetString(srvflags.Address)
	transport := ctx.Viper.GetString(srvflags.Transport)
	home := ctx.Viper.GetString(flags.FlagHome)

	db, err := openDB(AppDBName, server.GetAppDBBackend(ctx.Viper), home)
	if err != nil {
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			ctx.Logger.Error("error closing db", "err", err.Error())
		}
	}()

	traceWriterFile := ctx.Viper.GetString(srvflags.TraceStore)
	traceWriter, err := openTraceWriter(traceWriterFile)
	if err != nil {
		return err
	}

	app := appCreator(ctx.Logger, db, traceWriter, ctx.Viper)

	config, err := fxcfg.GetConfig(ctx.Viper)
	if err != nil {
		ctx.Logger.Error("failed to get server config", "error", err.Error())
		return err
	}

	if err := config.ValidateBasic(); err != nil {
		ctx.Logger.Error("invalid server config", "error", err.Error())
		return err
	}

	_, err = startTelemetry(config.Telemetry)
	if err != nil {
		return err
	}

	svr, err := abciserver.NewServer(addr, transport, app)
	if err != nil {
		return fmt.Errorf("error creating listener: %v", err)
	}

	svr.SetLogger(ctx.Logger.With("module", "abci-server"))

	err = svr.Start()
	if err != nil {
		tmos.Exit(err.Error())
	}

	defer func() {
		if err = svr.Stop(); err != nil {
			tmos.Exit(err.Error())
		}
	}()

	// Wait for SIGINT or SIGTERM signal
	return server.WaitForQuitSignals()
}

// legacyAminoCdc is used for the legacy REST API
//
//gocyclo:ignore
func startInProcess(ctx *server.Context, clientCtx client.Context, appCreator types.AppCreator) error {
	cfg := ctx.Config
	home := cfg.RootDir
	logger := ctx.Logger

	db, err := openDB(AppDBName, server.GetAppDBBackend(ctx.Viper), home)
	if err != nil {
		logger.Error("failed to open DB", "error", err.Error())
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			ctx.Logger.Error("error closing db", "err", err.Error())
		}
	}()

	traceWriterFile := ctx.Viper.GetString(srvflags.TraceStore)
	traceWriter, err := openTraceWriter(traceWriterFile)
	if err != nil {
		logger.Error("failed to open trace writer", "error", err.Error())
		return err
	}

	config, err := fxcfg.GetConfig(ctx.Viper)
	if err != nil {
		logger.Error("failed to get server config", "error", err.Error())
		return err
	}

	if err := config.ValidateBasic(); err != nil {
		logger.Error("invalid server config", "error", err.Error())
		return err
	}

	app := appCreator(ctx.Logger, db, traceWriter, ctx.Viper)

	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		logger.Error("failed load or gen node key", "error", err.Error())
		return err
	}

	var (
		tmNode   *node.Node
		gRPCOnly = ctx.Viper.GetBool(srvflags.GRPCOnly)
	)

	if gRPCOnly {
		logger.Info("starting node in query only mode; Tendermint is disabled")
		config.GRPC.Enable = true
		config.JSONRPC.EnableIndexer = false
	} else {
		logger.Info("starting node with ABCI Tendermint in-process")

		genDocProvider := node.DefaultGenesisDocProviderFunc(cfg)
		tmNode, err = node.NewNode(
			cfg,
			pvm.LoadOrGenFilePV(cfg.PrivValidatorKeyFile(), cfg.PrivValidatorStateFile()),
			nodeKey,
			proxy.NewLocalClientCreator(app),
			genDocProvider,
			node.DefaultDBProvider,
			node.DefaultMetricsProvider(cfg.Instrumentation),
			ctx.Logger.With("module", "node"),
		)
		if err != nil {
			logger.Error("failed init node", "error", err.Error())
			return err
		}

		if err := tmNode.Start(); err != nil {
			logger.Error("failed start tendermint server", "error", err.Error())
			return err
		}

		defer func() {
			if tmNode != nil && tmNode.IsRunning() {
				_ = tmNode.Stop()
			}
		}()
	}

	// Add the tx service to the gRPC router. We only need to register this
	// service if API or gRPC or JSONRPC is enabled, and avoid doing so in the general
	// case, because it spawns a new local tendermint RPC client.
	if (config.API.Enable || config.GRPC.Enable || config.JSONRPC.Enable || config.JSONRPC.EnableIndexer) && tmNode != nil {
		clientCtx = clientCtx.WithClient(local.New(tmNode))

		app.RegisterTxService(clientCtx)
		app.RegisterTendermintService(clientCtx)

		if a, ok := app.(types.ApplicationQueryService); ok {
			a.RegisterNodeService(clientCtx)
		}
	}

	metrics, err := startTelemetry(config.Telemetry)
	if err != nil {
		return err
	}

	// Enable metrics if JSONRPC is enabled and --metrics is passed
	// Flag not added in config to avoid user enabling in config without passing in CLI
	if config.JSONRPC.Enable && ctx.Viper.GetBool(srvflags.JSONRPCEnableMetrics) {
		ethmetricsexp.Setup(config.JSONRPC.MetricsAddress)
	}

	var idxer ethermint.EVMTxIndexer
	if config.JSONRPC.EnableIndexer {
		idxDB, err := ethermintserver.OpenIndexerDB(home, server.GetAppDBBackend(ctx.Viper))
		if err != nil {
			logger.Error("failed to open evm indexer DB", "error", err.Error())
			return err
		}

		idxLogger := ctx.Logger.With("module", "indexer-evm")
		idxer = indexer.NewKVIndexer(idxDB, idxLogger, clientCtx)
		indexerService := ethermintserver.NewEVMIndexerService(idxer, clientCtx.Client)
		indexerService.SetLogger(idxLogger)

		errCh := make(chan error)
		go func() {
			if err := indexerService.Start(); err != nil {
				errCh <- err
			}
		}()

		select {
		case err := <-errCh:
			return err
		case <-time.After(types.ServerStartTime): // assume server started successfully
		}

		defer func() {
			if indexerService != nil && indexerService.IsRunning() {
				_ = indexerService.Stop()
			}
		}()
	}

	if config.API.Enable || config.JSONRPC.Enable {
		clientCtx = clientCtx.WithHomeDir(home)

		// Set `GRPCClient` to `clientCtx` to enjoy concurrent grpc query.
		// only use it if gRPC server is enabled.
		if config.GRPC.Enable {
			_, port, err := net.SplitHostPort(config.GRPC.Address)
			if err != nil {
				return errorsmod.Wrapf(err, "invalid grpc address %s", config.GRPC.Address)
			}

			maxSendMsgSize := config.GRPC.MaxSendMsgSize
			if maxSendMsgSize == 0 {
				maxSendMsgSize = serverconfig.DefaultGRPCMaxSendMsgSize
			}

			maxRecvMsgSize := config.GRPC.MaxRecvMsgSize
			if maxRecvMsgSize == 0 {
				maxRecvMsgSize = serverconfig.DefaultGRPCMaxRecvMsgSize
			}

			grpcAddress := fmt.Sprintf("127.0.0.1:%s", port)

			// If grpc is enabled, configure grpc client for grpc gateway and json-rpc.
			grpcClient, err := grpc.Dial(
				grpcAddress,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithDefaultCallOptions(
					grpc.ForceCodec(codec.NewProtoCodec(clientCtx.InterfaceRegistry).GRPCCodec()),
					grpc.MaxCallRecvMsgSize(maxRecvMsgSize),
					grpc.MaxCallSendMsgSize(maxSendMsgSize),
				),
			)
			if err != nil {
				return err
			}

			clientCtx = clientCtx.WithGRPCClient(grpcClient)
			ctx.Logger.Debug("gRPC client assigned to client context", "address", grpcAddress)
		}
	}

	var apiSrv *api.Server
	if config.API.Enable {
		apiSrv = api.New(clientCtx, ctx.Logger.With("module", "api-server"))
		app.RegisterAPIRoutes(apiSrv, config.API)

		if config.Telemetry.Enabled {
			apiSrv.SetTelemetry(metrics)
		}

		errCh := make(chan error)
		go func() {
			if err := apiSrv.Start(config.Config); err != nil {
				errCh <- err
			}
		}()

		select {
		case err := <-errCh:
			return err
		case <-time.After(types.ServerStartTime): // assume server started successfully
		}

		defer func() {
			if apiSrv != nil {
				_ = apiSrv.Close()
			}
		}()
	}

	var (
		grpcSrv    *grpc.Server
		grpcWebSrv *http.Server
	)

	if config.GRPC.Enable {
		grpcSrv, err = servergrpc.StartGRPCServer(clientCtx, app, config.GRPC)
		if err != nil {
			return err
		}
		defer grpcSrv.Stop()
		if config.GRPCWeb.Enable {
			grpcWebSrv, err = servergrpc.StartGRPCWeb(grpcSrv, config.Config)
			if err != nil {
				ctx.Logger.Error("failed to start grpc-web http server", "error", err.Error())
				return err
			}

			defer func() {
				if err := grpcWebSrv.Close(); err != nil {
					logger.Error("failed to close the grpc-web http server", "error", err.Error())
				}
			}()
		}
	}

	var (
		httpSrv     *http.Server
		httpSrvDone chan struct{}
	)

	if config.JSONRPC.Enable {
		ethClientCtx := clientCtx.WithChainID(fxtypes.ChainIdWithEIP155())

		tmEndpoint := "/websocket"
		tmRPCAddr := cfg.RPC.ListenAddress
		httpSrv, httpSrvDone, err = StartJSONRPC(ctx, ethClientCtx, tmRPCAddr, tmEndpoint, config.ToEthermintConfig(), idxer)
		if err != nil {
			return err
		}
		defer func() {
			shutdownCtx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelFn()
			if err := httpSrv.Shutdown(shutdownCtx); err != nil {
				logger.Error("HTTP server shutdown produced a warning", "error", err.Error())
			} else {
				logger.Info("HTTP server shut down, waiting 5 sec")
				select {
				case <-time.Tick(5 * time.Second):
				case <-httpSrvDone:
				}
			}
		}()
	}

	// At this point it is safe to block the process if we're in query only mode as
	// we do not need to start Rosetta or handle any Tendermint related processes.
	if gRPCOnly {
		// wait for signal capture and gracefully return
		return server.WaitForQuitSignals()
	}

	var rosettaSrv crgserver.Server
	if config.Rosetta.Enable {
		offlineMode := config.Rosetta.Offline

		// If GRPC is not enabled rosetta cannot work in online mode, so it works in
		// offline mode.
		if !config.GRPC.Enable {
			offlineMode = true
		}

		minGasPrices, err := sdk.ParseDecCoins(config.MinGasPrices)
		if err != nil {
			ctx.Logger.Error("failed to parse minimum-gas-prices", "error", err.Error())
			return err
		}

		conf := &rosetta.Config{
			Blockchain:          config.Rosetta.Blockchain,
			Network:             config.Rosetta.Network,
			TendermintRPC:       ctx.Config.RPC.ListenAddress,
			GRPCEndpoint:        config.GRPC.Address,
			Addr:                config.Rosetta.Address,
			Retries:             config.Rosetta.Retries,
			Offline:             offlineMode,
			GasToSuggest:        config.Rosetta.GasToSuggest,
			EnableFeeSuggestion: config.Rosetta.EnableFeeSuggestion,
			GasPrices:           minGasPrices.Sort(),
			Codec:               clientCtx.Codec.(*codec.ProtoCodec),
			InterfaceRegistry:   clientCtx.InterfaceRegistry,
		}

		rosettaSrv, err = rosetta.ServerFromConfig(conf)
		if err != nil {
			return err
		}

		errCh := make(chan error)
		go func() {
			if err := rosettaSrv.Start(); err != nil {
				errCh <- err
			}
		}()

		select {
		case err := <-errCh:
			return err
		case <-time.After(types.ServerStartTime): // assume server started successfully
		}
	}
	// Wait for SIGINT or SIGTERM signal
	return server.WaitForQuitSignals()
}

func openTraceWriter(traceWriterFile string) (w io.Writer, err error) {
	if traceWriterFile == "" {
		return
	}

	filePath := filepath.Clean(traceWriterFile)
	return os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0o600,
	)
}

func startTelemetry(cfg telemetry.Config) (*telemetry.Metrics, error) {
	if !cfg.Enabled {
		return nil, nil
	}
	return telemetry.New(cfg)
}

// wrapCPUProfile runs callback in a goroutine, then wait for quit signals.
func wrapCPUProfile(ctx *server.Context, callback func() error) error {
	if cpuProfile := ctx.Viper.GetString(srvflags.CPUProfile); cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			return err
		}

		ctx.Logger.Info("starting CPU profiler", "profile", cpuProfile)
		if err := pprof.StartCPUProfile(f); err != nil {
			return err
		}

		defer func() {
			ctx.Logger.Info("stopping CPU profiler", "profile", cpuProfile)
			pprof.StopCPUProfile()
			if err := f.Close(); err != nil {
				ctx.Logger.Info("failed to close cpu-profile file", "profile", cpuProfile, "err", err.Error())
			}
		}()
	}

	errCh := make(chan error)
	go func() {
		errCh <- callback()
	}()

	select {
	case err := <-errCh:
		return err

	case <-time.After(types.ServerStartTime):
	}

	return server.WaitForQuitSignals()
}

func checkMainnetAndBlock(genesisDoc *tmtypes.GenesisDoc, config *tmcfg.Config) error {
	if genesisDoc.InitialHeight > 1 || genesisDoc.ChainID != fxtypes.MainnetChainId || config.StateSync.Enable {
		return nil
	}
	genesisTime, err := time.Parse("2006-01-02T15:04:05Z", "2021-07-05T04:00:00Z")
	if err != nil {
		return err
	}
	blockStoreDB, err := node.DefaultDBProvider(&node.DBContext{ID: "blockstore", Config: config})
	if err != nil {
		return err
	}
	defer blockStoreDB.Close()
	blockStore := store.NewBlockStore(blockStoreDB)
	if genesisDoc.GenesisTime.Equal(genesisTime) {
		genesisBytes, _ := tmjson.Marshal(genesisDoc)
		if fxtypes.Sha256Hex(genesisBytes) != fxtypes.MainnetGenesisHash {
			return nil
		}
		if blockStore.Height() < fxtypes.MainnetBlockHeightV2 {
			return errors.New("invalid version: The current block height is less than the fxv2 upgrade height(8_756_000), " +
				"sync block from scratch please use use fxcored v1.x.x")
		}
		if blockStore.Height() < fxtypes.MainnetBlockHeightV3 {
			return errors.New("invalid version: The current block height is less than the fxv3 upgrade height(8_756_000)," +
				" please use the v2.x.x version to synchronize the block or download the latest snapshot")
		}
		return errors.New("invalid version: The current version is not released, please use the corresponding version")
	}
	return nil
}
