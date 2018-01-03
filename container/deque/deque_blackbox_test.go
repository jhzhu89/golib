package deque

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeque(t *testing.T) {
	t.Run(`NewDeque`, func(t *testing.T) {
		d := NewDeque()
		assert.NotNil(t, d.map_)
		assert.Equal(t, initialMapSize, len(*d.map_))

		var numNodes = 1
		var nonEmptyNode = (d.mapSize - numNodes) / 2
		for i := 0; i < initialMapSize; i++ {
			if i == nonEmptyNode {
				assert.Equal(t, dequeBufSize, len(*(*d.map_)[i]))
			} else {
				assert.Nil(t, (*d.map_)[i])
			}
		}

		assert.NotNil(t, d.start)
		assert.Equal(t, 0, d.start.cur)
		assert.Equal(t, nonEmptyNode, d.start.node)
		assert.Equal(t, &d.map_, d.start.map_)

		assert.NotNil(t, d.finish)
		assert.Equal(t, 0, d.finish.cur)
		assert.Equal(t, nonEmptyNode, d.finish.node)
		assert.Equal(t, &d.map_, d.start.map_)

		assert.Equal(t, 0, d.Size())
		assert.True(t, d.Empty())
	})

	t.Run(`NewDequeN`, func(t *testing.T) {
		test := func(elems, numNonEmptyNodes, startCur, startNode, finishCur,
			finishNode, size int, empty bool) {
			d := NewDequeN(elems)
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
		test(dequeBufSize-1, 1, 0, 3, dequeBufSize-1, 3, dequeBufSize-1, false)
		test(dequeBufSize, 2, 0, 3, 0, 4, dequeBufSize, false)
		test(8*dequeBufSize, 9, 0, 1, 0, 9, dequeBufSize*8, false)
	})

	t.Run(`NewDequeFromRange`, func(t *testing.T) {
		var v = make(vector, 0, 8*dequeBufSize)
		for i := 0; i < 8*dequeBufSize; i++ {
			v = append(v, i)
		}

		test := func(first, last InputIter, numNonEmptyNodes, startCur, startNode, finishCur,
			finishNode, size int, empty bool) {
			d := NewDequeFromRange(first, last)
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

		var newIt = func(i int) *vectorIter { return &vectorIter{i, &v} }

		test(newIt(0), newIt(1), 1, 0, 3, 1, 3, 1, false)
		test(newIt(0), newIt(dequeBufSize-1), 1, 0, 3, dequeBufSize-1, 3, dequeBufSize-1, false)
		test(newIt(0), newIt(dequeBufSize), 2, 0, 3, 0, 4, dequeBufSize, false)
		test(newIt(0), newIt(8*dequeBufSize), 9, 0, 1, 0, 9, dequeBufSize*8, false)
	})
}

func TestDequeMethods(t *testing.T) {
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
		test(dequeBufSize-1, 1, 0, 3, dequeBufSize-1, 3, dequeBufSize-1, false)
		test(dequeBufSize, 1, 0, 3, 0, 4, dequeBufSize, false)
		// cause reallocate map
		test(8*dequeBufSize, 8, 0, 4, 0, 12, dequeBufSize*8, false)
	})

	t.Run(`ResizeAssign`, func(t *testing.T) {
		test := func(toSize, numNonEmptyNodes, startCur, startNode, finishCur,
			finishNode, size, mapSize int, empty bool) {
			d := NewDeque()
			d.ResizeAssign(toSize, 1)
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

			for it := d.Begin(); !it.Equal(d.End()); it.Next() {
				assert.Equal(t, it.Deref(), 1)
			}
		}

		test(1, 2, dequeBufSize-1, 2, 0, 3, 1, initialMapSize, false)
		test(dequeBufSize, 2, 0, 2, 0, 3, dequeBufSize, initialMapSize, false)
		test(8*dequeBufSize, 9, 0, 4, 0, 12, 8*dequeBufSize, 2*initialMapSize, false)
	})

	t.Run(`ShrinkToFit`, func(t *testing.T) {
		var d = NewDequeN(8 * dequeBufSize)
		assert.Equal(t, 11, len(*d.map_))
		d.Clear()
		assert.Equal(t, 0, d.Size())
		d.PushBack(1)
		assert.Equal(t, 1, d.Size())
		assert.Equal(t, 11, len(*d.map_))
		assert.True(t, d.ShrinkToFit())
		assert.Equal(t, 8, len(*d.map_))
	})

	t.Run(`ElementAccess`, func(t *testing.T) {
		var d = NewDeque()
		for i := 0; i < 10; i++ {
			d.PushBack(i)
		}

		for i := 0; i < 10; i++ {
			assert.Equal(t, i, d.At(i))
		}

		assert.Equal(t, 0, d.Front())
		assert.Equal(t, 9, d.Back())
	})

	t.Run(`FillAssign`, func(t *testing.T) {
		var d = NewDeque()
		for i := 0; i < 10; i++ {
			d.PushBack(i)
		}

		d.FillAssign(512, 2)
		for i := 0; i < 512; i++ {
			assert.Equal(t, 2, d.At(i))
		}
	})

	t.Run(`AssignRange`, func(t *testing.T) {
		var d = NewDeque()
		var n = 10
		for i := 0; i < n; i++ {
			d.PushBack(n - i)
		}

		var v = make(vector, 0, dequeBufSize)
		for i := 0; i < dequeBufSize; i++ {
			v = append(v, i)
		}

		d.AssignRange(&vectorIter{0, &v}, &vectorIter{dequeBufSize, &v})
		for i := 0; i < dequeBufSize; i++ {
			assert.Equal(t, i, d.At(i))
		}
	})

	t.Run(`PushBack`, func(t *testing.T) {
		test := func(elems, startCur, startNode, finishCur, finishNode, size int, empty bool) {
			d := NewDequeN(elems)
			d.PushBack(1)

			assert.Equal(t, startCur, d.start.cur)
			assert.Equal(t, startNode, d.start.node)

			assert.Equal(t, finishCur, d.finish.cur)
			assert.Equal(t, finishNode, d.finish.node)

			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())
		}

		test(0, 0, 3, 1, 3, 1, false)
		test(dequeBufSize-1, 0, 3, 0, 4, dequeBufSize, false)
	})
}
