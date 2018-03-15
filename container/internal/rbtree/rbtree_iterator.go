// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rbtree

var _ BidirectIter = (*RbTreeIter)(nil)

// RbTreeIter implements a bidirectional iterator.
type RbTreeIter struct {
	//node *listNode
}

// CanMultiPass indicates this is a bidirectional iterator.
func (it *RbTreeIter) CanMultiPass() {}

// CopyAssign assgins given iterator to this iterator.
func (it *RbTreeIter) CopyAssign(r IterCRef) {
	*it = *(r.(*RbTreeIter))
}

// Swap swaps data with another iterator.
func (it *RbTreeIter) Swap(r IterCRef) {
	var r_ = r.(*RbTreeIter)
	*it, *r_ = *r_, *it
}

// Deref is like the derefference operation (*it) in c++ to get the value at this position.
func (it *RbTreeIter) Deref() Value {
	return it.node.val
}

// DerefSet is like the derefference operation (*it=val) in c++ to
// set value to the element at this position.
func (it *RbTreeIter) DerefSet(val Value) {
	it.node.val = val
}

// Next is like the pointer arithmetic it++ in c++ to increment the iterator.
func (it *RbTreeIter) Next() {
	it.node = it.node.next
}

// Next2 moves an iterator forward.
func (it *RbTreeIter) Next2() *RbTreeIter {
	it.Next()
	return it
}

// Prev is like the pointer arithmetic it-- in c++ to decrement the iterator.
func (it *RbTreeIter) Prev() {
	it.node = it.node.prev
}

// Prev2 moves an iterator backward.
func (it *RbTreeIter) Prev2() *RbTreeIter {
	it.Prev()
	return it
}

// Clone returns a copy of this iterator.
func (it *RbTreeIter) Clone() IterRef {
	return &RbTreeIter{it.node}
}

// Clone2 returns a copy of it.
func (it *RbTreeIter) Clone2() *RbTreeIter {
	return &RbTreeIter{it.node}
}

// EqualTo checks if given iterator is equal to this iterator.
func (it *RbTreeIter) EqualTo(r IterCRef) bool {
	return *it == *r.(*RbTreeIter)
}

//func (it *RbTreeIter) next() *RbTreeIter {
//	var it_ = &RbTreeIter{}
//	if it.node != nil {
//		it_.node = it.node.next
//	}
//	return it_
//}
