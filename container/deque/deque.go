package deque

import (
	"github.com/jhzhu89/golib/algorithm"
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
	dequeImpl
}

// Iterators

func (d *Deque) Begin() *DequeIter {
	return d.start.Clone().(*DequeIter)
}

func (d *Deque) End() *DequeIter {
	return d.finish.Clone().(*DequeIter)
}

func (d *Deque) RBegin() *ReverseIter {
	return iterator.NewReverseIterator(d.finish)
}

func (d *Deque) REnd() *ReverseIter {
	return iterator.NewReverseIterator(d.start)
}

// Capacity

func (d *Deque) Size() int {
	return d.start.Distance(d.finish)
}

func (d *Deque) Resize(newSize int) {
	var len = d.Size()
	if newSize > len {
		d.deafultAppend(newSize - len)
	} else if newSize < len {
		var it = d.start.Clone().(*DequeIter)
		it.NextN(newSize)
		d.eraseAtEnd(it)
	}
}

//func (d *Deque) ResizeAssign(n int, val Value) {
//
//}

func (d *Deque) ShrinkToFit() bool {
	// make a new deque and swap it.
	return false
}

func (d *Deque) Empty() bool {
	return d.start.Equal(d.finish)
}

// Element access
func (d *Deque) At(n int) Value {
	var it = d.Begin()
	it.NextN(n)
	return (*(*d.map_)[it.node])[it.cur]
}

func (d *Deque) Front() Value {
	return d.At(0)
}

func (d *Deque) Back() Value {
	var it = d.End()
	it.Prev()
	return (*(*d.map_)[it.node])[it.cur]
}

// Modifiers

func (d *Deque) FillAssign(size int, val Value) {
	if size > d.Size() {
		algorithm.Fill(d.Begin(), d.End(), val)
		d.FillInsert(d.End(), size-d.Size(), val)
	} else {
		var it = d.Begin()
		it.NextN(size)
		d.eraseAtEnd(it)
		algorithm.Fill(d.Begin(), d.End(), val)
	}
}

func (d *Deque) AssignRange(first, last InputIter) {
	var size = iterator.Distance(first, last)
	if size > d.Size() {
		var mid = first.Clone().(InputIter)
		iterator.Advance(mid, d.Size())
		algorithm.Copy(first, mid, d.Begin())
		d.InsertRange(d.End(), mid, last)
	} else {
		d.eraseAtEnd(algorithm.Copy(first, last, d.Begin()).(*DequeIter))
	}
}

func (d *Deque) PushBack(val Value) {
	if d.finish.cur != dequeBufSize-1 {
		(*(*d.map_)[d.finish.node])[d.finish.cur] = val
		d.finish.cur++
	} else {
		d.reserveMapAtBack(1)
		(*d.map_)[d.finish.node+1] = d.allocateNode()
		(*(*d.map_)[d.finish.node])[d.finish.cur] = val
		d.finish.setNode(d.finish.node + 1)
		d.finish.cur = 0
	}
}

func (d *Deque) PushFront(val Value) {
	if d.start.cur != 0 {
		(*(*d.map_)[d.start.node])[d.start.cur-1] = val
		d.start.cur--
	} else {
		d.reserveMapAtFront(1)
		(*d.map_)[d.start.node-1] = d.allocateNode()
		d.start.setNode(d.start.node - 1)
		d.start.cur = dequeBufSize - 1
		(*(*d.map_)[d.start.node])[d.start.cur] = val
	}
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

func (d *Deque) FillInsert(pos *DequeIter, size int, val Value) *DequeIter {
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

func (d *Deque) deafultAppend(n int) {
	if n > 0 {
		d.finish = d.reserveElementsAtBack(n)
	}
}

func (d *Deque) reserveElementsAtBack(n int) *DequeIter {
	var vacancies = dequeBufSize - d.finish.cur - 1
	if n > vacancies {
		d.newElementsAtBack(n - vacancies)
	}
	var it = d.finish.Clone().(*DequeIter)
	it.NextN(n)
	return it
}

func (d *Deque) newElementsAtBack(newElems int) {
	// TODO: add size limit?
	var newNodes = (newElems + dequeBufSize - 1) / dequeBufSize
	d.reserveMapAtBack(newNodes)
	for i := 0; i < newNodes; i++ {
		(*d.map_)[d.finish.node+1] = d.allocateNode()
	}
}

func (d *Deque) reserveMapAtBack(nodesToAdd int) {
	if nodesToAdd+1 > d.mapSize-d.finish.node {
		d.reallocateMap(nodesToAdd, false)
	}
}

func (d *Deque) reserveMapAtFront(nodesToAdd int) {
	if nodesToAdd > d.start.node {
		d.reallocateMap(nodesToAdd, true)
	}
}

func (d *Deque) reallocateMap(nodesToAdd int, addAtFront bool) {
	var oldNumNodes = d.finish.node - d.start.node + 1
	var newNumNodes = oldNumNodes + nodesToAdd

	var newStart int
	if d.mapSize > 2*newNumNodes {
		newStart = (d.mapSize - newNumNodes) / 2
		if addAtFront {
			newStart += nodesToAdd
		}
		d.map_.copy(d.start.node, d.finish.node+1, newStart)
	} else {
		var newMapSize = d.mapSize + max(d.mapSize, nodesToAdd) + 2
		var newMap = d.allocateMap(newMapSize)
		newStart = (newMapSize - newNumNodes) / 2
		if addAtFront {
			newStart += nodesToAdd
		}
		copy((*newMap)[newStart:], (*d.map_)[d.start.node:d.finish.node+1])
		d.deallocateMap(d.map_)
		d.map_ = newMap
		d.mapSize = newMapSize
	}
	d.start.setNode(newStart)
	d.finish.setNode(newStart + oldNumNodes - 1)
}

func (d *Deque) eraseAtEnd(pos *DequeIter) {
	pos = pos.Clone().(*DequeIter)
	d.destroyData(pos, d.End())
	d.destroyNodes(pos.node+1, d.finish.node+1)
}

func (d *Deque) destroyData(first, last *DequeIter) {
	var destroy = func(nodeSlice node) {
		for i := range nodeSlice {
			nodeSlice[i] = nil
		}
	}

	for node := first.node + 1; node < last.node; node++ {
		destroy(*(*d.map_)[node])
	}

	if first.node != last.node {
		destroy((*(*d.map_)[first.node])[first.cur:])
		destroy((*(*d.map_)[last.node])[:last.cur])
	} else {
		destroy((*(*d.map_)[first.node])[first.cur:last.cur])
	}
}

func (d *Deque) fillAssign(n int, val Value) {

}

type node []Value

type nodeMap []*node

func (nm *nodeMap) copy(start, finish, newStart int) {
	copy((*nm)[newStart:], (*nm)[start:finish])
}

type dequeImpl struct {
	map_          *nodeMap
	mapSize       int
	start, finish *DequeIter
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
	var n = make(node, dequeBufSize, dequeBufSize)
	return &n
}

func (i *dequeImpl) deallocateNode(n *node) {
	*n = nil
}

func (i *dequeImpl) allocateMap(size int) *nodeMap {
	var map_ = make(nodeMap, 0, size)
	return &map_
}

func (i *dequeImpl) deallocateMap(map_ *nodeMap) {
	*map_ = nil
}
