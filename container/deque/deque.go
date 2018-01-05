// Copyright 2017-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package deque implements a c++ STL-like deque.
//
// To iterate over a deque (where d is a *Deque):
//	for it := d.Begin(); !it.Equal(d.End()); it.Next()) {
//		val := it.Deref()
//		// do something with val
//	}
package deque

import (
	"github.com/jhzhu89/golib/algorithm"
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/iterator"
)

const (
	// The size of underlying node which stores data.
	DequeBufSize = 512
	// The initial size of underlying map which holds pointers to nodes.
	InitialMapSize = 8

	nodeEnd = DequeBufSize
)

// Type aliases.
type (
	Value = container.Value

	Iter        = iterator.Iter
	IterRef     = iterator.IterRef
	IterCRef    = iterator.IterCRef
	InputIter   = iterator.InputIter
	RandIter    = iterator.RandIter
	ForwardIter = iterator.ForwardIter
	ReverseIter = iterator.ReverseIterator
)

// Deque represents a c++ STL-like deque.
// The zero value for Deque cannot be directly used. Call NewDeque*()
// to get a valid Deque instance.
type Deque struct {
	dequeImpl
}

func newDeque() *Deque {
	d := new(Deque)
	d.start = &DequeIter{map_: &d.map_}
	d.finish = &DequeIter{map_: &d.map_}
	return d
}

// New creates a Deque with no elements.
func New() *Deque {
	d := newDeque()
	d.initializeMap(0)
	return d
}

// NewN creates a Deque with n nil elements.
func NewN(n int) *Deque {
	d := newDeque()
	d.initializeMap(n)
	return d
}

// NewFromRange creates a Deque consisting of copies of the
// elements from [first, last).
func NewFromRange(first, last InputIter) *Deque {
	d := newDeque()
	d.rangeInitialize(first, last)
	return d
}

// Iterators

// Begin returns a read/write iterator that points to the first element in the
// Deque. Iteration is done in ordinary element order.
func (d *Deque) Begin() *DequeIter {
	return clone(d.start)
}

// End returns a read/write iterator that points one past the last
// element in the Deque. Iteration is done in ordinary
// element order.
func (d *Deque) End() *DequeIter {
	return clone(d.finish)
}

// RBegin returns a read/write reverse iterator that points to the
// last element in the Deque. Iteration is done in reverse
// element order.
func (d *Deque) RBegin() *ReverseIter {
	return iterator.NewReverseIterator(d.finish)
}

// REnd returns a read/write reverse iterator that points to one
// before the first element in the Deque. Iteration is done
// in reverse element order.
func (d *Deque) REnd() *ReverseIter {
	return iterator.NewReverseIterator(d.start)
}

// Capacity

// Size returns the number of elements in the Deque.
func (d *Deque) Size() int {
	return d.start.Distance(d.finish)
}

// Resize resizes the Deque to the specified number of elements.
func (d *Deque) Resize(newSize int) {
	var len = d.Size()
	if newSize > len {
		d.deafultAppend(newSize - len)
	} else if newSize < len {
		d.eraseAtEnd(nextN(d.Begin(), newSize))
	}
}

// ResizeAssign resizes the Deque to the specified number of elements.
// val is the data with which new elements should be populated.
func (d *Deque) ResizeAssign(newSize int, val Value) {
	var len = d.Size()
	if newSize > len {
		d.FillInsert(d.finish, newSize-len, val)
	} else if newSize < len {
		d.eraseAtEnd(nextN(d.Begin(), newSize))
	}
}

// ShrinkToFit shrinks deque to reduce memory use.
func (d *Deque) ShrinkToFit() bool {
	var frontCapacity = d.start.cur
	if frontCapacity == 0 {
		return false
	}

	var backCapacity = DequeBufSize - d.finish.cur
	if frontCapacity+backCapacity < DequeBufSize {
		return false
	}

	var x = NewN(d.Size())
	x.rangeInitialize(d.start, d.finish)
	d.Swap(x)
	return true
}

// Empty returns true if the Deuqe is empty.
func (d *Deque) Empty() bool {
	return d.start.Equal(d.finish)
}

// Element access

// At accesses data contained in the Deque by subscript.
func (d *Deque) At(n int) Value {
	var it = d.Begin()
	it.NextN(n)
	return (*(*d.map_)[it.node])[it.cur]
}

// Front returns the data at the first element of the Deuqe.
func (d *Deque) Front() Value {
	return d.At(0)
}

// Back returns the data at the last element of the Deuqe.
func (d *Deque) Back() Value {
	var it = d.End()
	it.Prev()
	return (*(*d.map_)[it.node])[it.cur]
}

// Modifiers

// FillAssign assigns a given value to a Deque.
func (d *Deque) FillAssign(size int, val Value) {
	if size > d.Size() {
		algorithm.Fill(d.start, d.finish, val)
		d.FillInsert(d.finish, size-d.Size(), val)
	} else {
		d.eraseAtEnd(nextN(clone(d.start), size))
		algorithm.Fill(d.start, d.finish, val)
	}
}

// AssignRange assigns a range to a Deque.
func (d *Deque) AssignRange(first, last InputIter) {
	var size = iterator.Distance(first, last)
	if size > d.Size() {
		var mid = first.Clone().(InputIter)
		iterator.Advance(mid, d.Size())
		algorithm.Copy(first, mid, d.start)
		d.InsertRange(d.finish, mid, last)
	} else {
		d.eraseAtEnd(algorithm.Copy(first, last, d.start).(*DequeIter))
	}
}

// PushBack adds data to the end of the Deque.
func (d *Deque) PushBack(val Value) {
	if d.finish.cur != DequeBufSize-1 {
		(*(*d.map_)[d.finish.node])[d.finish.cur] = val
		d.finish.cur++
	} else {
		d.reserveMapAtBack(1)
		(*d.map_)[d.finish.node+1] = d.allocateNode()
		(*(*d.map_)[d.finish.node])[d.finish.cur] = val
		d.finish.setNode(d.finish.node + 1)
		d.finish.cur = 0
	}
}

// PushFront adds data to the front of the Deque.
func (d *Deque) PushFront(val Value) {
	if d.start.cur != 0 {
		(*(*d.map_)[d.start.node])[d.start.cur-1] = val
		d.start.cur--
	} else {
		d.reserveMapAtFront(1)
		(*d.map_)[d.start.node-1] = d.allocateNode()
		d.start.setNode(d.start.node - 1)
		d.start.cur = DequeBufSize - 1
		(*(*d.map_)[d.start.node])[d.start.cur] = val
	}
}

// PopBack removes last element.
func (d *Deque) PopBack() {
	if d.finish.cur != 0 {
		d.finish.cur--
	} else {
		d.deallocateNode((*d.map_)[d.finish.node])
		d.finish.setNode(d.finish.node - 1)
		d.finish.cur = DequeBufSize - 1
	}
	(*(*d.map_)[d.finish.node])[d.finish.cur] = nil
}

// PopFront removes first element.
func (d *Deque) PopFront() {
	(*(*d.map_)[d.start.node])[d.start.cur] = nil
	if d.start.cur != DequeBufSize-1 {
		d.start.cur++
	} else {
		d.deallocateNode((*d.map_)[d.start.node])
		d.start.setNode(d.start.node + 1)
		d.start.cur = 0
	}
}

// Insert inserts given value into Deque before specified iterator.
func (d *Deque) Insert(pos *DequeIter, val Value) *DequeIter {
	pos = clone(pos)
	if pos.cur == d.start.cur {
		d.PushFront(val)
		return d.Begin()
	} else if pos.cur == d.finish.cur {
		d.PushBack(val)
		return prev(d.End())
	} else {
		var index = d.start.Distance(pos)
		if index < d.Size()/2 {
			d.PushFront(d.Front())
			var front1 = next(d.Begin())
			var front2 = next(clone(front1))
			pos = nextN(d.Begin(), index)
			var pos1 = next(clone(pos))
			algorithm.Copy(front2, pos1, front1)
		} else {
			d.PushBack(d.Back())
			var back1 = prev(d.End())
			var back2 = prev(clone(back1))
			pos = nextN(d.Begin(), index)
			algorithm.CopyBackward(pos, back2, back1)
		}
		(*(*d.map_)[pos.node])[pos.cur] = val
		return pos
	}
}

type insertFunc func(it Iter, val Value) Iter

func (i insertFunc) Insert(it Iter, val Value) Iter {
	return i(it, val)
}

// InsertRange inserts a range into the Deque.
func (d *Deque) InsertRange(pos *DequeIter, first, last InputIter) *DequeIter {
	pos = clone(pos)
	var offset = d.start.Distance(pos)

	switch first.(type) {
	case ForwardIter:
		var n = iterator.Distance(first, last)
		if pos.cur == d.start.cur {
			var newStart = d.reserveElementsAtFront(n)
			algorithm.Copy(first, last, newStart)
			d.start = newStart
		} else if pos.cur == d.finish.cur {
			var newFinish = d.reserveElementsAtBack(n)
			algorithm.Copy(first, last, d.finish)
			d.finish = newFinish
		} else {
			var elemsBefore = d.start.Distance(pos)
			var len = d.Size()
			if elemsBefore < len/2 {
				var newStart = d.reserveElementsAtFront(n)
				pos = nextN(clone(d.start), elemsBefore)
				algorithm.Copy(d.start, pos, newStart)
				d.start = newStart
				algorithm.Copy(first, last, prevN(clone(pos), n))
			} else {
				var newFinish = d.reserveElementsAtBack(n)
				var elemsAfter = len - elemsBefore
				pos = prevN(clone(d.finish), elemsAfter)
				algorithm.CopyBackward(pos, d.finish, newFinish)
				d.finish = newFinish
				algorithm.Copy(first, last, pos)
			}
		}

	default:
		algorithm.Copy(first, last,
			iterator.NewInsertIterator(
				insertFunc(func(it Iter, val Value) Iter {
					return d.Insert(it.(*DequeIter), val)
				}), pos,
			),
		)
	}
	return nextN(clone(d.start), offset)
}

//FillInsert inserts a number of copies of given data into the Deque.
func (d *Deque) FillInsert(pos *DequeIter, n int, val Value) *DequeIter {
	var offset = d.start.Distance(pos)
	d.fillInsert(pos, n, val)
	return nextN(clone(d.start), offset)
}

// Erase removes element at given position.
func (d *Deque) Erase(pos *DequeIter) *DequeIter {
	var next = next(clone(pos))
	var index = d.start.Distance(pos)
	if index < d.Size()>>1 {
		algorithm.CopyBackward(d.start, pos, next)
		d.PopFront()
	} else {
		if !next.Equal(d.finish) {
			algorithm.Copy(next, d.finish, pos)
		}
		d.PopBack()
	}
	return nextN(clone(d.start), index)
}

// EraseRange removes a range of elements.
func (d *Deque) EraseRange(first, last *DequeIter) *DequeIter {
	if first.Equal(last) {
		return first
	} else if first.Equal(d.start) && last.Equal(d.finish) {
		d.Clear()
		return d.End()
	} else {
		var n = first.Distance(last)
		var elemsBefore = d.start.Distance(first)
		if elemsBefore <= (d.Size()-n)/2 {
			if !first.Equal(d.start) {
				algorithm.CopyBackward(d.start, first, last)
			}
			d.eraseAtBegin(nextN(clone(d.start), n))
		} else {
			if !last.Equal(d.finish) {
				algorithm.Copy(last, d.finish, first)
			}
			d.eraseAtEnd(prevN(clone(d.finish), n))
		}
		return nextN(clone(d.start), elemsBefore)
	}
}

// Swap swaps data with another Deque.
func (d *Deque) Swap(x *Deque) {
	d.start, x.start = x.start, d.start
	d.finish, x.finish = x.finish, d.finish
	d.map_, x.map_ = x.map_, d.map_
	d.mapSize, x.mapSize = x.mapSize, d.mapSize
}

// Clear erases all the elements.
func (d *Deque) Clear() {
	d.eraseAtEnd(d.start)
}

func (d *Deque) deafultAppend(n int) {
	if n > 0 {
		d.finish = d.reserveElementsAtBack(n)
	}
}

func (d *Deque) reserveElementsAtBack(n int) *DequeIter {
	var vacancies = DequeBufSize - d.finish.cur - 1
	if n > vacancies {
		d.newElementsAtBack(n - vacancies)
	}
	return nextN(clone(d.finish), n)
}

func (d *Deque) newElementsAtBack(newElems int) {
	// TODO: add size limit?
	var newNodes = (newElems + DequeBufSize - 1) / DequeBufSize
	d.reserveMapAtBack(newNodes)
	for i := 1; i <= newNodes; i++ {
		(*d.map_)[d.finish.node+i] = d.allocateNode()
	}
}

func (d *Deque) reserveMapAtBack(nodesToAdd int) {
	if nodesToAdd+1 > d.mapSize-d.finish.node {
		d.reallocateMap(nodesToAdd, false)
	}
}

func (d *Deque) reserveElementsAtFront(n int) *DequeIter {
	var vacancies = d.start.cur
	if n > vacancies {
		d.newElementsAtFront(n - vacancies)
	}
	return prevN(clone(d.start), n)
}

func (d *Deque) newElementsAtFront(newElems int) {
	// TODO: add size limit?
	var newNodes = (newElems + DequeBufSize - 1) / DequeBufSize
	d.reserveMapAtFront(newNodes)
	for i := 1; i <= newNodes; i++ {
		(*d.map_)[d.start.node-i] = d.allocateNode()
	}
}

func (d *Deque) reserveMapAtFront(nodesToAdd int) {
	if nodesToAdd > d.start.node {
		d.reallocateMap(nodesToAdd, true)
	}
}

func (d *Deque) reallocateMap(nodesToAdd int, addAtFront bool) {
	var oldNumNodes = d.finish.node - d.start.node + 1
	var newNumNodes = oldNumNodes + nodesToAdd

	var newStart int
	if d.mapSize > 2*newNumNodes {
		newStart = (d.mapSize - newNumNodes) / 2
		if addAtFront {
			newStart += nodesToAdd
		}
		copy((*d.map_)[newStart:], (*d.map_)[d.start.node:d.finish.node+1])
	} else {
		var newMapSize = d.mapSize + max(d.mapSize, nodesToAdd) + 2
		var newMap = d.allocateMap(newMapSize)
		newStart = (newMapSize - newNumNodes) / 2
		if addAtFront {
			newStart += nodesToAdd
		}
		copy((*newMap)[newStart:], (*d.map_)[d.start.node:d.finish.node+1])
		d.deallocateMap(d.map_)
		d.map_ = newMap
		d.mapSize = newMapSize
	}
	d.start.setNode(newStart)
	d.finish.setNode(newStart + oldNumNodes - 1)
}

func (d *Deque) eraseAtEnd(pos *DequeIter) {
	d.destroyData(pos, d.finish)
	d.destroyNodes(pos.node+1, d.finish.node+1)
	d.finish = clone(pos)
}

func (d *Deque) eraseAtBegin(pos *DequeIter) {
	d.destroyData(d.start, pos)
	d.destroyNodes(d.start.node, pos.node)
	d.start = clone(pos)
}

func (d *Deque) destroyData(first, last *DequeIter) {
	var destroy = func(nodeSlice node) {
		for i := range nodeSlice {
			nodeSlice[i] = nil
		}
	}

	for node := first.node + 1; node < last.node; node++ {
		destroy(*(*d.map_)[node])
	}

	if first.node != last.node {
		destroy((*(*d.map_)[first.node])[first.cur:])
		destroy((*(*d.map_)[last.node])[:last.cur])
	} else {
		destroy((*(*d.map_)[first.node])[first.cur:last.cur])
	}
}

func (d *Deque) fillInsert(pos *DequeIter, n int, val Value) {
	pos = clone(pos)
	if pos.cur == d.start.cur {
		var newStart = d.reserveElementsAtFront(n)
		algorithm.Fill(newStart, d.start, val)
		d.start = newStart
	} else if pos.cur == d.finish.cur {
		var newFinish = d.reserveElementsAtBack(n)
		algorithm.Fill(d.finish, newFinish, val)
		d.finish = newFinish
	} else {
		var elemsBefore = d.start.Distance(pos)
		var len = d.Size()
		if elemsBefore < len/2 {
			var newStart = d.reserveElementsAtFront(n)
			pos = nextN(clone(d.start), elemsBefore)
			algorithm.Copy(d.start, pos, newStart)
			d.start = newStart
			algorithm.Fill(prevN(clone(pos), n), pos, val)
		} else {
			var newFinish = d.reserveElementsAtBack(n)
			var elemsAfter = len - elemsBefore
			pos = prevN(clone(d.finish), elemsAfter)
			algorithm.CopyBackward(pos, d.finish, newFinish)
			d.finish = newFinish
			algorithm.Fill(pos, nextN(clone(pos), n), val)
		}
	}
}

func (d *Deque) rangeInitialize(first, last InputIter) {
	switch first.(type) {
	case ForwardIter:
		var n = iterator.Distance(first, last)
		d.initializeMap(n)
		algorithm.Copy(first, last, d.start)

	default:
		d.initializeMap(0)
		for first = first.Clone().(InputIter); !first.Equal(last); first.Next() {
			d.PushBack(first.Deref())
		}
	}
}

type node []Value

type nodeMap []*node

type dequeImpl struct {
	map_          *nodeMap
	mapSize       int
	start, finish *DequeIter
}

func (i *dequeImpl) initializeMap(numElements int) {
	var numNodes = numElements/DequeBufSize + 1
	i.mapSize = max(InitialMapSize, numNodes+2)
	i.map_ = i.allocateMap(i.mapSize)

	var nstart = (i.mapSize - numNodes) / 2
	var nfinish = nstart + numNodes

	i.createNodes(nstart, nfinish)
	i.start.setNode(nstart)
	i.finish.setNode(nfinish - 1)
	i.finish.cur = numElements % DequeBufSize
}

func (i *dequeImpl) createNodes(start, finish int) {
	for cur := start; cur < finish; cur++ {
		(*i.map_)[cur] = i.allocateNode()
	}
}

func (i *dequeImpl) destroyNodes(start, finish int) {
	for n := start; n < finish; n++ {
		i.deallocateNode((*i.map_)[n])
		(*i.map_)[n] = (*node)(nil)
	}
}

func (i *dequeImpl) allocateNode() *node {
	var n = make(node, DequeBufSize, DequeBufSize)
	return &n
}

func (i *dequeImpl) deallocateNode(n *node) {
	*n = nil
}

func (i *dequeImpl) allocateMap(size int) *nodeMap {
	var map_ = make(nodeMap, size, size)
	return &map_
}

func (i *dequeImpl) deallocateMap(map_ *nodeMap) {
	*map_ = nil
}
