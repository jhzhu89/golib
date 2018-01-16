// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package fwdlist implements a c++ STL-like forward list.
//
// To iterate over a vector (where fl is a *ForwardList):
//	for it := fl.Begin(); !it.Equal(fl.End()); it.Next()) {
//		val := it.Deref()
//		// do something with val
//	}
package fwdlist

import (
	"github.com/jhzhu89/golib/algorithm"
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/iterator"
	"github.com/jhzhu89/golib/util"
)

// Type aliases.
type (
	Value = container.Value

	IterRef     = iterator.IterRef
	IterCRef    = iterator.IterCRef
	InputIter   = iterator.InputIter
	RandIter    = iterator.RandIter
	ForwardIter = iterator.ForwardIter
	ReverseIter = iterator.ReverseIterator

	Any = interface{}
)

type Comparator interface {
	compare(a, b Any) bool
}

//type Predicator func(v Any) bool
//
//func (p Predicator) Predicate(v Any) bool {
//	return p.Predicate(v)
//}
//
//type BinaryPredicator func(a, b Any) bool
//
//func (p BinaryPredicator) Predicate(a, b Any) bool {
//	return p.Predicate(a, b)
//}

// ForwardList represents a c++ STL-like vector. It is a standard container
// which offers fixed time access to individual elements in any order.
type ForwardList struct {
	fwdListImpl
}

// New creates a ForwardList with no elements.
func New() *ForwardList {

}

// NewN creates a ForwardList with n nil elements.
func NewN(n int) *ForwardList {

}

// NewNValues fills the ForwardList with n copies of val.
func NewNValues(n int, val Value) *ForwardList {

}

// NewFromRange creates a ForwardList consisting of copies of the
// elements from [first, last).
func NewFromRange(first, last InputIter) *ForwardList {

}

// Iterators

// Begin returns a read/write iterator that points to the first element in the
// ForwardList. Iteration is done in ordinary element order.
func (fl *ForwardList) Begin() *ForwardListIter {

}

// End returns a read/write iterator that points one past the last
// element in the ForwardList. Iteration is done in ordinary
// element order.
func (fl *ForwardList) End() *ForwardListIter {

}

// BeforeBegin returns a read/write iterator that points before the first element in the
// ForwardList. Iteration is done in ordinary element order.
func (fl *ForwardList) BeforeBegin() *ForwardListIter {

}

// Element access

// At accesses data contained in the ForwardList by subscript.
func (fl *ForwardList) At(n int) Value {

}

// Front returns the data at the first element of the ForwardList.
func (fl *ForwardList) Front() Value {

}

// Capacity

// Empty returns true if the ForwardList is empty.
func (fl *ForwardList) Empty() bool {

}

// Size returns the number of elements in the ForwardList.
func (fl *ForwardList) Size() int {

}

// Modifiers

// Clear erases all the elements.
func (fl *ForwardList) Clear() {

}

// InsertAfter inserts given value into ForwardList after specified iterator.
func (fl *ForwardList) InsertAfter(pos *ForwardListIter, val Value) *ForwardListIter {

}

// RangeInsertAfter inserts a range into the ForwardList.
func (fl *ForwardList) RangeInsertAfter(pos *ForwardListIter, first, last InputIter) *ForwardListIter {

}

// FillInsertAfter inserts a number of copies of given data into the ForwardList.
func (fl *ForwardList) FillInsertAfter(pos *ForwardListIter, n int, val Value) *ForwardListIter {

}

// EraseAfter removes the element pointed to by the iterator following pos.
func (fl *ForwardList) EraseAfter(pos *ForwardListIter) *ForwardListIter {

}

// RangeEraseAfter removes a range of elements.
// first is the iterator pointing before the first elements to be erased.
// last is the iterator pointing to one past the last elements to be erased.
func (fl *ForwardList) RangeEraseAfter(first, last *ForwardListIter) *ForwardListIter {

}

// Swap swaps data with another ForwardList.
func (fl *ForwardList) Swap(x *ForwardList) {

}

// PushFront adds data to the front of the ForwardList.
func (fl *ForwardList) PushFront(val Value) {

}

// PopFront removes first element.
func (fl *ForwardList) PopFront() {

}

// Resize resizes the ForwardList to the specified number of elements.
func (fl *ForwardList) Resize(newSize int) {

}

// FillResize resizes the ForwardList to the specified number of elements.
// val is the data with which new elements should be populated.
func (fl *ForwardList) FillResize(newSize int, val Value) {

}

// FillAssign assigns a given value to a ForwardList.
func (fl *ForwardList) FillAssign(size int, val Value) {

}

// RangeAssign assigns a range to a ForwardList.
func (fl *ForwardList) RangeAssign(first, last InputIter) {

}

// Operations

// Merge merges sorted lists according to comparison function.
func (fl *ForwardList) Merge(list *ForwardList, comp Comp) {

}

// SpliceAfter inserts contents of another ForwardList.
func (fl *ForwardList) SpliceAfter(pos *ForwardListIter, list *ForwardList) {

}

// RangeSpliceAfter inserts range from another ForwardList.
func (fl *ForwardList) RangeSpliceAfter(pos, before, last *ForwardListIter) {

}

// Remove removes all elements equal to value.
func (fl *ForwardList) Remove(val Value) {

}

// RemoveIf removes all elements satisfying a predicate.
func (fl *ForwardList) RemoveIf(pred Pred) {

}

// Unique removes consecutive duplicate elements.
func (fl *ForwardList) Unique() {
}

// UniqueIf removes consecutive elements satisfying a predicate.
func (fl *ForwardList) UniqueIf(binPred BinPred) {
}

// Reverse reverses the elements in list.
func (fl *ForwardList) Reverse() {
}

// Sort reverses the elements in list.
func (fl *ForwardList) Sort(comp Comp) {
}

//func (fl *ForwardList) fillAssign(n int, val Value) {
//
//}
//
//func (fl *ForwardList) fillInsert(pos *ForwardListIter, n int, val Value) {
//
//}
//
//func (fl *ForwardList) fillInitialize(n int, val Value) {
//
//}
//
//func (fl *ForwardList) rangeInitialize(first, last InputIter) {
//
//}
//
//func (fl *ForwardList) rangeInsert(pos *ForwardListIter, first, last InputIter) {
//
//}
//
//func (fl *ForwardList) checkLen(n int) int {
//
//}
//
//func (fl *ForwardList) insertAux(pos *ForwardListIter, val Value) {
//
//}
//
//func (fl *ForwardList) defaultAppend(n int) {
//
//}
//
//func (fl *ForwardList) eraseAtEnd(pos *ForwardListIter) {
//
//}
//
//func (fl *ForwardList) erase(pos *ForwardListIter) *ForwardListIter {
//
//}
//
//func (fl *ForwardList) rangeErase(first, last *ForwardListIter) *ForwardListIter {
//
//}
//
//func (fl *ForwardList) assignAux(first, last InputIter) {
//
//}

// fwdListImpl handles vector allocation.
type fwdListImpl struct {
}

type fwdListNode struct {
	next *fwdListNode
	val  Value
}

func (fn *fwdListNode) transferAfter(begin, end *fwdListNode) *fwdListNode {
}

func (fn *fwdListNode) reverseAfter() {
}
