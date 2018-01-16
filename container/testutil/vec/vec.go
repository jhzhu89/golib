package vec

import (
	"github.com/jhzhu89/golib/iterator"
)

type Vec []interface{}

var _ iterator.InputIter = (*VecIter)(nil)

type VecIter struct {
	i    int
	data *Vec
}

func (it *VecIter) Swap(r iterator.IterCRef) {
	var r_ = r.(*VecIter)
	it.i, r_.i = r_.i, it.i
}
func (it *VecIter) CopyAssign(r iterator.IterCRef) { it.i = r.(*VecIter).i }
func (it *VecIter) Clone() iterator.IterRef        { return &VecIter{it.i, it.data} }
func (it *VecIter) Deref() iterator.Value          { return (*it.data)[it.i] }
func (it *VecIter) Next()                          { it.i++ }
func (it *VecIter) Equal(r iterator.IterCRef) bool { return it.i == r.(*VecIter).i }
func (it *VecIter) CanMultiPass()                  {}

func NewIt(i int, v *Vec) *VecIter { return &VecIter{i, v} }
