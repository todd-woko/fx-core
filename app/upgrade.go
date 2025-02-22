package app

import (
	"fmt"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/functionx/fx-core/v4/app/upgrades"
	v4 "github.com/functionx/fx-core/v4/app/upgrades/v4"
)

var upgradeArray = upgrades.Upgrades{v4.Upgrade}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *App) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range upgradeArray {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, upgrade.StoreUpgrades()))
		}
	}
}

func (app *App) setupUpgradeHandlers() {
	for _, upgrade := range upgradeArray {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.mm,
				app.configurator,
				app.AppKeepers,
			),
		)
	}
}

func GetUpgrades() upgrades.Upgrades {
	return upgradeArray
}
