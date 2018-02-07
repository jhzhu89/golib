package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvance(t *testing.T) {
	t.Run(`Panic`, func(t *testing.T) {
		test := func(it InputIterator, n int) {
			defer func() {
				r := recover()
				assert.Contains(t, fmt.Sprintf("%s", r), `can move backward`)
			}()

			Advance(it, n)
		}

		test(&inputIt{}, -1)
		test(&forwardIt{}, -1)
	})

	t.Run(`Ok`, func(t *testing.T) {
		test := func(it InputIterator, n int) {
			defer func() {
				r := recover()
				assert.Empty(t, r)
			}()

			Advance(it, n)
		}

		test(&inputIt{}, 1)
		test(&forwardIt{}, 1)
		test(&bidirectIt{}, 1)
		test(new(randIt), 1)

		test(&bidirectIt{}, -1)
		test(new(randIt), -1)
	})
}

func TestDistance(t *testing.T) {
	t.Run(`Ok`, func(t *testing.T) {
		test := func(a, b InputIter) {
			defer func() {
				r := recover()
				assert.Empty(t, r)
			}()

			Distance(a, b)
		}

		test(&inputIt{}, &inputIt{})
		test(&forwardIt{}, &forwardIt{})
		test(&bidirectIt{}, &bidirectIt{})
		test(new(randIt), new(randIt))
	})
}
