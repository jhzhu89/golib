// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package deque_test

import (
	"fmt"

	"github.com/jhzhu89/golib/container/deque"
)

func Example() {
	// Create a new deque and put some values in it.
	d := deque.New()
	d.PushBack(1)
	d.PushBack("2")
	d.PushBack(struct{}{})
	d.PushBack(nil)

	// Iterate through deque and print its contents.
	for it := d.Begin(); !it.EqualTo(d.End()); it.Next() {
		fmt.Println(it.Deref())
	}
	// Output: 1
	// 2
	// {}
	// <nil>
}
