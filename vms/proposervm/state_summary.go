// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/vms/proposervm/summary"
)

var _ block.StateSummary = &stateSummary{}

// stateSummary implements block.StateSummary by layering three objects:
// 1. [statelessSummary] carries all summary marshallable content along with
//    data immediately retrievable from it.
// 2. [innerSummary] reports the height of the summary as well as notifying the
//    inner vm of the summary's acceptance.
// 3. [block] is used to update the proposervm's last accepted block upon
//    Accept.
//
// Note: summary.StatelessSummary contains the data to build both [innerSummary]
//       and [block].
type stateSummary struct {
	statelessSummary summary.StateSummary

	// inner summary, retrieved via Parse
	innerSummary block.StateSummary

	// block associated with the summary
	block Block
}

func (s *stateSummary) ID() ids.ID {
	return s.statelessSummary.ID()
}

func (s *stateSummary) Height() uint64 {
	return s.innerSummary.Height()
}

func (s *stateSummary) Bytes() []byte {
	return s.statelessSummary.Bytes()
}

func (s *stateSummary) Accept() (bool, error) {
	// a statefulSummary carries the full proposerVM block associated
	// with the summary. We store this block and update height index with it,
	// so that state sync could resume after a shutdown.
	if err := s.block.acceptOuterBlk(); err != nil {
		return false, err
	}
	return s.innerSummary.Accept()
}