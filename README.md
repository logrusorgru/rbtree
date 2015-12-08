Ebony
=====

[![GoDoc](https://godoc.org/github.com/logrusorgru/ebony?status.svg)](https://godoc.org/github.com/logrusorgru/ebony)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/ebony.svg)](https://travis-ci.org/logrusorgru/ebony)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/ebony/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/ebony?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/ebony)](http://goreportcard.com/report/logrusorgru/ebony)

Golang red-black tree with uint index, not thread safe

[Nice visualization](http://www.cs.usfca.edu/~galles/visualization/RedBlack.html)

### Types

The tree configured to store values of type `interface{}` index type `uint`.
If you want to use other types, switch to the branch called "typed".

### Methods

| Method name | Time |
|:-----------:|:----:|
| Set   | O(log*n*) |
| Del   | O(log*n*) |
| Get   | O(log*n*) |
| Exist | O(log*n*) |
| Count | O(1) |
| Move  | O(2log*n*) |
| Max   | O(log*n*) |
| Min   | O(log*n*) |
| Flush | O(1) |
| Walk  |  -   |
| Range |  -   |

### Memory usage

O(*n*&times;node),

```go
node := 3*sizeof(uintptr) +
          sizeof(uint) +
          sizeof(bool) +
          sizeof(interface{}) // data
```

### Install

Get or update

```bash
go get github.com/logrusorgru/ebony
```

Test

```bash
cd $GOPATH/src/github.com/logrusorgru/ebony
go test
```

Run benchmark

_expensive_

```bash
cd $GOPATH/src/github.com/logrusorgru/ebony
go test -test.bench .
```
_limited_

```bash
cd $GOPATH/src/github.com/logrusorgru/ebony
go test -test.benchtime=0.1s -test.bench .
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/logrusorgru/ebony"
)

func main() {
	// create new red-black tree
	tr := ebony.New()

	// append value to tree
	tr.Set(0, "zero")
	tr.Set(12, "some value")
	var t uint = 98
	tr.Set(t, "don't forget about uint indexing")
	tr.Set(199, "ebony distributed under WTFPL feel free to fork it")
	tr.Set(2, "trash")

	// delete value from tree
	tr.Del(2)

	// get it
	fmt.Println(tr.Get(2))
	fmt.Println(tr.Get(t))

	// check existence only
	fmt.Println(tr.Exist(12))
	fmt.Println(tr.Exist(590))

	// get count
	fmt.Println(tr.Count())

	// change index of value
	tr.Move(12, 13)
	fmt.Println(tr.Get(13))
	fmt.Println(tr.Exist(12))

	// take min index and value
	fmt.Println(tr.Min())

	// max
	fmt.Println(tr.Max())

	// flush
	tr.Flush()
	fmt.Println(tr.Count())
	fmt.Println(tr.Min())
	fmt.Println(tr.Max())
	fmt.Println(tr.Get(0))
}
```

### Usage Notes

A `nil` is the value. Use `Del()` to delete value. But if value doesn't exist
method `Get()` returns `nil`. You can to use `struct{}` as an emty value to
avoid confusions. `Walk()` doesn't support `Tree` manipulations, yet (`Set()`
and `Del()` ops.). See [godoc](https://godoc.org/github.com/logrusorgru/ebony)
for any details. If you want to lookup the tree much more than change it,
take a look at LLRB (if memory usage are critical)
([read](http://www.read.seas.harvard.edu/~kohler/notes/llrb.html) |
[source](https://github.com/petar/GoLLRB))

### Licensing

Copyright &copy; 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.
