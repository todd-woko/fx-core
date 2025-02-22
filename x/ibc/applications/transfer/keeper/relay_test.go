package keeper_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v6/testing"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/functionx/fx-core/v4/testutil/helpers"
	fxtypes "github.com/functionx/fx-core/v4/types"
	avalanchetypes "github.com/functionx/fx-core/v4/x/avalanche/types"
	bsctypes "github.com/functionx/fx-core/v4/x/bsc/types"
	erc20types "github.com/functionx/fx-core/v4/x/erc20/types"
	ethtypes "github.com/functionx/fx-core/v4/x/eth/types"
	fxtransfer "github.com/functionx/fx-core/v4/x/ibc/applications/transfer"
	fxtransfertypes "github.com/functionx/fx-core/v4/x/ibc/applications/transfer/types"
	fxibctesting "github.com/functionx/fx-core/v4/x/ibc/testing"
	polygontypes "github.com/functionx/fx-core/v4/x/polygon/types"
	trontypes "github.com/functionx/fx-core/v4/x/tron/types"
)

func (suite *KeeperTestSuite) TestSendTransfer() {
	var (
		coin          sdk.Coin
		path          *ibctesting.Path
		sender        sdk.AccAddress
		timeoutHeight clienttypes.Height
		memo          string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"successful transfer with native token",
			func() {}, true,
		},
		{
			"successful transfer from source chain with memo",
			func() {
				memo = "memo"
			}, true,
		},
		{
			"successful transfer with IBC token",
			func() {
				// send IBC token back to chainB
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount)
			}, true,
		},
		{
			"successful transfer with IBC token and memo",
			func() {
				// send IBC token back to chainB
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount)
				memo = "memo"
			}, true,
		},
		{
			"source channel not found",
			func() {
				// channel references wrong ID
				path.EndpointA.ChannelID = ibctesting.InvalidID
			}, false,
		},
		{
			"transfer failed - sender account is blocked",
			func() {
				sender = suite.GetApp(suite.chainA.App).AccountKeeper.GetModuleAddress(types.ModuleName)
			}, false,
		},
		{
			"send coin failed",
			func() {
				coin = sdk.NewCoin("randomdenom", sdk.NewInt(100))
			}, false,
		},
		{
			"failed to parse coin denom",
			func() {
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, "randomdenom", coin.Amount)
			}, false,
		},
		{
			"send from module account failed, insufficient balance",
			func() {
				coin = types.GetTransferCoin(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coin.Denom, coin.Amount.Add(sdk.NewInt(1)))
			}, false,
		},
		{
			"channel capability not found",
			func() {
				capability := suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)

				// Release channel capability
				err := suite.GetApp(suite.chainA.App).ScopedTransferKeeper.ReleaseCapability(suite.chainA.GetContext(), capability)
				suite.Require().NoError(err)
			}, false,
		},
		{
			"SendPacket fails, timeout height and timeout timestamp are zero",
			func() {
				timeoutHeight = clienttypes.ZeroHeight()
			}, false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = fxibctesting.NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			coin = sdk.NewCoin(fxtypes.DefaultDenom, sdk.NewInt(100))
			sender = suite.chainA.SenderAccount.GetAddress()
			memo = ""
			timeoutHeight = suite.chainB.GetTimeoutHeight()

			// create IBC token on chainA
			transferMsg := types.NewMsgTransfer(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, coin, suite.chainB.SenderAccount.GetAddress().String(), suite.chainA.SenderAccount.GetAddress().String(), suite.chainA.GetTimeoutHeight(), 0, "")
			result, err := suite.chainB.SendMsgs(transferMsg)
			suite.Require().NoError(err) // message committed

			packet, err := ibctesting.ParsePacketFromEvents(result.GetEvents())
			suite.Require().NoError(err)

			err = path.RelayPacket(packet)
			suite.Require().NoError(err)

			tc.malleate()

			msg := types.NewMsgTransfer(
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				coin, sender.String(), suite.chainB.SenderAccount.GetAddress().String(),
				timeoutHeight, 0, // only use timeout height
				memo,
			)

			res, err := suite.GetApp(suite.chainA.App).IBCTransferKeeper.Transfer(sdk.WrapSDKContext(suite.chainA.GetContext()), msg)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	baseDenom := "stake"
	senderAddr := sdk.AccAddress(tmrand.Bytes(20))
	receiveAddr := sdk.AccAddress(tmrand.Bytes(20))
	transferAmount := sdkmath.NewInt(tmrand.Int63n(100000000000))
	ibcDenomTrace := transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: baseDenom,
	}

	testCases := []struct {
		name          string
		malleate      func(fxIbcTransferMsg *channeltypes.Packet)
		expPass       bool
		errorStr      string
		checkBalance  bool
		checkCoinAddr sdk.AccAddress
		expCoins      sdk.Coins
	}{
		{
			"pass - normal - ibc transfer packet",
			func(packet *channeltypes.Packet) {
			},
			true,
			"",
			true,
			receiveAddr,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), transferAmount)),
		},
		{
			"pass - normal - fx ibc transfer packet",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = ""
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()
			},
			true,
			"",
			true,
			receiveAddr,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), transferAmount)),
		},
		{
			"pass - normal - receive address is 0xAddress, coin is DefaultCoin",
			func(packet *channeltypes.Packet) {
				protID := "transfer"
				channelID := "channel-0"
				coins := sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, transferAmount))
				err := suite.GetApp(suite.chainA.App).BankKeeper.MintCoins(suite.chainA.GetContext(), types.ModuleName, coins)
				suite.Require().NoError(err)
				portChannelAddr := transfertypes.GetEscrowAddress(protID, channelID)
				err = suite.GetApp(suite.chainA.App).BankKeeper.SendCoinsFromModuleToAccount(suite.chainA.GetContext(), types.ModuleName, portChannelAddr, coins)
				suite.Require().NoError(err)

				packetData := transfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Receiver = common.BytesToAddress(receiveAddr.Bytes()).String()
				packetData.Denom = transfertypes.DenomTrace{
					BaseDenom: fxtypes.DefaultDenom,
					Path:      fmt.Sprintf("%s/%s", protID, channelID),
				}.GetFullDenomPath()
				packet.Data = packetData.GetBytes()
			},
			true,
			"",
			true,
			receiveAddr,
			sdk.NewCoins(sdk.NewCoin(fxtypes.DefaultDenom, transferAmount)),
		},
		{
			"pass - normal - receive address is 0xAddress",
			func(packet *channeltypes.Packet) {
				packetData := transfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Receiver = common.BytesToAddress(receiveAddr.Bytes()).String()
				packet.Data = packetData.GetBytes()
				meta := banktypes.Metadata{
					// token -> contracts
					Base: ibcDenomTrace.IBCDenom(),
					// evm token: name
					Name: ibcDenomTrace.GetFullDenomPath(),
					// evm token: symbol
					Symbol: strings.ToUpper(baseDenom),
					// evm token decimal - denomunits.denom == symbol
					DenomUnits: []*banktypes.DenomUnit{
						{
							Denom:    strings.ToLower(baseDenom),
							Exponent: 0,
						},
						{
							Denom:    strings.ToUpper(baseDenom),
							Exponent: uint32(rand.Int31n(19)),
						},
					},
				}
				_, err := suite.GetApp(suite.chainA.App).Erc20Keeper.RegisterNativeCoin(suite.chainA.GetContext(), meta)
				suite.Require().NoError(err)
			},
			true,
			"",
			true,
			receiveAddr,
			sdk.NewCoins(),
		},
		{
			"pass - normal - sender is 0xAddress",
			func(packet *channeltypes.Packet) {
				bscKeeper := suite.GetApp(suite.chainA.App).BscKeeper
				bscKeeper.AddBridgeToken(suite.chainA.GetContext(), common.BytesToAddress(tmrand.Bytes(20)).String(), ibcDenomTrace.IBCDenom())
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Sender = common.BytesToAddress(senderAddr.Bytes()).String()
				packetData.Router = bsctypes.ModuleName
				packetData.Fee = sdkmath.ZeroInt().String()
				packetData.Receiver = common.BytesToAddress(receiveAddr).String()
				packet.Data = packetData.GetBytes()
			},
			true,
			"",
			true,
			senderAddr,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), sdkmath.ZeroInt())),
		},
		{
			name: "error - normal - transferAfter return error, receive address is error",
			malleate: func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Receiver = tmrand.Str(20)
				routes := []string{ethtypes.ModuleName, bsctypes.ModuleName, trontypes.ModuleName, polygontypes.ModuleName, avalanchetypes.ModuleName, erc20types.ModuleName}
				packetData.Router = routes[tmrand.Int63n(int64(len(routes)))]
				packet.Data = packetData.GetBytes()
			},
			expPass:       false,
			errorStr:      "ABCI code: 7: error handling packet: see events for details",
			checkBalance:  true,
			checkCoinAddr: senderAddr,
			expCoins:      sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), sdkmath.ZeroInt())),
		},
		{
			"error - normal - router not exists",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = tmrand.Str(8)
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()
			},
			false,
			// 103: router not found error
			"ABCI code: 103: error handling packet: see events for details",
			true,
			senderAddr,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), sdkmath.ZeroInt())),
		},
		{
			"error - normal - receive address is 0xAddress but coin not registered",
			func(packet *channeltypes.Packet) {
				packetData := transfertypes.FungibleTokenPacketData{}
				transfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Receiver = common.BytesToAddress(receiveAddr.Bytes()).String()
				packet.Data = packetData.GetBytes()
			},
			false,
			// 4: token pair not found
			"ABCI code: 4: error handling packet: see events for details",
			true,
			senderAddr,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), sdkmath.ZeroInt())),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			chain := suite.GetApp(suite.chainA.App)
			transferIBCModule := transfer.NewIBCModule(chain.IBCTransferKeeper)
			fxIBCMiddleware := fxtransfer.NewIBCMiddleware(chain.FxTransferKeeper, transferIBCModule)
			packetData := transfertypes.NewFungibleTokenPacketData(baseDenom, transferAmount.String(), senderAddr.String(), receiveAddr.String(), "")
			// only use timeout height
			packet := channeltypes.NewPacket(packetData.GetBytes(), 1, ibcgotesting.TransferPort, "channel-0", ibcgotesting.TransferPort, "channel-0", clienttypes.Height{
				RevisionNumber: 100,
				RevisionHeight: 100000,
			}, 0)
			tc.malleate(&packet)

			cacheCtx, writeFn := suite.chainA.GetContext().CacheContext()
			ackI := fxIBCMiddleware.OnRecvPacket(cacheCtx, packet, nil)
			if ackI == nil || ackI.Success() {
				// write application state changes for asynchronous and successful acknowledgements
				writeFn()
			}
			suite.Require().NotNil(ackI)

			ack, ok := ackI.(channeltypes.Acknowledgement)
			suite.chainA.GetContext().EventManager().EmitEvents(cacheCtx.EventManager().Events())

			if tc.expPass {
				suite.Require().Truef(ack.Success(), "ackError:%s,causeError:%s,packetData:%s", ack.GetError(), getOnRecvPacketErrorByEvent(cacheCtx), string(packet.GetData()))
			} else {
				suite.Require().False(ack.Success())
				suite.Require().True(ok)
				suite.Require().Equalf(tc.errorStr, ack.GetError(), "packetData:%s", string(packet.GetData()))
			}

			if tc.checkBalance {
				bankKeeper := suite.GetApp(suite.chainA.App).BankKeeper
				actualCoins := bankKeeper.GetAllBalances(suite.chainA.GetContext(), tc.checkCoinAddr)
				suite.Require().True(tc.expCoins.IsEqual(actualCoins), "exp:%s,actual:%s", tc.expCoins, actualCoins)
			}
		})
	}
}

func getOnRecvPacketErrorByEvent(ctx sdk.Context) string {
	events := ctx.EventManager().Events()
	for _, event := range events {
		if event.Type == transfertypes.EventTypePacket {
			for _, attr := range event.Attributes {
				if string(attr.Key) == fxtransfertypes.AttributeKeyRecvError {
					return string(attr.Value)
				}
			}
		}
	}
	return ""
}

func (suite *KeeperTestSuite) TestOnAcknowledgementPacket() {
	baseDenom := "stake"
	senderAddr := sdk.AccAddress(tmrand.Bytes(20))
	receiveAddr := sdk.AccAddress(tmrand.Bytes(20))
	transferAmount := sdkmath.NewInt(tmrand.Int63n(100000000000))
	testCases := []struct {
		name         string
		malleate     func(fxIbcTransferMsg *channeltypes.Packet, ack *channeltypes.Acknowledgement)
		expPass      bool
		errorStr     string
		checkBalance bool
		expCoins     sdk.Coins
	}{
		{
			"pass - success ack - ibc transfer packet",
			func(packet *channeltypes.Packet, ack *channeltypes.Acknowledgement) {
				escrowAddress := transfertypes.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)))
			},
			true,
			"",
			true,
			sdk.NewCoins(),
		},
		{
			"pass - error ack - ibc transfer packet",
			func(packet *channeltypes.Packet, ack *channeltypes.Acknowledgement) {
				*ack = channeltypes.NewErrorAcknowledgement(fmt.Errorf("test"))

				escrowAddress := transfertypes.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)))
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			chain := suite.GetApp(suite.chainA.App)
			transferIBCModule := transfer.NewIBCModule(chain.IBCTransferKeeper)
			fxIBCMiddleware := fxtransfer.NewIBCMiddleware(chain.FxTransferKeeper, transferIBCModule)
			packetData := transfertypes.NewFungibleTokenPacketData(baseDenom, transferAmount.String(), senderAddr.String(), receiveAddr.String(), "")
			// only use timeout height
			packet := channeltypes.NewPacket(packetData.GetBytes(), 1, ibcgotesting.TransferPort, "channel-0", ibcgotesting.TransferPort, "channel-0", clienttypes.Height{
				RevisionNumber: 100,
				RevisionHeight: 100000,
			}, 0)

			ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
			tc.malleate(&packet, &ack)

			err := fxIBCMiddleware.OnAcknowledgementPacket(suite.chainA.GetContext(), packet, ack.Acknowledgement(), nil)
			if tc.expPass {
				suite.Require().NoError(err, "packetData:%s", string(packet.GetData()))
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equalf(tc.errorStr, err.Error(), "packetData:%s", string(packet.GetData()))
			}

			if tc.checkBalance {
				bankKeeper := suite.GetApp(suite.chainA.App).BankKeeper
				senderAddrCoins := bankKeeper.GetAllBalances(suite.chainA.GetContext(), senderAddr)
				suite.Require().True(tc.expCoins.IsEqual(senderAddrCoins), "exp:%s,actual:%s", tc.expCoins, senderAddrCoins)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnTimeoutPacket() {
	baseDenom := "stake"
	senderAddr := sdk.AccAddress(tmrand.Bytes(20))
	receiveAddr := sdk.AccAddress(tmrand.Bytes(20))
	transferAmount := sdkmath.NewInt(tmrand.Int63n(100000000000))
	ibcDenomTrace := transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: tmrand.Str(6),
	}
	testCases := []struct {
		name         string
		malleate     func(fxIbcTransferMsg *channeltypes.Packet)
		expPass      bool
		errorStr     string
		checkBalance bool
		expCoins     sdk.Coins
	}{
		{
			"pass - normal - ibc transfer packet",
			func(packet *channeltypes.Packet) {
				escrowAddress := transfertypes.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)))
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)),
		},
		{
			"pass - normal - fx ibc transfer packet",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = ""
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()

				escrowAddress := transfertypes.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)))
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)),
		},
		{
			"pass - normal - ibc mint token - router is empty",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Denom = ibcDenomTrace.GetFullDenomPath()
				packetData.Router = ""
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(ibcDenomTrace.IBCDenom(), transferAmount)),
		},
		{
			"pass - router not empty | amount + zero fee",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = tmrand.Str(4)
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()

				amount := sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount))
				escrowAddress := transfertypes.GetEscrowAddress(packet.GetDestPort(), packet.GetDestChannel())
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, amount)
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount)),
		},
		{
			"pass - router not empty | amount + fee",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = tmrand.Str(4)
				fee := sdkmath.NewInt(50)
				packetData.Fee = fee.String()
				packet.Data = packetData.GetBytes()

				amount := sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount.Add(fee)))
				escrowAddress := transfertypes.GetEscrowAddress(packet.GetDestPort(), packet.GetDestChannel())
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, amount)
			},
			true,
			"",
			true,
			sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount.Add(sdkmath.NewInt(50)))),
		},
		{
			"error - escrow address insufficient 10coin",
			func(packet *channeltypes.Packet) {
				packetData := fxtransfertypes.FungibleTokenPacketData{}
				fxtransfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &packetData)
				packetData.Router = ""
				packetData.Fee = sdkmath.ZeroInt().String()
				packet.Data = packetData.GetBytes()

				escrowAddress := transfertypes.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
				mintCoin(suite.T(), suite.chainA.GetContext(), suite.GetApp(suite.chainA.App), escrowAddress, sdk.NewCoins(sdk.NewCoin(baseDenom, transferAmount.Sub(sdkmath.NewInt(10)))))
			},
			false,
			fmt.Sprintf("unable to unescrow tokens, this may be caused by a malicious counterparty module or a bug: please open an issue on counterparty module: %d%s is smaller than %d%s: insufficient funds", transferAmount.Sub(sdkmath.NewInt(10)).Uint64(), baseDenom, transferAmount.Uint64(), baseDenom),
			true,
			sdk.NewCoins(),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			chain := suite.GetApp(suite.chainA.App)
			transferIBCModule := transfer.NewIBCModule(chain.IBCTransferKeeper)
			fxIBCMiddleware := fxtransfer.NewIBCMiddleware(chain.FxTransferKeeper, transferIBCModule)
			packetData := transfertypes.NewFungibleTokenPacketData(baseDenom, transferAmount.String(), senderAddr.String(), receiveAddr.String(), "")
			// only use timeout height
			packet := channeltypes.NewPacket(packetData.GetBytes(), 1, ibcgotesting.TransferPort, "channel-0", ibcgotesting.TransferPort, "channel-0", clienttypes.Height{
				RevisionNumber: 100,
				RevisionHeight: 100000,
			}, 0)
			tc.malleate(&packet)

			err := fxIBCMiddleware.OnTimeoutPacket(suite.chainA.GetContext(), packet, nil)
			if tc.expPass {
				suite.Require().NoError(err, "packetData:%s", string(packet.GetData()))
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equalf(tc.errorStr, err.Error(), "packetData:%s", string(packet.GetData()))
			}

			if tc.checkBalance {
				bankKeeper := suite.GetApp(suite.chainA.App).BankKeeper
				senderAddrCoins := bankKeeper.GetAllBalances(suite.chainA.GetContext(), senderAddr)
				suite.Require().True(tc.expCoins.IsEqual(senderAddrCoins), "exp:%s,actual:%s", tc.expCoins, senderAddrCoins)
			}
		})
	}
}

func mintCoin(t *testing.T, ctx sdk.Context, chain *helpers.TestingApp, address sdk.AccAddress, coins sdk.Coins) {
	require.NoError(t, chain.BankKeeper.MintCoins(ctx, transfertypes.ModuleName, coins))
	require.NoError(t, chain.BankKeeper.SendCoinsFromModuleToAccount(ctx, transfertypes.ModuleName, address, coins))
}
