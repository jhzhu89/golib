package vector

import (
	"github.com/jhzhu89/golib/algorithm"
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/iterator"
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

type Vector struct {
	vectorImpl
}

// Iterators

func (v *Vector) Begin() *VectorIter {
	return clone(v.start)
}

func (v *Vector) End() *VectorIter {
	return clone(v.finish)
}

func (v *Vector) RBegin() *ReverseIter {
	return iterator.NewReverseIterator(v.finish)
}

func (v *Vector) REnd() *ReverseIter {
	return iterator.NewReverseIterator(v.start)
}

// Element access

func (v *Vector) At(n int) Value {
	return nil
}

func (v *Vector) Front() Value {
	return v.At(0)
}

func (v *Vector) Back() Value {
	return nil
}

// Capacity

// Empty returns true if the Deuqe is empty.
func (v *Vector) Empty() bool {
	return true
}

func (v *Vector) Size() int {
	return v.finish.cur - v.start.cur
}

//func (v *Vector) MaxSize() int {
//	return 0
//}

func (v *Vector) Reserve(newCap int) {
}

func (v *Vector) Capacity() int {
	return 0
}

func (v *Vector) ShrinkToFit() bool {
	return false
}

// Modifiers
func (v *Vector) Clear() {
}

func (v *Vector) Insert(pos *VectorIter, val Value) *VectorIter {
	return nil
}

//func (i insertFunc) Insert(it Iter, val Value) Iter {
//	return nil
//}

func (v *Vector) InsertRange(pos *VectorIter, first, last InputIter) *VectorIter {
	return nil
}

func (v *Vector) FillInsert(pos *VectorIter, n int, val Value) *VectorIter {
	return nil
}

func (v *Vector) Erase(pos *VectorIter) *VectorIter {
	return nil
}

func (v *Vector) EraseRange(first, last *VectorIter) *VectorIter {
	return nil
}

func (v *Vector) Swap(x *Vector) {

}

func (v *Vector) PushBack(val Value) {

}

func (v *Vector) PopBack(val Value) {

}

func (v *Vector) Resize(newSize int) {
	var len = v.Size()
	if newSize > len {
		v.defaultAppend(newSize - len)
	} else if newSize < len {
		v.eraseAtEnd(nextN(clone(v.start), newSize))
	}
}

func (v *Vector) ResizeFill(newSize int, val Value) {
	var len = v.Size()
	if newSize > len {
		v.fillInsert(v.End(), newSize-len, val)
	} else if newSize < len {
		v.eraseAtEnd(nextN(clone(v.start), newSize))
	}
}

// FillAssign assigns a given value to a Deque.
func (v *Vector) FillAssign(size int, val Value) {
	v.fillAssign(size, val)
}

// AssignRange assigns a range to a Deque.
func (v *Vector) AssignRange(first, last InputIter) {

}

func (v *Vector) fillAssign(n int, val Value) {
	if n > v.Capacity() {
	} else if n > v.Size() {
		algorithm.Fill(v.start, v.endOfStorage, val)
	} else {
		v.eraseAtEnd(algorithm.FillN(v.Begin(), n, val).(*VectorIter))
	}
}

func (v *Vector) fillInsert(pos *VectorIter, n int, val Value) {
	if n != 0 {
		if v.endOfStorage.cur-v.finish.cur >= n {
			var elemsAfter = pos.Distance(v.finish)
			var oldFinish = clone(v.finish)
			if elemsAfter > n {

			} else {

			}
		} else {
		}
	}
}

func (v *Vector) fillInitialize(n int, val Value) {
	v.createStorage(n)
	v.finish = algorithm.FillN(v.start, n, val).(*VectorIter)
}

func (v *Vector) rangeInitialize(first, last InputIter) {
	switch first.(type) {
	case ForwardIter:
		var n = iterator.Distance(first, last)
		v.createStorage(n)
		var it = first.Clone().(InputIter)
		for i := 0; i < n; i++ {
			(*v.data)[i] = it.Deref()
			it.Next()
		}

	default:
		for first = first.Clone().(InputIter); !first.Equal(last); first.Next() {
			v.PushBack(first.Deref())
		}
	}
}

func (v *Vector) defaultAppend(n int) {
}

func (v *Vector) eraseAtEnd(it *VectorIter) {
}

type vec []Value

// vectorImpl handles vector allocation.
type vectorImpl struct {
	data *vec

	start, finish, endOfStorage *VectorIter
}

func (v *vectorImpl) newIter(i int) *VectorIter {
	return &VectorIter{i, &v.data}
}

func (v *vectorImpl) allocate(n int) {
	var vec_ = make(vec, n, n)
	v.data = &vec_
}

func (v *vectorImpl) createStorage(n int) {
	v.allocate(n)
	v.start, v.finish, v.endOfStorage = v.newIter(0), v.newIter(0), v.newIter(len(*v.data))
}
