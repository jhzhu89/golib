// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package fwdlist

var _ ForwardIter = (*ForwardListIter)(nil)

// ForwardListIter implements a forward iterator.
type ForwardListIter struct {
	node *fwdListNode
}

// CanMultiPass indicates this is a forward iterator.
func (it *ForwardListIter) CanMultiPass() {}

// CopyAssign assgins given iterator to this iterator.
func (it *ForwardListIter) CopyAssign(r IterCRef) {
	*it = *(r.(*ForwardListIter))
}

// Swap swaps data with another iterator.
func (it *ForwardListIter) Swap(r IterCRef) {
	var r_ = r.(*ForwardListIter)
	*it, *r_ = *r_, *it
}

// Deref is like the derefference operation (*it) in c++ to get the value at this position.
func (it *ForwardListIter) Deref() Value {
	return it.node.val
}

// DerefSet is like the derefference operation (*it=val) in c++ to
// set value to the element at this position.
func (it *ForwardListIter) DerefSet(val Value) {
	it.node.val = val
}

// Next is like the pointer arithmetic it++ in c++ to increment the iterator.
func (it *ForwardListIter) Next() {
	it.node = it.node.next
}

// Next2 moves an iterator forward.
func (it *ForwardListIter) Next2() *ForwardListIter {
	it.Next()
	return it
}

// Clone returns a copy of this iterator.
func (it *ForwardListIter) Clone() IterRef {
	return &ForwardListIter{it.node}
}

// Clone2 returns a copy of it.
func (it *ForwardListIter) Clone2() *ForwardListIter {
	return &ForwardListIter{it.node}
}

// Equal checks if given iterator is equal to this iterator.
func (it *ForwardListIter) Equal(r IterCRef) bool {
	return *it == *r.(*ForwardListIter)
}
