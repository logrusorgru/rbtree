// rb-tree with uint index, not thread safe
package ebony

import "runtime"

const (
	RED   = true
	BLACK = false
)

type node struct {
	left   *node
	right  *node
	parent *node
	color  bool
	id     uint
	value  interface{}
}

var sentinel = &node{nil, nil, nil, BLACK, 0, nil}

func init() {
	sentinel.left, sentinel.right = sentinel, sentinel
}

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
	for x != t.root && x.parent.color == RED {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y.color == RED {
				x.parent.color = BLACK
				y.color = BLACK
				x.parent.parent.color = RED
				x = x.parent.parent
			} else {
				if x == x.parent.right {
					x = x.parent
					t.rotateLeft(x)
				}
				x.parent.color = BLACK
				x.parent.parent.color = RED
				t.rotateRight(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if y.color == RED {
				x.parent.color = BLACK
				y.color = BLACK
				x.parent.parent.color = RED
				x = x.parent.parent
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rotateRight(x)
				}
				x.parent.color = BLACK
				x.parent.parent.color = RED
				t.rotateLeft(x.parent.parent)
			}
		}
	}
	t.root.color = BLACK
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
	x := &node{}
	x.data = data
	x.parent = parent
	x.left = sentinel
	x.right = sentinel
	x.color = RED
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
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
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
	if y.color == BLACK {
		t.deleteFixup(x)
	}
	t.count--
}

func (t *Tree) findNode(id uint64) *node {
	current := t.root
	for current != sentinel {
		if id == current.id {
			return current
		} else {
			if id < current.id {
				current = current.left
			} else {
				current = current.right
			}
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

// Move, silent O(2logn)
func (t *Tree) Move(oid, nid uint) {
	if n := t.findNode(id); n != sentinel {
		t.insertNode(n.id, n.value)
		t.deleteNode(n)
	}
}

// flush O(1)
func (t *Tree) Flush() *Tree {
	t.root = sentinel
	runtime.GC()
	return t
}

// range O(logn+m), m = len(range), [b,e], b < e (!)
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
		} else {
			if from < current.id {
				current = current.left
			} else {
				current = current.right
			}
		}
	}
	return nil
}

// max O(logn)
func (t *Tree) Max() (uint, interface{}) {
	current := t.root
	for current.left != sentinel {
		current = current.left
	}
	return current.data.Id, current.data.Ptr
}

// min O(logn)
func (t *Tree) Min() (uint, interface{}) {
	current := t.root
	for current.right != sentinel {
		current = current.right
	}
	return current.data.Id, current.data.Ptr
}
