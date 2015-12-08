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
