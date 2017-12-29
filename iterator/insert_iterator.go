package iterator

type Inserter interface {
	Insert(it Iter, val Value) Iter
}

var _ OutputIter = (*InsertIterator)(nil)

// An output interator.
type InsertIterator struct {
	ins  Inserter
	iter Iter
}

func NewInsertIterator(c Inserter, iter Iter) *InsertIterator {
	return &InsertIterator{c, iter.Clone().(Iter)}
}

func (it *InsertIterator) Clone() IterRef {
	return &InsertIterator{it.ins, it.iter.Clone().(Iter)}
}

func (it *InsertIterator) CopyAssign(r IterCRef) {
	var r_ = r.(*InsertIterator)
	it.ins = r_.ins
	it.iter = r_.iter.Clone().(Iter)
}

func (it *InsertIterator) Swap(r IterCRef) {
	var r_ = r.(*InsertIterator)
	it.ins, r_.ins = r_.ins, it.ins
	it.iter, r_.iter = r_.iter, it.iter
}

func (it *InsertIterator) DerefSet(val Value) {
	it.iter = it.ins.Insert(it.iter, val)
	it.iter.Next()
}

func (it *InsertIterator) Next() {}

type BackPusher interface {
	PushBack(val Value)
}

var _ OutputIter = (*BackInsertIterator)(nil)

type BackInsertIterator struct {
	bp BackPusher
}

func NewBackInsertIterator(c BackPusher) *BackInsertIterator {
	return &BackInsertIterator{c}
}

func (it *BackInsertIterator) Clone() IterRef {
	return &BackInsertIterator{it.bp}
}

func (it *BackInsertIterator) CopyAssign(r IterCRef) {
	var r_ = r.(*BackInsertIterator)
	it.bp = r_.bp
}

func (it *BackInsertIterator) Swap(r IterCRef) {
	var r_ = r.(*BackInsertIterator)
	it.bp, r_.bp = r_.bp, it.bp
}

func (it *BackInsertIterator) DerefSet(val Value) {
	it.bp.PushBack(val)
}

func (it *BackInsertIterator) Next() {}

type FrontPusher interface {
	PushFront(val Value)
}

var _ OutputIter = (*FrontInsertIterator)(nil)

type FrontInsertIterator struct {
	fp FrontPusher
}

func NewFrontInsertIterator(c FrontPusher) *FrontInsertIterator {
	return &FrontInsertIterator{c}
}

func (it *FrontInsertIterator) Clone() IterRef {
	return &FrontInsertIterator{it.fp}
}

func (it *FrontInsertIterator) CopyAssign(r IterCRef) {
	var r_ = r.(*FrontInsertIterator)
	it.fp = r_.fp
}

func (it *FrontInsertIterator) Swap(r IterCRef) {
	var r_ = r.(*FrontInsertIterator)
	it.fp, r_.fp = r_.fp, it.fp
}

func (it *FrontInsertIterator) DerefSet(val Value) {
	it.fp.PushFront(val)
}

func (it *FrontInsertIterator) Next() {}
