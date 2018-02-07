package vector_test

import (
	"testing"

	"github.com/jhzhu89/golib/container/vector"
	"github.com/stretchr/testify/assert"
)

func TestExportedMethods(t *testing.T) {
	t.Run(`At`, func(t *testing.T) {
		var v = vector.New()
		t.Run(`Panic`, func(t *testing.T) {
			defer func() {
				r := recover()
				assert.NotNil(t, r)
			}()
			v.At(0)
		})

		v.PushBack(1)
		assert.Equal(t, 1, v.At(0))
	})

	t.Run(`FrontAndBack`, func(t *testing.T) {
		var v = vector.New()
		v.PushBack(1)
		assert.Equal(t, 1, v.Back())
		assert.Equal(t, v.Front(), v.Back())

		v.PushBack(2)
		assert.Equal(t, 2, v.Back())
		assert.Equal(t, 1, v.Front())
	})

	t.Run(`Empty`, func(t *testing.T) {
		var v = vector.New()
		assert.True(t, v.Empty())
		v.PushBack(1)
		assert.False(t, v.Empty())
	})

	t.Run(`Capacity`, func(t *testing.T) {
		var v = vector.New()
		assert.Equal(t, 0, v.Size())
		var size = 512
		v.Reserve(size)
		assert.Equal(t, 0, v.Size())
		assert.Equal(t, size, v.Capacity())
		v.FillAssign(size, 1)
		assert.Equal(t, size, v.Size())
		assert.Equal(t, size, v.Capacity())
		v.Clear()
		assert.Equal(t, 0, v.Size())
		assert.Equal(t, size, v.Capacity())
		assert.True(t, v.ShrinkToFit())
		assert.Equal(t, 0, v.Size())
		assert.Equal(t, 0, v.Capacity())
	})

	t.Run(`PushAndPopBack`, func(t *testing.T) {
		var v = vector.New()
		var size = 512
		for i := 0; i < size; i++ {
			v.PushBack(i)
		}

		for i := 0; i < size; i++ {
			assert.Equal(t, i, v.At(i))
		}
		assert.Equal(t, size, v.Size())
		assert.Equal(t, size, v.Capacity())

		for i := 0; i < size; i++ {
			v.PopBack()
			assert.Equal(t, size-i-1, v.Size())
			assert.Equal(t, size, v.Capacity())
		}
	})

	t.Run(`Resize`, func(t *testing.T) {
		var v = vector.New()
		var size = 512
		v.Resize(size)
		assert.Equal(t, size, v.Size())
		assert.Equal(t, size, v.Capacity())
		for it := v.Begin(); !it.EqualTo(v.End()); it.Next() {
			assert.Nil(t, it.Deref())
		}

		v.Resize(size / 2)
		assert.Equal(t, size/2, v.Size())
		assert.Equal(t, size, v.Capacity())
	})

	t.Run(`FillResize`, func(t *testing.T) {
		var v = vector.New()
		var size = 512
		v.FillResize(size, 1)
		assert.Equal(t, size, v.Size())
		assert.Equal(t, size, v.Capacity())
		for it := v.Begin(); !it.EqualTo(v.End()); it.Next() {
			assert.Equal(t, 1, it.Deref())
		}

		v.FillResize(size/2, 1)
		assert.Equal(t, size/2, v.Size())
		assert.Equal(t, size, v.Capacity())
		for it := v.Begin(); !it.EqualTo(v.End()); it.Next() {
			assert.Equal(t, 1, it.Deref())
		}
	})

	t.Run(`FillAssign`, func(t *testing.T) {
		var v = vector.New()
		var size = 512
		v.FillAssign(size, 1)
		assert.Equal(t, size, v.Size())
		for it := v.Begin(); !it.EqualTo(v.End()); it.Next() {
			assert.Equal(t, 1, it.Deref())
		}
	})

	t.Run(`RangeAssign`, func(t *testing.T) {
		var size = 512
		var v1 = vector.New()
		var v2 = vector.NewNValues(size, 1)

		v1.RangeAssign(v2.Begin(), v2.End())
		assert.Equal(t, size, v1.Size())
		for it := v1.Begin(); !it.EqualTo(v1.End()); it.Next() {
			assert.Equal(t, 1, it.Deref())
		}
	})
}
