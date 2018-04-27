// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package deque

var _ RandIter = (*DequeIter)(nil)

// DequeIter implements a random access iterator.
type DequeIter struct {
	cur  int
	node int
	map_ **nodeMap
}

// CanMultiPass indicates this is a forward iterator.
func (it *DequeIter) CanMultiPass() {}

// Clone returns a copy of this iterator.
func (it *DequeIter) Clone() IterRef {
	return &DequeIter{it.cur, it.node, it.map_}
}

// CopyAssign assgins given iterator to this iterator.
func (it *DequeIter) CopyAssign(r IterCRef) {
	*it = *(r.(*DequeIter))
}

// Swap swaps data with another iterator.
func (it *DequeIter) Swap(r IterCRef) {
	var r_ = r.(*DequeIter)
	*it, *r_ = *r_, *it
}

// Deref is like the derefference operation (*it) in c++ to get the value at this position.
func (it *DequeIter) Deref() Value {
	return (*(*(*it.map_))[it.node])[it.cur]
}

// DerefSet is like the derefference operation (*it=val) in c++ to
// set value to the element at this position.
func (it *DequeIter) DerefSet(val Value) {
	(*(*(*it.map_))[it.node])[it.cur] = val
}

func (it *DequeIter) setNode(newNode int) {
	it.node = newNode
}

// Next is like the pointer arithmetic it++ in c++ to increment the iterator.
func (it *DequeIter) Next() {
	it.cur++
	if it.cur == nodeEnd {
		it.setNode(it.node + 1)
		it.cur = 0
	}
}

// Prev is like the pointer arithmetic it-- in c++ to decrement the iterator.
func (it *DequeIter) Prev() {
	if it.cur == 0 {
		it.setNode(it.node - 1)
		it.cur = nodeEnd
	}
	it.cur--
}

// NextN is like the pointer arithmetic it+=n in c++ to increment the iterator by n.
func (it *DequeIter) NextN(n int) {
	var offset = n + it.cur
	if offset >= 0 && offset < DequeBufSize {
		it.cur += n
	} else {
		var nodeOffset int
		if offset > 0 {
			nodeOffset = offset / DequeBufSize
		} else {
			nodeOffset = -(-offset-1)/DequeBufSize - 1
		}
		it.setNode(it.node + nodeOffset)
		it.cur = offset - nodeOffset*DequeBufSize
	}
}

// PrevN is like the pointer arithmetic it-=n in c++ to decrement the iterator by n.
func (it *DequeIter) PrevN(n int) {
	it.NextN(-n)
}

// Next2 moves an iterator forward.
func (it *DequeIter) Next2() *DequeIter {
	it.Next()
	return it
}

// Prev2 moves an iterator backward.
func (it *DequeIter) Prev2() *DequeIter {
	it.Prev()
	return it
}

// NextN2 moves an iterator forward by n.
func (it *DequeIter) NextN2(n int) *DequeIter {
	it.NextN(n)
	return it
}

// PrevN2 moves an iterator backward by n.
func (it *DequeIter) PrevN2(n int) *DequeIter {
	it.PrevN(n)
	return it
}

// clone returns a copy of it.
func (it *DequeIter) clone() *DequeIter {
	return &DequeIter{it.cur, it.node, it.map_}
}

// EqualTo checks if given iterator is equal to this iterator.
func (it *DequeIter) EqualTo(r IterCRef) bool {
	var r_ = r.(*DequeIter)
	return it.map_ == r_.map_ && it.node == r_.node && it.cur == r_.cur
}

// LessThan checks if this iterator is less than given iterator.
func (it *DequeIter) LessThan(r IterCRef) bool {
	var r_ = r.(*DequeIter)
	return it.map_ == r_.map_ &&
		((it.node == r_.node && it.cur < r_.cur) ||
			it.node < r_.node)
}

// Distance counts distance between given iterator and this one.
// 	d = r - it.
func (it *DequeIter) Distance(r IterCRef) int {
	var r_ = r.(*DequeIter)
	return (r_.node-it.node)*DequeBufSize + r_.cur - it.cur
}
