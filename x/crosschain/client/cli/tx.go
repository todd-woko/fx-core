package cli

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govv1betal "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	gethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	troncommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/spf13/cobra"

	"github.com/functionx/fx-core/v4/x/crosschain/types"
)

func GetTxCmd(subCmd ...*cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Crosschain transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(subCmd...)
	return cmd
}

func GetTxSubCmds(chainName string) []*cobra.Command {
	cmds := []*cobra.Command{
		CmdUpdateChainOraclesProposal(chainName),

		// set bridger address
		CmdCreateOracleBridger(chainName),
		// add oracle stake
		CmdAddOracleDelegate(chainName),
		// send to external chain
		CmdSendToExternal(chainName),
		CmdCancelSendToExternal(chainName),
		CmdIncreaseBridgeFee(chainName),
		CmdRequestBatch(chainName),

		// oracle consensus confirm
		CmdOracleSetConfirm(chainName),
		CmdRequestBatchConfirm(chainName),
	}
	for _, command := range cmds {
		flags.AddTxFlagsToCmd(command)
	}
	return cmds
}

// CmdUpdateChainOraclesProposal
// nolint:staticcheck
func CmdUpdateChainOraclesProposal(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-crosschain-oracles [oracles]",
		Short: fmt.Sprintf("update %s oracles", chainName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			oracles := strings.Split(args[0], ",")
			for i, oracle := range oracles {
				oracleAddr, err := sdk.AccAddressFromBech32(oracle)
				if err != nil {
					return err
				}
				oracles[i] = oracleAddr.String()
			}
			proposal := &types.UpdateChainOraclesProposal{
				Title:       title,
				Description: description,
				Oracles:     oracles,
				ChainName:   chainName,
			}
			fromAddress := cliCtx.GetFromAddress()
			msg, err := govv1betal.NewMsgSubmitProposal(proposal, deposit, fromAddress)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)
	_ = cmd.MarkFlagRequired(cli.FlagDeposit)
	return cmd
}

func CmdCreateOracleBridger(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-oracle-bridger [validator-address] [bridger-address] [external-address] [delegate-amount]",
		Short: "Allows oracle to delegate their voting responsibilities to a given key.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			bridgerAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			externalAddress, err := getContractAddr(args[2])
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			msg := types.MsgBondedOracle{
				OracleAddress:    cliCtx.GetFromAddress().String(),
				BridgerAddress:   bridgerAddr.String(),
				ExternalAddress:  externalAddress,
				ValidatorAddress: valAddr.String(),
				DelegateAmount:   amount,
				ChainName:        chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	return cmd
}

func CmdAddOracleDelegate(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-oracle-delegate [delegate-amount]",
		Short: "Allows oracle add delegate.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			msg := types.MsgAddDelegate{
				OracleAddress: cliCtx.GetFromAddress().String(),
				Amount:        amount,
				ChainName:     chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	return cmd
}

func CmdSendToExternal(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-to-external [external-dest] [amount] [bridge-fee]",
		Short: "Adds a new entry to the transaction pool to withdraw an amount from the bridge contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			externalDestAddr, err := getContractAddr(args[0])
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "amount")
			}
			bridgeFee, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return errorsmod.Wrap(err, "bridge fee")
			}

			msg := types.MsgSendToExternal{
				Sender:    cliCtx.GetFromAddress().String(),
				Dest:      externalDestAddr,
				Amount:    amount,
				BridgeFee: bridgeFee,
				ChainName: chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	return cmd
}

func CmdCancelSendToExternal(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-send-to-external [tx-ID]",
		Short: "Cancel transaction send to external",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			msg := &types.MsgCancelSendToExternal{
				TransactionId: txId,
				Sender:        cliCtx.GetFromAddress().String(),
				ChainName:     chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func CmdIncreaseBridgeFee(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "increase-bridge-fee [tx-ID] [add-bridge-fee]",
		Short: "Increase bridge fee for send to external transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			addBridgeFee, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "add bridge fee")
			}

			msg := &types.MsgIncreaseBridgeFee{
				ChainName:     chainName,
				TransactionId: txId,
				Sender:        cliCtx.GetFromAddress().String(),
				AddBridgeFee:  addBridgeFee,
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func CmdRequestBatch(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build-batch [token-denom] [minimum-fee] [base-fee] [external-fee-receive]",
		Short: "Build a new batch on the fx side for pooled withdrawal transactions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denom := args[0]

			minimumFee, ok := sdkmath.NewIntFromString(args[1])
			if !ok || minimumFee.IsNegative() {
				return fmt.Errorf("miniumu fee is valid, %v", args[1])
			}
			baseFee := sdkmath.ZeroInt()
			if len(args[2]) > 0 {
				baseFee, ok = sdkmath.NewIntFromString(args[2])
				if !ok {
					return fmt.Errorf("invalid base fee: %v", args[2])
				}
			}
			feeReceive := args[3]
			if strings.HasPrefix(feeReceive, "0x") {
				if !gethcommon.IsHexAddress(feeReceive) {
					return fmt.Errorf("invalid feeReceive address: %v", feeReceive)
				}
				feeReceive = gethcommon.HexToAddress(feeReceive).Hex()
			}
			msg := &types.MsgRequestBatch{
				Sender:     clientCtx.GetFromAddress().String(),
				Denom:      denom,
				MinimumFee: minimumFee,
				FeeReceive: feeReceive,
				ChainName:  chainName,
				BaseFee:    baseFee,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func CmdRequestBatchConfirm(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-batch-confirm [contract-address] [nonce] [private-key]",
		Short: "Send batch confirm msg",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()

			tokenContract, err := getContractAddr(args[0])
			if err != nil {
				return err
			}
			nonce, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			privateKey, err := recoveryPrivateKeyByKeystore(args[2])
			if err != nil {
				return err
			}
			externalAddress := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

			queryClient := types.NewQueryClient(clientCtx)
			batchRequestByNonceResp, err := queryClient.BatchRequestByNonce(cmd.Context(), &types.QueryBatchRequestByNonceRequest{
				Nonce:         nonce,
				TokenContract: tokenContract,
				ChainName:     chainName,
			})
			if err != nil {
				return err
			}
			if batchRequestByNonceResp.Batch == nil {
				return fmt.Errorf("not found batch request by nonce, tokenContract: %v, nonce: %v", tokenContract, nonce)
			}
			// Determine whether it has been confirmed
			batchConfirmResp, err := queryClient.BatchConfirm(cmd.Context(), &types.QueryBatchConfirmRequest{
				Nonce:          nonce,
				TokenContract:  tokenContract,
				BridgerAddress: fromAddress.String(),
				ChainName:      chainName,
			})
			if err != nil {
				return err
			}
			if batchConfirmResp.GetConfirm() != nil {
				confirm := batchConfirmResp.GetConfirm()
				return clientCtx.PrintProto(confirm)
			}
			paramsResp, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{
				ChainName: chainName,
			})
			if err != nil {
				return err
			}
			checkpoint, err := batchRequestByNonceResp.GetBatch().GetCheckpoint(paramsResp.Params.GetGravityId())
			if err != nil {
				return err
			}
			signature, err := types.NewEthereumSignature(checkpoint, privateKey)
			if err != nil {
				return err
			}
			msg := &types.MsgConfirmBatch{
				Nonce:           nonce,
				TokenContract:   tokenContract,
				ExternalAddress: externalAddress.String(),
				BridgerAddress:  fromAddress.String(),
				Signature:       hex.EncodeToString(signature),
				ChainName:       chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func CmdOracleSetConfirm(chainName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-set-confirm [nonce] [private-key]",
		Short: "Send oracle-set confirm msg",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()

			nonce, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			privateKey, err := recoveryPrivateKeyByKeystore(args[1])
			if err != nil {
				return err
			}
			externalAddress := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

			queryClient := types.NewQueryClient(clientCtx)
			oracleSetRequestResp, err := queryClient.OracleSetRequest(cmd.Context(), &types.QueryOracleSetRequestRequest{
				Nonce: nonce, ChainName: chainName,
			})
			if err != nil {
				return err
			}
			// Determine whether it has been confirmed
			oracleSetConfirmResp, err := queryClient.OracleSetConfirm(cmd.Context(), &types.QueryOracleSetConfirmRequest{
				Nonce:          nonce,
				BridgerAddress: fromAddress.String(),
				ChainName:      chainName,
			})
			if err != nil {
				return err
			}
			if oracleSetConfirmResp.GetConfirm() != nil {
				confirm := oracleSetConfirmResp.GetConfirm()
				return clientCtx.PrintProto(confirm)
			}
			paramsResp, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{
				ChainName: chainName,
			})
			if err != nil {
				return err
			}
			checkpoint, err := oracleSetRequestResp.GetOracleSet().GetCheckpoint(paramsResp.Params.GetGravityId())
			if err != nil {
				return err
			}
			signature, err := types.NewEthereumSignature(checkpoint, privateKey)
			if err != nil {
				return err
			}
			msg := &types.MsgOracleSetConfirm{
				Nonce:           nonce,
				BridgerAddress:  fromAddress.String(),
				ExternalAddress: externalAddress.String(),
				Signature:       hex.EncodeToString(signature),
				ChainName:       chainName,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func recoveryPrivateKeyByKeystore(privateKey string) (*ecdsa.PrivateKey, error) {
	var ethPrivateKey *ecdsa.PrivateKey
	if _, err := os.Stat(privateKey); err == nil {
		file, err := os.ReadFile(privateKey)
		if err != nil {
			return nil, err
		}
		stdinReader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return nil, err
		}
		password := strings.TrimSpace(stdinReader)
		key, err := keystore.DecryptKey(file, password)
		if err != nil {
			return nil, err
		}
		ethPrivateKey = key.PrivateKey
	} else {
		key, err := ethcrypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, fmt.Errorf("invalid eth private key: %s", err.Error())
		}
		ethPrivateKey = key
	}
	return ethPrivateKey, nil
}

func getContractAddr(addr string) (string, error) {
	if strings.HasPrefix(addr, "0x") {
		if !gethcommon.IsHexAddress(addr) {
			return "", fmt.Errorf("invalid address: %s", addr)
		}
		addr = gethcommon.HexToAddress(addr).Hex()
	} else {
		tronAddr, err := troncommon.DecodeCheck(addr)
		if err != nil {
			return "", fmt.Errorf("doesn't pass format validation: %s", addr)
		}
		addr = troncommon.EncodeCheck(tronAddr)
	}
	return addr, nil
}
