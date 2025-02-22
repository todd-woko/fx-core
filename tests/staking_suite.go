package tests

import (
	"context"
	"math/big"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/functionx/fx-core/v4/client"
	testscontract "github.com/functionx/fx-core/v4/tests/contract"
	fxtypes "github.com/functionx/fx-core/v4/types"
	precompilesstaking "github.com/functionx/fx-core/v4/x/evm/precompiles/staking"
)

type StakingSuite struct {
	Erc20TestSuite
	abi abi.ABI
}

func NewStakingSuite(ts *TestSuite) StakingSuite {
	return StakingSuite{
		Erc20TestSuite: NewErc20TestSuite(ts),
		abi:            precompilesstaking.GetABI(),
	}
}

func (suite *StakingSuite) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(suite.privKey.PubKey().Address())
}

func (suite *StakingSuite) Address() common.Address {
	return common.BytesToAddress(suite.privKey.PubKey().Address())
}

func (suite *StakingSuite) StakingQuery() stakingtypes.QueryClient {
	return suite.GRPCClient().StakingQuery()
}

func (suite *StakingSuite) TransactionOpts(privateKey cryptotypes.PrivKey) *bind.TransactOpts {
	ecdsa, err := crypto.ToECDSA(privateKey.Bytes())
	suite.Require().NoError(err)

	chainId, err := suite.EthClient().ChainID(suite.ctx)
	suite.Require().NoError(err)

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsa, chainId)
	suite.Require().NoError(err)

	auth.GasTipCap = big.NewInt(1e9)
	auth.GasFeeCap = big.NewInt(6e12)
	return auth
}

// DeployStakingContract deploy staking contract
// solc version 0.8.19 https://github.com/ethereum/solidity/releases
// abigen version 1.11.5-stable https://github.com/ethereum/go-ethereum/releases
//
//go:generate solc --abi ../x/evm/precompiles/staking/solidity/staking_test.sol -o ./artifacts --overwrite
//go:generate solc --bin ../x/evm/precompiles/staking/solidity/staking_test.sol -o ./artifacts --overwrite
//go:generate abigen --abi ./artifacts/staking_test.abi --bin ./artifacts/staking_test.bin --type TestStaking --pkg contract --out ./contract/test_staking.go
func (suite *StakingSuite) DeployStakingContract(privKey cryptotypes.PrivKey) (common.Address, common.Hash) {
	stakingBin := fxtypes.MustDecodeHex(testscontract.TestStakingMetaData.Bin)
	return suite.DeployContract(privKey, stakingBin)
}

// DelegationRewards Get delegatorAddress rewards
func (suite *StakingSuite) DelegationRewards(delAddr, valAddr string) sdk.DecCoins {
	response, err := suite.GRPCClient().DistrQuery().DelegationRewards(suite.ctx, &distrtypes.QueryDelegationRewardsRequest{DelegatorAddress: delAddr, ValidatorAddress: valAddr})
	suite.Require().NoError(err)
	return response.Rewards
}

func (suite *StakingSuite) SetWithdrawAddress(delAddr, withdrawAddr sdk.AccAddress) {
	setWithdrawAddress := distrtypes.NewMsgSetWithdrawAddress(delAddr, withdrawAddr)
	txResponse := suite.BroadcastTx(suite.privKey, setWithdrawAddress)
	suite.Require().True(txResponse.Code == 0)
	response, err := suite.GRPCClient().DistrQuery().DelegatorWithdrawAddress(suite.ctx, &distrtypes.QueryDelegatorWithdrawAddressRequest{DelegatorAddress: delAddr.String()})
	suite.Require().NoError(err)
	suite.Require().EqualValues(response.WithdrawAddress, withdrawAddr.String())
}

func (suite *StakingSuite) Delegate(privateKey cryptotypes.PrivKey, valAddr string, delAmount *big.Int) *ethtypes.Receipt {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack("delegate", valAddr)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, delAmount, pack)
	suite.Require().NoError(err)
	return suite.SendTransaction(transaction)
}

func (suite *StakingSuite) DelegateByContract(privateKey cryptotypes.PrivKey, contract common.Address, valAddr string, delAmount *big.Int) *ethtypes.Receipt {
	stakingContract, err := testscontract.NewTestStaking(contract, suite.EthClient())
	suite.Require().NoError(err)

	auth := suite.TransactionOpts(privateKey)
	auth.Value = delAmount

	tx, err := stakingContract.Delegate(auth, valAddr)
	suite.Require().NoError(err)

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, suite.EthClient(), tx)
	suite.Require().NoError(err)
	suite.Require().Equal(receipt.Status, ethtypes.ReceiptStatusSuccessful)
	return receipt
}

func (suite *StakingSuite) WithdrawByContract(privateKey cryptotypes.PrivKey, contract common.Address, valAddr string) *ethtypes.Receipt {
	stakingContract, err := testscontract.NewTestStaking(contract, suite.EthClient())
	suite.Require().NoError(err)

	auth := suite.TransactionOpts(privateKey)

	tx, err := stakingContract.Withdraw(auth, valAddr)
	suite.Require().NoError(err)

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, suite.EthClient(), tx)
	suite.Require().NoError(err)
	suite.Require().Equal(receipt.Status, ethtypes.ReceiptStatusSuccessful)
	return receipt
}

func (suite *StakingSuite) UndelegateByContract(privateKey cryptotypes.PrivKey, contract common.Address, valAddr string, shares *big.Int) *ethtypes.Receipt {
	stakingContract, err := testscontract.NewTestStaking(contract, suite.EthClient())
	suite.Require().NoError(err)

	auth := suite.TransactionOpts(privateKey)

	tx, err := stakingContract.Undelegate(auth, valAddr, shares)
	suite.Require().NoError(err)

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, suite.EthClient(), tx)
	suite.Require().NoError(err)
	suite.Require().Equal(receipt.Status, ethtypes.ReceiptStatusSuccessful)
	return receipt
}

func (suite *StakingSuite) UnDelegate(privateKey cryptotypes.PrivKey, valAddr string, shares *big.Int) {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.UndelegateMethodName, valAddr, shares)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, nil, pack)
	suite.Require().NoError(err)
	suite.SendTransaction(transaction)
}

func (suite *StakingSuite) WithdrawReward(privateKey cryptotypes.PrivKey, valAddr string) {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.WithdrawMethodName, valAddr)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, nil, pack)
	suite.Require().NoError(err)
	suite.SendTransaction(transaction)
}

func (suite *StakingSuite) Delegation(valAddr string, delAddr common.Address) (*big.Int, *big.Int) {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.DelegationMethodName, valAddr, delAddr)
	suite.Require().NoError(err)
	output, err := suite.EthClient().CallContract(suite.ctx, ethereum.CallMsg{To: &stakingContract, Data: pack}, nil)
	suite.Require().NoError(err)
	var out []interface{}
	res, err := suite.abi.Unpack(precompilesstaking.DelegationMethodName, output)
	suite.Require().NoError(err)
	out = res
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	return out0, out1
}

func (suite *StakingSuite) Rewards(valAddr string, delAddr common.Address) *big.Int {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.DelegationRewardsMethodName, valAddr, delAddr)
	suite.Require().NoError(err)
	output, err := suite.EthClient().CallContract(suite.ctx, ethereum.CallMsg{To: &stakingContract, Data: pack}, nil)
	suite.Require().NoError(err)
	var out []interface{}
	res, err := suite.abi.Unpack(precompilesstaking.DelegationRewardsMethodName, output)
	suite.Require().NoError(err)
	out = res
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0
}

func (suite *StakingSuite) TransferShares(privateKey cryptotypes.PrivKey, valAddr string, receipt common.Address, shares *big.Int) *ethtypes.Receipt {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.TransferSharesMethodName, valAddr, receipt, shares)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, nil, pack)
	suite.Require().NoError(err)
	return suite.SendTransaction(transaction)
}

func (suite *StakingSuite) TransferSharesByContract(privateKey cryptotypes.PrivKey, valAddr string, contract, to common.Address, shares *big.Int) *ethtypes.Receipt {
	stakingContract, err := testscontract.NewTestStaking(contract, suite.EthClient())
	suite.Require().NoError(err)

	auth := suite.TransactionOpts(privateKey)

	tx, err := stakingContract.TransferShares(auth, valAddr, to, shares)
	suite.Require().NoError(err)

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, suite.EthClient(), tx)
	suite.Require().NoError(err)
	suite.Require().Equal(receipt.Status, ethtypes.ReceiptStatusSuccessful)
	return receipt
}

func (suite *StakingSuite) TransferFromShares(privateKey cryptotypes.PrivKey, valAddr string, from, receipt common.Address, shares *big.Int) *ethtypes.Receipt {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.TransferFromSharesMethodName, valAddr, from, receipt, shares)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, nil, pack)
	suite.Require().NoError(err)
	return suite.SendTransaction(transaction)
}

func (suite *StakingSuite) TransferFromSharesByContract(privateKey cryptotypes.PrivKey, valAddr string, contract, from, to common.Address, shares *big.Int) *ethtypes.Receipt {
	stakingContract, err := testscontract.NewTestStaking(contract, suite.EthClient())
	suite.Require().NoError(err)

	auth := suite.TransactionOpts(privateKey)

	tx, err := stakingContract.TransferFromShares(auth, valAddr, from, to, shares)
	suite.Require().NoError(err)

	ctx, cancel := context.WithTimeout(suite.ctx, 5*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, suite.EthClient(), tx)
	suite.Require().NoError(err)
	suite.Require().Equal(receipt.Status, ethtypes.ReceiptStatusSuccessful)
	return receipt
}

func (suite *StakingSuite) ApproveShares(privateKey cryptotypes.PrivKey, valAddr string, spender common.Address, shares *big.Int) *ethtypes.Receipt {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.ApproveSharesMethodName, valAddr, spender, shares)
	suite.Require().NoError(err)
	transaction, err := client.BuildEthTransaction(suite.ctx, suite.EthClient(), privateKey, &stakingContract, nil, pack)
	suite.Require().NoError(err)
	return suite.SendTransaction(transaction)
}

func (suite *StakingSuite) AllowanceShares(valAddr string, owner, spender common.Address) *big.Int {
	stakingContract := precompilesstaking.GetAddress()
	pack, err := precompilesstaking.GetABI().Pack(precompilesstaking.AllowanceSharesMethodName, valAddr, owner, spender)
	suite.Require().NoError(err)
	output, err := suite.EthClient().CallContract(suite.ctx, ethereum.CallMsg{To: &stakingContract, Data: pack}, nil)
	suite.Require().NoError(err)
	var out []interface{}
	res, err := suite.abi.Unpack(precompilesstaking.AllowanceSharesMethodName, output)
	suite.Require().NoError(err)
	out = res
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0
}

func (suite *StakingSuite) LogReward(logs []*ethtypes.Log, valAddr string, addr common.Address) *big.Int {
	for _, log := range logs {
		if log.Address == precompilesstaking.GetAddress() &&
			log.Topics[0] == precompilesstaking.WithdrawEvent.ID &&
			log.Topics[1] == addr.Hash() {
			unpack, err := precompilesstaking.WithdrawEvent.Inputs.NonIndexed().Unpack(log.Data)
			suite.Require().NoError(err)
			suite.Require().Equal(unpack[0].(string), valAddr)
			return unpack[1].(*big.Int)
		}
	}
	return big.NewInt(0)
}
