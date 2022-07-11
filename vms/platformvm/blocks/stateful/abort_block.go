// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stateful

// var _ Block = &AbortBlock{}

// // AbortBlock being accepted results in the proposal of its parent (which must
// // be a proposal block) being rejected.
// type AbortBlock struct {
// 	*stateless.AbortBlock
// 	*commonBlock
// }

// // NewAbortBlock returns a new *AbortBlock where the block's parent, a proposal
// // block, has ID [parentID]. Additionally the block will track if it was
// // originally preferred or not for metrics.
// func NewAbortBlock(
// 	manager Manager,
// 	parentID ids.ID,
// 	height uint64,
// ) (*AbortBlock, error) {
// 	statelessBlk, err := stateless.NewAbortBlock(parentID, height)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return toStatefulAbortBlock(
// 		statelessBlk,
// 		manager,
// 		choices.Processing,
// 	)
// }

// func toStatefulAbortBlock(
// 	statelessBlk *stateless.AbortBlock,
// 	manager Manager,
// 	status choices.Status,
// ) (*AbortBlock, error) {
// 	abort := &AbortBlock{
// 		AbortBlock: statelessBlk,
// 		commonBlock: &commonBlock{
// 			Manager: manager,
// 			baseBlk: &statelessBlk.CommonBlock,
// 		},
// 	}

// 	return abort, nil
// }

// func (a *AbortBlock) Verify() error {
// 	return a.VerifyAbortBlock(a.AbortBlock)
// }

// func (a *AbortBlock) Accept() error {
// 	return a.AcceptAbortBlock(a.AbortBlock)
// }

// func (a *AbortBlock) Reject() error {
// 	return a.RejectAbortBlock(a.AbortBlock)
// }

// func (a *AbortBlock) conflicts(s ids.Set) (bool, error) {
// 	return a.conflictsAbortBlock(a, s)
// }

// func (a *AbortBlock) setBaseState() {
// 	a.setBaseStateAbortBlock(a)
// }