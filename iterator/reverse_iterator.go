package iterator

var _ BidirectIter = (*ReverseIterator)(nil)

type ReverseIterator struct {
	// it should be a bidirectional iterator at least.
	iter                     BidirectIter
	canOutput, canRandAccess bool
}

func NewReverseIterator(iter BidirectIter) *ReverseIterator {
	var r = &ReverseIterator{iter: iter.Clone().(BidirectIter)}
	if _, ok := r.iter.(OutputIter); ok {
		r.canOutput = true
	}
	if _, ok := r.iter.(RandIter); ok {
		r.canRandAccess = true
	}
	return r
}

func (it *ReverseIterator) CanMultiPass() {}

func (it *ReverseIterator) Clone() IterRef {
	return &ReverseIterator{it.iter, it.canOutput, it.canRandAccess}
}

func (it *ReverseIterator) CopyAssign(r IterCRef) {
	var r_ = r.(*ReverseIterator)
	it.iter = r_.iter.(BidirectIter).Clone().(BidirectIter)
	it.canOutput = r_.canOutput
	it.canRandAccess = r_.canRandAccess
}

func (it *ReverseIterator) Swap(r IterCRef) {
	var r_ = r.(*ReverseIterator)
	it.iter, r_.iter = r_.iter, it.iter
	it.canOutput, r_.canOutput = r_.canOutput, it.canOutput
	it.canRandAccess, r_.canRandAccess = r_.canRandAccess, it.canRandAccess
}

func (it *ReverseIterator) CanOutput() bool {
	return it.canOutput
}

func (it *ReverseIterator) CanRandAccess() bool {
	return it.canRandAccess
}

func (it *ReverseIterator) Deref() Value {
	var tmp = it.iter.Clone().(BidirectIter)
	tmp.Prev()
	return tmp.Deref()
}

func (it *ReverseIterator) DerefSet(val Value) {
	if !it.canOutput {
		panic("not an OutputIterator")
	}
	it.iter.(OutputIter).DerefSet(val)
}

func (it *ReverseIterator) EqualTo(r IterCRef) bool {
	return it.iter.EqualTo(r.(*ReverseIterator).iter)
}

func (it *ReverseIterator) Next() {
	it.iter.Prev()
}

func (it *ReverseIterator) Next2() *ReverseIterator {
	it.Next()
	return it
}

func (it *ReverseIterator) Prev() {
	it.iter.Next()
}

func (it *ReverseIterator) Prev2() *ReverseIterator {
	it.Prev()
	return it
}

func (it *ReverseIterator) LessThan(r IterCRef) bool {
	if !it.canRandAccess {
		panic("not a RandomAccessIterator")
	}
	return !it.iter.(RandIter).LessThan(r.(*ReverseIterator).iter) && !it.EqualTo(r)
}

func (it *ReverseIterator) NextN(n int) {
	if !it.canRandAccess {
		panic("not a RandomAccessIterator")
	}
	it.iter.(RandIter).PrevN(n)
}

func (it *ReverseIterator) NextN2(n int) *ReverseIterator {
	it.NextN(n)
	return it
}

func (it *ReverseIterator) PrevN(n int) {
	if !it.canRandAccess {
		panic("not a RandomAccessIterator")
	}
	it.iter.(RandomAccessIterator).NextN(n)
}

func (it *ReverseIterator) PrevN2(n int) *ReverseIterator {
	it.PrevN(n)
	return it
}

// r - it
func (it *ReverseIterator) Distance(r IterCRef) int {
	if !it.canRandAccess {
		panic("not a RandomAccessIterator")
	}
	return -it.iter.(RandIter).Distance(r.(*ReverseIterator).iter)
}
