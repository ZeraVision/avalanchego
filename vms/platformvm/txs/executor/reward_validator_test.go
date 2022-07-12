// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/stretchr/testify/assert"
)

func TestRewardValidatorTxExecuteOnCommit(t *testing.T) {
	h := newTestHelpersCollection()
	defer func() {
		if err := internalStateShutdown(h); err != nil {
			t.Fatal(err)
		}
	}()
	dummyHeight := uint64(1)

	currentStakers := h.tState.CurrentStakers()
	toRemoveTx, _, err := currentStakers.GetNextStaker()
	if err != nil {
		t.Fatal(err)
	}
	toRemoveTxID := toRemoveTx.ID()
	toRemove := toRemoveTx.Unsigned.(*txs.AddValidatorTx)

	// Case 1: Chain timestamp is wrong
	tx, err := h.txBuilder.NewRewardValidatorTx(toRemoveTxID)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor := ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err == nil {
		t.Fatalf("should have failed because validator end time doesn't match chain timestamp")
	}

	// Advance chain timestamp to time that next validator leaves
	h.tState.SetTimestamp(toRemove.EndTime())

	// Case 2: Wrong validator
	tx, err = h.txBuilder.NewRewardValidatorTx(ids.GenerateTestID())
	if err != nil {
		t.Fatal(err)
	}

	txExecutor = ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err == nil {
		t.Fatalf("should have failed because validator ID is wrong")
	}

	// Case 3: Happy path
	tx, err = h.txBuilder.NewRewardValidatorTx(toRemoveTxID)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor = ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err != nil {
		t.Fatal(err)
	}

	onCommitCurrentStakers := txExecutor.OnCommit.CurrentStakers()
	nextToRemoveTx, _, err := onCommitCurrentStakers.GetNextStaker()
	if err != nil {
		t.Fatal(err)
	}
	if toRemoveTxID == nextToRemoveTx.ID() {
		t.Fatalf("Should have removed the previous validator")
	}

	// check that stake/reward is given back
	stakeOwners := toRemove.Stake[0].Out.(*secp256k1fx.TransferOutput).AddressesSet()

	// Get old balances
	oldBalance, err := avax.GetBalance(h.tState, stakeOwners)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor.OnCommit.Apply(h.tState)
	h.tState.SetHeight(dummyHeight)
	if err := h.tState.Commit(); err != nil {
		t.Fatal(err)
	}

	onCommitBalance, err := avax.GetBalance(h.tState, stakeOwners)
	if err != nil {
		t.Fatal(err)
	}

	if onCommitBalance != oldBalance+toRemove.Validator.Weight()+27 {
		t.Fatalf("on commit, should have old balance (%d) + staked amount (%d) + reward (%d) but have %d",
			oldBalance, toRemove.Validator.Weight(), 27, onCommitBalance)
	}
}

func TestRewardValidatorTxExecuteOnAbort(t *testing.T) {
	h := newTestHelpersCollection()
	defer func() {
		if err := internalStateShutdown(h); err != nil {
			t.Fatal(err)
		}
	}()
	dummyHeight := uint64(1)

	currentStakers := h.tState.CurrentStakers()
	toRemoveTx, _, err := currentStakers.GetNextStaker()
	if err != nil {
		t.Fatal(err)
	}
	toRemoveTxID := toRemoveTx.ID()
	toRemove := toRemoveTx.Unsigned.(*txs.AddValidatorTx)

	// Case 1: Chain timestamp is wrong
	tx, err := h.txBuilder.NewRewardValidatorTx(toRemoveTxID)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor := ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err == nil {
		t.Fatalf("should have failed because validator end time doesn't match chain timestamp")
	}

	// Advance chain timestamp to time that next validator leaves
	h.tState.SetTimestamp(toRemove.EndTime())

	// Case 2: Wrong validator
	tx, err = h.txBuilder.NewRewardValidatorTx(ids.GenerateTestID())
	if err != nil {
		t.Fatal(err)
	}

	txExecutor = ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err == nil {
		t.Fatalf("should have failed because validator ID is wrong")
	}

	// Case 3: Happy path
	tx, err = h.txBuilder.NewRewardValidatorTx(toRemoveTxID)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor = ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	if err != nil {
		t.Fatal(err)
	}

	onAbortCurrentStakers := txExecutor.OnAbort.CurrentStakers()
	nextToRemoveTx, _, err := onAbortCurrentStakers.GetNextStaker()
	if err != nil {
		t.Fatal(err)
	}
	if toRemoveTxID == nextToRemoveTx.ID() {
		t.Fatalf("Should have removed the previous validator")
	}

	// check that stake/reward isn't given back
	stakeOwners := toRemove.Stake[0].Out.(*secp256k1fx.TransferOutput).AddressesSet()

	// Get old balances
	oldBalance, err := avax.GetBalance(h.tState, stakeOwners)
	if err != nil {
		t.Fatal(err)
	}

	txExecutor.OnAbort.Apply(h.tState)
	h.tState.SetHeight(dummyHeight)
	if err := h.tState.Commit(); err != nil {
		t.Fatal(err)
	}

	onAbortBalance, err := avax.GetBalance(h.tState, stakeOwners)
	if err != nil {
		t.Fatal(err)
	}

	if onAbortBalance != oldBalance+toRemove.Validator.Weight() {
		t.Fatalf("on abort, should have old balance (%d) + staked amount (%d) but have %d",
			oldBalance, toRemove.Validator.Weight(), onAbortBalance)
	}
}

func TestRewardDelegatorTxExecuteOnCommit(t *testing.T) {
	assert := assert.New(t)
	h := newTestHelpersCollection()
	defer func() {
		if err := internalStateShutdown(h); err != nil {
			t.Fatal(err)
		}
	}()
	dummyHeight := uint64(1)

	vdrRewardAddress := ids.GenerateTestShortID()
	delRewardAddress := ids.GenerateTestShortID()

	vdrStartTime := uint64(defaultValidateStartTime.Unix()) + 1
	vdrEndTime := uint64(defaultValidateStartTime.Add(2 * defaultMinStakingDuration).Unix())
	vdrNodeID := ids.GenerateTestNodeID()

	vdrTx, err := h.txBuilder.NewAddValidatorTx(
		h.cfg.MinValidatorStake, // stakeAmt
		vdrStartTime,
		vdrEndTime,
		vdrNodeID,        // node ID
		vdrRewardAddress, // reward address
		reward.PercentDenominator/4,
		[]*crypto.PrivateKeySECP256K1R{preFundedKeys[0]},
		ids.ShortEmpty,
	)
	assert.NoError(err)

	delStartTime := vdrStartTime
	delEndTime := vdrEndTime

	delTx, err := h.txBuilder.NewAddDelegatorTx(
		h.cfg.MinDelegatorStake,
		delStartTime,
		delEndTime,
		vdrNodeID,
		delRewardAddress,
		[]*crypto.PrivateKeySECP256K1R{preFundedKeys[0]},
		ids.ShortEmpty, // Change address
	)
	assert.NoError(err)

	h.tState.AddCurrentStaker(vdrTx, 0)
	h.tState.AddTx(vdrTx, status.Committed)
	h.tState.AddCurrentStaker(delTx, 1000000)
	h.tState.AddTx(delTx, status.Committed)
	h.tState.SetTimestamp(time.Unix(int64(delEndTime), 0))
	h.tState.SetHeight(dummyHeight)
	assert.NoError(h.tState.Commit())
	err = h.tState.Load()
	assert.NoError(err)
	// test validator stake
	set, ok := h.cfg.Validators.GetValidators(constants.PrimaryNetworkID)
	assert.True(ok)
	stake, ok := set.GetWeight(vdrNodeID)
	assert.True(ok)
	assert.Equal(h.cfg.MinValidatorStake+h.cfg.MinDelegatorStake, stake)

	tx, err := h.txBuilder.NewRewardValidatorTx(delTx.ID())
	assert.NoError(err)

	txExecutor := ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	assert.NoError(err)

	vdrDestSet := ids.ShortSet{}
	vdrDestSet.Add(vdrRewardAddress)
	delDestSet := ids.ShortSet{}
	delDestSet.Add(delRewardAddress)

	expectedReward := uint64(1000000)

	oldVdrBalance, err := avax.GetBalance(h.tState, vdrDestSet)
	assert.NoError(err)
	oldDelBalance, err := avax.GetBalance(h.tState, delDestSet)
	assert.NoError(err)

	txExecutor.OnCommit.Apply(h.tState)
	h.tState.SetHeight(dummyHeight)
	assert.NoError(h.tState.Commit())

	// If tx is committed, delegator and delegatee should get reward
	// and the delegator's reward should be greater because the delegatee's share is 25%
	commitVdrBalance, err := avax.GetBalance(h.tState, vdrDestSet)
	assert.NoError(err)
	vdrReward, err := math.Sub64(commitVdrBalance, oldVdrBalance)
	assert.NoError(err)
	assert.NotZero(vdrReward, "expected delegatee balance to increase because of reward")

	commitDelBalance, err := avax.GetBalance(h.tState, delDestSet)
	assert.NoError(err)
	delReward, err := math.Sub64(commitDelBalance, oldDelBalance)
	assert.NoError(err)
	assert.NotZero(delReward, "expected delegator balance to increase because of reward")

	assert.Less(vdrReward, delReward, "the delegator's reward should be greater than the delegatee's because the delegatee's share is 25%")
	assert.Equal(expectedReward, delReward+vdrReward, "expected total reward to be %d but is %d", expectedReward, delReward+vdrReward)

	stake, ok = set.GetWeight(vdrNodeID)
	assert.True(ok)
	assert.Equal(h.cfg.MinValidatorStake, stake)
}

func TestRewardDelegatorTxExecuteOnAbort(t *testing.T) {
	assert := assert.New(t)
	h := newTestHelpersCollection()
	defer func() {
		if err := internalStateShutdown(h); err != nil {
			t.Fatal(err)
		}
	}()
	dummyHeight := uint64(1)

	initialSupply := h.tState.GetCurrentSupply()

	vdrRewardAddress := ids.GenerateTestShortID()
	delRewardAddress := ids.GenerateTestShortID()

	vdrStartTime := uint64(defaultValidateStartTime.Unix()) + 1
	vdrEndTime := uint64(defaultValidateStartTime.Add(2 * defaultMinStakingDuration).Unix())
	vdrNodeID := ids.GenerateTestNodeID()

	vdrTx, err := h.txBuilder.NewAddValidatorTx(
		h.cfg.MinValidatorStake, // stakeAmt
		vdrStartTime,
		vdrEndTime,
		vdrNodeID,        // node ID
		vdrRewardAddress, // reward address
		reward.PercentDenominator/4,
		[]*crypto.PrivateKeySECP256K1R{preFundedKeys[0]},
		ids.ShortEmpty,
	)
	assert.NoError(err)

	delStartTime := vdrStartTime
	delEndTime := vdrEndTime
	delTx, err := h.txBuilder.NewAddDelegatorTx(
		h.cfg.MinDelegatorStake,
		delStartTime,
		delEndTime,
		vdrNodeID,
		delRewardAddress,
		[]*crypto.PrivateKeySECP256K1R{preFundedKeys[0]},
		ids.ShortEmpty,
	)
	assert.NoError(err)

	h.tState.AddCurrentStaker(vdrTx, 0)
	h.tState.AddTx(vdrTx, status.Committed)
	h.tState.AddCurrentStaker(delTx, 1000000)
	h.tState.AddTx(delTx, status.Committed)
	h.tState.SetTimestamp(time.Unix(int64(delEndTime), 0))
	h.tState.SetHeight(dummyHeight)
	assert.NoError(h.tState.Commit())
	err = h.tState.Load()
	assert.NoError(err)

	tx, err := h.txBuilder.NewRewardValidatorTx(delTx.ID())
	assert.NoError(err)

	txExecutor := ProposalTxExecutor{
		Backend:     &h.execBackend,
		ParentState: h.tState,
		Tx:          tx,
	}
	err = tx.Unsigned.Visit(&txExecutor)
	assert.NoError(err)

	vdrDestSet := ids.ShortSet{}
	vdrDestSet.Add(vdrRewardAddress)
	delDestSet := ids.ShortSet{}
	delDestSet.Add(delRewardAddress)

	expectedReward := uint64(1000000)

	oldVdrBalance, err := avax.GetBalance(h.tState, vdrDestSet)
	assert.NoError(err)
	oldDelBalance, err := avax.GetBalance(h.tState, delDestSet)
	assert.NoError(err)

	txExecutor.OnAbort.Apply(h.tState)
	h.tState.SetHeight(dummyHeight)
	assert.NoError(h.tState.Commit())

	// If tx is aborted, delegator and delegatee shouldn't get reward
	newVdrBalance, err := avax.GetBalance(h.tState, vdrDestSet)
	assert.NoError(err)
	vdrReward, err := math.Sub64(newVdrBalance, oldVdrBalance)
	assert.NoError(err)
	assert.Zero(vdrReward, "expected delegatee balance not to increase")

	newDelBalance, err := avax.GetBalance(h.tState, delDestSet)
	assert.NoError(err)
	delReward, err := math.Sub64(newDelBalance, oldDelBalance)
	assert.NoError(err)
	assert.Zero(delReward, "expected delegator balance not to increase")

	newSupply := h.tState.GetCurrentSupply()
	assert.Equal(initialSupply-expectedReward, newSupply, "should have removed un-rewarded tokens from the potential supply")
}