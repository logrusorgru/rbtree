Ebony
=====

[![GoDoc](https://godoc.org/github.com/logrusorgru/ebony?status.svg)](https://godoc.org/github.com/logrusorgru/ebony)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/ebony.svg)](https://travis-ci.org/logrusorgru/ebony)
covers-100%

Golang red-black tree with uint index, not thread safe

[Nice visualization](http://www.cs.usfca.edu/~galles/visualization/RedBlack.html)

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

### Memory usage

O(*n*&times;node),

node = 3&times;ptr_size +
       uint_size +
       bool_size +
       data_size

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

### Licensing

Copyright &copy; 2015 Konstantin Ivanov <ivanov.konstantin@logrus.org.ru>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.
