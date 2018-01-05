package deque

var _ RandIter = (*DequeIter)(nil)

// implement a random access iterator.
type DequeIter struct {
	cur  int
	node int
	map_ **nodeMap
}

func (it *DequeIter) CanMultiPass() {}

func (it *DequeIter) Clone() IterRef {
	return &DequeIter{it.cur, it.node, it.map_}
}

func (it *DequeIter) CopyAssign(r IterCRef) {
	*it = *(r.(*DequeIter))
}

func (it *DequeIter) Swap(r IterCRef) {
	var r_ = r.(*DequeIter)
	*it, *r_ = *r_, *it
}

func (it *DequeIter) Deref() Value {
	return (*(*(*it.map_))[it.node])[it.cur]
}

func (it *DequeIter) DerefSet(val Value) {
	(*(*(*it.map_))[it.node])[it.cur] = val
}

func (it *DequeIter) setNode(newNode int) {
	it.node = newNode
}

func (it *DequeIter) Next() {
	it.cur++
	if it.cur == nodeEnd {
		it.setNode(it.node + 1)
		it.cur = 0
	}
}

func (it *DequeIter) Prev() {
	if it.cur == 0 {
		it.setNode(it.node - 1)
		it.cur = nodeEnd
	}
	it.cur--
}

func (it *DequeIter) NextN(n int) {
	var offset = n + it.cur
	if offset >= 0 && offset < DequeBufSize {
		it.cur += n
	} else {
		var nodeOffset int
		if offset > 0 {
			nodeOffset = offset / DequeBufSize
		} else {
			nodeOffset = -(-offset-1)/DequeBufSize - 1
		}
		it.setNode(it.node + nodeOffset)
		it.cur = offset - nodeOffset*DequeBufSize
	}
}

func (it *DequeIter) PrevN(n int) {
	it.NextN(-n)
}

func (it *DequeIter) Equal(r IterCRef) bool {
	var r_ = r.(*DequeIter)
	return it.map_ == r_.map_ && it.node == r_.node && it.cur == r_.cur
}

func (it *DequeIter) LessThan(r IterCRef) bool {
	var r_ = r.(*DequeIter)
	return it.map_ == r_.map_ &&
		((it.node == r_.node && it.cur < r_.cur) ||
			it.node < r_.node)
}

func (it *DequeIter) Distance(r IterCRef) int {
	var r_ = r.(*DequeIter)
	return (r_.node-it.node)*DequeBufSize + r_.cur - it.cur
}

// util funcs
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func nextN(it *DequeIter, n int) *DequeIter {
	it.NextN(n)
	return it
}

func prevN(it *DequeIter, n int) *DequeIter {
	it.PrevN(n)
	return it
}

func next(it *DequeIter) *DequeIter {
	it.Next()
	return it
}

func prev(it *DequeIter) *DequeIter {
	it.Prev()
	return it
}

func clone(it *DequeIter) *DequeIter {
	return &DequeIter{it.cur, it.node, it.map_}
}
