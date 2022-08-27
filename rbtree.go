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

// Package rbtree is the red-black tree.
package rbtree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type color bool

const (
	red   color = true
	black color = false
)

type node[Key constraints.Ordered, Value any] struct {
	left   *node[Key, Value]
	right  *node[Key, Value]
	parent *node[Key, Value]
	color  color
	key    Key
	value  Value
}

// Tree is the RB-tree
type Tree[Key constraints.Ordered, Value any] struct {
	sentinel *node[Key, Value]
	root     *node[Key, Value]
	len      int
}

func (t *Tree[Key, Value]) rotateLeft(x *node[Key, Value]) {

	var y = x.right

	x.right = y.left

	if y.left != t.sentinel {
		y.left.parent = x
	}

	if y != t.sentinel {
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

	if x != t.sentinel {
		x.parent = y
	}
}

func (t *Tree[Key, Value]) rotateRight(x *node[Key, Value]) {

	var y = x.left

	x.left = y.right

	if y.right != t.sentinel {
		y.right.parent = x
	}

	if y != t.sentinel {
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

	if x != t.sentinel {
		x.parent = y
	}
}

func (t *Tree[Key, Value]) insertFixup(x *node[Key, Value]) {

	for x != t.root && x.parent.color == red {

		if x.parent == x.parent.parent.left {

			var y = x.parent.parent.right

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

			var y = x.parent.parent.left

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
func (t *Tree[Key, Value]) insertNode(key Key, value Value) {

	var (
		current = t.root

		parent *node[Key, Value]
	)

	for current != t.sentinel {
		if key == current.key {
			current.value = value
			return
		}
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	var x = &node[Key, Value]{
		value:  value,
		parent: parent,
		left:   t.sentinel,
		right:  t.sentinel,
		color:  red,
		key:    key,
	}

	if parent != nil {
		if key < parent.key {
			parent.left = x
		} else {
			parent.right = x
		}
	} else {
		t.root = x
	}

	t.insertFixup(x)
	t.len++
}

func (t *Tree[Key, Value]) deleteFixup(x *node[Key, Value]) {

	for x != t.root && x.color == black {

		if x == x.parent.left {
			var w = x.parent.right

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

			var w = x.parent.left

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
func (t *Tree[Key, Value]) deleteNode(z *node[Key, Value]) {

	var x, y *node[Key, Value]
	if z == nil || z == t.sentinel {
		return
	}

	if z.left == t.sentinel || z.right == t.sentinel {
		y = z
	} else {
		y = z.right
		for y.left != t.sentinel {
			y = y.left
		}
	}

	if y.left != t.sentinel {
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
		z.key = y.key
		z.value = y.value
	}

	if y.color == black {
		t.deleteFixup(x)
	}

	t.len--
}

func (t *Tree[Key, Value]) findNode(key Key) *node[Key, Value] {

	var current = t.root

	for current != t.sentinel {
		if key == current.key {
			return current
		}
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	return t.sentinel
}

// New creates the new RB-Tree
func New[Key constraints.Ordered, Value any]() *Tree[Key, Value] {

	var (
		zeroKey   Key
		zeroValue Value

		sentinel = &node[Key, Value]{
			left:   nil,
			right:  nil,
			parent: nil,
			color:  black,
			key:    zeroKey,
			value:  zeroValue,
		}
	)

	sentinel.left, sentinel.right = sentinel, sentinel

	return &Tree[Key, Value]{
		sentinel: sentinel,
		root:     sentinel,
	}
}

// Set the value. Silent O(logn). This will overwrite the existing value.
// To simulate SetNx() method use:
//
//    if !tr.IsExist(key) {
//        tr.Set(key, value)
//    }
//
// Its complexity from O(logn) to O(2logn)
func (t *Tree[Key, Value]) Set(key Key, value Value) {
	t.insertNode(key, value)
}

// Del deletes the value. Silent O(logn)
func (t *Tree[Key, Value]) Del(key Key) {
	t.deleteNode(t.findNode(key))
}

// Get O(logn)
func (t *Tree[Key, Value]) Get(key Key) Value {
	return t.findNode(key).value
}

// IsExist O(logn)
func (t *Tree[Key, Value]) IsExist(key Key) bool {
	return t.findNode(key) != t.sentinel
}

// Len O(1)
func (t *Tree[Key, Value]) Len() int {
	return t.len
}

// Move moves the value from one index to another. Silent.
// It just changes index of value O(2logn)
func (t *Tree[Key, Value]) Move(oldKey, newKey Key) {
	if n := t.findNode(oldKey); n != t.sentinel {
		t.insertNode(newKey, n.value)
		t.deleteNode(n)
	}
}

// Empty makes the tree empty O(1). It returns the Tree itself.
func (t *Tree[Key, Value]) Empty() *Tree[Key, Value] {
	t.root = t.sentinel
	t.len = 0
	return t
}

// Max returns maximum index and its value O(logn)
func (t *Tree[Key, Value]) Max() (Key, Value) {
	var current = t.root

	for current.right != t.sentinel {
		current = current.right
	}

	return current.key, current.value
}

// Min returns minimum indedx and its value O(logn)
func (t *Tree[Key, Value]) Min() (Key, Value) {

	var current = t.root

	for current.left != t.sentinel {
		current = current.left
	}

	return current.key, current.value
}

// WalkFunc is a walker function type
type WalkFunc[Key constraints.Ordered, Value any] func(key Key,
	value Value) error

// ErrStop is the error for stop walking
var ErrStop = errors.New("stop a walking")

func (n *node[Key, Value]) walkLeft(sentinel *node[Key, Value], from, to Key,
	wl WalkFunc[Key, Value]) error {

	if n.key > from {
		if n.left != sentinel {
			if err := n.left.walkLeft(sentinel, from, to, wl); err != nil {
				return err
			}
		}
	}
	if n.key >= from && n.key <= to {
		if err := wl(n.key, n.value); err != nil {
			return err
		}
	}
	if n.key < to {
		if n.right != sentinel {
			if err := n.right.walkLeft(sentinel, from, to, wl); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *node[Key, Value]) walkRight(sentinel *node[Key, Value], from, to Key,
	wl WalkFunc[Key, Value]) error {

	if n.key < from {
		if n.right != sentinel {
			if err := n.right.walkRight(sentinel, from, to, wl); err != nil {
				return err
			}
		}
	}
	if n.key <= from && n.key >= to {
		if err := wl(n.key, n.value); err != nil {
			return err
		}
	}
	if n.key > to {
		if n.left != sentinel {
			if err := n.left.walkRight(sentinel, from, to, wl); err != nil {
				return err
			}
		}
	}
	return nil
}

// Walk on the Tree.
//
// Any error returned by the WalkFunc stops a walking.
// Also, there is special ErrStop, for example:
//
//    if err := tr.Walk(0, 500, walkFunc); err != nil && err != rbtree.ErrStop {
//        log.Println(err) // real error
//    }
//
// To pass through the entire tree, use the minimum possible and
// maximum possible values of the index. For example:
//
//
//    tr.Walk(math.MinUint, math.MaxUint, walkFunc)
//
func (t *Tree[Key, Value]) Walk(from, to Key,
	wl WalkFunc[Key, Value]) error {

	if from == to {
		node := t.findNode(from)
		if node != t.sentinel {
			return wl(node.key, node.value)
		}
		return nil
	} else if from < to {
		return t.root.walkLeft(t.sentinel, from, to, wl)
	}

	// else if to < from
	return t.root.walkRight(t.sentinel, from, to, wl)
}

// Slice returns all values in given range if any.
// O(logn+m), m = len(range), [b,e] order dependent of cpm(b, e)
// Recursive. The required stack size is proportional to the height of the tree.
// To simulate GraterThen and LaterThen methods use the minimum possible and
// maximum possible values of the index. For example:
//
//    gt78 := tr.Slice(78, math.MaxUint)
//
// To take k-v pairs use Walk method with custom WalkFunc like this:
//
//    type Pair struct {
//        Key   uint
//        Value interface{}
//    }
//
//    func SliceKV(tr *rbtree.Tree, from, to uint) []Pair {
//        pr := []Pair{}
//        walkFunc := func(key uint, value interface{}) error {
//            pr = append(pr, Pair{key, value})
//            return nil
//        }
//        tr.Walk(from, to, walkFunc)
//        if len(pr) == 0 {
//            return nil
//        }
//        return pr
//    }
//
func (t *Tree[Key, Value]) Slice(from, to Key) (vals []Value) {
	var walkFunc = func(_ Key, value Value) error {
		vals = append(vals, value)
		return nil
	}
	t.Walk(from, to, walkFunc)
	return
}
