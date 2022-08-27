//
// Copyright (c) 2022 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Unlicense.
// See LICENSE file for more details or see below.
//

//
// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>
//

package rbtree

import (
	"fmt"
	"math"
)

func ExampleNew() {
	var tr = New[int, string]()
	fmt.Printf("%T", tr)
	// Output:
	// *rbtree.Tree[int,string]
}

func ExampleTree_Set() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleTree_Del() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	tr.Del(0)
	fmt.Println(tr.IsExist(0))
	// Output:
	// false
}

func ExampleTree_Get() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleTree_IsExist() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	fmt.Println(tr.IsExist(0))
	// Output:
	// true
}

func ExampleTree_Len() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	fmt.Println(tr.Len())
	tr.Set(0, "hello")
	fmt.Println(tr.Len())
	tr.Set(1, "hi")
	fmt.Println(tr.Len())
	// Output:
	// 1
	// 1
	// 2
}

func ExampleTree_Move() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	tr.Move(0, 1)
	fmt.Println(tr.Get(1))
	// Output:
	// hello
}

func ExampleTree_Min() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Min())
	// Output:
	// 0 hello
}

func ExampleTree_Max() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Max())
	// Output:
	// 1 hi
}

func ExampleTree_Empty() {
	var tr = New[int, string]()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	tr.Empty()
	fmt.Println(tr.Len())
	// Output:
	// 0
}

func ExampleTree_Range() {
	var tr = New[int, string]()
	tr.Set(0, "zero")
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	fmt.Println(tr.Slice(1, 2))
	fmt.Println(tr.Slice(2, 1))
	// Output:
	// [one two]
	// [two one]
}

func ExampleTree_Walk() {
	var tr = New[int, string]()
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	var walkFunc = func(key int, value string) error {
		fmt.Println(key, "-", value)
		return nil
	}
	tr.Walk(0, math.MaxInt, walkFunc)
	// Output:
	// 1 - one
	// 2 - two
	// 3 - three
}
