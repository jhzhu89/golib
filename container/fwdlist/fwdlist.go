// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package fwdlist implements a c++ STL-like forward list.
//
// To iterate over a fwdlist (where fl is a *ForwardList):
//	for it := fl.Begin(); !it.Equal(fl.End()); it.Next()) {
//		val := it.Deref()
//		// do something with val
//	}
package fwdlist

import (
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/fn"
	"github.com/jhzhu89/golib/iterator"
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
)

// ForwardList represents a c++ STL-like forward list. It is a standard container with linear
// time access to elements, and fixed time insertion/deletion at any point in the sequence.
type ForwardList struct {
	head *fwdListNode
}

// New creates a ForwardList with no elements.
func New() *ForwardList {
	return &ForwardList{&fwdListNode{}}
}

// NewN creates a ForwardList with n nil elements.
func NewN(n int) *ForwardList {
	fl := &ForwardList{&fwdListNode{}}
	fl.defaultInitialize(n)
	return fl
}

// NewNValues fills the ForwardList with n copies of val.
func NewNValues(n int, val Value) *ForwardList {
	fl := &ForwardList{&fwdListNode{}}
	fl.fillInitialize(n, val)
	return fl
}

// NewFromRange creates a ForwardList consisting of copies of the
// elements from [first, last).
func NewFromRange(first, last InputIter) *ForwardList {
	fl := &ForwardList{&fwdListNode{}}
	fl.rangeInitialize(first, last)
	return fl
}

// Iterators

// Begin returns a read/write iterator that points to the first element in the
// ForwardList. Iteration is done in ordinary element order.
func (fl *ForwardList) Begin() *ForwardListIter {
	return &ForwardListIter{fl.head.next}
}

// End returns a read/write iterator that points one past the last
// element in the ForwardList. Iteration is done in ordinary
// element order.
func (fl *ForwardList) End() *ForwardListIter {
	return &ForwardListIter{}
}

// BeforeBegin returns a read/write iterator that points before the first element in the
// ForwardList. Iteration is done in ordinary element order.
func (fl *ForwardList) BeforeBegin() *ForwardListIter {
	return &ForwardListIter{fl.head}
}

// Element access

// Front returns the data at the first element of the ForwardList.
func (fl *ForwardList) Front() Value {
	return fl.head.next.val
}

// Capacity

// Empty returns true if the ForwardList is empty.
func (fl *ForwardList) Empty() bool {
	return fl.head.next == nil
}

// Modifiers

// Clear erases all the elements.
func (fl *ForwardList) Clear() {
	fl.rangeEraseAfter(fl.head, nil)
}

// InsertAfter inserts given value into ForwardList after specified iterator.
func (fl *ForwardList) InsertAfter(pos *ForwardListIter, val Value) *ForwardListIter {
	return &ForwardListIter{fl.insertAfter(pos, val)}
}

// RangeInsertAfter inserts a range into the ForwardList.
func (fl *ForwardList) RangeInsertAfter(pos *ForwardListIter, first, last InputIter) *ForwardListIter {
	var tmp = NewFromRange(first, last)
	if !tmp.Empty() {
		return fl.spliceAfter(pos, tmp.BeforeBegin(), tmp.End())
	} else {
		return pos.clone()
	}
}

// FillInsertAfter inserts a number of copies of given data into the ForwardList.
func (fl *ForwardList) FillInsertAfter(pos *ForwardListIter, n int, val Value) *ForwardListIter {
	if n > 0 {
		var tmp = NewNValues(n, val)
		return fl.spliceAfter(pos, tmp.BeforeBegin(), tmp.End())
	} else {
		return &ForwardListIter{pos.node}
	}
}

// EraseAfter removes the element pointed to by the iterator following pos.
func (fl *ForwardList) EraseAfter(pos *ForwardListIter) *ForwardListIter {
	return &ForwardListIter{fl.eraseAfter(pos.node)}
}

// RangeEraseAfter removes a range of elements.
// first is the iterator pointing before the first elements to be erased.
// last is the iterator pointing to one past the last elements to be erased.
func (fl *ForwardList) RangeEraseAfter(pos, last *ForwardListIter) *ForwardListIter {
	return &ForwardListIter{fl.rangeEraseAfter(pos.node, last.node)}
}

// Swap swaps data with another ForwardList.
func (fl *ForwardList) Swap(x *ForwardList) {
	fl.head.next, x.head.next = x.head.next, fl.head.next
}

// PushFront adds data to the front of the ForwardList.
func (fl *ForwardList) PushFront(val Value) {
	fl.insertAfter(fl.BeforeBegin(), val)
}

// PopFront removes first element.
func (fl *ForwardList) PopFront() {
	fl.eraseAfter(fl.head)
}

// Resize resizes the ForwardList to the specified number of elements.
func (fl *ForwardList) Resize(newSize int) {
	var k = fl.BeforeBegin()

	var len = 0
	for *k.makeNext() != *fl.End() && len < newSize {
		k.Next()
		len++
	}
	if len == newSize {
		fl.RangeEraseAfter(k, fl.End())
	} else {
		fl.defaultInsertAfter(k, newSize-len)
	}
}

// FillResize resizes the ForwardList to the specified number of elements.
// val is the data with which new elements should be populated.
func (fl *ForwardList) FillResize(newSize int, val Value) {
	var k = fl.BeforeBegin()

	var len = 0
	for *k.makeNext() != *fl.End() && len < newSize {
		k.Next()
		len++
	}
	if len == newSize {
		fl.RangeEraseAfter(k, fl.End())
	} else {
		fl.FillInsertAfter(k, newSize-len, val)
	}
}

// FillAssign assigns a given value to a ForwardList.
func (fl *ForwardList) FillAssign(n int, val Value) {
	fl.assignN(n, val)
}

// RangeAssign assigns a range to a ForwardList.
func (fl *ForwardList) RangeAssign(first, last InputIter) {
	fl.rangeAssign(first, last)
}

// Operations

// Merge merges sorted lists according to comparison function.
func (fl *ForwardList) Merge(list *ForwardList, comp fn.Compatator) {
	fl.merge(list, comp)
}

// SpliceAfter inserts contents of another ForwardList.
func (fl *ForwardList) SpliceAfter(pos *ForwardListIter, list *ForwardList) {
	if !list.Empty() {
		fl.spliceAfter(pos, list.BeforeBegin(), list.End())
	}
}

// SpliceElementAfter removes the element in list referenced by i and inserts it
// into the current list after pos.
func (fl *ForwardList) SpliceElementAfter(pos *ForwardListIter, i *ForwardListIter) {
	fl.spliceElementAfter(pos, i)
}

// RangeSpliceAfter inserts range from another ForwardList.
func (fl *ForwardList) RangeSpliceAfter(pos, before, last *ForwardListIter) {
	fl.spliceAfter(pos, before, last)
}

// Remove removes all elements equal to value.
func (fl *ForwardList) Remove(val Value) {
	var curr = fl.head
	var extra *fwdListNode

	for curr.next != nil {
		if curr.next.val == val {
			if &(curr.next.val) != &val {
				fl.eraseAfter(curr)
				continue
			} else {
				extra = curr
			}
		}
		curr = curr.next
	}

	if extra != nil {
		fl.eraseAfter(extra)
	}
}

// RemoveIf removes all elements satisfying a predicate.
func (fl *ForwardList) RemoveIf(pred fn.Predicator) {
	var curr = fl.head
	for curr.next != nil {
		if pred.Predicate(curr.next.val) {
			fl.eraseAfter(curr)
		} else {
			curr = curr.next
		}
	}
}

// Unique removes consecutive duplicate elements.
func (fl *ForwardList) Unique() {
	fl.UniqueIf(fn.BinaryPredicateFunc(func(a, b interface{}) bool { return a == b }))
}

// UniqueIf removes consecutive elements satisfying a predicate.
func (fl *ForwardList) UniqueIf(binPred fn.BinaryPredicator) {
	var first, last = fl.Begin(), fl.End()
	if first.EqualTo(last) {
		return
	}
	var next = first.clone()
	// ++next != last
	for !next.next().EqualTo(last) {
		if binPred.Predicate(first.Deref(), next.Deref()) {
			fl.eraseAfter(first.node)
		} else {
			first = next.clone()
		}
		next = first.clone()
	}
}

// Reverse reverses the elements in list.
func (fl *ForwardList) Reverse() {
	fl.head.reverseAfter()
}

// Sort sorts the elements in list.
func (fl *ForwardList) Sort(comp fn.Compatator) {
	fl.sort(comp)
}

func (fl *ForwardList) defaultInitialize(n int) {
	var to = fl.head
	for ; n > 0; n-- {
		to.next = fl.createNode(nil)
		to = to.next
	}
}

func (fl *ForwardList) fillInitialize(n int, val Value) {
	var to = fl.head
	for ; n > 0; n-- {
		to.next = fl.createNode(val)
		to = to.next
	}
}

func (fl *ForwardList) rangeInitialize(first, last InputIter) {
	var to = fl.head
	for first = first.Clone().(InputIter); !first.EqualTo(last); first.Next() {
		to.next = fl.createNode(first.Deref())
		to = to.next
	}
}

func (fl *ForwardList) createNode(val Value) *fwdListNode {
	return &fwdListNode{val: val}
}

func (fl *ForwardList) eraseAfter(pos *fwdListNode) *fwdListNode {
	var curr = pos.next
	pos.next = curr.next
	curr.val = nil
	return pos.next
}

func (fl *ForwardList) rangeEraseAfter(pos, last *fwdListNode) *fwdListNode {
	var curr = pos.next
	for curr != last {
		var temp = curr
		curr = curr.next
		temp.val = nil
	}
	pos.next = last
	return last
}

func (fl *ForwardList) insertAfter(pos *ForwardListIter, val Value) *fwdListNode {
	var to = pos.node
	var node = fl.createNode(val)
	node.next = to.next
	to.next = node
	return to.next
}

func (fl *ForwardList) spliceAfter(pos, before, last *ForwardListIter) *ForwardListIter {
	var tmp, b, e = pos.node, before.node, before.node

	for e != nil && e.next != last.node {
		e = e.next
	}

	if b != e {
		return &ForwardListIter{tmp.transferAfter(b, e)}
	} else {
		return &ForwardListIter{tmp}
	}
}

func (fl *ForwardList) spliceElementAfter(pos *ForwardListIter, i *ForwardListIter) {
	var j = i.clone()
	j.Next()

	if pos.EqualTo(i) || pos.EqualTo(j) {
		return
	}

	var tmp = pos.node
	tmp.transferAfter(i.node, j.node)
}

func (fl *ForwardList) defaultInsertAfter(pos *ForwardListIter, n int) {
	for ; n > 0; n-- {
		pos = fl.InsertAfter(pos, nil)
	}
}

func (fl *ForwardList) assignN(n int, val Value) {
	var prev, curr, end = fl.BeforeBegin(), fl.Begin(), fl.End()
	for !curr.EqualTo(end) && n > 0 {
		curr.node.val = val
		prev.Next()
		curr.Next()
		n--
	}
	if n > 0 {
		fl.FillInsertAfter(prev, n, val)
	} else if !curr.EqualTo(end) {
		fl.RangeEraseAfter(prev, end)
	}
}

func (fl *ForwardList) rangeAssign(first, last InputIter) {
	first = first.Clone().(InputIter)
	var prev, curr, end = fl.BeforeBegin(), fl.Begin(), fl.End()
	for !curr.EqualTo(end) && !first.EqualTo(last) {
		curr.node.val = first.Deref()
		prev.Next()
		curr.Next()
		first.Next()
	}
	if !first.EqualTo(last) {
		fl.RangeInsertAfter(prev, first, last)
	} else if !curr.EqualTo(end) {
		fl.RangeEraseAfter(prev, end)
	}
}

func (fl *ForwardList) merge(list *ForwardList, comp fn.Compatator) {
	var node = fl.head
	for node.next != nil && list.head.next != nil {
		if comp.Compare(list.head.next.val, node.next.val) {
			node.transferAfter(list.head, list.head.next)
		}
		node = node.next
	}
	if list.head.next != nil {
		node.next = list.head.next
		list.head.next = nil
	}
}

func (fl *ForwardList) sort(comp fn.Compatator) {
	// If `next' is 0, return immediately.
	var list = fl.head.next
	if list == nil {
		return
	}

	var insize = 1

	for {
		var p = list
		list = nil
		var tail *fwdListNode = nil

		// Count number of merges we do in this pass.
		var nmerges = 0

		for p != nil {
			nmerges++
			// There exists a merge to be done.
			// Step `insize' places along from p.
			var q = p
			var psize = 0
			for i := 0; i < insize; i++ {
				psize++
				q = q.next
				if q == nil {
					break
				}
			}

			// If q hasn't fallen off end, we have two lists to merge.
			var qsize = insize

			// Now we have two lists; merge them.
			for psize > 0 || (qsize > 0 && q != nil) {
				// Decide whether next node of merge comes from p or q.
				var e *fwdListNode
				if psize == 0 {
					// p is empty; e must come from q.
					e = q
					q = q.next
					qsize--
				} else if qsize == 0 || q == nil {
					// q is empty; e must come from p.
					e = p
					p = p.next
					psize--
				} else if comp.Compare(p.val, q.val) {
					// First node of p is lower; e must come from p.
					e = p
					p = p.next
					psize--
				} else {
					// First node of q is lower; e must come from q.
					e = q
					q = q.next
					qsize--
				}

				// Add the next node to the merged list.
				if tail != nil {
					tail.next = e
				} else {
					list = e
				}
				tail = e
			}

			// Now p has stepped `insize' places along, and q has too.
			p = q
		}
		tail.next = nil

		// If we have done only one merge, we're finished.
		// Allow for nmerges == 0, the empty list case.
		if nmerges <= 1 {
			fl.head.next = list
			return
		}

		// Otherwise repeat, merging lists twice the size.
		insize *= 2
	}
}

type fwdListNode struct {
	next *fwdListNode
	val  Value
}

func (n *fwdListNode) transferAfter(begin, end *fwdListNode) *fwdListNode {
	var keep = begin.next
	if end != nil {
		begin.next = end.next
		end.next = n.next
	} else {
		begin.next = nil
	}
	n.next = keep
	return end
}

func (n *fwdListNode) reverseAfter() {
	var tail = n.next
	if tail == nil {
		return
	}

	for tail.next != nil {
		var keep = n.next
		n.next = tail.next
		tail.next = tail.next.next
		n.next.next = keep
	}
}
