package deque_test

import (
	"testing"

	. "github.com/jhzhu89/golib/container/deque"
	"github.com/stretchr/testify/assert"
)

type vector []interface{}

var _ InputIter = (*vectorIter)(nil)

type vectorIter struct {
	i    int
	data *vector
}

func (it *vectorIter) Swap(r IterCRef) {
	var r_ = r.(*vectorIter)
	it.i, r_.i = r_.i, it.i
}
func (it *vectorIter) CopyAssign(r IterCRef) { it.i = r.(*vectorIter).i }
func (it *vectorIter) Clone() IterRef        { return &vectorIter{it.i, it.data} }
func (it *vectorIter) Deref() Value          { return (*it.data)[it.i] }
func (it *vectorIter) Next()                 { it.i++ }
func (it *vectorIter) Equal(r IterCRef) bool { return it.i == r.(*vectorIter).i }
func (it *vectorIter) CanMultiPass()         {}

var dequeBufSize = DequeBufSize

func TestNewDeque(t *testing.T) {
	t.Run(`NewDeque`, func(t *testing.T) {
		d := NewDeque()
		assert.Equal(t, 0, d.Size())
		assert.True(t, d.Empty())
	})

	t.Run(`NewDequeN`, func(t *testing.T) {
		test := func(elems, size int, empty bool) {
			d := NewDequeN(elems)
			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())
		}

		test(1, 1, false)
		test(dequeBufSize-1, dequeBufSize-1, false)
		test(dequeBufSize, dequeBufSize, false)
		test(8*dequeBufSize, 8*dequeBufSize, false)
	})

	t.Run(`NewDequeFromRange`, func(t *testing.T) {
		var v = make(vector, 0, 8*dequeBufSize)
		for i := 0; i < 8*dequeBufSize; i++ {
			v = append(v, i)
		}

		test := func(first, last InputIter, size int, empty bool) {
			d := NewDequeFromRange(first, last)
			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())
		}

		var newIt = func(i int) *vectorIter { return &vectorIter{i, &v} }

		test(newIt(0), newIt(1), 1, false)
		test(newIt(0), newIt(dequeBufSize-1), dequeBufSize-1, false)
		test(newIt(0), newIt(dequeBufSize), dequeBufSize, false)
		test(newIt(0), newIt(8*dequeBufSize), 8*dequeBufSize, false)
	})
}

func TestDequeMethodsBlackbox(t *testing.T) {
	t.Run(`ResizeAssign`, func(t *testing.T) {
		test := func(toSize, size int, empty bool) {
			d := NewDeque()
			d.ResizeAssign(toSize, 1)
			assert.Equal(t, size, d.Size())
			assert.Equal(t, empty, d.Empty())

			for it := d.Begin(); !it.Equal(d.End()); it.Next() {
				assert.Equal(t, it.Deref(), 1)
			}
		}

		test(1, 1, false)
		test(dequeBufSize, dequeBufSize, false)
		test(8*dequeBufSize, 8*dequeBufSize, false)
	})

	t.Run(`ShrinkToFit`, func(t *testing.T) {
		var d = NewDequeN(8 * dequeBufSize)
		d.Clear()
		assert.Equal(t, 0, d.Size())
		d.PushBack(1)
		assert.Equal(t, 1, d.Size())
		assert.False(t, d.ShrinkToFit())

		d.FillAssign(dequeBufSize, 1)
		for i := 0; i < dequeBufSize*2/3; i++ {
			d.PopFront()
		}
		assert.True(t, d.ShrinkToFit())
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

	t.Run(`Insert`, func(t *testing.T) {
		var d = NewDeque()
		var it = d.Insert(d.Begin(), 1)
		assert.True(t, it.Equal(d.Begin()))
		assert.Equal(t, 1, d.Size())

		it = d.Insert(d.End(), 2)
		var end = d.End()
		end.Prev()
		assert.True(t, it.Equal(end))
		assert.Equal(t, 2, d.Size())

		it = d.Insert(it, 3)
		assert.Equal(t, 3, d.Size())

		it = d.Begin()
		for _, n := range []int{1, 3, 2} {
			assert.Equal(t, n, it.Deref())
			it.Next()
		}
	})

	t.Run(`InsertRange`, func(t *testing.T) {
		var v = make(vector, 0, 8*dequeBufSize)
		for i := 0; i < 8*dequeBufSize; i++ {
			v = append(v, i)
		}

		var newIt = func(i int) *vectorIter { return &vectorIter{i, &v} }

		var d = NewDequeN(10)
		var begin = d.Begin()
		var pos = d.Begin()
		pos.NextN(5)

		var newPos = d.InsertRange(pos, newIt(0), newIt(8*dequeBufSize))
		assert.Equal(t, begin.Distance(pos), d.Begin().Distance(newPos))

		for i := 0; i < 8*dequeBufSize; i++ {
			assert.Equal(t, v[i], newPos.Deref())
			newPos.Next()
		}

		assert.Equal(t, 5, newPos.Distance(d.End()))
		assert.Equal(t, 8*dequeBufSize+10, d.Size())
	})

	t.Run(`FillInsert`, func(t *testing.T) {
		var d = NewDequeN(10)
		var begin = d.Begin()
		var pos = d.Begin()
		pos.NextN(5)

		var newPos = d.FillInsert(pos, 8*dequeBufSize, 1)
		assert.Equal(t, begin.Distance(pos), d.Begin().Distance(newPos))

		for i := 0; i < 8*dequeBufSize; i++ {
			assert.Equal(t, 1, newPos.Deref())
			newPos.Next()
		}

		assert.Equal(t, 5, newPos.Distance(d.End()))
		assert.Equal(t, 8*dequeBufSize+10, d.Size())
	})

	t.Run(`Erase`, func(t *testing.T) {
		var d = NewDeque()
		for i := 0; i < 8*dequeBufSize; i++ {
			d.PushBack(i)
		}

		var pos = d.Begin()
		pos.NextN(4 * dequeBufSize)
		var distance = d.Begin().Distance(pos)

		var newPos = d.Erase(pos)

		assert.Equal(t, distance, d.Begin().Distance(newPos))
		assert.Equal(t, 8*dequeBufSize-1, d.Size())
		for it := d.Begin(); !it.Equal(d.End()); it.Next() {
			assert.NotEqual(t, 4*dequeBufSize, it.Deref())
		}
	})

	t.Run(`EraseRange`, func(t *testing.T) {
		var d = NewDeque()
		for i := 0; i < 8*dequeBufSize; i++ {
			d.PushBack(i)
		}

		var first = d.Begin()
		first.NextN(2 * dequeBufSize)
		var last = d.Begin()
		last.NextN(3 * dequeBufSize)
		var distance = d.Begin().Distance(first)

		var newPos = d.EraseRange(first, last)

		assert.Equal(t, distance, d.Begin().Distance(newPos))
		assert.Equal(t, 7*dequeBufSize, d.Size())

		for i := 0; i < 7*dequeBufSize; i++ {
			if i < 2*dequeBufSize {
				assert.Equal(t, i, d.At(i))
			} else {
				assert.Equal(t, i+dequeBufSize, d.At(i))
			}
		}
	})

	t.Run(`Clear`, func(t *testing.T) {
		var d = NewDequeN(10000)
		d.Clear()
		assert.Equal(t, 0, d.Size())
		assert.Equal(t, d.Begin(), d.End())
	})
}
