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

func ExampleTree_Flush() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	tr.Flush()
	fmt.Println(tr.Count())
	// Output:
	// 0
}
