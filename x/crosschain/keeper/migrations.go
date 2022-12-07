package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	crosschainv3 "github.com/functionx/fx-core/v3/x/crosschain/legacy/v3"
	"github.com/functionx/fx-core/v3/x/crosschain/types"
)

func (k Keeper) Migrate2to3(ctx sdk.Context) error {
	// update params
	k.paramSpace.Set(ctx, types.ParamsStoreKeySignedWindow, uint64(30_000))
	k.paramSpace.Set(ctx, types.ParamsStoreSlashFraction, sdk.NewDecWithPrec(8, 1))

	// fix oracle delegate
	validatorsByPower := k.stakingKeeper.GetBondedValidatorsByPower(ctx)
	if len(validatorsByPower) <= 0 {
		panic("no found bonded validator")
	}
	validator := validatorsByPower[0].GetOperator()
	oracles := k.GetAllOracles(ctx, false)
	proposalOracle, _ := k.GetProposalOracle(ctx)
	return crosschainv3.MigrateDepositToStaking(ctx, k.moduleName, k.stakingKeeper, k.bankKeeper, oracles, proposalOracle, validator)
}
