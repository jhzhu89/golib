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
