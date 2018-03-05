package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
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

	t.Run(`1`, func(t *testing.T) {
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

	t.Run(`2`, func(t *testing.T) {
		var h1 = list1()
		var h2 = list2()

		h1.next.transfer(h2.next, h2.next.next)
		assert.Equal(t, 2, countElements(h2))
		assert.Equal(t, 4, countElements(h1))
		check(h1, []int{4, 1, 2, 3})
		check(h2, []int{5, 6})
	})
}
