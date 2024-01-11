package v19_test

import (
	"fmt"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/CosmosContracts/juno/v19/app/apptesting"
	decorators "github.com/CosmosContracts/juno/v19/app/decorators"
	v19 "github.com/CosmosContracts/juno/v19/app/upgrades/v19"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) setupMockCore1MultisigAccount() *vestingtypes.PeriodicVestingAccount {
	core1Multisig := v19.CreateMainnetVestingAccount(s.Ctx, s.App.AppKeepers)
	s.App.AppKeepers.AccountKeeper.SetAccount(s.Ctx, core1Multisig)

	return core1Multisig
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()
	preUpgradeChecks(s)

	// Core-1 account mock up
	va := s.setupMockCore1MultisigAccount()
	// fmt.Println(va)

	v := s.App.AppKeepers.StakingKeeper.GetAllValidators(s.Ctx)
	fmt.Printf("Validators: %d", len(v))

	val1 := v[0]
	s.AllocateRewardsToValidator(val1.GetOperator(), sdk.NewInt(1_000_000))

	s.Ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + 10)
	s.Require().NotPanics(func() {
		s.App.BeginBlocker(s.Ctx, abci.RequestBeginBlock{})
	})

	// delegate to a validator
	s.DelegateToValidator(val1.GetOperator(), va.GetAddress(), sdk.NewInt(1))

	// get delegations
	dels := s.App.AppKeepers.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, va.GetAddress())
	s.Require().Equal(1, len(dels))

	// upgrade
	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v19.UpgradeName, upgradeHeight)
	postUpgradeChecks(s)

	// TODO: check it was modified
	updatedAcc := s.App.AppKeepers.AccountKeeper.GetAccount(s.Ctx, va.GetAddress())
	fmt.Println(updatedAcc)
}

func preUpgradeChecks(s *UpgradeTestSuite) {
	// Change Rate Decorator Test
	// Create a validator with a max change rate of 20%
	for i := 0; i < 100; i++ {

		// Create validator keys & desc
		valPub := secp256k1.GenPrivKey().PubKey()
		valAddr := sdk.ValAddress(valPub.Address())
		description := stakingtypes.NewDescription(fmt.Sprintf("test_moniker%d", i), "", "", "", "")
		validator, err := stakingtypes.NewValidator(
			valAddr,
			valPub,
			description,
		)
		s.Require().NoError(err)

		// Set validator commission
		changeRate := "1.00"
		if i < 100 {
			changeRate = fmt.Sprintf("0.%02d", i)
		}
		validator.Commission.MaxChangeRate.Set(sdk.MustNewDecFromStr(changeRate))

		// Set validator in kv store
		s.App.AppKeepers.StakingKeeper.SetValidator(s.Ctx, validator)
	}
}

func postUpgradeChecks(s *UpgradeTestSuite) {
	// Change Rate Decorator Test
	// Ensure all validators have a max change rate of 5%
	validators := s.App.AppKeepers.StakingKeeper.GetAllValidators(s.Ctx)
	for _, validator := range validators {
		s.Require().True(validator.Commission.MaxChangeRate.LTE(sdk.MustNewDecFromStr(decorators.MaxChangeRate)))
	}

}
