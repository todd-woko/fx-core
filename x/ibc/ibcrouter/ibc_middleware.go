package ibcrouter

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"

	"github.com/functionx/fx-core/v2/x/ibc/ibcrouter/parser"
	"github.com/functionx/fx-core/v2/x/ibc/ibcrouter/types"

	"github.com/functionx/fx-core/v2/x/ibc"

	"github.com/cosmos/ibc-go/v3/modules/core/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"

	fxtransfertypes "github.com/functionx/fx-core/v2/x/ibc/applications/transfer/types"
)

const (
	ForwardPacketTimeHour time.Duration = 12
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 interface for transfer given the transfer keeper.
type IBCMiddleware struct {
	*ibc.Module
	keeper      transferkeeper.Keeper
	ics4Wrapper porttypes.ICS4Wrapper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(k transferkeeper.Keeper, ics4Wrapper porttypes.ICS4Wrapper, app porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		Module:      ibc.NewModule(app),
		keeper:      k,
		ics4Wrapper: ics4Wrapper,
	}
}

// OnRecvPacket implements the IBCModule interface
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	var data fxtransfertypes.FungibleTokenPacketData
	if err := fxtransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement("cannot unmarshal ICS-20 transfer packet data")
	}

	// check the packet has router
	if len(data.Router) > 0 {
		return im.Module.OnRecvPacket(ctx, packet, relayer)
	}

	ack, err := handlerForwardTransferPacket(ctx, im, packet, transfertypes.NewFungibleTokenPacketData(data.GetDenom(), data.GetAmount(), data.GetSender(), data.GetReceiver()), relayer)
	if err != nil {
		ack = transfertypes.NewErrorAcknowledgement(err)
	}

	if err != nil {
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeRouter,
			sdk.NewAttribute(types.AttributeKeyRouteError, err.Error()),
		))
	}
	return ack
}

func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.Module.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.Module.OnTimeoutPacket(ctx, packet, relayer)
}

func (im IBCMiddleware) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI) error {
	return im.ics4Wrapper.SendPacket(ctx, chanCap, packet)
}

func (im IBCMiddleware) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, ack exported.Acknowledgement) error {
	return im.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func handlerForwardTransferPacket(ctx sdk.Context, im IBCMiddleware, packet channeltypes.Packet, data transfertypes.FungibleTokenPacketData, relayer sdk.AccAddress) (exported.Acknowledgement, error) {
	// parse out any forwarding info
	parsedReceiver, err := parser.ParseReceiverData(data.Receiver)
	if err != nil {
		return nil, err
	}

	if !parsedReceiver.ShouldForward {
		return im.Module.OnRecvPacket(ctx, packet, relayer), nil
	}

	newData := data
	newData.Receiver = parsedReceiver.HostAccAddr.String()
	bz, err := transfertypes.ModuleCdc.MarshalJSON(&newData)
	if err != nil {
		return nil, err
	}
	newPacket := packet
	newPacket.Data = bz

	ack := im.Module.OnRecvPacket(ctx, newPacket, relayer)
	if ack.Success() {
		// recalculate denom, skip checks that were already done in app.OnRecvPacket
		denom := GetDenomByIBCPacket(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetDestPort(), packet.GetDestChannel(), newData.GetDenom())
		// parse the transfer amount
		transferAmount, ok := sdk.NewIntFromString(data.Amount)
		if !ok {
			return nil, sdkerrors.Wrapf(transfertypes.ErrInvalidAmount, "unable to parse forward transfer amount (%s) into sdk.Int", data.Amount)
		}

		var token = sdk.NewCoin(denom, transferAmount)
		err = im.keeper.SendTransfer(ctx, parsedReceiver.Port, parsedReceiver.Channel, token, parsedReceiver.HostAccAddr, parsedReceiver.Destination, clienttypes.Height{}, uint64(ctx.BlockTime().Add(ForwardPacketTimeHour*time.Hour).UnixNano()))
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	}
	return ack, nil
}

func GetDenomByIBCPacket(sourcePort, sourceChannel, destPort, destChannel, packetDenom string) string {
	var denom string

	if transfertypes.ReceiverChainIsSource(sourcePort, sourceChannel, packetDenom) {
		voucherPrefix := transfertypes.GetDenomPrefix(sourcePort, sourceChannel)
		unPrefixedDenom := packetDenom[len(voucherPrefix):]

		// coin denomination used in sending from the escrow address
		denom = unPrefixedDenom

		// The denomination used to send the coins is either the native denom or the hash of the path
		// if the denomination is not native.
		denomTrace := transfertypes.ParseDenomTrace(unPrefixedDenom)
		if denomTrace.Path != "" {
			denom = denomTrace.IBCDenom()
		}
	} else {
		// since SendPacket did not prefix the denomination, we must prefix denomination here
		sourcePrefix := transfertypes.GetDenomPrefix(destPort, destChannel)
		// NOTE: sourcePrefix contains the trailing "/"
		prefixedDenom := sourcePrefix + packetDenom

		// construct the denomination trace from the full raw denomination
		denom = transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	}
	return denom
}
