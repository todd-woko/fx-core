package types_test

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/functionx/fx-core/v4/testutil/helpers"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

func TestIsEmptyHash(t *testing.T) {
	testCases := []struct {
		name     string
		hash     string
		expEmpty bool
	}{
		{
			"empty string", "", true,
		},
		{
			"zero hash", common.Hash{}.String(), true,
		},

		{
			"non-empty hash", common.BytesToHash([]byte{1, 2, 3, 4}).String(), false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expEmpty, fxtypes.IsEmptyHash(tc.hash), tc.name)
	}
}

func TestIsZeroEthereumAddress(t *testing.T) {
	testCases := []struct {
		name     string
		address  string
		expEmpty bool
	}{
		{
			"empty string", "", true,
		},
		{
			"zero address", common.Address{}.String(), true,
		},

		{
			"non-empty address", common.BytesToAddress([]byte{1, 2, 3, 4}).String(), false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expEmpty, fxtypes.IsZeroEthereumAddress(tc.address), tc.name)
	}
}

func TestValidateEthereumAddress(t *testing.T) {
	testCases := []struct {
		name     string
		address  string
		expError bool
	}{
		{
			"empty string", "", true,
		},
		{
			"invalid address", "0x", true,
		},
		{
			"zero address", common.Address{}.String(), false,
		},
		{
			"valid address", helpers.GenerateAddress().Hex(), false,
		},
		{
			"invalid address - upper address", strings.ToUpper(helpers.GenerateAddress().Hex()), true,
		},
		{
			"invalid address - lower address", strings.ToLower(helpers.GenerateAddress().Hex()), true,
		},
	}

	for _, tc := range testCases {
		err := fxtypes.ValidateEthereumAddress(tc.address)

		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}
