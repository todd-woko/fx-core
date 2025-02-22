package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tranfsertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"

	fxtypes "github.com/functionx/fx-core/v4/types"
)

type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	GetUnbondingDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (ubd stakingtypes.UnbondingDelegation, found bool)

	RemoveDelegation(ctx sdk.Context, delegation stakingtypes.Delegation) error
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
}

type StakingMsgServer interface {
	Delegate(goCtx context.Context, msg *stakingtypes.MsgDelegate) (*stakingtypes.MsgDelegateResponse, error)
	BeginRedelegate(goCtx context.Context, msg *stakingtypes.MsgBeginRedelegate) (*stakingtypes.MsgBeginRedelegateResponse, error)
	Undelegate(goCtx context.Context, msg *stakingtypes.MsgUndelegate) (*stakingtypes.MsgUndelegateResponse, error)
}

type DistributionKeeper interface {
	WithdrawDelegatorReward(goCtx context.Context, msg *distributiontypes.MsgWithdrawDelegatorReward) (*distributiontypes.MsgWithdrawDelegatorRewardResponse, error)
}

// BankKeeper defines the expected bank keeper methods
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetDenomMetaData(ctx sdk.Context, denom string) (bank.Metadata, bool)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	IterateAllDenomMetaData(ctx sdk.Context, cb func(bank.Metadata) bool)
}

type Erc20Keeper interface {
	TransferAfter(ctx sdk.Context, sender sdk.AccAddress, receive string, coin, fee sdk.Coin, _ bool) error
	ConvertDenomToTarget(ctx sdk.Context, from sdk.AccAddress, coin sdk.Coin, fxTarget fxtypes.FxTarget) (sdk.Coin, error)
	HookOutgoingRefund(ctx sdk.Context, txID uint64, sender sdk.AccAddress, totalCoin sdk.Coin) error
	SetOutgoingTransferRelation(ctx sdk.Context, txID uint64)
	HasOutgoingTransferRelation(ctx sdk.Context, txID uint64) bool
	IsOriginDenom(ctx sdk.Context, denom string) bool
	ToTargetDenom(ctx sdk.Context, denom, base string, aliases []string, fxTarget fxtypes.FxTarget) string
}

type IBCTransferKeeper interface {
	Transfer(goCtx context.Context, msg *tranfsertypes.MsgTransfer) (*tranfsertypes.MsgTransferResponse, error)

	SetDenomTrace(ctx sdk.Context, denomTrace tranfsertypes.DenomTrace)
}

type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
}

type (
	ParamSet = paramtypes.ParamSet
	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	Subspace interface {
		GetParamSet(ctx sdk.Context, ps ParamSet)
		HasKeyTable() bool
		WithKeyTable(table paramtypes.KeyTable) paramtypes.Subspace
	}
)
