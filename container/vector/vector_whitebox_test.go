package vector

import (
	"testing"

	"github.com/jhzhu89/golib/container/testutil/vec"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run(`New`, func(t *testing.T) {
		var v = New()
		assert.NotNil(t, v.data)
		assert.Equal(t, 0, len(*v.data))
		assert.Equal(t, 0, v.finish.cur)
		assert.Equal(t, 0, v.endOfStorage.cur)
	})

	t.Run(`NewN`, func(t *testing.T) {
		var n = 512
		var v = NewN(512)
		assert.NotNil(t, v.data)
		assert.Equal(t, n, len(*v.data))
		assert.Equal(t, n, v.finish.cur)
		assert.Equal(t, n, v.endOfStorage.cur)
	})

	t.Run(`NewNValue`, func(t *testing.T) {
		var n = 512
		var v = NewNValue(512, 1)
		assert.NotNil(t, v.data)
		assert.Equal(t, n, len(*v.data))
		assert.Equal(t, n, v.finish.cur)
		assert.Equal(t, n, v.endOfStorage.cur)
		for i := 0; i < n; i++ {
			assert.Equal(t, 1, v.At(i))
		}
	})

	t.Run(`NewFromRange`, func(t *testing.T) {
		var n = 512
		var tv = make(vec.Vec, 0, n)
		for i := 0; i < n; i++ {
			tv = append(tv, i)
		}

		var v = NewFromRange(vec.NewIt(0, &tv), vec.NewIt(n, &tv))
		assert.NotNil(t, v.data)
		assert.Equal(t, n, len(*v.data))
		assert.Equal(t, n, v.finish.cur)
		assert.Equal(t, n, v.endOfStorage.cur)
	})
}

func TestFillAssign(t *testing.T) {
	var v = New()
	var n = 40960
	v.fillAssign(n, 1)
	assert.Equal(t, n, len(*v.data))
	for i := 0; i < n; i++ {
		assert.Equal(t, 1, v.At(i))
	}
	assert.Equal(t, n, v.finish.cur)
	assert.Equal(t, n, v.endOfStorage.cur)

	v.fillAssign(n/2, 2)
	assert.Equal(t, n, len(*v.data))
	for i := 0; i < n; i++ {
		if i < n/2 {
			assert.Equal(t, 2, v.At(i))
		} else {
			assert.Nil(t, v.At(i))
		}
	}
	assert.Equal(t, n/2, v.finish.cur)
	assert.Equal(t, n, v.endOfStorage.cur)
}

func TestFillInsert(t *testing.T) {
	var v = New()
	var n1, v1 = 1024, 1
	v.fillInsert(v.Begin(), n1, v1)
	assert.Equal(t, n1, v.Size())
	assert.Equal(t, n1, v.Capacity())
	assert.Equal(t, n1, len(*v.data))
	assert.Equal(t, n1, v.finish.cur)
	assert.Equal(t, n1, v.endOfStorage.cur)
	for i := 0; i < n1; i++ {
		assert.Equal(t, v1, v.At(i))
	}

	var n2, v2 = 512, 2
	v.fillInsert(v.Begin(), n2, v2)
	assert.Equal(t, n1+n2, v.Size())
	assert.Equal(t, 2*n1, v.Capacity())
	assert.Equal(t, 2*n1, len(*v.data))
	assert.Equal(t, n1+n2, v.finish.cur)
	assert.Equal(t, 2*n1, v.endOfStorage.cur)
	for i := 0; i < n1+n2; i++ {
		if i < n2 {
			assert.Equal(t, v2, v.At(i))
		} else {
			assert.Equal(t, v1, v.At(i))
		}
	}
}

func TestFillInitialize(t *testing.T) {
	var n = 512
	var v = NewN(n)
	v.Clear()
	var val = 10
	v.fillInitialize(n/2, val)

	assert.Equal(t, n/2, v.Size())
	assert.Equal(t, n, v.Capacity())
	assert.Equal(t, n, len(*v.data))
	assert.Equal(t, n/2, v.finish.cur)
	assert.Equal(t, n, v.endOfStorage.cur)
	for i := 0; i < n; i++ {
		if i < n/2 {
			assert.Equal(t, val, v.At(i))
		} else {
			assert.Nil(t, v.At(i))
		}
	}
}

func TestRangeInsert(t *testing.T) {
	var n1 = 512
	var tv = make(vec.Vec, 0, n1)
	for i := 0; i < n1; i++ {
		tv = append(tv, i)
	}

	var n2 = 10
	var v = NewN(n2)
	var pos = nextN(v.Begin(), n2/2)
	v.rangeInsert(pos, vec.NewIt(0, &tv), vec.NewIt(n1, &tv))
	assert.Equal(t, n1+n2, len(*v.data))
	assert.Equal(t, n1+n2, v.finish.cur)
	assert.Equal(t, n1+n2, v.endOfStorage.cur)

	for i := 0; i < n1+n2; i++ {
		if i < n2/2 || i >= n1+n2/2 {
			assert.Nil(t, v.At(i))
		} else {
			assert.Equal(t, i-n2/2, v.At(i))
		}
	}
}

func TestInsertAux(t *testing.T) {
	var n = 512
	var v = NewN(n)
	v.PushBack(nil) // double its capacity
	var pos = nextN(clone(v.start), n/2)

	var val = 1
	v.insertAux(pos, val)
	assert.Equal(t, n+2, v.Size())
	assert.Equal(t, n*2, v.Capacity())
	assert.Equal(t, n*2, len(*v.data))
	assert.Equal(t, n+2, v.finish.cur)
	assert.Equal(t, n*2, v.endOfStorage.cur)
	for i := 0; i < v.Size(); i++ {
		if i == pos.cur {
			assert.Equal(t, val, v.At(i))
		} else {
			assert.Nil(t, v.At(i))
		}
	}
}

func TestDefaultAppend(t *testing.T) {
	var n = 512
	var v = NewN(n)

	v.defaultAppend(n / 2)
	assert.Equal(t, n/2+n, v.Size())
	assert.Equal(t, n*2, v.Capacity())
	assert.Equal(t, n*2, len(*v.data))
	assert.Equal(t, n/2+n, v.finish.cur)
	assert.Equal(t, n*2, v.endOfStorage.cur)
}

func TestEraseAtEnd(t *testing.T) {
	var n, val = 512, 1
	var v = NewNValue(n, val)
	var pos = nextN(v.Begin(), n/2)

	v.eraseAtEnd(pos)
	assert.Equal(t, n/2, v.Size())
	assert.Equal(t, n, v.Capacity())
	assert.Equal(t, n, len(*v.data))
	assert.Equal(t, n/2, v.finish.cur)
	assert.Equal(t, n, v.endOfStorage.cur)
	for i := 0; i < n; i++ {
		if i < n/2 {
			assert.Equal(t, val, v.At(i))
		} else {
			assert.Nil(t, v.At(i))
		}
	}

	v.eraseAtEnd(v.Begin())
	assert.Equal(t, 0, v.Size())
	assert.Equal(t, n, v.Capacity())
	assert.Equal(t, n, len(*v.data))
	assert.Equal(t, 0, v.finish.cur)
	assert.Equal(t, n, v.endOfStorage.cur)
}

func TestErase(t *testing.T) {
	var v = New()
	var n = 5
	for i := 0; i < n; i++ {
		v.PushBack(i)
	}

	var test = func(pos *VectorIter, vals []int) {
		v.Erase(pos)
		assert.Equal(t, len(vals), v.Size())
		for i, n := range vals {
			assert.Equal(t, n, v.At(i))
		}
	}

	test(nextN(v.Begin(), 2), []int{0, 1, 3, 4})
	test(v.Begin(), []int{1, 3, 4})
	test(v.End(), []int{1, 3})
}

func TestEraseRange(t *testing.T) {
	var v = New()
	var n = 10
	for i := 0; i < n; i++ {
		v.PushBack(i)
	}

	var test = func(first, last *VectorIter, vals []int) {
		v.EraseRange(first, last)
		assert.Equal(t, len(vals), v.Size())
		for i, n := range vals {
			assert.Equal(t, n, v.At(i))
		}
	}

	test(nextN(v.Begin(), 2), prevN(v.End(), 2), []int{0, 1, 8, 9})
	test(v.Begin(), nextN(v.Begin(), 1), []int{1, 8, 9})
	test(v.Begin(), v.End(), nil)
}

func TestAssignAux(t *testing.T) {
	var n = 512
	var tv = make(vec.Vec, 0, n)
	for i := 0; i < n; i++ {
		tv = append(tv, i)
	}

	var v = New()
	for i := 0; i < n/2; i++ {
		v.PushBack(n - i)
	}

	v.assignAux(vec.NewIt(0, &tv), vec.NewIt(n, &tv))

	for i := 0; i < n; i++ {
		assert.Equal(t, i, v.At(i))
	}

	v.assignAux(vec.NewIt(0, &tv), vec.NewIt(1, &tv))
	assert.Equal(t, 1, v.Size())
	for i := 0; i < n; i++ {
		if i < 1 {
			assert.Equal(t, i, v.At(i))
		} else {
			assert.Nil(t, v.At(i))
		}
	}
}
