// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package vector_test

import (
	"fmt"

	"github.com/jhzhu89/golib/container/vector"
)

func Example() {
	// Create a new vector and put some values in it.
	v := vector.New()
	v.PushBack(1)
	v.PushBack("2")
	v.PushBack(struct{}{})
	v.PushBack(nil)

	for !v.Empty() {
		fmt.Println(v.Back())
		v.PopBack()
	}
	// Output: <nil>
	// {}
	// 2
	// 1
}
