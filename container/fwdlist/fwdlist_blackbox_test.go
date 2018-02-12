package fwdlist_test

import (
	"testing"

	"github.com/jhzhu89/golib/container/fwdlist"
	"github.com/jhzhu89/golib/fn"
	"github.com/stretchr/testify/assert"
)

func TestFront(t *testing.T) {
	fl := fwdlist.NewN(1)
	assert.Nil(t, fl.Front())
	fl.InsertAfter(fl.BeforeBegin(), 1)
	assert.Equal(t, 1, fl.Front())
}

func TestEmpty(t *testing.T) {
	fl := fwdlist.New()
	assert.True(t, fl.Empty())
	fl.InsertAfter(fl.BeforeBegin(), 1)
	assert.False(t, fl.Empty())
}

func TestClear(t *testing.T) {
	fl := fwdlist.NewNValues(1024, 1)
	assert.False(t, fl.Empty())
	fl.Clear()
	assert.True(t, fl.Empty())
}

func TestPushAndPopFront(t *testing.T) {
	fl := fwdlist.NewNValues(1024, 1)
	fl.PushFront(2)
	assert.Equal(t, 2, fl.Front())
	fl.PushFront(3)
	assert.Equal(t, 3, fl.Front())

	fl.PopFront()
	assert.Equal(t, 2, fl.Front())
	fl.PopFront()
	assert.Equal(t, 1, fl.Front())
}

func TestResize(t *testing.T) {
	fl := fwdlist.New()
	assert.Equal(t, 0, fwdlist.Count(fl))
	fl.Resize(1024)
	assert.Equal(t, 1024, fwdlist.Count(fl))
	fl.Resize(512)
	assert.Equal(t, 512, fwdlist.Count(fl))
}

func TestFillResize(t *testing.T) {
	fl := fwdlist.New()
	assert.Equal(t, 0, fwdlist.Count(fl))
	fl.FillResize(1024, 1)
	assert.Equal(t, 1024, fwdlist.Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, 1, it.Deref())
	}
	fl.FillResize(512, 2)
	assert.Equal(t, 512, fwdlist.Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, 1, it.Deref())
	}
}

func TestFillAssign(t *testing.T) {
	fl := fwdlist.New()
	fl.AssignN(1024, 1)
	assert.Equal(t, 1024, fwdlist.Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, 1, it.Deref())
	}

	fl.AssignN(512, 2)
	assert.Equal(t, 512, fwdlist.Count(fl))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.Equal(t, 2, it.Deref())
	}
}

func TestRemove(t *testing.T) {
	var n, val = 1024, 1
	fl := fwdlist.NewNValues(n, val)
	fl.Remove(val)
	assert.True(t, fl.Empty())

	for i := 0; i < 10; i++ {
		fl.InsertAfter(fl.BeforeBegin(), i)
	}
	fl.Remove(5)
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.NotEqual(t, 5, it.Deref())
	}
}

func TestRemoveIf(t *testing.T) {
	fl := fwdlist.New()
	for i := 0; i < 10; i++ {
		fl.InsertAfter(fl.BeforeBegin(), i)
	}
	fl.RemoveIf(fn.PredicateFunc(func(v interface{}) bool { return v.(int) > 5 }))
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		assert.True(t, it.Deref().(int) <= 5)
	}
}

func TestUnique(t *testing.T) {
	fl := fwdlist.NewNValues(10, 1)
	fl.Unique()
	assert.Equal(t, 1, fwdlist.Count(fl))

	fl.InsertNAfter(fl.BeforeBegin(), 10, 2)
	fl.InsertNAfter(fl.BeforeBegin(), 10, 3)
	fl.Unique()
	assert.Equal(t, 3, fwdlist.Count(fl))

	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		var i = it.Deref().(int)
		assert.True(t, i == 1 || i == 2 || i == 3)
	}
}

func TestUniqueIf(t *testing.T) {
	fl := fwdlist.New()
	n := 10
	for i := 0; i < n; i++ {
		fl.InsertAfter(fl.BeforeBegin(), i)
	}

	assert.Equal(t, n, fwdlist.Count(fl))

	fl.UniqueIf(fn.BinaryPredicateFunc(func(a, b interface{}) bool { return a.(int)-b.(int) == 1 }))
	assert.Equal(t, n/2, fwdlist.Count(fl))

	var nums []int
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		nums = append(nums, it.Deref().(int))
	}

	assert.Equal(t, []int{9, 7, 5, 3, 1}, nums)
}
