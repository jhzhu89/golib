package deque

import (
	_ "github.com/jhzhu89/golib/algorithm"
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/iterator"
)

const (
	dequeBufSize   = 512
	nodeEnd        = dequeBufSize
	initialMapSize = 8
)

type (
	Value = container.Value

	IterRef  = iterator.IterRef
	IterCRef = iterator.IterCRef

	InputIter = iterator.InputIter
	RandIter  = iterator.RandIter

	ReverseIter = iterator.ReverseIterator
)

type Deque struct {
}

// Iterators

func (d *Deque) Begin() *DequeIter {
	return nil
}

func (d *Deque) End() *DequeIter {
	return nil
}

func (d *Deque) RBegin() *ReverseIter {
	return nil
}

func (d *Deque) REnd() *ReverseIter {
	return nil
}

// Capacity

func (d *Deque) Size() int {
	return 0
}

func (d *Deque) Resize(n int, val Value) {

}

func (d *Deque) Empty() bool {
	return false
}

func (d *Deque) ShrinkToFit() {
}

// Element access
func (d *Deque) At(n int) Value {
	return nil
}

func (d *Deque) Front() Value {
	return nil
}

func (d *Deque) Back() Value {
	return nil
}

// Modifiers

func (d *Deque) eraseAtEnd(pos *DequeIter) {
}

func (d *Deque) AssignRange(first, last InputIter) {
	//var size = iterator.Distance(first, last)
	//if size > d.Size() {
	//	var mid = first
	//	iterator.Advance(iterator.Ref(mid).(InputIterator), d.Size())
	//	algorithm.Copy(first, mid, d.Begin())
	//	// insert(end(), __mid, __last)
	//} else {
	//	d.eraseAtEnd(algorithm.Copy(first, last, d.Begin()).(DequeIter))
	//}
}

func (d *Deque) AssignFill(size int, val Value) {
}

func (d *Deque) PushBack(val Value) {
}

func (d *Deque) PushFront(val Value) {
}

func (d *Deque) PopBack() {
}

func (d *Deque) PopFront() {
}

func (d *Deque) Insert(pos *DequeIter, val Value) *DequeIter {
	return nil
}

func (d *Deque) InsertRange(pos *DequeIter, first, last InputIter) *DequeIter {
	return nil
}

func (d *Deque) InsertFill(pos *DequeIter, size int, val Value) *DequeIter {
	return nil
}

func (d *Deque) Erase(pos *DequeIter) *DequeIter {
	return nil
}

func (d *Deque) EraseRange(first, last *DequeIter) *DequeIter {
	return nil
}

func (d *Deque) Swap(x *Deque) {
}

func (d *Deque) Clear() {
}

type node []Value

type dequeImpl struct {
	map_          *[]*node
	mapSize       int
	start, finish DequeIter
}

func (i *dequeImpl) initializeMap(numElements int) {
	var numNodes = numElements/dequeBufSize + 1
	i.mapSize = max(initialMapSize, numNodes+2)
	i.map_ = i.allocateMap(i.mapSize)

	var nstart = (i.mapSize - numNodes) / 2
	var nfinish = nstart + numNodes

	i.createNodes(nstart, nfinish)
	i.start.setNode(nstart)
	i.finish.setNode(nfinish - 1)
	i.finish.cur = numElements % dequeBufSize
}

func (i *dequeImpl) createNodes(start, finish int) {
	for cur := start; cur < finish; cur++ {
		(*i.map_)[cur] = i.allocateNode()
	}
}

func (i *dequeImpl) destroyNodes(start, finish int) {
	for n := start; n < finish; n++ {
		i.deallocateNode((*i.map_)[n])
		(*i.map_)[n] = (*node)(nil)
	}
}

func (i *dequeImpl) allocateNode() *node {
	var n = make(node, 0, dequeBufSize)
	return &n
}

func (i *dequeImpl) deallocateNode(n *node) {
	*n = nil
}

func (i *dequeImpl) allocateMap(size int) *[]*node {
	var map_ = make([]*node, 0, size)
	return &map_
}

func (i *dequeImpl) deallocateMap(map_ *[]*node) {
	*map_ = nil
}

var _ RandIter = (*DequeIter)(nil)

// implement a random access iterator.
type DequeIter struct {
	cur  int
	node int
	map_ *[]*node
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
	it.cur, r_.cur = r_.cur, it.cur
	it.node, r_.node = r_.node, it.node
	it.map_, r_.map_ = r_.map_, it.map_
}

func (it *DequeIter) Deref() Value {
	return (*(*it.map_)[it.node])[it.cur]
}

func (it *DequeIter) DerefSet(val Value) {
	(*(*it.map_)[it.node])[it.cur] = val
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
	if offset >= 0 && offset < dequeBufSize {
		it.cur += n
	} else {
		var nodeOffset int
		if offset > 0 {
			nodeOffset = offset / dequeBufSize
		} else {
			nodeOffset = -(-offset-1)/dequeBufSize - 1
		}
		it.setNode(it.node + nodeOffset)
		it.cur = offset - nodeOffset*dequeBufSize
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
	return (r_.node-it.node)*dequeBufSize + r_.cur - it.cur
}

// util funcs
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
