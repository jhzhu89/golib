// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package fwdlist_test

import (
	"fmt"

	"github.com/jhzhu89/golib/container/fwdlist"
	"github.com/jhzhu89/golib/fn"
)

func Example() {
	// Create a new forward list and put some values in it.
	fl := fwdlist.New()
	fl.PushFront(1)
	fl.PushFront(2)
	fl.PushFront(3)
	fl.PushFront(4)
	fl.PushFront(4)
	fl.PushFront(5)

	// Remove consecutive duplicates.
	fl.Unique()

	// Iterate through forward list and print its contents.
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		fmt.Println(it.Deref())
	}

	// Sort elements in asc order.
	fl.Sort(fn.CompareFunc(func(a, b interface{}) bool { return a.(int) < b.(int) }))

	// Iterate through forward list and print its contents.
	for it := fl.Begin(); !it.EqualTo(fl.End()); it.Next() {
		fmt.Println(it.Deref())
	}
	// Output: 5
	// 4
	// 3
	// 2
	// 1
	// 1
	// 2
	// 3
	// 4
	// 5
}
