package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/functionx/fx-core/app"

	sdk "github.com/cosmos/cosmos-sdk/types"

	appCmd "github.com/functionx/fx-core/app/cmd"
	fxtypes "github.com/functionx/fx-core/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/spf13/cobra"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/types"
	bip39 "github.com/tyler-smith/go-bip39"
)

const (
	// FlagOverwrite defines a flag to overwrite an existing genesis JSON file.
	FlagOverwrite = "overwrite"

	// FlagRecover defines a flag to initialize the private validator key from a specific seed.
	FlagRecover = "recover"

	// FlagDenom defines a flag to set the default coin denomination
	FlagDenom = "denom"
)

// initCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func initCmd(nodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [moniker]",
		Short: "Initialize private validator, p2p, genesis, application and client configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			chainID, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}

			// Get bip39 mnemonic
			var mnemonic string
			flagRecover, err := cmd.Flags().GetBool(FlagRecover)
			if err != nil {
				return err
			}
			if flagRecover {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				mnemonic, err = input.GetString("Enter your bip39 mnemonic", inBuf)
				if err != nil {
					return err
				}

				if !bip39.IsMnemonicValid(mnemonic) {
					return errors.New("invalid mnemonic")
				}
			}

			nodeID, _, err := genutil.InitializeNodeValidatorFilesFromMnemonic(config, mnemonic)
			if err != nil {
				return err
			}

			config.Moniker = args[0]

			genFile := config.GenesisFile()
			overwrite, _ := cmd.Flags().GetBool(FlagOverwrite)

			if !overwrite && tmos.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}
			flagDenom, err := cmd.Flags().GetString(FlagDenom)
			if err != nil || flagDenom == "" {
				return fmt.Errorf("invalid staking denom: %v", err)
			}
			appState, err := json.MarshalIndent(app.NewDefAppGenesisByDenom(flagDenom, clientCtx.Codec), "", " ")
			if err != nil {
				return fmt.Errorf("failed to marshall default genesis state: %s", err.Error())
			}

			genDoc := &types.GenesisDoc{}
			if _, err := os.Stat(genFile); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
				genDoc.ConsensusParams = app.CustomConsensusParams()
			} else {
				genDoc, err = types.GenesisDocFromFile(genFile)
				if err != nil {
					return fmt.Errorf("failed to read genesis doc from file: %s", err.Error())
				}
			}

			genDoc.ChainID = chainID
			genDoc.Validators = nil
			genDoc.AppState = appState
			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return fmt.Errorf("failed to export gensis file: %s", err.Error())
			}

			toPrint := appCmd.NewPrintInfo(config.Moniker, chainID, nodeID, "", appState)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

			out, err := json.MarshalIndent(toPrint, "", " ")
			if err != nil {
				return err
			}
			return clientCtx.PrintBytes(sdk.MustSortJSON(out))
		},
	}

	cmd.Flags().String(cli.HomeFlag, nodeHome, "node's home directory")
	cmd.Flags().Bool(FlagOverwrite, false, "overwrite the genesis.json file")
	cmd.Flags().Bool(FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(FlagDenom, fxtypes.DefaultDenom, "set the default coin denomination")
	cmd.Flags().StringP(cli.OutputFlag, "o", "json", "Output format (text|json)")
	return cmd
}
