package algorithm

func Fill(first, last MForwardIter, val Value) {
	for f := first.Clone().(MForwardIter); !f.Equal(last); f.Next() {
		f.DerefSet(val)
	}
}

func Copy(first, last InputIter, result OutputIter) OutputIter {
	first = first.Clone().(InputIter)
	result = result.Clone().(OutputIter)
	switch first.(type) {
	case RandIter:
		for n := first.(RandIter).Distance(last); n > 0; n-- {
			result.DerefSet(first.Deref())
			first.Next()
			result.Next()
		}

	default:
		for !first.Equal(last) {
			result.DerefSet(first.Deref())
			first.Next()
			result.Next()
		}
	}

	return result
}

func CopyBackward(first, last BidirectIter, result MBidirectIter) MBidirectIter {
	last = last.Clone().(BidirectIter)
	result = result.Clone().(MBidirectIter)

	switch last.(type) {
	case RandIter:
		for n := -last.(RandIter).Distance(first); n > 0; n-- {
			last.Prev()
			result.Prev()
			result.DerefSet(last.Deref())
		}

	default:
		for !last.Equal(first) {
			last.Prev()
			result.Prev()
			result.DerefSet(last.Deref())
		}
	}

	return result
}
