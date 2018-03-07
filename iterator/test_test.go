package iterator

type dummyBase struct{}

func (it *dummyBase) CopyAssign(_ IterCRef) {}
func (it *dummyBase) Swap(_ IterCRef)       {}

var _ InputIterator = (*inputIt)(nil)

type inputIt struct {
	*dummyBase
}

func (it *inputIt) Deref() Value            { return nil }
func (it *inputIt) Next()                   {}
func (it *inputIt) EqualTo(r IterCRef) bool { return true }
func (it *inputIt) Clone() IterRef          { return &inputIt{it.dummyBase} }

var _ ForwardIterator = (*forwardIt)(nil)

type forwardIt struct {
	*dummyBase
}

func (it *forwardIt) CanMultiPass()           {}
func (it *forwardIt) DerefSet(_ Value)        {}
func (it *forwardIt) Deref() Value            { return nil }
func (it *forwardIt) Next()                   {}
func (it *forwardIt) EqualTo(_ IterCRef) bool { return true }
func (it *forwardIt) Clone() IterRef          { return &forwardIt{it.dummyBase} }

var _ BidirectionalIterator = (*bidirectIt)(nil)

type bidirectIt struct {
	*dummyBase
}

func (it *bidirectIt) CanMultiPass()           {}
func (it *bidirectIt) DerefSet(_ Value)        {}
func (it *bidirectIt) Deref() Value            { return nil }
func (it *bidirectIt) Next()                   {}
func (it *bidirectIt) Prev()                   {}
func (it *bidirectIt) EqualTo(_ IterCRef) bool { return true }
func (it *bidirectIt) Clone() IterRef          { return &bidirectIt{it.dummyBase} }

var _ RandomAccessIterator = (*randIt)(nil)

type randIt int

func (it *randIt) CopyAssign(_ IterCRef) {}
func (it *randIt) Swap(_ IterCRef)       {}
func (it *randIt) Clone() IterRef {
	var r = new(randIt)
	*r = *it
	return r
}
func (it *randIt) CanMultiPass()            {}
func (it *randIt) DerefSet(_ Value)         {}
func (it *randIt) Deref() Value             { return nil }
func (it *randIt) Next()                    { *it++ }
func (it *randIt) Prev()                    { *it-- }
func (it *randIt) EqualTo(r IterCRef) bool  { return *it == *(r.(*randIt)) }
func (it *randIt) LessThan(r IterCRef) bool { return *it < *(r.(*randIt)) }
func (it *randIt) Distance(r IterCRef) int  { return int(*(r.(*randIt)) - *it) }
func (it *randIt) NextN(n int)              { *it += randIt(n) }
func (it *randIt) PrevN(n int)              { *it -= randIt(n) }
