package algorithm

import (
	"github.com/jhzhu89/golib/container"
	"github.com/jhzhu89/golib/iterator"
)

type (
	// An meaningful name instead of just interface{}.
	Value = container.Value

	InputIter  = iterator.InputIterator
	MInputIter = iterator.MutableInputIterator

	OutputIter = iterator.OutputIterator

	ForwardIter  = iterator.ForwardIterator
	MForwardIter = iterator.MutableForwardIterator

	BidirectIter  = iterator.BidirectionalIterator
	MBidirectIter = iterator.MutableBidirectionalIterator

	RandIter = iterator.RandomAccessIterator
)
