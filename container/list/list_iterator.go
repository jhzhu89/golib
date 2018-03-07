// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package list

var _ BidirectIter = (*ListIter)(nil)

// ListIter implements a bidirectional iterator.
type ListIter struct {
	node *listNode
}

// CanMultiPass indicates this is a bidirectional iterator.
func (it *ListIter) CanMultiPass() {}

// CopyAssign assgins given iterator to this iterator.
func (it *ListIter) CopyAssign(r IterCRef) {
	*it = *(r.(*ListIter))
}

// Swap swaps data with another iterator.
func (it *ListIter) Swap(r IterCRef) {
	var r_ = r.(*ListIter)
	*it, *r_ = *r_, *it
}

// Deref is like the derefference operation (*it) in c++ to get the value at this position.
func (it *ListIter) Deref() Value {
	return it.node.val
}

// DerefSet is like the derefference operation (*it=val) in c++ to
// set value to the element at this position.
func (it *ListIter) DerefSet(val Value) {
	it.node.val = val
}

// Next is like the pointer arithmetic it++ in c++ to increment the iterator.
func (it *ListIter) Next() {
	it.node = it.node.next
}

// Next2 moves an iterator forward.
func (it *ListIter) Next2() *ListIter {
	it.Next()
	return it
}

// Prev is like the pointer arithmetic it-- in c++ to decrement the iterator.
func (it *ListIter) Prev() {
	it.node = it.node.prev
}

// Prev2 moves an iterator backward.
func (it *ListIter) Prev2() *ListIter {
	it.Prev()
	return it
}

// Clone returns a copy of this iterator.
func (it *ListIter) Clone() IterRef {
	return &ListIter{it.node}
}

// Clone2 returns a copy of it.
func (it *ListIter) Clone2() *ListIter {
	return &ListIter{it.node}
}

// EqualTo checks if given iterator is equal to this iterator.
func (it *ListIter) EqualTo(r IterCRef) bool {
	return *it == *r.(*ListIter)
}

func (it *ListIter) next() *ListIter {
	var it_ = &ListIter{}
	if it.node != nil {
		it_.node = it.node.next
	}
	return it_
}
