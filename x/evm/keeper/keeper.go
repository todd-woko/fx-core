package keeper

import (
	"fmt"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/types"

	fxserverconfig "github.com/functionx/fx-core/v4/server/config"
	fxtypes "github.com/functionx/fx-core/v4/types"
	fxevmtypes "github.com/functionx/fx-core/v4/x/evm/types"
)

type Keeper struct {
	*evmkeeper.Keeper

	// access to account state
	accountKeeper fxevmtypes.AccountKeeper

	// has evm hooks
	hasHooks bool

	module common.Address

	authority string
}

func NewKeeper(ek *evmkeeper.Keeper, ak fxevmtypes.AccountKeeper, authority string) *Keeper {
	ctx := sdk.Context{}.WithChainID(fxtypes.ChainIdWithEIP155())
	ek.WithChainID(ctx)
	addr := ak.GetModuleAddress(types.ModuleName)
	if addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	return &Keeper{
		Keeper:        ek,
		accountKeeper: ak,
		authority:     authority,
		module:        common.BytesToAddress(addr),
	}
}

// SetHooks sets the hooks for the EVM module
// It should be called only once during initialization, it panic if called more than once.
func (k *Keeper) SetHooks(eh types.EvmHooks) *Keeper {
	k.Keeper.SetHooks(eh)
	k.hasHooks = true
	return k
}

// CallEVMWithoutGas performs a smart contract method call using contract data without gas
func (k *Keeper) CallEVMWithoutGas(
	ctx sdk.Context,
	from common.Address,
	contract *common.Address,
	data []byte,
	commit bool,
) (*types.MsgEthereumTxResponse, error) {
	gasMeter := ctx.GasMeter()
	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	nonce, err := k.accountKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasLimit := fxserverconfig.DefaultGasCap
	params := ctx.ConsensusParams()
	if params != nil && params.Block != nil && params.Block.MaxGas > 0 {
		gasLimit = uint64(params.Block.MaxGas)
	}

	msg := ethtypes.NewMessage(
		from,
		contract,
		nonce,
		big.NewInt(0), // amount
		gasLimit,      // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.ApplyMessage(ctx, msg, types.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		errStr := res.VmError
		if res.VmError == vm.ErrExecutionReverted.Error() {
			if reason, err := abi.UnpackRevert(common.CopyBytes(res.Ret)); err == nil {
				errStr = reason
			}
		}
		return res, errorsmod.Wrap(types.ErrVMExecution, errStr)
	}

	ctx.WithGasMeter(gasMeter)

	return res, nil
}
