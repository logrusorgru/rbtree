rbtree
=====

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/logrusorgru/rbtree?utm_source=godoc)
[![Unlicense](https://img.shields.io/badge/license-unlicense-blue.svg)](http://unlicense.org/)
[![Build Status](https://github.com/logrusorgru/rbtree/workflows/test/badge.svg)](https://github.com/logrusorgru/rbtree/actions?workflow=run%20tests)
[![Coverage Status](https://coveralls.io/repos/github/logrusorgru/rbtree/badge.svg?branch=master)](https://coveralls.io/github/logrusorgru/rbtree?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/rbtree)](http://goreportcard.com/report/logrusorgru/rbtree)

Golang red-black tree.

[Nice visualization](http://www.cs.usfca.edu/~galles/visualization/RedBlack.html)

### Types

It uses a comparable type as a key and any type as a value.

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
| Empty | O(1) |
| Walk  |  -   |
| Range |  -   |

### Memory usage

O(*n*&times;node),

where
```go
node = 3*sizeof(uintptr) +
          sizeof(Key) +
          sizeof(bool) +
          sizeof(Value) // data
```

### Install

Get or update

```bash
go get github.com/logrusorgru/rbtree
```

Test

```bash
cd $GOPATH/src/github.com/logrusorgru/rbtree
go test
```

Run benchmark

_expensive_

```bash
cd $GOPATH/src/github.com/logrusorgru/rbtree
go test -test.bench .
```
_limited_

```bash
cd $GOPATH/src/github.com/logrusorgru/rbtree
go test -test.benchtime=0.1s -test.bench .
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/logrusorgru/rbtree"
)

func main() {
	// create new red-black tree
	tr := rbtree.New()

	// append value to tree
	tr.Set(0, "zero")
	tr.Set(12, "some value")
	var t uint = 98
	tr.Set(t, "don't forget about uint indexing")
	tr.Set(199, "rbtree distributed under WTFPL feel free to fork it")
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

	// empty
	tr.Empty()
	fmt.Println(tr.Count())
	fmt.Println(tr.Min())
	fmt.Println(tr.Max())
	fmt.Println(tr.Get(0))
}
```

### See also

If you want to lookup the tree much more than change it,
take a look at LLRB (if memory usage are critical)
([read](http://www.read.seas.harvard.edu/~kohler/notes/llrb.html) |
[source](https://github.com/petar/GoLLRB))

## Licensing

Copyright © 2016-2022 Konstantin Ivanov. This work is free. It comes without any
warranty, to the extent permitted by applicable law. You can redistribute it
and/or modify it under the terms of the the Unlicense. See the LICENSE file for
more details.
