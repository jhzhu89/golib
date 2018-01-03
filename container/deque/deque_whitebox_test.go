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
	test(d, d.Begin(), 1, dequeBufSize-1, 2, 0, 3, 1, 2, false)
	d = NewDeque()
	test(d, d.Begin(), dequeBufSize, 0, 2, 0, 3, dequeBufSize, 2, false)

	d = NewDequeN(1)
	test(d, d.End(), 1, 0, 3, 2, 3, 2, 1, false)
	d = NewDequeN(1)
	test(d, d.End(), dequeBufSize, 0, 3, 1, 4, dequeBufSize+1, 2, false)

	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), 1, 0, 3, 11, 3, 11, 1, false)
	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), dequeBufSize, 0, 3, 10, 4, dequeBufSize+10, 2, false)

	d = NewDequeN(10)
	test(d, nextN(d.Begin(), 5), 8*dequeBufSize, 0, 4, 10, 12, 8*dequeBufSize+10, 9, false)
}
