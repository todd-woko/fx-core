package staking

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/functionx/fx-core/v4/x/evm/types"
)

type Contract struct {
	ctx           sdk.Context
	evm           *vm.EVM
	bankKeeper    BankKeeper
	distrKeeper   DistrKeeper
	stakingKeeper StakingKeeper
	evmKeeper     EvmKeeper
}

func NewPrecompiledContract(
	ctx sdk.Context,
	evm *vm.EVM,
	bankKeeper BankKeeper,
	stakingKeeper StakingKeeper,
	distrKeeper DistrKeeper,
	evmKeeper EvmKeeper,
) *Contract {
	return &Contract{
		ctx:           ctx,
		evm:           evm,
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
		distrKeeper:   distrKeeper,
		evmKeeper:     evmKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return stakingAddress
}

func (c *Contract) IsStateful() bool {
	return true
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	if len(input) <= 4 {
		return 0
	}
	switch string(input[:4]) {
	case string(DelegateMethod.ID):
		return DelegateGas
	case string(UndelegateMethod.ID):
		return UndelegateGas
	case string(WithdrawMethod.ID):
		return WithdrawGas
	case string(DelegationMethod.ID):
		return DelegationGas
	case string(DelegationRewardsMethod.ID):
		return DelegationRewardsGas
	case string(TransferSharesMethod.ID):
		return TransferSharesGas
	case string(ApproveSharesMethod.ID):
		return ApproveSharesGas
	case string(AllowanceSharesMethod.ID):
		return AllowanceSharesGas
	case string(TransferFromSharesMethod.ID):
		return TransferFromSharesGas
	default:
		return 0
	}
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) <= 4 {
		return types.PackRetError("invalid input")
	}

	cacheCtx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()

	// parse input
	switch string(contract.Input[:4]) {
	case string(DelegateMethod.ID):
		ret, err = c.Delegate(cacheCtx, evm, contract, readonly)
	case string(UndelegateMethod.ID):
		ret, err = c.Undelegate(cacheCtx, evm, contract, readonly)
	case string(WithdrawMethod.ID):
		ret, err = c.Withdraw(cacheCtx, evm, contract, readonly)
	case string(DelegationMethod.ID):
		ret, err = c.Delegation(cacheCtx, evm, contract, readonly)
	case string(DelegationRewardsMethod.ID):
		ret, err = c.DelegationRewards(cacheCtx, evm, contract, readonly)
	case string(TransferSharesMethod.ID):
		ret, err = c.TransferShares(cacheCtx, evm, contract, readonly)
	case string(ApproveSharesMethod.ID):
		ret, err = c.ApproveShares(cacheCtx, evm, contract, readonly)
	case string(AllowanceSharesMethod.ID):
		ret, err = c.AllowanceShares(cacheCtx, evm, contract, readonly)
	case string(TransferFromSharesMethod.ID):
		ret, err = c.TransferFromShares(cacheCtx, evm, contract, readonly)
	default:
		err = errors.New("unknown method")
	}

	if err != nil {
		// revert evm state
		evm.StateDB.RevertToSnapshot(snapshot)
		return types.PackRetError(err.Error())
	}

	// commit and append events
	commit()
	return ret, nil
}

func (c *Contract) AddLog(event abi.Event, topics []common.Hash, args ...interface{}) error {
	data, err := event.Inputs.NonIndexed().Pack(args...)
	if err != nil {
		return fmt.Errorf("pack %s event error: %s", event.Name, err.Error())
	}
	newTopic := []common.Hash{event.ID}
	if len(topics) > 0 {
		newTopic = append(newTopic, topics...)
	}
	c.evm.StateDB.AddLog(&ethtypes.Log{
		Address:     c.Address(),
		Topics:      newTopic,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	})
	return nil
}

func ParseMethodParams(method abi.Method, v interface{}, data []byte) error {
	unpacked, err := method.Inputs.Unpack(data)
	if err != nil {
		return err
	}
	return method.Inputs.Copy(v, unpacked)
}
