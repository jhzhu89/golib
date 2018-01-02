package deque

import (
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
}
