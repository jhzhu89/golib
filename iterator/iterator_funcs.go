package iterator

import (
	"fmt"
)

func Advance(it InputIter, n int) {
	if n >= 0 {
		switch v := it.(type) {
		case RandIter:
			v.NextN(n)

		default:
			for i := n; i > 0; i-- {
				v.Next()
			}
		}
	} else {
		switch v := it.(type) {
		case RandIter:
			v.PrevN(n)

		case BidirectIter:
			for i := n; i > 0; i-- {
				v.Prev()
			}

		default:
			panic(fmt.Sprintf("only %v, %v can move backward", Bidirectional, RandomAccess))
		}
	}
}

func Distance(it1, it2 InputIter) int {
	switch it1.(type) {
	case RandIter:
		return it1.(RandIter).Distance(it2)

	default:
		var n int
		it1 = it1.Clone().(InputIter)
		it2 = it2.Clone().(InputIter)
		for !it1.Equal(it2) {
			it1.Next()
			n++
		}
		return n
	}
}
