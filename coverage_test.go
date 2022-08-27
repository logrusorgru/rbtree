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

// gain coverage to 100%
package rbtree

import (
	"fmt"
	"testing"
)

func TestNilRange(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	if vls := tr.Slice(5, 10); vls != nil {
		t.Errorf("[nil range] range is not nil, expected 'nil', got '%v'", vls)
	}
}

func TestNilWalk(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	var walkFunc = func(key int, value string) error {
		return fmt.Errorf("[nil walk] synthetic error, you should not see it,"+
			" key %d, value '%v'", key, value)
	}
	var err error
	if err = tr.Walk(5, 10, walkFunc); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
	if err = tr.Walk(10, 5, walkFunc); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
}

func TestOneNilWalk(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	var walkFunc = func(key int, value string) error {
		return fmt.Errorf("[nil walk] synthetic error, you should not see it,"+
			" key %d, value '%v'", key, value)
	}
	var err error
	if err = tr.Walk(10, 10, walkFunc); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
}

func TestDelNil(t *testing.T) {
	const x = "x"
	var tr = New[int, string]()
	tr.Set(0, x)
	tr.Del(1)
	if tr.Len() != 1 {
		t.Errorf("[del nil] wrong count after del, expected 1, got %d",
			tr.Len())
	}
}
