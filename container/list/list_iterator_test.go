package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListIter(t *testing.T) {
	t.Run(`Clone`, func(t *testing.T) {
		var it = &ListIter{}
		var c = it.Clone().(*ListIter)
		assert.NotEqual(t, fmt.Sprintf("%p", it), fmt.Sprintf("%p", c))
	})

	t.Run(`Swap`, func(t *testing.T) {
		var node1, node2 = &listNode{}, &listNode{}
		var it1 = &ListIter{node1}
		var it2 = &ListIter{node2}
		it1.Swap(it2)

		assert.Equal(t, node2, it1.node)
		assert.Equal(t, node1, it2.node)
	})

	t.Run(`Equal`, func(t *testing.T) {
		var node1, node2 = &listNode{}, &listNode{}
		var it1 = &ListIter{node1}
		var it2 = &ListIter{node2}
		assert.False(t, it1.EqualTo(it2))

		var it3 = &ListIter{node1}
		assert.True(t, it1.EqualTo(it3))
	})
}
