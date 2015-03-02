// rb-tree with uint index, not thread safe
package ebony

import "runtime"

const (
	red   = true
	black = false
)

type node struct {
	left   *node
	right  *node
	parent *node
	color  bool
	id     uint
	value  interface{}
}

var sentinel = &node{nil, nil, nil, black, 0, nil}

func init() {
	sentinel.left, sentinel.right = sentinel, sentinel
}

// Tree
type Tree struct {
	root  *node
	count uint
}

func (t *Tree) rotateLeft(x *node) {
	y := x.right
	x.right = y.left
	if y.left != sentinel {
		y.left.parent = x
	}
	if y != sentinel {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	} else {
		t.root = y
	}
	y.left = x
	if x != sentinel {
		x.parent = y
	}
}

func (t *Tree) rotateRight(x *node) {
	y := x.left
	x.left = y.right
	if y.right != sentinel {
		y.right.parent = x
	}
	if y != sentinel {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	} else {
		t.root = y
	}
	y.right = x
	if x != sentinel {
		x.parent = y
	}
}

func (t *Tree) insertFixup(x *node) {
	for x != t.root && x.parent.color == red {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				if x == x.parent.right {
					x = x.parent
					t.rotateLeft(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateRight(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rotateRight(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateLeft(x.parent.parent)
			}
		}
	}
	t.root.color = black
}

// silent rewrite if exist
func (t *Tree) insertNode(id uint, value interface{}) {
	current := t.root
	var parent *node
	for current != sentinel {
		if id == current.id {
			current.value = value
			return
		}
		parent = current
		if id < current.id {
			current = current.left
		} else {
			current = current.right
		}
	}
	x := &node{
		value:  value,
		parent: parent,
		left:   sentinel,
		right:  sentinel,
		color:  red,
		id:     id,
	}
	if parent != nil {
		if id < parent.id {
			parent.left = x
		} else {
			parent.right = x
		}
	} else {
		t.root = x
	}
	t.insertFixup(x)
	t.count++
}

func (t *Tree) deleteFixup(x *node) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					t.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}

// silent
func (t *Tree) deleteNode(z *node) {
	var x, y *node
	if z == nil || z == sentinel {
		return
	}
	if z.left == sentinel || z.right == sentinel {
		y = z
	} else {
		y = z.right
		for y.left != sentinel {
			y = y.left
		}
	}
	if y.left != sentinel {
		x = y.left
	} else {
		x = y.right
	}
	x.parent = y.parent
	if y.parent != nil {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	} else {
		t.root = x
	}
	if y != z {
		z.id = y.id
		z.value = y.value
	}
	if y.color == black {
		t.deleteFixup(x)
	}
	t.count--
}

func (t *Tree) findNode(id uint) *node {
	current := t.root
	for current != sentinel {
		if id == current.id {
			return current
		}
		if id < current.id {
			current = current.left
		} else {
			current = current.right
		}
	}
	return sentinel
}

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
	t.count = 0
	runtime.GC()
	return t
}

// Moved to draft
//
// Range returns all values in given range if any.
// O(logn+m), m = len(range), [b,e] order dependent of cpm(b, e)
//func (t *Tree) Range(from, to uint) []interface{} {
//}

// Max returns maximum index and its value O(logn)
func (t *Tree) Max() (uint, interface{}) {
	current := t.root
	for current.right != sentinel {
		current = current.right
	}
	return current.id, current.value
}

// Min returns minimum indedx and its value O(logn)
func (t *Tree) Min() (uint, interface{}) {
	current := t.root
	for current.left != sentinel {
		current = current.left
	}
	return current.id, current.value
}
