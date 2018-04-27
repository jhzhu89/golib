// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package vector

var _ RandIter = (*VectorIter)(nil)

// VectorIter implements a random access iterator.
type VectorIter struct {
	cur  int
	data **node
}

// CanMultiPass indicates this is a forward iterator.
func (it *VectorIter) CanMultiPass() {}

// Clone returns a copy of this iterator.
func (it *VectorIter) Clone() IterRef {
	return &VectorIter{it.cur, it.data}
}

// CopyAssign assgins given iterator to this iterator.
func (it *VectorIter) CopyAssign(r IterCRef) {
	*it = *(r.(*VectorIter))
}

// Swap swaps data with another iterator.
func (it *VectorIter) Swap(r IterCRef) {
	var r_ = r.(*VectorIter)
	*it, *r_ = *r_, *it
}

// Deref is like the derefference operation (*it) in c++ to get the value at this position.
func (it *VectorIter) Deref() Value {
	return (*(*it.data))[it.cur]
}

// DerefSet is like the derefference operation (*it=val) in c++ to
// set value to the element at this position.
func (it *VectorIter) DerefSet(val Value) {
	(*(*it.data))[it.cur] = val
}

// Next is like the pointer arithmetic it++ in c++ to increment the iterator.
func (it *VectorIter) Next() {
	it.cur++
}

// Prev is like the pointer arithmetic it-- in c++ to decrement the iterator.
func (it *VectorIter) Prev() {
	it.cur--
}

// NextN is like the pointer arithmetic it+=n in c++ to increment the iterator by n.
func (it *VectorIter) NextN(n int) {
	it.cur += n
}

// PrevN is like the pointer arithmetic it-=n in c++ to decrement the iterator by n.
func (it *VectorIter) PrevN(n int) {
	it.NextN(-n)
}

// next moves an iterator forward.
func (it *VectorIter) next() *VectorIter {
	it.Next()
	return it
}

// prev moves an iterator backward.
func (it *VectorIter) prev() *VectorIter {
	it.Prev()
	return it
}

// nextN moves an iterator forward by n.
func (it *VectorIter) nextN(n int) *VectorIter {
	it.NextN(n)
	return it
}

// prevN moves an iterator backward by n.
func (it *VectorIter) prevN(n int) *VectorIter {
	it.PrevN(n)
	return it
}

// clone returns a copy of it.
func (it *VectorIter) clone() *VectorIter {
	return &VectorIter{it.cur, it.data}
}

// EqualTo checks if given iterator is equal to this iterator.
func (it *VectorIter) EqualTo(r IterCRef) bool {
	return *it == *r.(*VectorIter)
}

// LessThan checks if this iterator is less than given iterator.
func (it *VectorIter) LessThan(r IterCRef) bool {
	var r_ = r.(*VectorIter)
	return it.data == r_.data && it.cur < r_.cur
}

// Distance counts distance between given iterator and this one.
// 	d = r - it.
func (it *VectorIter) Distance(r IterCRef) int {
	return r.(*VectorIter).cur - it.cur
}
