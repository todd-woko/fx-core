package parser_test

import (
	"testing"

	"github.com/functionx/fx-core/v2/x/ibc/applications/transfer/parser"

	"github.com/stretchr/testify/require"
)

func TestParseReceiverDataTransfer(t *testing.T) {
	data := "cosmos1vzxkv3lxccnttr9rs0002s93sgw72h7ghukuhs|transfer/channel-0:cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k"
	pt, err := parser.ParseReceiverData(data)

	require.NoError(t, err)
	require.True(t, pt.ShouldForward)
	require.Equal(t, pt.HostAccAddr.String(), "cosmos1vzxkv3lxccnttr9rs0002s93sgw72h7ghukuhs")
	require.Equal(t, pt.Destination, "cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k")
	require.Equal(t, pt.Port, "transfer")
	require.Equal(t, pt.Channel, "channel-0")
}

func TestParseReceiverDataNoTransfer(t *testing.T) {
	data := "cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k"
	pt, err := parser.ParseReceiverData(data)

	require.NoError(t, err)
	require.False(t, pt.ShouldForward)
}

func TestParseReceiverDataErrors(t *testing.T) {
	testCases := []struct {
		name          string
		data          string
		errStartsWith string
	}{
		{
			"unparsable transfer field",
			"",
			"unparsable receiver",
		},
		{
			"unparsable transfer field",
			"abc:def:",
			"unparsable receiver",
		},
		{
			"missing pipe",
			"transfer/channel-0:cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k",
			"formatting incorrect",
		},
		{
			"invalid this chain address",
			"somm16plylpsgxechajltx9yeseqexzdzut9g8vla4k|transfer/channel-0:cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k",
			"decoding bech32 failed",
		},
		{
			"missing slash",
			"cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k|transfer\\channel-0:cosmos16plylpsgxechajltx9yeseqexzdzut9g8vla4k",
			"formatting incorrect",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.ParseReceiverData(tc.data)
			require.Error(t, err)
			require.Equal(t, err.Error()[:len(tc.errStartsWith)], tc.errStartsWith)
		})
	}
}
