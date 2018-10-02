//
// Copyright (c) 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
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

// gain coverage to 100%
package ebony

import (
	"fmt"
	"testing"
)

func TestNilRange(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	if vls := tr.Range(5, 10); vls != nil {
		t.Errorf("[nil range] range is not nil, expected 'nil', got '%v'", vls)
	}
}

func TestNilWalk(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	wl := func(key uint, value interface{}) error {
		return fmt.Errorf("[nil walk] synthetic error, you should not see it,"+
			" key %d, value '%v'", key, value)
	}
	if err := tr.Walk(5, 10, wl); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
	if err := tr.Walk(10, 5, wl); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
}

func TestOneNilWalk(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	wl := func(key uint, value interface{}) error {
		return fmt.Errorf("[nil walk] synthetic error, you should not see it,"+
			" key %d, value '%v'", key, value)
	}
	if err := tr.Walk(10, 10, wl); err != nil {
		t.Errorf("[nil walk] unexpected error '%v'", err)
	}
}

func TestDelNil(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	tr.Del(1)
	if tr.count != 1 {
		t.Errorf("[del nil] wrong count after del, expected 1, got %d",
			tr.count)
	}
}
