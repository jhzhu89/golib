package iterator

import (
	"github.com/jhzhu89/golib/container"
)

// Type aliases, just to make code more readable.
type (
	Value = container.Value

	Iter          = Iterator
	OutputIter    = OutputIterator
	InputIter     = InputIterator
	MInputIter    = MutableInputIterator
	ForwardIter   = ForwardIterator
	MForwardIter  = MutableForwardIterator
	BidirectIter  = BidirectionalIterator
	MBidirectIter = MutableBidirectionalIterator
	RandIter      = RandomAccessIterator
	MRandIter     = MutableRandomAccessIterator

	// A common tag which means that this si a container.
	Cont    = interface{}
	ContRef = interface{}

	// *IterVal:  the argument should be a value of an iterator.
	// *IterRef:  the argument should be a reference of an iterator.
	// *IterCRef: the argument should be a refference of an iterator, but the function
	//            should not modify it.
	// M*Iter*:   the iterator additionally satisfies OutputIterator.
	//

	// A common tag which means that this iterator could be any type of iterator.
	IterVal   = interface{}
	IterRef   = interface{}
	IterCRef  = interface{}
	MIterVal  = interface{}
	MIterRef  = interface{}
	MIterCRef = interface{}
)
