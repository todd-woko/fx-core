package crosschain_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/v6/modules/core/23-commitment/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	localhosttypes "github.com/cosmos/ibc-go/v6/modules/light-clients/09-localhost/types"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/functionx/fx-core/v4/app"
	fxserverconfig "github.com/functionx/fx-core/v4/server/config"
	"github.com/functionx/fx-core/v4/testutil/helpers"
	fxtypes "github.com/functionx/fx-core/v4/types"
	crosschaintypes "github.com/functionx/fx-core/v4/x/crosschain/types"
	"github.com/functionx/fx-core/v4/x/erc20/types"
)

type PrecompileTestSuite struct {
	suite.Suite
	ctx        sdk.Context
	app        *app.App
	signer     *helpers.Signer
	crosschain common.Address
}

func TestPrecompileTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

// Test helpers
func (suite *PrecompileTestSuite) SetupTest() {
	// account key
	priv, err := ethsecp256k1.GenerateKey()
	require.NoError(suite.T(), err)
	suite.signer = helpers.NewSigner(priv)

	set, accs, balances := helpers.GenerateGenesisValidator(tmrand.Intn(10)+1, nil)
	suite.app = helpers.SetupWithGenesisValSet(suite.T(), set, accs, balances...)

	suite.ctx = suite.app.NewContext(false, tmproto.Header{
		Height:          suite.app.LastBlockHeight(),
		ChainID:         fxtypes.ChainId(),
		ProposerAddress: set.Proposer.Address,
		Time:            time.Now().UTC(),
	})
	suite.ctx = suite.ctx.WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin(fxtypes.DefaultDenom, sdkmath.OneInt())))
	suite.ctx = suite.ctx.WithBlockGasMeter(sdk.NewGasMeter(1e18))

	helpers.AddTestAddr(suite.app, suite.ctx, suite.signer.AccAddress(), sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, sdkmath.NewInt(10000).Mul(sdkmath.NewInt(1e18)))))

	stakingContract, err := suite.app.EvmKeeper.DeployContract(suite.ctx, suite.signer.Address(), fxtypes.MustABIJson(CrosschainTestABI), fxtypes.MustDecodeHex(CrosschainTestBin))
	suite.Require().NoError(err)
	suite.crosschain = stakingContract
}

func (suite *PrecompileTestSuite) PackEthereumTx(signer *helpers.Signer, contract common.Address, amount *big.Int, data []byte) (*evmtypes.MsgEthereumTx, error) {
	fromAddr := signer.Address()
	value := hexutil.Big(*amount)
	args, err := json.Marshal(&evmtypes.TransactionArgs{To: &contract, From: &fromAddr, Data: (*hexutil.Bytes)(&data), Value: &value})
	suite.Require().NoError(err)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
	res, err := evmtypes.NewQueryClient(queryHelper).EstimateGas(sdk.WrapSDKContext(suite.ctx),
		&evmtypes.EthCallRequest{
			Args:    args,
			GasCap:  fxserverconfig.DefaultGasCap,
			ChainId: suite.app.EvmKeeper.ChainID().Int64(),
		},
	)
	if err != nil {
		return nil, err
	}

	ethTx := evmtypes.NewTx(
		fxtypes.EIP155ChainID(),
		suite.app.EvmKeeper.GetNonce(suite.ctx, signer.Address()),
		&contract,
		amount,
		res.Gas,
		nil,
		suite.app.FeeMarketKeeper.GetBaseFee(suite.ctx),
		big.NewInt(1),
		data,
		nil,
	)
	ethTx.From = signer.Address().Hex()
	err = ethTx.Sign(ethtypes.LatestSignerForChainID(fxtypes.EIP155ChainID()), signer)
	return ethTx, err
}

func (suite *PrecompileTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.app.EndBlock(abci.RequestEndBlock{
		Height: header.Height,
	})
	suite.app.Commit()
	// after commit ctx header
	header.Height += 1

	// begin block
	header.Time = time.Now().UTC()
	header.Height += 1
	suite.app.BeginBlock(abci.RequestBeginBlock{
		Header: header,
	})
	suite.ctx = suite.ctx.WithBlockHeight(header.Height)
}

func (suite *PrecompileTestSuite) RandSigner() *helpers.Signer {
	privKey := helpers.NewEthPrivKey()
	// helpers.AddTestAddr(suite.app, suite.ctx, privKey.PubKey().Address().Bytes(), sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, sdkmath.NewInt(1000).Mul(sdkmath.NewInt(1e18)))))
	signer := helpers.NewSigner(privKey)
	suite.app.AccountKeeper.SetAccount(suite.ctx, suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, signer.AccAddress()))
	return signer
}

func (suite *PrecompileTestSuite) MintFeeCollector(coins sdk.Coins) {
	err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, types.ModuleName, authtypes.FeeCollectorName, coins)
	suite.Require().NoError(err)
}

func (suite *PrecompileTestSuite) DeployContract(from common.Address) (common.Address, error) {
	contract, err := suite.app.Erc20Keeper.DeployUpgradableToken(suite.ctx, suite.app.Erc20Keeper.ModuleAddress(), "Test token", "TEST", 18)
	suite.Require().NoError(err)

	_, err = suite.app.EvmKeeper.ApplyContract(suite.ctx, suite.app.Erc20Keeper.ModuleAddress(), contract, fxtypes.GetERC20().ABI, "transferOwnership", from)
	suite.Require().NoError(err)
	return contract, nil
}

func (suite *PrecompileTestSuite) DeployFXRelayToken() (types.TokenPair, banktypes.Metadata) {
	fxToken := fxtypes.GetFXMetaData(fxtypes.DefaultDenom)

	pair, err := suite.app.Erc20Keeper.RegisterNativeCoin(suite.ctx, fxToken)
	suite.Require().NoError(err)
	return *pair, fxToken
}

func (suite *PrecompileTestSuite) CrossChainKeepers() map[string]CrossChainKeeper {
	value := reflect.ValueOf(suite.app.CrossChainKeepers)
	keepers := make(map[string]CrossChainKeeper)
	for i := 0; i < value.NumField(); i++ {
		res := value.Field(i).MethodByName("GetGravityID").Call([]reflect.Value{reflect.ValueOf(suite.ctx)})
		gravityID := res[0].String()
		chainName := strings.TrimSuffix(strings.TrimPrefix(gravityID, "fx-"), "-bridge")
		cck := value.Field(i).Interface().(CrossChainKeeper)
		if chainName == "bridge-eth" {
			// keepers["gravity"] = cck
			keepers["eth"] = cck
		} else {
			keepers[chainName] = cck
		}
	}
	return keepers
}

func (suite *PrecompileTestSuite) GenerateCrossChainDenoms(addDenoms ...string) Metadata {
	keepers := suite.CrossChainKeepers()
	modules := make([]string, 0, len(keepers))
	for m := range keepers {
		modules = append(modules, m)
	}
	count := tmrand.Intn(len(modules)-1) + 1

	denoms := make([]string, len(modules))
	denomModules := make([]string, len(modules))
	for index, m := range modules {
		address := helpers.GenerateAddressByModule(m)

		denom := fmt.Sprintf("%s%s", m, address)
		denoms[index] = denom
		denomModules[index] = m

		k := keepers[m]
		k.AddBridgeToken(suite.ctx, address, fmt.Sprintf("%s%s", m, address))
	}
	if count >= len(modules) {
		count = len(modules) - 1
	}
	metadata := fxtypes.GetCrossChainMetadata("Test Token", helpers.NewRandSymbol(), 18, append(denoms[:count], addDenoms...)...)
	return Metadata{metadata: metadata, modules: denomModules[:count], notModules: denomModules[count:]}
}

func (suite *PrecompileTestSuite) MintLockNativeTokenToModule(md banktypes.Metadata, amt sdkmath.Int) sdk.Coin {
	generateAddress := helpers.GenerateAddress()

	count := 1
	if len(md.DenomUnits) > 0 && len(md.DenomUnits[0].Aliases) > 0 {
		// add alias to erc20 module
		for _, alias := range md.DenomUnits[0].Aliases {
			// add alias for erc20 module
			coins := sdk.NewCoins(sdk.NewCoin(alias, amt))
			helpers.AddTestAddr(suite.app, suite.ctx, generateAddress.Bytes(), coins)
			err := suite.app.BankKeeper.SendCoinsFromAccountToModule(suite.ctx, generateAddress.Bytes(), types.ModuleName, coins)
			suite.Require().NoError(err)
		}
		count = len(md.DenomUnits[0].Aliases)
	}

	// add denom to erc20 module
	coin := sdk.NewCoin(md.Base, amt.Mul(sdkmath.NewInt(int64(count))))
	helpers.AddTestAddr(suite.app, suite.ctx, generateAddress.Bytes(), sdk.NewCoins(coin))
	err := suite.app.BankKeeper.SendCoinsFromAccountToModule(suite.ctx, generateAddress.Bytes(), types.ModuleName, sdk.NewCoins(coin))
	suite.Require().NoError(err)

	return coin
}

func (suite *PrecompileTestSuite) BalanceOf(contract, account common.Address) *big.Int {
	var balanceRes struct{ Value *big.Int }
	err := suite.app.EvmKeeper.QueryContract(suite.ctx, account, contract, fxtypes.GetERC20().ABI, "balanceOf", &balanceRes, account)
	suite.Require().NoError(err)
	return balanceRes.Value
}

func (suite *PrecompileTestSuite) MintERC20Token(signer *helpers.Signer, contractAddr, to common.Address, amount *big.Int) *evmtypes.MsgEthereumTxResponse {
	erc20 := fxtypes.GetERC20()
	transferData, err := erc20.ABI.Pack("mint", to, amount)
	suite.Require().NoError(err)
	return suite.sendEvmTx(signer, contractAddr, transferData)
}

func (suite *PrecompileTestSuite) ModuleMintERC20Token(contractAddr, to common.Address, amount *big.Int) {
	erc20 := fxtypes.GetERC20()
	rsp, err := suite.app.EvmKeeper.ApplyContract(suite.ctx, suite.app.Erc20Keeper.ModuleAddress(), contractAddr, erc20.ABI, "mint", to, amount)
	suite.Require().NoError(err)
	suite.Require().Empty(rsp.VmError)
}

func (suite *PrecompileTestSuite) TransferERC20Token(signer *helpers.Signer, contractAddr, to common.Address, amount *big.Int) *evmtypes.MsgEthereumTxResponse {
	erc20 := fxtypes.GetERC20()
	transferData, err := erc20.ABI.Pack("transfer", to, amount)
	suite.Require().NoError(err)
	return suite.sendEvmTx(signer, contractAddr, transferData)
}

func (suite *PrecompileTestSuite) ERC20Approve(signer *helpers.Signer, contractAddr, to common.Address, amount *big.Int) *evmtypes.MsgEthereumTxResponse {
	erc20 := fxtypes.GetERC20()
	transferData, err := erc20.ABI.Pack("approve", to, amount)
	suite.Require().NoError(err)
	return suite.sendEvmTx(signer, contractAddr, transferData)
}

func (suite *PrecompileTestSuite) ERC20Allowance(contract, owner, spender common.Address) *big.Int {
	var allowanceRes struct{ Value *big.Int }
	err := suite.app.EvmKeeper.QueryContract(suite.ctx, owner, contract, fxtypes.GetERC20().ABI, "allowance", &allowanceRes, owner, spender)
	suite.Require().NoError(err)
	return allowanceRes.Value
}

func (suite *PrecompileTestSuite) TransferERC20TokenToModule(signer *helpers.Signer, contractAddr common.Address, amount *big.Int) *evmtypes.MsgEthereumTxResponse {
	erc20 := fxtypes.GetERC20()
	moduleAddress := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
	transferData, err := erc20.ABI.Pack("transfer", common.BytesToAddress(moduleAddress.Bytes()), amount)
	suite.Require().NoError(err)
	return suite.sendEvmTx(signer, contractAddr, transferData)
}

func (suite *PrecompileTestSuite) TransferERC20TokenToModuleWithoutHook(contractAddr, from common.Address, amount *big.Int) {
	erc20 := fxtypes.GetERC20()
	moduleAddress := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
	_, err := suite.app.EvmKeeper.ApplyContract(suite.ctx, from, contractAddr, erc20.ABI, "transfer", common.BytesToAddress(moduleAddress.Bytes()), amount)
	suite.Require().NoError(err)
}

func (suite *PrecompileTestSuite) RandPrefixAndAddress() (string, string) {
	if tmrand.Intn(10)%2 == 0 {
		return "0x", helpers.GenerateAddress().Hex()
	}
	prefix := strings.ToLower(tmrand.Str(5))
	accAddress, err := bech32.ConvertAndEncode(prefix, suite.RandSigner().AccAddress().Bytes())
	suite.Require().NoError(err)
	return prefix, accAddress
}

func (suite *PrecompileTestSuite) RandTransferChannel() (portID, channelID string) {
	portID = "transfer"
	channelID = fmt.Sprintf("channel-%d", tmrand.Intn(100))
	connectionID := connectiontypes.FormatConnectionIdentifier(uint64(tmrand.Intn(100)))
	clientID := clienttypes.FormatClientIdentifier(exported.Localhost, uint64(tmrand.Intn(100)))

	revision := clienttypes.ParseChainID(suite.ctx.ChainID())
	localHostClient := localhosttypes.NewClientState(
		suite.ctx.ChainID(), clienttypes.NewHeight(revision, uint64(suite.ctx.BlockHeight())),
	)
	suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, clientID, localHostClient)

	prevConsState := &ibctmtypes.ConsensusState{
		Timestamp:          suite.ctx.BlockTime(),
		NextValidatorsHash: suite.ctx.BlockHeader().NextValidatorsHash,
	}
	height := clienttypes.NewHeight(0, uint64(suite.ctx.BlockHeight()))
	suite.app.IBCKeeper.ClientKeeper.SetClientConsensusState(suite.ctx, clientID, height, prevConsState)

	channelCapability, err := suite.app.ScopedIBCKeeper.NewCapability(suite.ctx, host.ChannelCapabilityPath(portID, channelID))
	suite.Require().NoError(err)
	err = suite.app.ScopedTransferKeeper.ClaimCapability(suite.ctx, capabilitytypes.NewCapability(channelCapability.Index), host.ChannelCapabilityPath(portID, channelID))
	suite.Require().NoError(err)

	connectionEnd := connectiontypes.NewConnectionEnd(connectiontypes.OPEN, clientID, connectiontypes.Counterparty{ClientId: "clientId", ConnectionId: "connection-1", Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, []*connectiontypes.Version{ibctesting.ConnectionVersion}, 500)
	suite.app.IBCKeeper.ConnectionKeeper.SetConnection(suite.ctx, connectionID, connectionEnd)

	channel := channeltypes.NewChannel(channeltypes.OPEN, channeltypes.ORDERED, channeltypes.NewCounterparty(portID, channelID), []string{connectionID}, ibctesting.DefaultChannelVersion)
	suite.app.IBCKeeper.ChannelKeeper.SetChannel(suite.ctx, portID, channelID, channel)
	suite.app.IBCKeeper.ChannelKeeper.SetNextSequenceSend(suite.ctx, portID, channelID, uint64(tmrand.Intn(10000)))

	return portID, channelID
}

func (suite *PrecompileTestSuite) AddIBCToken(portID, channelID string) string {
	denomTrace := ibctransfertypes.DenomTrace{
		Path:      fmt.Sprintf("%s/%s", portID, channelID),
		BaseDenom: "test",
	}
	suite.app.IBCTransferKeeper.SetDenomTrace(suite.ctx, denomTrace)
	return denomTrace.IBCDenom()
}

func (suite *PrecompileTestSuite) sendEvmTx(signer *helpers.Signer, contractAddr common.Address, data []byte) *evmtypes.MsgEthereumTxResponse {
	from := signer.Address()

	args, err := json.Marshal(&evmtypes.TransactionArgs{To: &contractAddr, From: &from, Data: (*hexutil.Bytes)(&data)})
	suite.Require().NoError(err)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
	res, err := evmtypes.NewQueryClient(queryHelper).EstimateGas(sdk.WrapSDKContext(suite.ctx),
		&evmtypes.EthCallRequest{
			Args:    args,
			GasCap:  fxserverconfig.DefaultGasCap,
			ChainId: suite.app.EvmKeeper.ChainID().Int64(),
		},
	)
	suite.Require().NoError(err)

	// Mint the max gas to the FeeCollector to ensure balance in case of refund
	// suite.MintFeeCollector(sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, sdkmath.NewInt(suite.app.FeeMarketKeeper.GetBaseFee(suite.ctx).Int64()*int64(res.Gas)))))

	msg := ethtypes.NewMessage(
		signer.Address(),
		&contractAddr,
		suite.app.EvmKeeper.GetNonce(suite.ctx, signer.Address()),
		big.NewInt(0),
		res.Gas,
		suite.app.FeeMarketKeeper.GetBaseFee(suite.ctx),
		nil,
		nil,
		data,
		nil,
		true,
	)

	rsp, err := suite.app.EvmKeeper.ApplyMessage(suite.ctx, msg, nil, true)
	suite.Require().NoError(err)
	suite.Require().False(rsp.Failed(), rsp.VmError)
	return rsp
}

type Metadata struct {
	metadata   banktypes.Metadata
	modules    []string
	notModules []string
}

func (m Metadata) RandModule() string {
	return m.modules[tmrand.Intn(len(m.modules))]
}

func (m Metadata) GetModules() []string {
	return m.modules
}

func (m Metadata) GetDenom(moduleName string) string {
	for _, denom := range m.metadata.DenomUnits[0].Aliases {
		if strings.HasPrefix(denom, moduleName) {
			return denom
		}
	}
	return ""
}

func (m Metadata) GetMetadata() banktypes.Metadata {
	return m.metadata
}

type CrossChainKeeper interface {
	AddBridgeToken(ctx sdk.Context, token, denom string)
	GetDenomByBridgeToken(ctx sdk.Context, denom string) *crosschaintypes.BridgeToken
	SetIbcDenomTrace(ctx sdk.Context, token, channelIBC string) (string, error)
	GetPendingSendToExternal(c context.Context, req *crosschaintypes.QueryPendingSendToExternalRequest) (*crosschaintypes.QueryPendingSendToExternalResponse, error)
	AddToOutgoingPool(ctx sdk.Context, sender sdk.AccAddress, receiver string, amount sdk.Coin, fee sdk.Coin) (uint64, error)
}

const testJsonABI = `
[
    {
        "type":"function",
        "name":"fip20CrossChainV2",
        "inputs":[
            {
                "name":"sender",
                "type":"address"
            },
            {
                "name":"refunder",
                "type":"address"
            },
            {
                "name":"receipt",
                "type":"string"
            },
            {
                "name":"amount",
                "type":"uint256"
            },
            {
                "name":"fee",
                "type":"uint256"
            },
            {
                "name":"target",
                "type":"bytes32"
            },
            {
                "name":"memo",
                "type":"string"
            }
        ],
        "outputs":[
            {
                "name":"result",
                "type":"bool"
            }
        ],
        "payable":false,
        "stateMutability":"nonpayable"
    },
	{
        "type":"function",
        "name":"fip20CrossChain",
        "inputs":[
            {
                "name":"sender",
                "type":"address"
            },
            {
                "name":"refunder",
                "type":"address"
            },
            {
                "name":"receipt",
                "type":"string"
            },
            {
                "name":"amount",
                "type":"uint256"
            },
            {
                "name":"fee",
                "type":"uint256"
            },
            {
                "name":"memo",
                "type":"string"
            }
        ],
        "outputs":[
            {
                "name":"result",
                "type":"bool"
            }
        ],
        "payable":false,
        "stateMutability":"nonpayable"
    },
    {
        "type":"function",
        "name":"fip20CancelSendToExternal",
        "inputs":[
            {
                "name":"chain",
                "type":"string"
            },
            {
                "name":"txID",
                "type":"uint256"
            },
            {
                "name":"refunder",
                "type":"address"
            }
        ],
        "outputs":[
            {
                "name":"result",
                "type":"bool"
            }
        ],
        "payable":false,
        "stateMutability":"nonpayable"
    }
]`
