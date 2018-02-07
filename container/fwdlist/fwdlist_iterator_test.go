package fwdlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForwardListIter(t *testing.T) {
	t.Run(`Clone`, func(t *testing.T) {
		var it = &ForwardListIter{}
		var c = it.Clone().(*ForwardListIter)
		assert.NotEqual(t, fmt.Sprintf("%p", it), fmt.Sprintf("%p", c))
	})

	t.Run(`Swap`, func(t *testing.T) {
		var node1, node2 = &fwdListNode{}, &fwdListNode{}
		var it1 = &ForwardListIter{node1}
		var it2 = &ForwardListIter{node2}
		it1.Swap(it2)

		assert.Equal(t, node2, it1.node)
		assert.Equal(t, node1, it2.node)
	})

	t.Run(`Equal`, func(t *testing.T) {
		var node1, node2 = &fwdListNode{}, &fwdListNode{}
		var it1 = &ForwardListIter{node1}
		var it2 = &ForwardListIter{node2}
		assert.False(t, it1.EqualTo(it2))

		var it3 = &ForwardListIter{node1}
		assert.True(t, it1.EqualTo(it3))
	})
}
