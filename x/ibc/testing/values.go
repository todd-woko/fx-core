/*
This file contains the variables, constants, and default values
used in the testing package and commonly defined in tests.
*/
package ibctesting

import (
	"time"

	_ "github.com/functionx/fx-core/v2/app"
	fxtypes "github.com/functionx/fx-core/v2/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	"github.com/functionx/fx-core/v2/x/ibc/testing/mock"
)

const (
	FirstClientID     = "07-tendermint-0"
	FirstChannelID    = "channel-0"
	FirstConnectionID = "connection-0"

	// Default params constants used to create a TM client
	TrustingPeriod     time.Duration = time.Hour * 24 * 7 * 2
	UnbondingPeriod    time.Duration = time.Hour * 24 * 7 * 3
	MaxClockDrift      time.Duration = time.Second * 10
	DefaultDelayPeriod uint64        = 0

	DefaultChannelVersion = mock.Version
	InvalidID             = "IDisInvalid"

	// Application Ports
	TransferPort = transfertypes.ModuleName
	MockPort     = mock.ModuleName

	// used for testing proposals
	Title       = "title"
	Description = "description"

	LongString = "LoremipsumdolorsitameconsecteturadipiscingeliseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequDuisauteiruredolorinreprehenderitinvoluptateelitsseillumoloreufugiatnullaariaturEcepteurintoccaectupidatatonroidentuntnulpauifficiaeseruntmollitanimidestlaborum"
)

var (
	DefaultOpenInitVersion *connectiontypes.Version

	// Default params variables used to create a TM client
	DefaultTrustLevel ibctmtypes.Fraction = ibctmtypes.DefaultTrustLevel
	TestCoin                              = sdk.NewCoin(fxtypes.DefaultDenom, sdk.NewInt(100))

	UpgradePath = []string{"upgrade", "upgradedIBCState"}

	ConnectionVersion = connectiontypes.ExportedVersionsToProto(connectiontypes.GetCompatibleVersions())[0]

	MockAcknowledgement          = mock.MockAcknowledgement.Acknowledgement()
	MockPacketData               = mock.MockPacketData
	MockFailPacketData           = mock.MockFailPacketData
	MockRecvCanaryCapabilityName = mock.MockRecvCanaryCapabilityName

	prefix = commitmenttypes.NewMerklePrefix([]byte("ibc"))
)
