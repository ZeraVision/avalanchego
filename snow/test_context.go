// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snow

import (
	"testing"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/prometheus/client_golang/prometheus"
)

func DefaultContextTest() *Context {
	return &Context{
		NetworkID:    0,
		SubnetID:     ids.Empty,
		ChainID:      ids.Empty,
		NodeID:       ids.EmptyNodeID,
		Log:          logging.NoLog{},
		BCLookup:     ids.NewAliaser(),
		Metrics:      metrics.NewOptionalGatherer(),
		ChainDataDir: "",
	}
}

func DefaultConsensusContextTest(t *testing.T) *ConsensusContext {
	var currentState State = Initializing
	return &ConsensusContext{
		Context:             DefaultContextTest(),
		Registerer:          prometheus.NewRegistry(),
		AvalancheRegisterer: prometheus.NewRegistry(),
		DecisionAcceptor:    noOpAcceptor{},
		ConsensusAcceptor:   noOpAcceptor{},
		SubnetStateTracker: &SubnetStateTrackerTest{
			T: t,
			IsSubnetSyncedF: func() bool {
				return currentState == NormalOp
			},
			SetStateF: func(chainID ids.ID, state State) {
				currentState = state
			},
			GetStateF: func(chainID ids.ID) State {
				return currentState
			},
		},
	}
}