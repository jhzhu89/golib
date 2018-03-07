package list

import (
	"testing"

	"github.com/jhzhu89/golib/container/testutil/vec"
	"github.com/jhzhu89/golib/fn"
	"github.com/jhzhu89/golib/iterator"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run(`New`, func(t *testing.T) {
		l := New()
		assert.NotNil(t, l.node)
		assert.Equal(t, l.node, l.node.prev)
		assert.Equal(t, l.node, l.node.next)
		assert.Equal(t, 0, l.node.val)
	})

	t.Run(`NewN`, func(t *testing.T) {
		var n = 1024
		l := NewN(n)
		assert.NotNil(t, l.node)
		assert.NotNil(t, l.node.prev)
		assert.NotNil(t, l.node.next)
		assert.Equal(t, n, l.node.val)

		var it = l.Begin()
		for i := 0; i < n; i++ {
			assert.Nil(t, it.Deref())
			it.Next()
		}
	})

	t.Run(`NewNValues`, func(t *testing.T) {
		var n = 1024
		l := NewNValues(n, 1)
		assert.NotNil(t, l.node)
		assert.NotNil(t, l.node.prev)
		assert.NotNil(t, l.node.next)
		assert.Equal(t, n, l.node.val)

		var it = l.Begin()
		for i := 0; i < n; i++ {
			assert.Equal(t, 1, it.Deref())
			it.Next()
		}
	})

	t.Run(`NewFromRange`, func(t *testing.T) {
		var n = 1024
		var v = make(vec.Vec, 0, n)
		for i := 0; i < n; i++ {
			v = append(v, i)
		}

		l := NewFromRange(vec.NewIt(0, &v), vec.NewIt(n, &v))
		var it = l.Begin()
		for i := 0; i < n; i++ {
			assert.Equal(t, i, it.Deref())
			it.Next()
		}
	})
}

func TestClear(t *testing.T) {
	l := NewNValues(1024, 1)
	l.clear()
	assert.NotNil(t, l.node)
	assert.Equal(t, 0, l.node.val)
	assert.Equal(t, l.node, l.node.next)
	assert.Equal(t, l.node, l.node.prev)
}

func TestDefaultInitialize(t *testing.T) {
	n, l := 1024, New()
	l.defaultInitialize(n)
	assert.Equal(t, n, l.node.val)
	it := l.Begin()
	for i := 0; i < n; i++ {
		assert.Nil(t, it.Deref())
		it.Next()
	}
}

func TestFillInitialize(t *testing.T) {
	n, l := 1024, New()
	l.fillInitialize(n, 1)
	assert.Equal(t, n, l.node.val)
	it := l.Begin()
	for i := 0; i < n; i++ {
		assert.Equal(t, 1, it.Deref())
		it.Next()
	}
}

func TestErase(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	l.erase(l.Begin())
	assert.Equal(t, 4, l.Size())
	it := l.Begin()
	for i := 0; i < 4; i++ {
		assert.Equal(t, i+1, it.Deref())
		it.Next()
	}
}

func TestInsert(t *testing.T) {
	t.Run(``, func(t *testing.T) {
		l := New()
		l.insert(l.Begin(), 1)
		assert.Equal(t, 1, l.Size())
		assert.Equal(t, 1, l.Begin().Deref())

		l.insert(l.Begin(), 2)
		assert.Equal(t, 2, l.Size())
		assert.Equal(t, 2, l.Begin().Deref())

		l.insert(l.End(), 3)
		assert.Equal(t, 3, l.Size())
		assert.Equal(t, 3, l.End().Deref())
	})

	t.Run(``, func(t *testing.T) {
		l := New()
		l.insert(l.End(), 1)
		assert.Equal(t, 1, l.Size())
		assert.Equal(t, 1, l.Begin().Deref())
	})
}

func TestSpliceList(t *testing.T) {
	t.Run(`Begin`, func(t *testing.T) {
		l1, l2, n := New(), New(), 5
		for i := 0; i < n; i++ {
			l1.PushBack(i)
			l2.PushBack(i + n)
		}

		l1.spliceList(l1.Begin(), l2)
		assert.Equal(t, 0, l2.Size())
		assert.Equal(t, n*2, l1.Size())

		it := l1.Begin()
		for _, v := range []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4} {
			assert.Equal(t, v, it.Deref())
			it.Next()
		}
	})

	t.Run(`End`, func(t *testing.T) {
		l1, l2, n := New(), New(), 5
		for i := 0; i < n; i++ {
			l1.PushBack(i)
			l2.PushBack(i + n)
		}

		l1.spliceList(l1.End(), l2)
		assert.Equal(t, 0, l2.Size())
		assert.Equal(t, n*2, l1.Size())

		it := l1.Begin()
		for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			assert.Equal(t, v, it.Deref())
			it.Next()
		}
	})

	t.Run(`Mid`, func(t *testing.T) {
		l1, l2, n := New(), New(), 5
		for i := 0; i < n; i++ {
			l1.PushBack(i)
			l2.PushBack(i + n)
		}

		l1.spliceList(l1.Begin().Next2().Next2(), l2)
		assert.Equal(t, 0, l2.Size())
		assert.Equal(t, n*2, l1.Size())

		it := l1.Begin()
		for _, v := range []int{0, 1, 5, 6, 7, 8, 9, 2, 3, 4} {
			assert.Equal(t, v, it.Deref())
			it.Next()
		}
	})
}

func TestSplice(t *testing.T) {
	check := func(l1, l2 *List, s1, s2 int, e1, e2 []int) {
		assert.Equal(t, s1, l1.Size())
		assert.Equal(t, s2, l2.Size())

		it := l1.Begin()
		for _, v := range e1 {
			assert.Equal(t, v, it.Deref())
			it.Next()
		}

		it = l2.Begin()
		for _, v := range e2 {
			assert.Equal(t, v, it.Deref())
			it.Next()
		}
	}

	t.Run(`SpliceElement`, func(t *testing.T) {
		l1, l2, n := New(), New(), 5
		for i := 0; i < n; i++ {
			l1.PushBack(i)
			l2.PushBack(i + n)
		}

		l1.spliceElement(l1.Begin(), l2, l2.Begin())
		check(l1, l2, n+1, n-1, []int{5, 0, 1, 2, 3, 4}, []int{6, 7, 8, 9})

		l1.spliceElement(l1.End(), l2, l2.End().Prev2())
		check(l1, l2, n+2, n-2, []int{5, 0, 1, 2, 3, 4, 9}, []int{6, 7, 8})

		l1.spliceElement(l1.Begin().Next2().Next2(), l2, l2.End().Prev2().Prev2())
		check(l1, l2, n+3, n-3, []int{5, 0, 7, 1, 2, 3, 4, 9}, []int{6, 8})
	})

	t.Run(`RangeSplice`, func(t *testing.T) {
		l1, l2, n := New(), New(), 5
		for i := 0; i < n; i++ {
			l1.PushBack(i)
			l2.PushBack(i + n)
		}

		l1.rangeSplice(l1.Begin(), l2, l2.Begin(), l2.Begin().Next2())
		check(l1, l2, n+1, n-1, []int{5, 0, 1, 2, 3, 4}, []int{6, 7, 8, 9})

		l1.rangeSplice(l1.End(), l2, l2.Begin(), l2.End())
		check(l1, l2, n*2, 0, []int{5, 0, 1, 2, 3, 4, 6, 7, 8, 9}, []int{})
	})
}

func TestFillAssign(t *testing.T) {
	l, n := New(), 1024

	l.fillAssign(n, 1)
	assert.Equal(t, n, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, 1, it.Deref())
	}

	l.fillAssign(n*2, 2)
	assert.Equal(t, n*2, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, 2, it.Deref())
	}

	l.fillAssign(n/2, 3)
	assert.Equal(t, n/2, l.Size())
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, 3, it.Deref())
	}
}

func TestRangeAssign(t *testing.T) {
	l, n, l1 := New(), 1024, NewNValues(1024, 1)

	l.rangeAssign(l1.Begin(), l1.End().Prev2())
	assert.Equal(t, n-1, l.Size())
	assert.Equal(t, n, l1.Size())

	l.rangeAssign(l1.Begin(), l1.Begin().Next2())
	assert.Equal(t, 1, l.Size())
	assert.Equal(t, n, l1.Size())
}

func TestMerge(t *testing.T) {
	l1, l2, n := New(), New(), 1024
	for i := 0; i < n; i++ {
		l1.PushBack(i * 2)
		l2.PushBack(i*2 + 1)
	}

	l1.merge(l2, fn.CompareFunc(func(a, b interface{}) bool { return a.(int) < b.(int) }))

	assert.Equal(t, n*2, l1.Size())
	assert.Equal(t, 0, l2.Size())

	var i int
	for it := l1.Begin(); !it.EqualTo(l1.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}
}

func TestSort(t *testing.T) {
	l, n := New(), 1024
	for i := 0; i < n; i++ {
		l.PushBack(n - i - 1)
	}

	l.sort(fn.CompareFunc(func(a, b interface{}) bool { return a.(int) < b.(int) }))
	var i int
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, i, it.Deref())
		i++
	}

	l.sort(fn.CompareFunc(func(a, b interface{}) bool { return a.(int) > b.(int) }))
	i = 0
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		assert.Equal(t, n-i-1, it.Deref())
		i++
	}
}

func TestResizePos(t *testing.T) {
	t.Run(`Empty`, func(t *testing.T) {
		l, n := New(), 1024
		it, sz := l.resizePos(n)
		assert.True(t, l.End().EqualTo(it))
		assert.Equal(t, n, sz)
	})

	t.Run(`NonEmpty`, func(t *testing.T) {
		n := 1024
		l := NewN(n / 2)
		it, sz := l.resizePos(n)
		assert.True(t, l.End().EqualTo(it))
		assert.Equal(t, n/2, sz)
	})
}

func TestDefaultAppend(t *testing.T) {
	t.Run(`Empty`, func(t *testing.T) {
		l, n := New(), 1024
		l.defaultAppend(n)
		assert.Equal(t, n, l.Size())
		for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
			assert.Nil(t, it.Deref())
		}
	})

	t.Run(`NonEmpty`, func(t *testing.T) {
		n := 1024
		l := NewNValues(n/2, 1)
		l.defaultAppend(n)
		assert.Equal(t, n+n/2, l.Size())
		it := l.Begin()
		iterator.Advance(it, n/2)
		for ; !it.EqualTo(l.End()); it.Next() {
			assert.Nil(t, it.Deref())
		}
	})
}

func TestDistance(t *testing.T) {
	l := New()
	assert.Equal(t, 0, l.distance(l.Begin().node, l.End().node))
	l.PushBack(1)
	assert.Equal(t, 1, l.distance(l.Begin().node, l.End().node))
	l.PushBack(1)
	assert.Equal(t, 2, l.distance(l.Begin().node, l.End().node))
	assert.Equal(t, 1, l.distance(l.Begin().node, l.Begin().Next2().node))
}

func TestListNode(t *testing.T) {
	newNode := func(v interface{}) *listNode {
		return &listNode{val: v}
	}

	link := func(n1, n2 *listNode) {
		n1.next, n2.prev = n2, n1
	}

	list1 := func() *listNode {
		// 1->2->3
		var h1 = newNode(nil)
		var n1 = newNode(1)
		var n2 = newNode(2)
		var n3 = newNode(3)
		link(h1, n1)
		link(n1, n2)
		link(n2, n3)
		link(n3, h1)
		return h1
	}

	list2 := func() *listNode {
		// 4->5->6
		var h2 = newNode(nil)
		var n4 = newNode(4)
		var n5 = newNode(5)
		var n6 = newNode(6)
		link(h2, n4)
		link(n4, n5)
		link(n5, n6)
		link(n6, h2)
		return h2
	}

	countElements := func(l *listNode) int {
		var c int
		for it := l; it.next != l; it = it.next {
			c++
		}
		return c
	}

	check := func(l *listNode, expected []int) {
		var it = l
		for _, e := range expected {
			assert.Equal(t, e, it.next.val)
			it = it.next
		}
	}

	t.Run(`Transfer`, func(t *testing.T) {
		var h1 = list1()
		var h2 = list2()

		h1.next.transfer(h2.next, h2)
		assert.Equal(t, 0, countElements(h2))
		assert.Equal(t, 6, countElements(h1))
		check(h1, []int{4, 5, 6, 1, 2, 3})

		h1.next.transfer(h1.next.next, h1.next.next.next)
		assert.Equal(t, 6, countElements(h1))
		check(h1, []int{5, 4, 6, 1, 2, 3})
	})

	t.Run(`Transfer`, func(t *testing.T) {
		var h1 = list1()
		var h2 = list2()

		h1.next.transfer(h2.next, h2.next.next)
		assert.Equal(t, 2, countElements(h2))
		assert.Equal(t, 4, countElements(h1))
		check(h1, []int{4, 1, 2, 3})
		check(h2, []int{5, 6})
	})

	t.Run(`Reverse`, func(t *testing.T) {
		var h1 = list1()
		h1.reverse()
		assert.Equal(t, 3, countElements(h1))
		check(h1, []int{3, 2, 1})
	})

	t.Run(`Hook`, func(t *testing.T) {
		var h1 = list1()
		var n = &listNode{val: 4}
		n.hook(h1.next)
		assert.Equal(t, 4, countElements(h1))
		check(h1, []int{4, 1, 2, 3})
	})

	t.Run(`Unhook`, func(t *testing.T) {
		var h1 = list1()
		h1.next.unhook()
		assert.Equal(t, 2, countElements(h1))
		check(h1, []int{2, 3})

		h1.next.next.unhook()
		assert.Equal(t, 1, countElements(h1))
		check(h1, []int{2})

		h1.next.unhook()
		assert.Equal(t, 0, countElements(h1))
	})
}
