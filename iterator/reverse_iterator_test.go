package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	var it = new(randIt)
	var rit = NewReverseIterator(it)

	rit.Next()
	assert.Equal(t, -1, *(*int)(rit.iter.(*randIt)))

	rit.Prev()
	assert.Equal(t, 0, *(*int)(rit.iter.(*randIt)))

	rit.NextN(100)
	assert.Equal(t, -100, *(*int)(rit.iter.(*randIt)))

	rit.PrevN(200)
	assert.Equal(t, 100, *(*int)(rit.iter.(*randIt)))

	var rit2 = NewReverseIterator(it)
	rit2.PrevN(100)
	assert.True(t, rit.EqualTo(rit2))

	rit2.Next()
	assert.True(t, rit.LessThan(rit2))

	rit2.Prev()
	assert.False(t, rit.LessThan(rit2))

	rit2.NextN(100)
	assert.Equal(t, 100, rit.Distance(rit2))
	assert.Equal(t, -100, rit2.Distance(rit))
}
