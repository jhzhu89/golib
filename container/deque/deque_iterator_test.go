package deque

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeIter(t *testing.T) {
	t.Run(`Clone`, func(t *testing.T) {
		var it = &DequeIter{}
		var c = it.Clone().(*DequeIter)
		assert.NotEqual(t, fmt.Sprintf("%p", it), fmt.Sprintf("%p", c))
	})

	t.Run(`Swap`, func(t *testing.T) {
		var it1 = &DequeIter{}
		var it2 = &DequeIter{10, 10, nil}
		it1.Swap(it2)

		assert.Equal(t, 10, it1.cur)
		assert.Equal(t, 10, it1.node)
		assert.Equal(t, 0, it2.cur)
		assert.Equal(t, 0, it2.node)
	})

	t.Run(`Equal`, func(t *testing.T) {
		var it1 = &DequeIter{}
		var it2 = &DequeIter{10, 10, nil}
		assert.False(t, it1.EqualTo(it2))

		var it3 = &DequeIter{}
		assert.True(t, it1.EqualTo(it3))
	})

	t.Run(`LessThan`, func(t *testing.T) {
		var it1 = &DequeIter{}
		var it2 = &DequeIter{10, 10, nil}
		assert.True(t, it1.LessThan(it2))

		var it3 = &DequeIter{}
		assert.False(t, it1.LessThan(it3))
	})

	t.Run(`Distance`, func(t *testing.T) {
		var it1 = &DequeIter{}
		var it2 = &DequeIter{10, 10, nil}
		assert.Equal(t, 10+10*DequeBufSize, it1.Distance(it2))
		assert.Equal(t, -10-10*DequeBufSize, it2.Distance(it1))

		var it3 = &DequeIter{}
		assert.Equal(t, 0, it1.Distance(it3))
	})

	t.Run(`NextN`, func(t *testing.T) {
		var it = &DequeIter{}
		it.NextN(10)
		assert.Equal(t, 10, it.cur)
		assert.Equal(t, 0, it.node)

		it.NextN(DequeBufSize)
		assert.Equal(t, 10, it.cur)
		assert.Equal(t, 1, it.node)

		it.NextN(-DequeBufSize)
		assert.Equal(t, 10, it.cur)
		assert.Equal(t, 0, it.node)
	})
}
