rbtree
=====

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/logrusorgru/rbtree?utm_source=godoc)
[![Unlicense](https://img.shields.io/badge/license-unlicense-blue.svg)](http://unlicense.org/)
[![Build Status](https://github.com/logrusorgru/rbtree/workflows/build/badge.svg)](https://github.com/logrusorgru/rbtree/actions?workflow=build)
[![Coverage Status](https://coveralls.io/repos/github/logrusorgru/rbtree/badge.svg?branch=master)](https://coveralls.io/github/logrusorgru/rbtree?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/rbtree)](http://goreportcard.com/report/logrusorgru/rbtree)

Golang red-black tree.

[Nice visualization](http://www.cs.usfca.edu/~galles/visualization/RedBlack.html)

### Types

It uses a comparable type as a key and any type as a value.

### Methods

| Method name | Time   |
|:-----------:|:------:|
| Set     | O(log<sub>2</sub>*n*)  |
| SetNx   | O(log<sub>2</sub>*n*)  |
| Del     | O(log<sub>2</sub>*n*)  |
| Get     | O(log<sub>2</sub>*n*)  |
| GetEx   | O(log<sub>2</sub>*n*)  |
| IsExist | O(log<sub>2</sub>*n*)  |
| Len     | O(1)       |
| Move    | O(2log<sub>2</sub>*n*) |
| Max     | O(log<sub>2</sub>*n*)  |
| Min     | O(log<sub>2</sub>*n*)  |
| Empty   | O(1)       |
| Walk    | O(log<sub>2</sub>*n* + *m*)   |
| Slice   | O(log<sub>2</sub>*n* + *m*)   |

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
go get -u github.com/logrusorgru/rbtree
```

Test

```bash
go test -cover -race github.com/logrusorgru/rbtree
```

Run benchmark

_expensive_

```bash
go test -timeout=30m -test.bench . github.com/logrusorgru/rbtree
```
_limited_

```bash
go test -test.benchtime=0.1s -test.bench . github.com/logrusorgru/rbtree
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
	var tr = rbtree.New[int, string]()

	// append value to tree
	tr.Set(0, "zero")
	tr.Set(12, "some value")
	var t int = 98
	tr.Set(t, "don't forget about int indexing")
	tr.Set(199, "rbtree distributed under Unlicense, feel free to fork it")
	tr.Set(2, "trash")

	// delete value from tree
	tr.Del(2)

	// get it
	fmt.Println(tr.Get(2))
	fmt.Println(tr.Get(t))

	// check existence only
	fmt.Println(tr.IsExist(12))
	fmt.Println(tr.IsExist(590))

	// get count
	fmt.Println(tr.Len())

	// change index of value
	tr.Move(12, 13)
	fmt.Println(tr.Get(13))
	fmt.Println(tr.IsExist(12))

	// take min index and value
	fmt.Println(tr.Min())

	// max
	fmt.Println(tr.Max())

	// empty
	tr.Empty()
	fmt.Println(tr.Len())
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
