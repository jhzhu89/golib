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

// next moves an iterator forward.
func (it *ForwardListIter) next() *ForwardListIter {
	it.Next()
	return it
}

func (it *ForwardListIter) makeNext() *ForwardListIter {
	var it_ = &ForwardListIter{}
	if it.node != nil {
		it_.node = it.node.next
	}
	return it_
}

// Clone returns a copy of this iterator.
func (it *ForwardListIter) Clone() IterRef {
	return &ForwardListIter{it.node}
}

// clone returns a copy of it.
func (it *ForwardListIter) clone() *ForwardListIter {
	return &ForwardListIter{it.node}
}

// EqualTo checks if given iterator is equal to this iterator.
func (it *ForwardListIter) EqualTo(r IterCRef) bool {
	return *it == *r.(*ForwardListIter)
}
