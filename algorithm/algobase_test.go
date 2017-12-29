package algorithm

import (
	"testing"

	"github.com/jhzhu89/golib/iterator"
	"github.com/stretchr/testify/assert"
)

type (
	Iter     = iterator.Iter
	IterCRef = iterator.IterCRef
)

type intSlice struct {
	data []Value
}

func (s *intSlice) Insert(iter Iter, val Value) Iter {
	return s.insert(iter.(*intSliceIter), val)
}

func (s *intSlice) insert(iter *intSliceIter, val Value) *intSliceIter {
	iter = iter.Clone().(*intSliceIter)
	var i = iter.i
	s.data = append(s.data, nil)
	copy(s.data[i+1:], s.data[i:])
	s.data[i] = val
	return iter
}

func (s *intSlice) PushBack(val Value) {
	s.data = append(s.data, val)
}

func (s *intSlice) PushFront(val Value) {
	var size = len(s.data)
	s.data = append(s.data, nil)
	copy(s.data[1:], s.data[:size])
	s.data[0] = val
}

type intSliceIter struct {
	data *[]Value
	i    int
}

func (it *intSliceIter) CopyAssign(r IterCRef) {
	var r_ = r.(*intSliceIter)
	it.i, it.data = r_.i, r_.data
}

func (it *intSliceIter) Swap(r IterCRef) {
	var r_ = r.(*intSliceIter)
	it.i, r_.i = r_.i, it.i
	it.data, r_.data = r_.data, it.data
}

func (it *intSliceIter) Clone() IterCRef {
	return &intSliceIter{it.data, it.i}
}

func (it *intSliceIter) Deref() Value {
	return (*it.data)[it.i]
}

func (it *intSliceIter) DerefSet(val Value) {
	(*it.data)[it.i] = val
}

func (it *intSliceIter) Equal(r interface{}) bool {
	return it.i == r.(*intSliceIter).i
}

func (it *intSliceIter) Next() {
	it.i++
}

func (it *intSliceIter) Prev() {
	it.i--
}

func (it *intSliceIter) CanMultiPass() {}

func TestCopy(t *testing.T) {
	var is1 = new(intSlice)
	is1.data = make([]Value, 0)

	var begin = &intSliceIter{&is1.data, 0}
	is1.insert(begin, 1)
	is1.insert(begin, 2)
	is1.insert(begin, 3)
	is1.insert(begin, 4)
	assert.Equal(t, &intSlice{[]Value{4, 3, 2, 1}}, is1)

	var end = &intSliceIter{&is1.data, len(is1.data)}

	t.Run(`InsertIterator`, func(t *testing.T) {
		t.Run(`Empty`, func(t *testing.T) {
			var is2 = new(intSlice)
			is2.data = []Value{}
			var iit = iterator.NewInsertIterator(is2, &intSliceIter{&is2.data, 0})
			Copy(begin, end, iit)
			assert.Equal(t, &intSlice{[]Value{4, 3, 2, 1}}, is2)
		})

		t.Run(`Begin`, func(t *testing.T) {
			var is2 = new(intSlice)
			is2.data = []Value{5}
			var iit = iterator.NewInsertIterator(is2, &intSliceIter{&is2.data, 0})
			Copy(begin, end, iit)
			assert.Equal(t, &intSlice{[]Value{4, 3, 2, 1, 5}}, is2)
		})

		t.Run(`Middle`, func(t *testing.T) {
			var is2 = new(intSlice)
			is2.data = []Value{5, 6}
			var iit = iterator.NewInsertIterator(is2, &intSliceIter{&is2.data, 1})
			Copy(begin, end, iit)
			assert.Equal(t, &intSlice{[]Value{5, 4, 3, 2, 1, 6}}, is2)
		})

		t.Run(`End`, func(t *testing.T) {
			var is2 = new(intSlice)
			is2.data = []Value{5, 6}
			var iit = iterator.NewInsertIterator(is2, &intSliceIter{&is2.data, 2})
			Copy(begin, end, iit)
			assert.Equal(t, &intSlice{[]Value{5, 6, 4, 3, 2, 1}}, is2)
		})
	})

	t.Run(`BackInsertIterator`, func(t *testing.T) {
		var is2 = new(intSlice)
		is2.data = []Value{5, 6}
		var bit = iterator.NewBackInsertIterator(is2)
		Copy(begin, end, bit)
		assert.Equal(t, &intSlice{[]Value{5, 6, 4, 3, 2, 1}}, is2)
	})

	t.Run(`FrontInsertIterator`, func(t *testing.T) {
		var is2 = new(intSlice)
		is2.data = []Value{5, 6}
		var fit = iterator.NewFrontInsertIterator(is2)
		Copy(begin, end, fit)
		assert.Equal(t, &intSlice{[]Value{1, 2, 3, 4, 5, 6}}, is2)
	})

	t.Run(`CopyBackward`, func(t *testing.T) {
		var is = new(intSlice)
		is.data = make([]Value, 0)

		var begin = &intSliceIter{&is.data, 0}
		is.insert(begin, 1)
		is.insert(begin, 2)
		is.insert(begin, 3)
		is.insert(begin, 4)
		assert.Equal(t, &intSlice{[]Value{4, 3, 2, 1}}, is)
		var end = &intSliceIter{&is.data, 2}

		var it = begin.Clone().(*intSliceIter)
		it.i = 4
		var newit = CopyBackward(begin, end, it)
		assert.Equal(t, &intSlice{[]Value{4, 3, 4, 3}}, is)
		assert.Equal(t, 2, newit.(*intSliceIter).i)
	})
}

func TestFill(t *testing.T) {
	var is = new(intSlice)
	is.data = make([]Value, 5, 5)

	var begin = &intSliceIter{&is.data, 0}
	var end = &intSliceIter{&is.data, 5}

	Fill(begin, end, 5)
	assert.Equal(t, &intSlice{[]Value{5, 5, 5, 5, 5}}, is)
}
