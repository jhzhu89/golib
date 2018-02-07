package fwdlist

import (
	"testing"

	"github.com/jhzhu89/golib/container/testutil/vec"
	"github.com/jhzhu89/golib/fn"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	t.Run(`default`, func(t *testing.T) {
		var n = 1024
		fl := New()
		fl.defaultInitialize(n)
		var c int
		for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
			assert.Nil(t, it.Deref())
			c++
		}
		assert.Equal(t, n, c)
	})

	t.Run(`fill`, func(t *testing.T) {
		var n = 1024
		fl := New()
		fl.fillInitialize(n, 1)
		var c int
		for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
			assert.Equal(t, 1, it.Deref())
			c++
		}
		assert.Equal(t, n, c)
	})

	t.Run(`range`, func(t *testing.T) {
		var n = 1024
		var v = make(vec.Vec, 0, n)
		for i := 0; i < n; i++ {
			v = append(v, i)
		}

		fl := New()
		fl.rangeInitialize(vec.NewIt(0, &v), vec.NewIt(n, &v))
		var c int
		for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
			assert.Equal(t, c, it.Deref())
			c++
		}
		assert.Equal(t, n, c)
	})
}

func TestEraseAfter(t *testing.T) {
	var newVec = func(n int) func(int) *vec.VecIter {
		var v = make(vec.Vec, 0, n)
		for i := 0; i < n; i++ {
			v = append(v, i)
		}
		return func(i int) *vec.VecIter { return vec.NewIt(i, &v) }
	}

	t.Run(`EraseAfter`, func(t *testing.T) {
		var n = 10
		var newIt = newVec(n)
		fl := New()
		fl.rangeInitialize(newIt(0), newIt(n))
		for i := 1; i < n; i++ {
			var node = fl.eraseAfter(fl.BeforeBegin().node)
			assert.Equal(t, i, node.val)
		}
	})

	t.Run(`RangeEraseAfter`, func(t *testing.T) {
		var n = 10
		var newIt = newVec(n)
		fl := New()
		fl.rangeInitialize(newIt(0), newIt(n))

		var node = fl.rangeEraseAfter(fl.BeforeBegin().node, fl.End().node)
		assert.Nil(t, node)
	})
}

func TestSpliceAfter(t *testing.T) {
	fl1 := New()
	fl2 := New()
	fl2.insertAfter(fl2.BeforeBegin(), 2)
	fl2.insertAfter(fl2.BeforeBegin(), 3)

	var it = fl1.spliceAfter(fl1.BeforeBegin(), fl2.BeforeBegin(), fl2.End())
	assert.Equal(t, 2, it.node.val)
	it = it.next()
	assert.Nil(t, it.node)

	assert.Equal(t, 3, fl1.Begin().Deref())
	assert.Equal(t, 2, fl1.Begin().Next2().Deref())

	fl3 := New()
	fl3.insertAfter(fl3.BeforeBegin(), 4)
	fl1.spliceElementAfter(fl1.BeforeBegin(), fl3.BeforeBegin())
	assert.Equal(t, 4, fl1.Begin().Deref())
}

func Count(fl *ForwardList) int {
	var c int
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		c++
	}
	return c
}

func TestAssignN(t *testing.T) {
	var n, v1, v2 = 1024, 1, 2
	fl := New()
	fl.assignN(n, v1)
	for i := 0; i < n; i++ {
		it := fl.Begin()
		for j := 0; j < i; j++ {
			it.Next()
		}
		assert.Equal(t, v1, it.Deref())
	}
	assert.Equal(t, n, Count(fl))

	fl.assignN(n/2, v2)
	for i := 0; i < n/2; i++ {
		it := fl.Begin()
		for j := 0; j < i; j++ {
			it.Next()
		}
		assert.Equal(t, v2, it.Deref())
	}
	assert.Equal(t, n/2, Count(fl))
}

func TestRangeAssign(t *testing.T) {
	var n, v1, v2 = 1024, 1, 2
	fl := New()
	fl1 := NewNValues(n, v1)
	assert.Equal(t, n, Count(fl1))
	fl.rangeAssign(fl1.Begin(), fl1.End())
	assert.Equal(t, n, Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, v1, it.Deref())
	}

	fl2 := NewNValues(n/2, v2)
	assert.Equal(t, n/2, Count(fl2))
	fl.rangeAssign(fl2.Begin(), fl2.End())
	assert.Equal(t, n/2, Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, v2, it.Deref())
	}
}

func TestMerge(t *testing.T) {
	fl1 := New()
	fl2 := New()
	for i := 5; i > 0; i-- {
		fl1.insertAfter(fl1.BeforeBegin(), i*2)
		fl2.insertAfter(fl2.BeforeBegin(), i*2-1)
	}

	fl1.merge(fl2, fn.CompareFunc(func(a, b interface{}) bool { return a.(int) < b.(int) }))

	var i = 1
	for it := fl1.Begin(); !it.EqualTo(fl1.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}
}

func TestSort(t *testing.T) {
	var n = 1024
	fl := New()
	for i := 0; i < n; i++ {
		fl.insertAfter(fl.BeforeBegin(), i)
	}

	fl.sort(fn.CompareFunc(func(a, b interface{}) bool { return a.(int) < b.(int) }))

	var i int
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}

	fl.sort(fn.CompareFunc(func(a, b interface{}) bool { return a.(int) > b.(int) }))
	i = 0
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, n-i-1, it.Deref())
		i++
	}
}

func TestTransferAfter(t *testing.T) {
	var n = 5
	fl1 := New()
	for i := 0; i < n; i++ {
		fl1.insertAfter(fl1.BeforeBegin(), n-i-1)
	}

	fl2 := New()
	var node = fl2.BeforeBegin().node.transferAfter(fl1.BeforeBegin().node, fl1.Begin().node)
	assert.NotNil(t, node)
	assert.Equal(t, 0, node.val)
	assert.Equal(t, n-1, Count(fl1))
	assert.Equal(t, 1, fl1.Begin().Deref())
	assert.Equal(t, 1, Count(fl2))
	assert.Equal(t, 0, fl2.Begin().Deref())

	node = fl2.BeforeBegin().node.transferAfter(fl1.Begin().node, fl1.End().node)
	assert.Nil(t, node)
	assert.Equal(t, 1, Count(fl1))
	// the 'end' node is null, so the one element in fl2 is dropped, the size should be n-2.
	assert.Equal(t, n-2, Count(fl2))
}

func TestReverseAfter(t *testing.T) {
	var n = 1024
	fl := New()
	for i := 0; i < n; i++ {
		fl.insertAfter(fl.BeforeBegin(), i)
	}

	fl.BeforeBegin().node.reverseAfter()
	var i int
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}
	assert.Equal(t, i, n)
}
