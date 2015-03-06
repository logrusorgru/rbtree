// gain coverage to 100%
package ebony

import (
	"errors"
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
		return errors.New(
			fmt.Sprintf("[nil walk] synthetic error, you should not see it, key %d, value '%v'", key, value),
		)
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
		return errors.New(
			fmt.Sprintf("[nil walk] synthetic error, you should not see it, key %d, value '%v'", key, value),
		)
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
		t.Errorf("[del nil] wrong count after del, expected 1, got %d", tr.count)
	}
}
