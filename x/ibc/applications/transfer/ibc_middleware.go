package transfer

import (
	"fmt"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	"github.com/functionx/fx-core/v2/x/ibc/applications/transfer/parser"

	"github.com/functionx/fx-core/v2/x/ibc"

	fxtypes "github.com/functionx/fx-core/v2/types"

	"github.com/cosmos/ibc-go/v3/modules/core/exported"

	"time"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	coretypes "github.com/cosmos/ibc-go/v3/modules/core/types"

	"github.com/functionx/fx-core/v2/x/ibc/applications/transfer/keeper"
	"github.com/functionx/fx-core/v2/x/ibc/applications/transfer/types"
)

const (
	ForwardPacketTimeHour time.Duration = 12
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 interface for transfer given the transfer keeper.
type IBCMiddleware struct {
	*ibc.Module
	keeper keeper.Keeper
}

func (im IBCMiddleware) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI) error {
	return im.keeper.SendPacket(ctx, chanCap, packet)
}

func (im IBCMiddleware) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, ack exported.Acknowledgement) error {
	//TODO implement me
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(k keeper.Keeper, app porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		Module: ibc.NewModule(app),
		keeper: k,
	}
}

// OnRecvPacket implements the IBCModule interface
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})

	var data types.FungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		ack = channeltypes.NewErrorAcknowledgement("cannot unmarshal ICS-20 transfer packet data")
	}

	// only attempt the application logic if the packet data
	// was successfully decoded
	var err error
	if ack.Success() {
		if len(data.GetFee()) == 0 {
			data.Fee = sdk.ZeroInt().String()
		}
		// if router set, route packet
		if ctx.BlockHeight() >= fxtypes.IBCRouteBlock() && data.Router != "" {
			err = im.keeper.FxOnRecvPacket(ctx, packet, data)
		} else {
			err = handlerForwardTransferPacket(ctx, im, packet, transfertypes.NewFungibleTokenPacketData(data.GetDenom(), data.GetAmount(), data.GetSender(), data.GetReceiver()))
		}

		if err != nil {
			ack = transfertypes.NewErrorAcknowledgement(err)
		}
	}

	event := sdk.NewEvent(
		transfertypes.EventTypePacket,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(transfertypes.AttributeKeyReceiver, data.Receiver),
		sdk.NewAttribute(transfertypes.AttributeKeyDenom, data.Denom),
		sdk.NewAttribute(transfertypes.AttributeKeyAmount, data.Amount),
		sdk.NewAttribute(transfertypes.AttributeKeyAckSuccess, fmt.Sprintf("%t", ack.Success())),
	)

	if err != nil {
		event = event.AppendAttributes(sdk.NewAttribute(types.AttributeKeyRecvError, err.Error()))
	}
	ctx.EventManager().EmitEvent(
		event,
	)

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}
	var data types.FungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}

	if err := im.keeper.OnAcknowledgementPacket(ctx, packet, data, ack); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			transfertypes.EventTypePacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(transfertypes.AttributeKeyReceiver, data.Receiver),
			sdk.NewAttribute(transfertypes.AttributeKeyDenom, data.Denom),
			sdk.NewAttribute(transfertypes.AttributeKeyAmount, data.Amount),
			sdk.NewAttribute(transfertypes.AttributeKeyAck, ack.String()),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				transfertypes.EventTypePacket,
				sdk.NewAttribute(transfertypes.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				transfertypes.EventTypePacket,
				sdk.NewAttribute(transfertypes.AttributeKeyAckError, resp.Error),
			),
		)
	}

	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	var data types.FungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}
	// refund tokens
	if err := im.keeper.OnTimeoutPacket(ctx, packet, data); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			transfertypes.EventTypeTimeout,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(transfertypes.AttributeKeyRefundReceiver, data.Sender),
			sdk.NewAttribute(transfertypes.AttributeKeyRefundDenom, data.Denom),
			sdk.NewAttribute(transfertypes.AttributeKeyRefundAmount, data.Amount),
		),
	)

	return nil
}

func handlerForwardTransferPacket(ctx sdk.Context, im IBCMiddleware, packet channeltypes.Packet, data transfertypes.FungibleTokenPacketData) error {
	// parse out any forwarding info
	parsedReceiver, err := parser.ParseReceiverData(data.Receiver)
	if err != nil {
		return err
	}

	if !parsedReceiver.ShouldForward {
		return im.keeper.OnRecvPacket(ctx, packet, data)
	}

	newData := data
	newData.Receiver = parsedReceiver.HostAccAddr.String()
	bz, err := types.ModuleCdc.MarshalJSON(&newData)
	if err != nil {
		return err
	}
	newPacket := packet
	newPacket.Data = bz

	if err = im.keeper.OnRecvPacket(ctx, newPacket, newData); err != nil {
		return err
	}
	// recalculate denom, skip checks that were already done in app.OnRecvPacket
	denom := GetDenomByIBCPacket(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetDestPort(), packet.GetDestChannel(), newData.GetDenom())
	// parse the transfer amount
	transferAmount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return sdkerrors.Wrapf(transfertypes.ErrInvalidAmount, "unable to parse forward transfer amount (%s) into sdk.Int", data.Amount)
	}

	var token = sdk.NewCoin(denom, transferAmount)
	err = im.keeper.SendTransfer(ctx, parsedReceiver.Port, parsedReceiver.Channel, token, parsedReceiver.HostAccAddr, parsedReceiver.Destination, clienttypes.Height{}, uint64(ctx.BlockTime().Add(ForwardPacketTimeHour*time.Hour).UnixNano()))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}
	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"ibc", types.ModuleName, "packet", "forward"},
			1,
			append(
				[]metrics.Label{
					telemetry.NewLabel(coretypes.LabelSourcePort, packet.GetSourcePort()),
					telemetry.NewLabel(coretypes.LabelSourceChannel, packet.GetSourceChannel()),
				},
				telemetry.NewLabel(coretypes.LabelSource, "true"),
			),
		)
	}()

	return nil
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
