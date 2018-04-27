package list_test

import (
	"testing"

	"github.com/jhzhu89/golib/container/list"
	"github.com/jhzhu89/golib/container/testutil/vec"
	"github.com/jhzhu89/golib/fn"
	"github.com/jhzhu89/golib/iterator"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	t.Run(`ReverseIter`, func(t *testing.T) {
		l := list.New()
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)

		assert.Equal(t, 3, l.RBegin().Deref())
		re := l.REnd()
		re.Prev()
		assert.Equal(t, 1, re.Deref())

		it := l.RBegin()
		for _, n := range []int{3, 2, 1} {
			assert.Equal(t, n, it.Deref())
			it.Next()
		}

		i := 3
		for it := l.RBegin(); !it.EqualTo(l.REnd()); it.Next() {
			assert.Equal(t, i, it.Deref())
			i--
		}
	})
}

func TestElementAccess(t *testing.T) {
	t.Run(`Front`, func(t *testing.T) {
		l := list.New()
		l.PushBack(1)
		assert.Equal(t, 1, l.Front())
		l.PushFront(2.0)
		assert.Equal(t, 2.0, l.Front())
		l.PushFront(nil)
		assert.Nil(t, l.Front())
	})

	t.Run(`Back`, func(t *testing.T) {
		l := list.New()
		l.PushBack(1)
		assert.Equal(t, 1, l.Back())
		l.PushBack(2.0)
		assert.Equal(t, 2.0, l.Back())
		l.PushBack(nil)
		assert.Nil(t, l.Back())
	})
}

func TestEmpty(t *testing.T) {
	l := list.New()
	assert.True(t, l.Empty())
	l.PushBack(1)
	assert.False(t, l.Empty())
	l.PopBack()
	assert.True(t, l.Empty())
}

func TestSize(t *testing.T) {
	l, n := list.New(), 1024
	for i := 0; i < n; i++ {
		l.PushBack(i)
		assert.Equal(t, i+1, l.Size())
	}
	for i := n; i > 0; i-- {
		l.PopBack()
		assert.Equal(t, i-1, l.Size())
	}
}

func TestRangeInsert(t *testing.T) {
	var n = 1024
	var v = make(vec.Vec, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, i)
	}

	l := list.New()
	l.RangeInsert(l.Begin(), vec.NewIt(0, &v), vec.NewIt(n, &v))
	assert.Equal(t, n, l.Size())
	i := 0
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}

	b := l.Begin()
	iterator.Advance(b, n/2)
	b = l.RangeInsert(b, vec.NewIt(0, &v), vec.NewIt(n, &v))
	assert.Equal(t, 0, b.Deref())
	assert.Equal(t, n*2, l.Size())

	for i := 0; i < n; i++ {
		assert.Equal(t, i, b.Deref())
		b.Next()
	}
}

func TestFillInsert(t *testing.T) {
	l, n, v1, v2 := list.New(), 1024, 1, 2
	l.FillInsert(l.Begin(), n, v1)
	assert.Equal(t, n, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, v1, it.Deref())
	}

	b := l.Begin()
	iterator.Advance(b, n/2)
	b = l.FillInsert(b, n, v2)
	assert.Equal(t, 2, b.Deref())
	assert.Equal(t, n*2, l.Size())

	for i := 0; i < n; i++ {
		assert.Equal(t, v2, b.Deref())
		b.Next()
	}
}

func TestErase(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Erase(l.Begin())
	assert.Equal(t, 2, l.Size())
	assert.Equal(t, 2, l.Begin().Deref())
	l.Erase(l.End().Prev2())
	assert.Equal(t, 1, l.Size())
	assert.Equal(t, 2, l.Begin().Deref())
	l.Erase(l.End().Prev2())
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.Empty())
}

func TestRangeErase(t *testing.T) {
	l, n := list.New(), 1024
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}

	b, e := l.Begin(), l.Begin()
	iterator.Advance(b, n/4)
	iterator.Advance(e, n/2)
	l.RangeErase(b, e)
	assert.Equal(t, 3*n/4, l.Size())
	it := l.Begin()
	for i := 0; i < n; i++ {
		if i >= n/4 && i < n/2 {
			continue
		}
		assert.Equal(t, i, it.Deref())
		it.Next()
	}

	l.RangeErase(l.Begin(), l.End())
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.Empty())
}

func TestPushPop(t *testing.T) {
	l := list.New()
	l.PushFront(1)
	l.PushBack(2)
	l.PushFront(3)
	assert.Equal(t, 3, l.Front())
	assert.Equal(t, 2, l.Back())
	l.PopBack()
	assert.Equal(t, 1, l.Back())
	l.PopFront()
	assert.Equal(t, 1, l.Front())
}

func TestResize(t *testing.T) {
	l, n := list.New(), 1024
	l.Resize(n)
	assert.Equal(t, n, l.Size())
	l.Resize(n / 2)
	assert.Equal(t, n/2, l.Size())
	l.Resize(0)
	assert.Equal(t, 0, l.Size())
	assert.True(t, l.Empty())

	v := 1
	l.FillResize(n, v)
	assert.Equal(t, n, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, v, it.Deref())
	}

	l.FillResize(n/2, v*2)
	assert.Equal(t, n/2, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, v, it.Deref())
	}

	l.FillResize(n, v*2)
	assert.Equal(t, n, l.Size())
	i := 0
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		if i < n/2 {
			assert.Equal(t, v, it.Deref())
		} else {
			assert.Equal(t, v*2, it.Deref())
		}
		i++
	}
}

func TestRemove(t *testing.T) {
	l, n := list.New(), 1024
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			l.PushBack('a')
		} else {
			l.PushBack(2)
		}
	}

	l.Remove('a')
	assert.Equal(t, n/2, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, 2, it.Deref())
	}
}

func TestRemoveIf(t *testing.T) {
	l, n := list.New(), 1024
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			l.PushBack('a')
		} else {
			l.PushBack(2)
		}
	}

	l.RemoveIf(fn.PredicateFunc(
		func(v interface{}) bool {
			_, ok := v.(int)
			return ok
		}))
	assert.Equal(t, n/2, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, 'a', it.Deref())
	}
}

func TestUnique(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack('c')
	l.PushBack('c')
	l.PushBack(nil)
	l.PushBack(nil)

	l.Unique()
	assert.Equal(t, 3, l.Size())
	assert.Equal(t, 1, l.Begin().Deref())
	b := l.Begin()
	b.Next()
	assert.Equal(t, 'c', b.Deref())
	assert.Nil(t, l.End().Prev2().Deref())
}

func TestUniqueIf(t *testing.T) {
	l := list.New()
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(4)
	l.PushBack(6)
	l.PushBack(3)
	l.PushBack(8)

	l.UniqueIf(fn.BinaryPredicateFunc(func(a, b interface{}) bool { return b.(int)%a.(int) == 0 }))
	assert.Equal(t, 4, l.Size())
	it := l.Begin()
	for _, n := range []int{3, 2, 3, 8} {
		assert.Equal(t, n, it.Deref())
		it.Next()
	}
}
