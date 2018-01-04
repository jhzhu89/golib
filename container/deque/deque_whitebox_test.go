package deque

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRellocateMap(t *testing.T) {
	var test = func(nodesToAdd int, addAtFront bool, mapSize, startNode int) {
		var d = NewDeque()
		d.reallocateMap(nodesToAdd, addAtFront)

		assert.Equal(t, mapSize, d.mapSize)
		assert.Equal(t, startNode, d.start.node)

		var nonNil *node
		for _, node := range *d.map_ {
			if node != nil {
				if nonNil == nil {
					nonNil = node
				} else {
					assert.Equal(t, fmt.Sprintf("%p", nonNil), fmt.Sprintf("%p", node))
				}
			}
		}
		assert.NotNil(t, nonNil)
	}

	test(1, false, 8, 3)
	test(1, true, 8, 4)
	test(11, false, 21, 4)
	test(11, true, 21, 15)
}

func TestFillInsert(t *testing.T) {
	var test = func(d *Deque, pos *DequeIter, n, startCur, startNode, finishCur,
		finishNode, size, numNonEmptyNodes int, empty bool) {
		d.fillInsert(pos, n, 1)

		assert.Equal(t, startCur, d.start.cur)
		assert.Equal(t, startNode, d.start.node)

		assert.Equal(t, finishCur, d.finish.cur)
		assert.Equal(t, finishNode, d.finish.node)

		assert.Equal(t, size, d.Size())
		assert.Equal(t, empty, d.Empty())

		var nonEmpty int
		for i := 0; i < len(*d.map_); i++ {
			if (*d.map_)[i] != nil {
				nonEmpty++
			}
		}
		assert.Equal(t, numNonEmptyNodes, nonEmpty)
	}

	var d = NewDeque()
	test(d, d.Begin(), 1, DequeBufSize-1, 2, 0, 3, 1, 2, false)
	d = NewDeque()
	test(d, d.Begin(), DequeBufSize, 0, 2, 0, 3, DequeBufSize, 2, false)

	d = NewDequeN(1)
	test(d, d.End(), 1, 0, 3, 2, 3, 2, 1, false)
	d = NewDequeN(1)
	test(d, d.End(), DequeBufSize, 0, 3, 1, 4, DequeBufSize+1, 2, false)

	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), 1, 0, 3, 11, 3, 11, 1, false)
	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), DequeBufSize, 0, 3, 10, 4, DequeBufSize+10, 2, false)

	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), 8*DequeBufSize, 0, 4, 10, 12, 8*DequeBufSize+10, 9, false)
}

func TestNewElementsAt(t *testing.T) {
	var test = func(front bool, elems, numNonEmptyNodes int) {
		var d = NewDeque()
		if front {
			d.newElementsAtFront(elems)
		} else {
			d.newElementsAtBack(elems)
		}

		var nonEmpty int
		for i := 0; i < len(*d.map_); i++ {
			if (*d.map_)[i] != nil {
				nonEmpty++
			}
		}
		assert.Equal(t, numNonEmptyNodes, nonEmpty)
	}

	t.Run(`Front`, func(t *testing.T) {
		test(true, 1, 2)
		test(true, DequeBufSize, 2)
		test(true, DequeBufSize+1, 3)
	})

	t.Run(`Back`, func(t *testing.T) {
		test(false, 1, 2)
		test(false, DequeBufSize, 2)
		test(false, DequeBufSize+1, 3)
	})
}

func TestDestroyData(t *testing.T) {
	var d = NewDequeN(2 * DequeBufSize)
	d.FillAssign(2*DequeBufSize, 10)

	var countNonNil = func() int {
		var c int
		for it := d.Begin(); !it.Equal(d.End()); it.Next() {
			if it.Deref() != nil {
				c++
			}
		}
		return c
	}

	assert.Equal(t, 2*DequeBufSize, countNonNil())

	d.destroyData(d.Begin(), d.End())
	assert.Equal(t, 2*DequeBufSize, d.Size())
	assert.Equal(t, 0, countNonNil())
}

func TestEraseAt(t *testing.T) {
	var countNonNilNodes = func(d *Deque) int {
		var n int
		for _, node := range *d.map_ {
			if node != nil {
				n++
			}
		}
		return n
	}

	t.Run(`Begin`, func(t *testing.T) {
		var d = NewDequeN(2 * DequeBufSize)
		d.FillAssign(2*DequeBufSize, 1)

		var pos = nextN(d.Begin(), DequeBufSize-1)
		d.eraseAtBegin(pos)
		assert.Equal(t, DequeBufSize-1, d.start.cur)
		assert.Equal(t, 3, countNonNilNodes(d))
		pos = next(pos)
		d.eraseAtBegin(pos)
		assert.Equal(t, 0, d.start.cur)
		assert.Equal(t, 2, countNonNilNodes(d))
	})

	t.Run(`End`, func(t *testing.T) {
		var d = NewDequeN(2 * DequeBufSize)
		d.FillAssign(2*DequeBufSize, 1)

		var pos = nextN(d.Begin(), DequeBufSize)
		d.eraseAtEnd(pos)
		assert.Equal(t, 0, d.finish.cur)
		assert.Equal(t, 2, countNonNilNodes(d))
		pos = prev(pos)
		d.eraseAtEnd(pos)
		assert.Equal(t, DequeBufSize-1, d.finish.cur)
		assert.Equal(t, 1, countNonNilNodes(d))
	})
}

func TestDequeMethodsWhitebox(t *testing.T) {
	t.Run(`BeginAndEnd`, func(t *testing.T) {
		var check = func(l, r *DequeIter) {
			assert.Equal(t, l.map_, r.map_)
			assert.Equal(t, l.cur, r.cur)
			assert.Equal(t, l.node, r.node)
			assert.NotEqual(t, fmt.Sprintf("%x", &l), fmt.Sprintf("%x", &r))
		}

		var d = NewDeque()
		check(d.start, d.Begin())
		check(d.finish, d.End())
	})

	t.Run(`Resize`, func(t *testing.T) {
		test := func(toSize, numNonEmptyNodes, startCur, startNode, finishCur,
			finishNode, size int, empty bool) {
			d := NewDeque()
			d.Resize(toSize)
			var nonEmpty int
			for i := 0; i < len(*d.map_); i++ {
				if (*d.map_)[i] != nil {
					nonEmpty++
				}
			}
			assert.Equal(t, numNonEmptyNodes, nonEmpty)

			assert.Equal(t, startCur, d.start.cur)
			assert.Equal(t, startNode, d.start.node)

			assert.Equal(t, finishCur, d.finish.cur)
			assert.Equal(t, finishNode, d.finish.node)

			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())
		}

		test(1, 1, 0, 3, 1, 3, 1, false)
		test(DequeBufSize-1, 1, 0, 3, DequeBufSize-1, 3, DequeBufSize-1, false)
		test(DequeBufSize, 2, 0, 3, 0, 4, DequeBufSize, false)
		test(2*DequeBufSize-1, 2, 0, 3, 511, 4, 2*DequeBufSize-1, false)
		test(2*DequeBufSize, 3, 0, 3, 0, 5, 2*DequeBufSize, false)
		// cause reallocate map
		test(8*DequeBufSize, 9, 0, 4, 0, 12, DequeBufSize*8, false)
	})

	t.Run(`PushOrPop`, func(t *testing.T) {
		const (
			pushBack = iota
			pushFront
			popBack
			popFront
		)

		test := func(d *Deque, op int, numNonEmptyNodes, startCur, startNode, finishCur, finishNode, size int, empty bool) {
			switch op {
			case pushBack:
				d.PushBack(1)
			case pushFront:
				d.PushFront(1)
			case popBack:
				d.PopBack()
			case popFront:
				d.PopFront()
			}

			assert.Equal(t, startCur, d.start.cur)
			assert.Equal(t, startNode, d.start.node)

			assert.Equal(t, finishCur, d.finish.cur)
			assert.Equal(t, finishNode, d.finish.node)

			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())

			var nonEmpty int
			for _, node := range *d.map_ {
				if node != nil && *node != nil {
					nonEmpty++
				}
			}
			assert.Equal(t, numNonEmptyNodes, nonEmpty)
		}

		t.Run(`Push`, func(t *testing.T) {
			t.Run(`Back`, func(t *testing.T) {
				test(NewDeque(), pushBack, 1, 0, 3, 1, 3, 1, false)
				test(NewDequeN(DequeBufSize-1), pushBack, 2, 0, 3, 0, 4, DequeBufSize, false)
			})

			t.Run(`Front`, func(t *testing.T) {
				test(NewDeque(), pushFront, 2, DequeBufSize-1, 2, 0, 3, 1, false)
			})
		})

		t.Run(`Pop`, func(t *testing.T) {
			t.Run(`Back`, func(t *testing.T) {
				test(NewDequeN(1), popBack, 1, 0, 3, 0, 3, 0, true)
				test(NewDequeN(DequeBufSize), popBack, 1, 0, 3, DequeBufSize-1, 3, DequeBufSize-1, false)
			})

			t.Run(`Front`, func(t *testing.T) {
				test(NewDequeN(1), popFront, 1, 1, 3, 1, 3, 0, true)

				var d = NewDequeN(2 * DequeBufSize)
				d.eraseAtBegin(nextN(d.Begin(), DequeBufSize-1))
				test(d, popFront, 2, 0, 3, 0, 4, DequeBufSize, false)
			})
		})
	})
}
