package ebony

import (
	//	"fmt"
	//	"github.com/kr/pretty"
	"testing"
)

// basic suit

func TestNew(t *testing.T) {
	tr := New()
	if tr.count != 0 {
		t.Error("[new] count != 0")
	}
	if tr.root != sentinel {
		t.Error("[new] root != sentinel")
	}
}

func TestSet(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	if tr.root.id != 0 {
		t.Error("[set] wrong id")
	}
	if tr.root.value.(string) != "x" {
		t.Error("[set] wrong value")
	}
	if tr.count != 1 {
		t.Error("[set] wrong count")
	}
}

func TestDel(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	tr.Del(0)
	if tr.count != 0 {
		t.Error("[del] wrong count after del")
	}
	if tr.root != sentinel {
		t.Error("[del] wrong tree state after del")
	}
}

func TestGet(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	val := tr.Get(0)
	switch v := val.(type) {
	case string:
		if v != x {
			t.Error("[get] wrong returned value")
		}
	default:
		t.Error("[get] wrong type of returned value")
	}
	if tr.count != 1 {
		t.Error("[get] wrong count")
	}
}

func TestExist(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	val := tr.Exist(0)
	if !val {
		t.Error("[exist] existing is not exist")
	}
	val = tr.Exist(12)
	if val {
		t.Error("[exist] not existing is exist")
	}
}

func TestCount(t *testing.T) {
	x := "x"
	tr := New()
	if tr.count != 0 {
		t.Errorf("[get] wrong count, expected 0, got %d", tr.count)
	}
	tr.Set(0, x)
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
	tr.Set(1, x)
	if tr.count != 2 {
		t.Errorf("[get] wrong count, expected 2, got %d", tr.count)
	}
	tr.Del(1)
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
	tr.Del(0)
	if tr.count != 0 {
		t.Errorf("[get] wrong count, expected 0, got %d", tr.count)
	}
}

/*
// Move, silent, changes index of value O(2logn)
func (t *Tree) Move(oid, nid uint) {
	if n := t.findNode(oid); n != sentinel {
		t.insertNode(nid, n.value)
		t.deleteNode(n)
	}
}

// Flush the tree O(1)
func (t *Tree) Flush() *Tree {
	t.root = sentinel
	runtime.GC()
	return t
}

// Range returns all values in given range if any.
// O(logn+m), m = len(range), [b,e], b < e (!)
func (t *Tree) Range(from, to uint) []interface{} {
	values := []interface{}{}
	current := t.root
	for current != sentinel {
		if from == current.id {
			values = append(values, current.value)
			current = current.left
			for current != sentinel {
				if current.id <= to {
					values = append(values, current.value)
					current = current.left
				} else {
					break
				}
			}
			return values
		}
		if from < current.id {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

// Max returns maximum index and its value O(logn)
func (t *Tree) Max() (uint, interface{}) {
	current := t.root
	for current.left != sentinel {
		current = current.left
	}
	return current.id, current.value
}

// Min returns minimum indedx and its value O(logn)
func (t *Tree) Min() (uint, interface{}) {
	current := t.root
	for current.right != sentinel {
		current = current.right
	}
	return current.id, current.value
}

*/
