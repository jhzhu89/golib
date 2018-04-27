// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package list implements a c++ STL-like list.
// To iterate over a list (where l is a *List):
//	for it := l.Begin(); !it.Equal(l.End()); it.Next()) {
//		val := it.Deref()
//		// do something with val
//	}
package list

import (
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/fn"
	"github.com/jhzhu89/golib/iterator"
)

// Type aliases.
type (
	Value = container.Value

	IterRef      = iterator.IterRef
	IterCRef     = iterator.IterCRef
	InputIter    = iterator.InputIter
	RandIter     = iterator.RandIter
	BidirectIter = iterator.BidirectIter
	ReverseIter  = iterator.ReverseIterator
)

// List represents a c++ STL-like list. It is a standard container with linear
// time access to elements, and fixed time insertion/deletion at any point in the sequence.
type List struct {
	node *listNode
}

// New creates a List with no elements.
func New() *List {
	l := &List{&listNode{}}
	l.node.next = l.node
	l.node.prev = l.node
	l.setSize(0)
	return l
}

// NewN creates a List with n nil elements.
func NewN(n int) *List {
	l := New()
	l.defaultInitialize(n)
	return l
}

// NewNValues fills the List with n copies of val.
func NewNValues(n int, val Value) *List {
	l := New()
	l.fillInitialize(n, val)
	return l
}

// NewFromRange creates a List consisting of copies of the
// elements from [first, last).
func NewFromRange(first, last InputIter) *List {
	l := New()
	for first = first.Clone().(InputIter); !first.EqualTo(last); first.Next() {
		l.PushBack(first.Deref())
	}
	return l
}

// Iterators

// Begin returns a read/write iterator that points to the first element in the
// List. Iteration is done in ordinary element order.
func (l *List) Begin() *ListIter {
	return &ListIter{l.node.next}
}

// End returns a read/write iterator that points one past the last
// element in the List. Iteration is done in ordinary
// element order.
func (l *List) End() *ListIter {
	return &ListIter{l.node}
}

// RBegin returns a read/write reverse iterator that points to the last element
// in the List. Iteration is done in reverse element order.
func (l *List) RBegin() *ReverseIter {
	return iterator.NewReverseIterator(l.End())
}

// REnd returns a read/write reverse iterator that points to one before the first
// element in the List. Iteration is done in reverse element order.
func (l *List) REnd() *ReverseIter {
	return iterator.NewReverseIterator(l.Begin())
}

// Element access

// Front returns the data at the first element of the List.
func (l *List) Front() Value {
	return l.Begin().Deref()
}

// Back returns the data at the last element of the List.
func (l *List) Back() Value {
	var tmp = l.End()
	tmp.Prev()
	return tmp.Deref()
}

// Capacity

// Empty returns true if the List is empty.
func (l *List) Empty() bool {
	return l.node.next == l.node
}

// Size returns the number of elements in the list.
func (l *List) Size() int {
	return l.nodeCount()
}

// Modifiers

// Clear erases all the elements.
func (l *List) Clear() {
	l.clear()
}

// Insert inserts given value into List before specified iterator.
func (l *List) Insert(pos *ListIter, val Value) *ListIter {
	return l.insert(pos, val)
}

// RangeInsert inserts a range into the List before the location specified by pos.
func (l *List) RangeInsert(pos *ListIter, first, last InputIter) *ListIter {
	var tmp = NewFromRange(first, last)
	if !tmp.Empty() {
		var it = tmp.Begin()
		l.spliceList(pos, tmp)
		return it
	}
	return pos.clone()
}

// FillInsert inserts a number of copies of given data before the location specified by pos.
func (l *List) FillInsert(pos *ListIter, n int, val Value) *ListIter {
	if n > 0 {
		var tmp = NewNValues(n, val)
		var it = tmp.Begin()
		l.spliceList(pos, tmp)
		return it
	}
	return pos.clone()
}

// Erase removes the element at the given pos and thus shorten the list by one.
func (l *List) Erase(pos *ListIter) *ListIter {
	var ret = &ListIter{pos.node.next}
	l.erase(pos)
	return ret
}

// RangeErase removes the elements in the range [first, last) and shorten the list accordingly.
// first is the iterator pointing to the first elements to be erased.
// last is the iterator pointing to one past the last elements to be erased.
func (l *List) RangeErase(first, last *ListIter) *ListIter {
	for !first.EqualTo(last) {
		first = l.Erase(first)
	}
	return last.clone()
}

// Swap swaps data with another List.
func (l *List) Swap(x *List) {
	l.node, x.node = x.node, l.node
}

// PushFront adds data to the front of the List.
func (l *List) PushFront(val Value) {
	l.insert(l.Begin(), val)
}

// PopFront removes first element.
func (l *List) PopFront() {
	l.erase(l.Begin())
}

// PushBack adds data to the end of the List.
func (l *List) PushBack(val Value) {
	l.insert(l.End(), val)
}

// PopBack removes last element.
func (l *List) PopBack() {
	l.erase(&ListIter{l.node.prev})
}

// Resize resizes the List to the specified number of elements.
func (l *List) Resize(newSize int) {
	var i *ListIter
	i, newSize = l.resizePos(newSize)
	if newSize > 0 {
		l.defaultAppend(newSize)
	} else {
		l.RangeErase(i, l.End())
	}
}

// FillResize resizes the List to the specified number of elements.
// val is the data with which new elements should be populated.
func (l *List) FillResize(newSize int, val Value) {
	var i *ListIter
	i, newSize = l.resizePos(newSize)
	if newSize > 0 {
		l.FillInsert(l.End(), newSize, val)
	} else {
		l.RangeErase(i, l.End())
	}
}

// FillAssign assigns a given value to a List.
func (l *List) FillAssign(n int, val Value) {
	l.fillAssign(n, val)
}

// RangeAssign assigns a range to a List.
func (l *List) RangeAssign(first, last InputIter) {
	l.rangeAssign(first, last)
}

// Operations

// Merge merges sorted lists according to comparison function.
func (l *List) Merge(list *List, comp fn.Compatator) {
	l.merge(list, comp)
}

// Splice inserts contents of another List.
func (l *List) Splice(pos *ListIter, list *List) {
	l.spliceList(pos, list)
}

// SpliceElementAfter removes the element in list referenced by i and inserts it
// into the current list after pos.
func (l *List) SpliceElement(pos *ListIter, list *List, i *ListIter) {
	l.spliceElement(pos, list, i)
}

// RangeSpliceAfter inserts range from another List.
func (l *List) RangeSplice(pos *ListIter, list *List, first, last *ListIter) {
	l.rangeSplice(pos, list, first, last)
}

// Remove removes all elements equal to value.
func (l *List) Remove(val Value) {
	var first, last = l.Begin(), l.End()
	for !first.EqualTo(last) {
		var next = first.clone().Next2()
		if first.Deref() == val {
			l.erase(first)
		}
		first = next
	}
}

// RemoveIf removes all elements satisfying a predicate.
func (l *List) RemoveIf(pred fn.Predicator) {
	var first, last = l.Begin(), l.End()
	for !first.EqualTo(last) {
		var next = first.clone().Next2()
		if pred.Predicate(first.Deref()) {
			l.erase(first)
		}
		first = next
	}
}

// Unique removes consecutive duplicate elements.
func (l *List) Unique() {
	var first, last = l.Begin(), l.End()
	if first.EqualTo(last) {
		return
	}

	var next = first.clone().Next2()
	for !next.EqualTo(last) {
		if first.Deref() == next.Deref() {
			l.erase(next)
		} else {
			first = next
		}
		next = first.clone().Next2()
	}
}

// UniqueIf removes consecutive elements satisfying a predicate.
func (l *List) UniqueIf(binPred fn.BinaryPredicator) {
	var first, last = l.Begin(), l.End()
	if first.EqualTo(last) {
		return
	}

	var next = first.clone().Next2()
	for !next.EqualTo(last) {
		if binPred.Predicate(first.Deref(), next.Deref()) {
			l.erase(next)
		} else {
			first = next
		}
		next = first.clone().Next2()
	}
}

// Reverse reverses the elements in list.
func (l *List) Reverse() {
	l.node.reverse()
}

// Sort sorts the elements in list.
func (l *List) Sort(comp fn.Compatator) {
	l.sort(comp)
}

func (l *List) getSize() int {
	return l.node.val.(int)
}

func (l *List) setSize(n int) {
	l.node.val = n
}

func (l *List) incSize(n int) {
	l.node.val = l.node.val.(int) + n
}

func (l *List) decSize(n int) {
	l.node.val = l.node.val.(int) - n
}

func (l *List) nodeCount() int {
	return l.node.val.(int)
}

func (l *List) clear() {
	var cur = l.node.next
	for cur != l.node {
		var tmp = cur
		cur = cur.next
		tmp.val = nil
		tmp = nil
	}

	l.node.next, l.node.prev = l.node, l.node
	l.setSize(0)
}

func (l *List) defaultInitialize(n int) {
	for ; n > 0; n-- {
		l.insert(l.End(), nil)
	}
}

func (l *List) fillInitialize(n int, val Value) {
	for ; n > 0; n-- {
		l.PushBack(val)
	}
}

func (l *List) createNode(val Value) *listNode {
	return &listNode{val: val}
}

func (l *List) erase(pos *ListIter) {
	l.decSize(1)
	pos.node.unhook()
	pos.node.val = nil
	pos.node = nil
}

func (l *List) insert(pos *ListIter, val Value) *ListIter {
	var tmp = l.createNode(val)
	tmp.hook(pos.node)
	l.incSize(1)
	return &ListIter{tmp}
}

func (l *List) spliceList(pos *ListIter, x *List) {
	if !x.Empty() {
		l.transfer(pos, x.Begin(), x.End())
		l.incSize(x.getSize())
		x.setSize(0)
	}
}

func (l *List) spliceElement(pos *ListIter, x *List, i *ListIter) {
	var j = i.clone()
	j.Next()
	if pos.EqualTo(i) || pos.EqualTo(j) {
		return
	}

	l.transfer(pos, i, j)
	l.incSize(1)
	x.decSize(1)
}

func (l *List) rangeSplice(pos *ListIter, x *List, first, last *ListIter) {
	if !first.EqualTo(last) {
		var n = l.distance(first.node, last.node)
		l.incSize(n)
		x.decSize(n)
		l.transfer(pos, first, last)
	}
}

func (l *List) fillAssign(n int, val Value) {
	var i = l.Begin()
	for !i.EqualTo(l.End()) && n > 0 {
		i.DerefSet(val)
		i.Next()
		n--
	}

	if n > 0 {
		l.FillInsert(l.End(), n, val)
	} else {
		l.RangeErase(i, l.End())
	}
}

func (l *List) rangeAssign(first2, last2 InputIter) {
	var first1, last1 = l.Begin(), l.End()
	first2 = first2.Clone().(InputIter)
	for !first1.EqualTo(last1) && !first2.EqualTo(last2) {
		first1.DerefSet(first2.Deref())
		first1.Next()
		first2.Next()
	}

	if first2.EqualTo(last2) {
		l.RangeErase(first1, last1)
	} else {
		l.RangeInsert(last1, first2, last2)
	}
}

func (l *List) merge(list *List, comp fn.Compatator) {
	var first1, last1 = l.Begin(), l.End()
	var first2, last2 = list.Begin(), list.End()

	for !first1.EqualTo(last1) && !first2.EqualTo(last2) {
		if comp.Compare(first2.Deref(), first1.Deref()) {
			var next = first2.clone().Next2()
			l.transfer(first1, first2, next)
			first2 = next
		} else {
			first1.Next()
		}
	}

	if !first2.EqualTo(last2) {
		l.transfer(last1, first2, last2)
	}

	l.incSize(list.getSize())
	list.setSize(0)
}

func (l *List) sort(comp fn.Compatator) {
	if l.node.next != l.node && l.node.next.next != l.node {
		var carry = New()
		var n = 64
		var tmp = make([]*List, n, n)
		for i := 0; i < n; i++ {
			tmp[i] = New()
		}
		var fill, counter = 0, 0

		for {
			carry.spliceElement(carry.Begin(), l, l.Begin())

			for counter = 0; counter != fill && !tmp[counter].Empty(); counter++ {
				tmp[counter].merge(carry, comp)
				carry.Swap(tmp[counter])
			}
			carry.Swap(tmp[counter])
			if counter == fill {
				fill++
			}

			if l.Empty() {
				break
			}
		}

		for counter = 1; counter != fill; counter++ {
			tmp[counter].merge(tmp[counter-1], comp)
		}
		l.Swap(tmp[fill-1])
	}
}

func (l *List) transfer(pos, first, last *ListIter) {
	pos.node.transfer(first.node, last.node)
}

func (l *List) resizePos(newSize int) (*ListIter, int) {
	var i *ListIter
	var len = l.Size()
	if newSize < len {
		if newSize <= len/2 {
			i = l.Begin()
			iterator.Advance(i, newSize)
		} else {
			i = l.End()
			var numErase = len - newSize
			iterator.Advance(i, -numErase)
		}
		return i, 0
	}

	i = l.End()
	return i, newSize - len
}

func (l *List) defaultAppend(n int) {
	for i := 0; i < n; i++ {
		l.insert(l.End(), nil)
	}
}

func (l *List) distance(first, last *listNode) int {
	var n = 0
	for first != last {
		first = first.next
		n++
	}
	return n
}

type listNode struct {
	next, prev *listNode
	val        Value
}

func (n *listNode) transfer(first, last *listNode) {
	if n != last {
		// Remove [first, last) from its old position.
		last.prev.next = n
		first.prev.next = last
		n.prev.next = first

		// Splice [first, last) into its new position.
		var tmp = n.prev
		n.prev = last.prev
		last.prev = first.prev
		first.prev = tmp
	}
}

func (n *listNode) reverse() {
	var tmp = n
	for {
		tmp.next, tmp.prev = tmp.prev, tmp.next
		tmp = tmp.prev
		if tmp == n {
			break
		}
	}
}

func (n *listNode) hook(pos *listNode) {
	n.next = pos
	n.prev = pos.prev
	pos.prev.next = n
	pos.prev = n
}

func (n *listNode) unhook() {
	var nextNode = n.next
	var prevNode = n.prev
	prevNode.next = nextNode
	nextNode.prev = prevNode
}
