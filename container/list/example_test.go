// Copyright 2018-present Jiahao Zhu. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package list_test

import (
	"fmt"

	"github.com/jhzhu89/golib/container/list"
	"github.com/jhzhu89/golib/fn"
)

func Example() {
	// Create a new list and put some values in it.
	l := list.New()
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(0.9)
	l.PushBack(0.9)

	// Remove consecutive duplicates.
	l.Unique()

	// Iterate through list and print its contents.
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		fmt.Println(it.Deref())
	}

	// Sort elements in asc order.
	l.Sort(fn.CompareFunc(
		func(a, b interface{}) bool {
			switch a.(type) {
			case int:
				return false
			default:
				return true
			}
		}))

	// Iterate through list and print its contents.
	for it := l.Begin(); !it.EqualTo(l.End()); it.Next() {
		fmt.Println(it.Deref())
	}
	// Output: 1
	// 0.9
	// 0.9
	// 1
}
