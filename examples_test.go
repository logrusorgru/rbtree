package ebony

import "fmt"

func ExampleNew() {
	tr := New()
	fmt.Printf("%T", tr)
	// Output:
	// *ebony.Tree
}

func ExampleSet() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleDel() {
	tr := New()
	tr.Set(0, "hello")
	tr.Del(0)
	fmt.Println(tr.Exist(0))
	// Output:
	// false
}

func ExampleGet() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Get(0))
	// Output:
	// hello
}

func ExampleExist() {
	tr := New()
	tr.Set(0, "hello")
	fmt.Println(tr.Exist(0))
	// Output:
	// true
}

func ExampleCount() {
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

func ExampleMove() {
	tr := New()
	tr.Set(0, "hello")
	tr.Move(0, 1)
	fmt.Println(tr.Get(1))
	// Output:
	// hello
}

func ExampleMin() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Min())
	// Output:
	// 0 hello
}

func ExampleMax() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Max())
	// Output:
	// 1 hi
}

func ExampleFlush() {
	tr := New()
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	tr.Flush()
	fmt.Println(tr.Count())
	// Output:
	// 0
}
