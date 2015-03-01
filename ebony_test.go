package ebony

import (
	"testing"
)

func TestNew(t *testing.T) {
	tr := New()
	if tr.count != 0 {
		t.Error("the Tree is new, but count != 0")
	}
	if tr.root != sentinel {
		t.Error("the Tree is new, but root != sentinel")
	}
}

/*
// create new RB-Tree
func New() *Tree {
	return &Tree{
		root: sentinel,
	}
}

// Set, silent O(logn)
func (t *Tree) Set(id uint, value interface{}) {
	t.insertNode(id, value)
}

// Del, silent O(logn)
func (t *Tree) Del(id uint) {
	t.deleteNode(t.findNode(id))
}

// Get O(logn)
func (t *Tree) Get(id uint) interface{} {
	return t.findNode(id).value
}

// Exist O(logn)
func (t *Tree) Exist(id uint) bool {
	return t.findNode(id) != sentinel
}

// Count O(1)
func (t *Tree) Count() uint {
	return t.count
}

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
