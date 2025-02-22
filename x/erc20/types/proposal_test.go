package types

import (
	"strings"
	"testing"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/evmos/ethermint/tests"
	"github.com/stretchr/testify/suite"
)

type ProposalTestSuite struct {
	suite.Suite
}

func TestProposalTestSuite(t *testing.T) {
	suite.Run(t, new(ProposalTestSuite))
}

func (suite *ProposalTestSuite) TestKeysTypes() {
	suite.Require().Equal("erc20", (&RegisterCoinProposal{}).ProposalRoute())
	suite.Require().Equal("RegisterCoin", (&RegisterCoinProposal{}).ProposalType())
	suite.Require().Equal("erc20", (&RegisterERC20Proposal{}).ProposalRoute())
	suite.Require().Equal("RegisterERC20", (&RegisterERC20Proposal{}).ProposalType())
	suite.Require().Equal("erc20", (&ToggleTokenConversionProposal{}).ProposalRoute())
	suite.Require().Equal("ToggleTokenConversion", (&ToggleTokenConversionProposal{}).ProposalType())
}

func (suite *ProposalTestSuite) TestCreateDenomDescription() {
	testCases := []struct {
		name      string
		denom     string
		expString string
	}{
		{
			"with valid address",
			"0xdAC17F958D2ee523a2206206994597C13D831ec7",
			"Function X coin token representation of 0xdAC17F958D2ee523a2206206994597C13D831ec7",
		},
		{
			"with empty string",
			"",
			"Function X coin token representation of ",
		},
	}
	for _, tc := range testCases {
		desc := CreateDenomDescription(tc.denom)
		suite.Require().Equal(desc, tc.expString)
	}
}

func (suite *ProposalTestSuite) TestRegisterERC20Proposal() {
	testCases := []struct {
		msg         string
		title       string
		description string
		pair        TokenPair
		expectPass  bool
	}{
		// Valid tests
		{msg: "Register token pair - valid pair enabled", title: "test", description: "test desc", pair: TokenPair{tests.GenerateAddress().String(), "test", true, OWNER_MODULE}, expectPass: true},
		{msg: "Register token pair - valid pair dissabled", title: "test", description: "test desc", pair: TokenPair{tests.GenerateAddress().String(), "test", false, OWNER_MODULE}, expectPass: true},
		// Missing params valid
		{msg: "Register token pair - invalid missing title ", title: "", description: "test desc", pair: TokenPair{tests.GenerateAddress().String(), "test", false, OWNER_MODULE}, expectPass: false},
		{msg: "Register token pair - invalid missing description ", title: "test", description: "", pair: TokenPair{tests.GenerateAddress().String(), "test", false, OWNER_MODULE}, expectPass: false},
		// Invalid address
		{msg: "Register token pair - invalid address (no hex)", title: "test", description: "test desc", pair: TokenPair{"0x5dCA2483280D9727c80b5518faC4556617fb19ZZ", "test", true, OWNER_MODULE}, expectPass: false},
		{msg: "Register token pair - invalid address (invalid length 1)", title: "test", description: "test desc", pair: TokenPair{"0x5dCA2483280D9727c80b5518faC4556617fb19", "test", true, OWNER_MODULE}, expectPass: false},
		{msg: "Register token pair - invalid address (invalid length 2)", title: "test", description: "test desc", pair: TokenPair{"0x5dCA2483280D9727c80b5518faC4556617fb194FFF", "test", true, OWNER_MODULE}, expectPass: false},
		{msg: "Register token pair - invalid address (invalid prefix)", title: "test", description: "test desc", pair: TokenPair{"1x5dCA2483280D9727c80b5518faC4556617fb19F", "test", true, OWNER_MODULE}, expectPass: false},
	}

	for i, tc := range testCases {
		tx := NewRegisterERC20Proposal(tc.title, tc.description, tc.pair.Erc20Address, []string{"eth0xaE0c78230711E7173D30353F76936a7dD7306b81"})
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}

func createFullMetadata(denom, symbol, name string) banktypes.Metadata {
	return banktypes.Metadata{
		Description: "desc",
		Base:        denom,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    denom,
				Exponent: 0,
			},
			{
				Denom:    symbol,
				Exponent: uint32(18),
			},
		},
		Name:    name,
		Symbol:  symbol,
		Display: denom,
	}
}

func createMetadata(denom, symbol string) banktypes.Metadata {
	return createFullMetadata(denom, symbol, denom)
}

func (suite *ProposalTestSuite) TestRegisterCoinProposal() {
	validMetadata := banktypes.Metadata{
		Description: "desc",
		Base:        "coin",
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "coin",
				Exponent: 0,
			},
			{
				Denom:    "coin2",
				Exponent: uint32(18),
			},
		},
		Name:    "coin",
		Symbol:  "coin2",
		Display: "coin",
	}

	validIBCDenom := "ibc/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2"
	validIBCSymbol := "ATOM"
	validIBCName := "Atom"

	testCases := []struct {
		msg         string
		title       string
		description string
		metadata    banktypes.Metadata
		expectPass  bool
	}{
		// Valid tests
		{msg: "Register token pair - valid pair enabled", title: "test", description: "test desc", metadata: validMetadata, expectPass: true},
		{msg: "Register token pair - valid pair dissabled", title: "test", description: "test desc", metadata: validMetadata, expectPass: true},

		// Invalid Regex (denom)
		{msg: "Register token pair - invalid starts with number", title: "test", description: "test desc", metadata: createMetadata("1test", "test"), expectPass: false},
		{msg: "Register token pair - invalid char '('", title: "test", description: "test desc", metadata: createMetadata("(test", "test"), expectPass: false},
		{msg: "Register token pair - invalid char '^'", title: "test", description: "test desc", metadata: createMetadata("^test", "test"), expectPass: false},
		// Invalid length
		{msg: "Register token pair - invalid length token (0)", title: "test", description: "test desc", metadata: createMetadata("", "test"), expectPass: false},
		{msg: "Register token pair - invalid length token (1)", title: "test", description: "test desc", metadata: createMetadata("a", "test"), expectPass: false},
		{msg: "Register token pair - invalid length token (128)", title: "test", description: "test desc", metadata: createMetadata(strings.Repeat("a", 129), "test"), expectPass: false},
		{msg: "Register token pair - invalid length title (140)", title: strings.Repeat("a", govv1beta1.MaxTitleLength+1), description: "test desc", metadata: validMetadata, expectPass: false},
		{msg: "Register token pair - invalid length description (5000)", title: "title", description: strings.Repeat("a", govv1beta1.MaxDescriptionLength+1), metadata: validMetadata, expectPass: false},

		// Ibc
		{msg: "Register token pair - ibc", title: "test", description: "test desc", metadata: createFullMetadata(validIBCDenom, validIBCSymbol, validIBCName), expectPass: true},
		{msg: "Register token pair - ibc invalid denom", title: "test", description: "test desc", metadata: createFullMetadata("ibc/", validIBCSymbol, validIBCName), expectPass: false},
	}

	for i, tc := range testCases {
		tx := NewRegisterCoinProposal(tc.title, tc.description, tc.metadata)
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}

func (suite *ProposalTestSuite) TestToggleTokenConversionProposal() {
	testCases := []struct {
		msg         string
		title       string
		description string
		token       string
		expectPass  bool
	}{
		{msg: "Enable token conversion proposal - valid denom", title: "test", description: "test desc", token: "test", expectPass: true},
		{msg: "Enable token conversion proposal - valid address", title: "test", description: "test desc", token: "0x5dCA2483280D9727c80b5518faC4556617fb194F", expectPass: true},
		{msg: "Enable token conversion proposal - invalid address", title: "test", description: "test desc", token: "0x123", expectPass: false},

		// Invalid missing params
		{msg: "Enable token conversion proposal - valid missing title", title: "", description: "test desc", token: "test", expectPass: false},
		{msg: "Enable token conversion proposal - valid missing description", title: "test", description: "", token: "test", expectPass: false},
		{msg: "Enable token conversion proposal - invalid missing token", title: "test", description: "test desc", token: "", expectPass: false},

		// Invalid regex
		{msg: "Enable token conversion proposal - invalid denom", title: "test", description: "test desc", token: "^test", expectPass: false},
		// Invalid length
		{msg: "Enable token conversion proposal - invalid length (1)", title: "test", description: "test desc", token: "a", expectPass: false},
		{msg: "Enable token conversion proposal - invalid length (128)", title: "test", description: "test desc", token: strings.Repeat("a", 129), expectPass: false},

		{msg: "Enable token conversion proposal - invalid length title (140)", title: strings.Repeat("a", govv1beta1.MaxTitleLength+1), description: "test desc", token: "test", expectPass: false},
		{msg: "Enable token conversion proposal - invalid length description (5000)", title: "title", description: strings.Repeat("a", govv1beta1.MaxDescriptionLength+1), token: "test", expectPass: false},
	}

	for i, tc := range testCases {
		tx := NewToggleTokenConversionProposal(tc.title, tc.description, tc.token)
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
		}
	}
}
