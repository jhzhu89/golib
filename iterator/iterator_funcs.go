package iterator

import (
	"fmt"
)

func Advance(it InputIter, n int) {
	switch v := it.(type) {
	case RandIter:
		v.NextN(n)

	case BidirectIter:
		if n > 0 {
			for ; n > 0; n-- {
				v.Next()
			}
		} else {
			for ; n < 0; n++ {
				v.Prev()
			}
		}

	default:
		if n < 0 {
			panic(fmt.Sprintf("%s only can move forward", Input))
		}
		for ; n > 0; n-- {
			v.Next()
		}
	}
}

func Distance(it1, it2 InputIter) int {
	switch it1.(type) {
	case RandIter:
		return it1.(RandIter).Distance(it2)

	default:
		var n int
		for it1 = it1.Clone().(InputIter); !it1.EqualTo(it2); it1.Next() {
			n++
		}
		return n
	}
}
