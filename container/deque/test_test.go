package deque

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
