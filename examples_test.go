//
// Copyright (c) 2015 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package ebony

import "fmt"

func ExampleNew() {
	tr := New()
	fmt.Printf("%T", tr)
	// Output:
	// *ebony.Tree
}

func ExampleTree_Set() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleTree_Del() {
	tr := New()
	tr.Set(0, "hello")
	tr.Del(0)
	fmt.Println(tr.Exist(0))
	// Output:
	// false
}

func ExampleTree_Get() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleTree_Exist() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Exist(0))
	// Output:
	// true
}

func ExampleTree_Count() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Count())
	tr.Set(0, "hello")
	fmt.Println(tr.Count())
	tr.Set(1, "hi")
	fmt.Println(tr.Count())
	// Output:
	// 1
	// 1
	// 2
}

func ExampleTree_Move() {
	tr := New()
	tr.Set(0, "hello")
	tr.Move(0, 1)
	fmt.Println(tr.Get(1))
	// Output:
	// hello
}

func ExampleTree_Min() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Min())
	// Output:
	// 0 hello
}

func ExampleTree_Max() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Max())
	// Output:
	// 1 hi
}

func ExampleTree_Empty() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	tr.Empty()
	fmt.Println(tr.Count())
	// Output:
	// 0
}

func ExampleTree_Range() {
	tr := New()
	tr.Set(0, "zero")
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	fmt.Println(tr.Range(1, 2))
	fmt.Println(tr.Range(2, 1))
	// Output:
	// [one two]
	// [two one]
}

func ExampleTree_Walk() {
	tr := New()
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	wl := func(key uint, value interface{}) error {
		fmt.Println(key, "-", value)
		return nil
	}
	const maxUint = ^uint(0)
	tr.Walk(0, maxUint, wl)
	// Output:
	// 1 - one
	// 2 - two
	// 3 - three
}
