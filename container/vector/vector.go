package vector

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
)

type Vector struct {
	vectorImpl
}

func New() *Vector {
	v := &Vector{}
	v.createStorage(0)
	return v
}

func NewN(n int) *Vector {
	v := &Vector{}
	v.createStorage(n)
	v.finish.cur = n
	return v
}

func NewNValues(n int, val Value) *Vector {
	v := &Vector{}
	v.createStorage(n)
	v.fillInitialize(n, val)
	return v
}

func NewFromRange(first, last InputIter) *Vector {
	v := New()
	v.rangeInitialize(first, last)
	return v
}

// Iterators

func (v *Vector) Begin() *VectorIter {
	return v.start.Clone2()
}

func (v *Vector) End() *VectorIter {
	return v.finish.Clone2()
}

func (v *Vector) RBegin() *ReverseIter {
	return iterator.NewReverseIterator(v.finish)
}

func (v *Vector) REnd() *ReverseIter {
	return iterator.NewReverseIterator(v.start)
}

// Element access

func (v *Vector) At(n int) Value {
	return (*v.data)[v.start.cur+n]
}

func (v *Vector) Front() Value {
	return v.At(0)
}

func (v *Vector) Back() Value {
	return (*v.data)[v.finish.cur-1]
}

// Capacity

// Empty returns true if the Vector is empty.
func (v *Vector) Empty() bool {
	return v.start.cur == v.finish.cur
}

func (v *Vector) Size() int {
	return v.finish.cur - v.start.cur
}

func (v *Vector) Reserve(n int) {
	if v.Capacity() < n {
		v.extend(n - v.Capacity())
	}
}

func (v *Vector) Capacity() int {
	return v.endOfStorage.cur - v.start.cur
}

func (v *Vector) ShrinkToFit() bool {
	if v.Capacity() == v.Size() {
		return false
	}
	var x = NewFromRange(v.start, v.finish)
	v.Swap(x)
	return true
}

// Modifiers
func (v *Vector) Clear() {
	v.eraseAtEnd(v.start)
}

func (v *Vector) Insert(pos *VectorIter, val Value) *VectorIter {
	var n = v.start.Distance(pos)
	if v.finish.cur == v.endOfStorage.cur {
		v.extend(v.checkLen(1) - v.Size())
	}
	if pos.cur == v.finish.cur {
		(*v.data)[v.finish.cur] = val
		v.finish.cur++
	} else {
		v.insertAux(v.start.Clone2().NextN2(n), val)
	}
	return v.start.Clone2().NextN2(n)
}

func (v *Vector) RangeInsert(pos *VectorIter, first, last InputIter) *VectorIter {
	v.rangeInsert(pos, first, last)
	return pos.Clone2()
}

func (v *Vector) FillInsert(pos *VectorIter, n int, val Value) *VectorIter {
	v.fillInsert(pos, n, val)
	return pos.Clone2()
}

func (v *Vector) Erase(pos *VectorIter) *VectorIter {
	return v.erase(pos)
}

func (v *Vector) RangeErase(first, last *VectorIter) *VectorIter {
	return v.rangeErase(first, last)
}

func (v *Vector) Swap(x *Vector) {
	v.data, x.data = x.data, v.data
	v.start, x.start = x.start, v.start
	v.finish, x.finish = x.finish, v.finish
	v.endOfStorage, x.endOfStorage = x.endOfStorage, v.endOfStorage
}

func (v *Vector) PushBack(val Value) {
	if v.finish.Equal(v.endOfStorage) {
		v.extend(v.checkLen(1) - v.Size())
	}
	(*v.data)[v.finish.cur] = val
	v.finish.cur++
}

func (v *Vector) PopBack() {
	v.finish.cur--
	(*v.data)[v.finish.cur] = nil
}

func (v *Vector) Resize(newSize int) {
	var len = v.Size()
	if newSize > len {
		v.defaultAppend(newSize - len)
	} else if newSize < len {
		v.eraseAtEnd(v.start.Clone2().NextN2(newSize))
	}
}

func (v *Vector) FillResize(newSize int, val Value) {
	var len = v.Size()
	if newSize > len {
		v.fillInsert(v.End(), newSize-len, val)
	} else if newSize < len {
		v.eraseAtEnd(v.start.Clone2().NextN2(newSize))
	}
}

// FillAssign assigns a given value to a Vector.
func (v *Vector) FillAssign(size int, val Value) {
	v.fillAssign(size, val)
}

// RangeAssign assigns a range to a Deque.
func (v *Vector) RangeAssign(first, last InputIter) {
	v.assignAux(first, last)
}

func (v *Vector) fillAssign(n int, val Value) {
	if n > v.Size() {
		if n > v.Capacity() {
			v.extend(n - v.Capacity())
		}
		algorithm.Fill(v.start, v.start.Clone2().NextN2(n), val)
		v.finish.cur += n
	} else {
		v.eraseAtEnd(algorithm.FillN(v.start, n, val).(*VectorIter))
	}
}

func (v *Vector) fillInsert(pos *VectorIter, n int, val Value) {
	if n != 0 {
		if v.endOfStorage.cur-v.finish.cur < n {
			v.extend(v.checkLen(n) - v.Size())
		}
		algorithm.CopyBackward(pos, v.finish, v.finish.Clone2().NextN2(n))
		v.finish.NextN(n)
		algorithm.Fill(pos, pos.Clone2().NextN2(n), val)
	}
}

func (v *Vector) fillInitialize(n int, val Value) {
	v.finish = algorithm.FillN(v.start, n, val).(*VectorIter)
}

func (v *Vector) rangeInitialize(first, last InputIter) {
	switch first.(type) {
	case ForwardIter:
		var n = iterator.Distance(first, last)
		v.extend(n - v.Size())
		v.finish = algorithm.Copy(first, last, v.start).(*VectorIter)

	default:
		for first = first.Clone().(InputIter); !first.Equal(last); first.Next() {
			v.PushBack(first.Deref())
		}
	}
}

func (v *Vector) rangeInsert(pos *VectorIter, first, last InputIter) {
	switch first.(type) {
	case ForwardIter:
		if !first.Equal(last) {
			var n = iterator.Distance(first, last)
			if v.endOfStorage.cur-v.finish.cur < n {
				v.extend(v.checkLen(n) - v.Size())
			}
			algorithm.CopyBackward(pos, v.finish, v.finish.Clone2().NextN2(n))
			v.finish.NextN(n)
			algorithm.Copy(first, last, pos)
		}

	default:
		for first = first.Clone().(InputIter); !first.Equal(last); first.Next() {
			pos = v.Insert(pos, first.Deref())
			pos.Next()
		}
	}
}

func (v *Vector) checkLen(n int) int {
	return v.Size() + util.Max(v.Size(), n)
}

func (v *Vector) insertAux(pos *VectorIter, val Value) {
	(*v.data)[v.finish.cur] = (*v.data)[v.finish.cur-1]
	v.finish.cur++
	algorithm.CopyBackward(pos, v.finish.Clone2().PrevN2(2), v.finish.Clone2().PrevN2(1))
	pos.DerefSet(val)
}

func (v *Vector) defaultAppend(n int) {
	if n != 0 {
		if v.endOfStorage.cur-v.finish.cur < n {
			v.extend(v.checkLen(n) - v.Size())
		}
		v.finish.NextN(n)
	}
}

func (v *Vector) eraseAtEnd(pos *VectorIter) {
	v.destroyData(pos)
	v.finish.cur = pos.cur
}

func (v *Vector) erase(pos *VectorIter) *VectorIter {
	var nextToPos = pos.Clone2().Next2()
	if !nextToPos.Equal(v.finish) {
		algorithm.Copy(nextToPos, v.finish, pos)
	}
	v.finish.cur--
	(*v.data)[v.finish.cur] = nil
	return pos.Clone2()
}

func (v *Vector) rangeErase(first, last *VectorIter) *VectorIter {
	if !first.Equal(last) {
		if !last.Equal(v.finish) {
			algorithm.Copy(last, v.finish, first)
		}
		v.eraseAtEnd(first.Clone2().NextN2(last.Distance(v.finish)))
	}
	return first.Clone2()
}

func (v *Vector) assignAux(first, last InputIter) {
	switch first.(type) {
	case ForwardIter:
		var len = iterator.Distance(first, last)
		if len > v.Size() {
			if len > v.Capacity() {
				v.extend(len - v.Capacity())
			}
			v.finish = algorithm.Copy(first, last, v.start).(*VectorIter)
		} else {
			v.eraseAtEnd(algorithm.Copy(first, last, v.start).(*VectorIter))
		}

	default:
		var cur = v.start.Clone2()
		for first = first.Clone().(InputIter); !first.Equal(last) &&
			!cur.Equal(v.finish); first.Next() {
			cur.DerefSet(first.Deref())
			cur.Next()
		}

		if first.Equal(last) {
			v.eraseAtEnd(cur)
		} else {
			v.rangeInsert(v.finish, first, last)
		}
	}
}

type node []Value

// vectorImpl handles vector allocation.
type vectorImpl struct {
	data *node

	start, finish, endOfStorage *VectorIter
}

func (v *vectorImpl) newIter(i int) *VectorIter {
	return &VectorIter{i, &v.data}
}

func (v *vectorImpl) allocate(n int) {
	var node_ = make(node, n, n)
	v.data = &node_
}

func (v *vectorImpl) extend(n int) {
	*v.data = append(*v.data, make(node, n, n)...)
	v.endOfStorage.cur = len(*v.data)
}

func (v *vectorImpl) destroyData(pos *VectorIter) {
	for i := pos.cur; i < v.finish.cur; i++ {
		(*v.data)[i] = nil
	}
}

func (v *vectorImpl) createStorage(n int) {
	v.allocate(n)
	v.start, v.finish, v.endOfStorage = v.newIter(0), v.newIter(0), v.newIter(n)
}
