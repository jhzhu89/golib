//
// Imlementations need follow these conventions.
//
// *Ref:  the argument should be a reference to an iterator;
// *CRef: the argument should be a reference to an iterator,
//        while the function cannot modify its value thru this
//        reference.
// *Val:  the argument should be a value of iterator.
//
// Iterators should have pointer receiver, since some of its
// methods need to change its internal states.
//
package iterator

type IteratorType uint8

func (i IteratorType) String() string {
	return iteratorType[i]
}

const (
	Output IteratorType = iota
	Input
	MutableInput
	Forward
	MutableForward
	Bidirectional
	MutableBidirectional
	RandomAccess
	MutableRandomAccess
)

var iteratorType = [...]string{
	`OutputIterator`,
	`InputIterator`,
	`MutableInputIterator`,
	`ForwardIterator`,
	`MutableForwardIterator`,
	`BidirectionalIterator`,
	`MutableBidirectionalIterator`,
	`RandomAccessIterator`,
	`MutableRandomAccessIterator`,
}

type iterator interface {
	// it = r
	CopyAssign(r IterCRef)
	Swap(r IterCRef)
	Clone() IterRef
	// it++
	Next()
}

type OutputIterator interface {
	iterator
	// *it = o
	DerefSet(Value)
}

type InputIterator interface {
	iterator
	// *it
	Deref() Value
	Equal(r IterCRef) bool
}

type MutableInputIterator interface {
	InputIterator
	DerefSet(Value)
}

type ForwardIterator interface {
	InputIterator
	// a tag to differentiate from InputIterator.
	CanMultiPass()
}

type MutableForwardIterator interface {
	ForwardIterator
	DerefSet(Value)
}

type BidirectionalIterator interface {
	ForwardIterator
	// it--
	Prev()
}

type MutableBidirectionalIterator interface {
	BidirectionalIterator
	DerefSet(Value)
}

type RandomAccessIterator interface {
	BidirectionalIterator
	// r must be a cref.
	LessThan(r IterCRef) bool
	NextN(n int)
	PrevN(n int)
	Distance(r IterCRef) int
}

type MutableRandomAccessIterator interface {
	RandomAccessIterator
	DerefSet(Value)
}
