package ante_test

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/ethermint/tests"

	"github.com/functionx/fx-core/v4/ante"
	fxtypes "github.com/functionx/fx-core/v4/types"
)

func (suite *AnteTestSuite) TestMempoolFeeDecorator() {
	clientCtx := NewClientCtx()
	txBuilder := clientCtx.TxConfig.NewTxBuilder()
	mfd := ante.NewMempoolFeeDecorator([]string{
		sdk.MsgTypeURL(&ibcchanneltypes.MsgRecvPacket{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgAcknowledgement{}),
		sdk.MsgTypeURL(&ibcclienttypes.MsgUpdateClient{}),
	}, 300_000)
	anteHandler := sdk.ChainAnteDecorators(mfd)
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	msg := testdata.NewTestMsg(addr1)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()

	suite.Require().NoError(txBuilder.SetMsgs(msg))
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}

	tx, err := suite.CreateEmptyTestTx(txBuilder, privs, accNums, accSeqs)
	suite.Require().NoError(err)

	// Set high gas price so standard test fee fails
	minGasPrice := []sdk.DecCoin{sdk.NewDecCoinFromDec(fxtypes.DefaultDenom, sdk.NewDec(200))}
	ctx := suite.ctx.WithMinGasPrices(minGasPrice).WithIsCheckTx(true)

	// anteHandler errors with insufficient fees
	_, err = anteHandler(ctx, tx, false)
	suite.Require().Error(err, "expected error due to low fee")

	// ensure no fees for certain IBC msgs
	suite.Require().NoError(txBuilder.SetMsgs(
		ibcchanneltypes.NewMsgRecvPacket(ibcchanneltypes.Packet{}, nil, ibcclienttypes.Height{}, sdk.AccAddress{}.String()),
	))

	oracleTx, err := suite.CreateEmptyTestTx(txBuilder, privs, accNums, accSeqs)
	suite.Require().NoError(err)
	_, err = anteHandler(ctx, oracleTx, false)
	suite.Require().NoError(err, "expected min fee bypass for IBC messages")

	ctx = ctx.WithIsCheckTx(false)

	// anteHandler should not error since we do not check min gas prices in DeliverTx
	_, err = anteHandler(ctx, tx, false)
	suite.Require().NoError(err, "unexpected error during DeliverTx")
}

var execTypes = []struct {
	name      string
	isCheckTx bool
	simulate  bool
}{
	{"deliverTx", false, false},
	{"deliverTxSimulate", false, true},
}

func (suite *AnteTestSuite) TestEthMinGasPriceDecorator() {
	denom := fxtypes.DefaultDenom
	from, privKey := tests.NewAddrKey()
	to := tests.GenerateAddress()
	emptyAccessList := ethtypes.AccessList{}

	testCases := []struct {
		name     string
		malleate func() sdk.Tx
		expPass  bool
		errMsg   string
	}{
		{
			"invalid tx type",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)
				return &invalidTx{}
			},
			false,
			"invalid message type",
		},
		{
			"wrong tx type",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)
				testMsg := banktypes.MsgSend{
					FromAddress: "fx15n7j0gwmywxwnyvwd4mrudcmxxvcehhn7djr7n",
					ToAddress:   "fx1mqruv0fx7sq06mlpa9gex5fcuxsyrz8dn9d02t",
					Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(10), Denom: denom}},
				}
				txBuilder := suite.CreateTestCosmosTxBuilder(sdkmath.NewInt(0), denom, &testMsg)
				return txBuilder.GetTx()
			},
			false,
			"invalid message type",
		},
		{
			"valid: invalid tx type with MinGasPrices = 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.ZeroDec()
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)
				return &invalidTx{}
			},
			true,
			"",
		},
		{
			"valid legacy tx with MinGasPrices = 0, gasPrice = 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.ZeroDec()
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), big.NewInt(0), nil, nil, nil)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"valid legacy tx with MinGasPrices = 0, gasPrice > 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.ZeroDec()
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), big.NewInt(10), nil, nil, nil)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"valid legacy tx with MinGasPrices = 10, gasPrice = 10",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), big.NewInt(10), nil, nil, nil)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"invalid legacy tx with MinGasPrices = 10, gasPrice = 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), big.NewInt(0), nil, nil, nil)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			false,
			"provided fee < minimum global fee",
		},
		{
			"valid dynamic tx with MinGasPrices = 0, EffectivePrice = 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.ZeroDec()
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(0), big.NewInt(0), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"valid dynamic tx with MinGasPrices = 0, EffectivePrice > 0",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.ZeroDec()
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(100), big.NewInt(50), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"valid dynamic tx with MinGasPrices < EffectivePrice",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(100), big.NewInt(100), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
		{
			"invalid dynamic tx with MinGasPrices > EffectivePrice",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(0), big.NewInt(0), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			false,
			"provided fee < minimum global fee",
		},
		{
			"invalid dynamic tx with MinGasPrices > BaseFee, MinGasPrices > EffectivePrice",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(100)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				feemarketParams := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				feemarketParams.BaseFee = sdkmath.NewInt(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, feemarketParams)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(1000), big.NewInt(0), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			false,
			"provided fee < minimum global fee",
		},
		{
			"valid dynamic tx with MinGasPrices > BaseFee, MinGasPrices < EffectivePrice (big GasTipCap)",
			func() sdk.Tx {
				params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				params.MinGasPrice = sdk.NewDec(100)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)

				feemarketParams := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
				feemarketParams.BaseFee = sdkmath.NewInt(10)
				suite.app.FeeMarketKeeper.SetParams(suite.ctx, feemarketParams)

				msg := suite.BuildTestEthTx(from, to, nil, make([]byte, 0), nil, big.NewInt(1000), big.NewInt(101), &emptyAccessList)
				return suite.CreateTestTx(msg, privKey, 1, false)
			},
			true,
			"",
		},
	}

	for _, et := range execTypes {
		for _, tc := range testCases {
			suite.Run(et.name+"_"+tc.name, func() {
				// suite.SetupTest(et.isCheckTx)
				suite.SetupTest()
				dec := ante.NewEthMinGasPriceDecorator(suite.app.FeeMarketKeeper, suite.app.EvmKeeper)
				_, err := dec.AnteHandle(suite.ctx, tc.malleate(), et.simulate, NextFn)

				if tc.expPass {
					suite.Require().NoError(err, tc.name)
				} else {
					suite.Require().Error(err, tc.name)
					suite.Require().Contains(err.Error(), tc.errMsg, tc.name)
				}
			})
		}
	}
}
