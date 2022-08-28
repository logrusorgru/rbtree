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

func (t *Tree[Key, Value]) insertNode(key Key, value Value, overwrite bool) (
	added bool) {

	var (
		current = t.root

		parent *node[Key, Value]
	)

	for current != t.sentinel {
		if key == current.key {
			if overwrite {
				current.value = value
				return
			}
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
	return true
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
	if z == t.sentinel {
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

	return current // root sentinel
}

func newSentinel[Key constraints.Ordered, Value any]() (
	sentinel *node[Key, Value]) {

	var (
		zeroKey   Key
		zeroValue Value
	)

	sentinel = &node[Key, Value]{
		left:   nil,
		right:  nil,
		parent: nil,
		color:  black,
		key:    zeroKey,
		value:  zeroValue,
	}

	sentinel.left, sentinel.right = sentinel, sentinel
	return
}

// New creates the new RB-Tree
func New[Key constraints.Ordered, Value any]() *Tree[Key, Value] {

	var sentinel = newSentinel[Key, Value]()

	return &Tree[Key, Value]{
		sentinel: sentinel,
		root:     sentinel,
	}
}

// Set the value. O(logn). This will overwrite the existing value.
func (t *Tree[Key, Value]) Set(key Key, value Value) (added bool) {
	return t.insertNode(key, value, true)
}

// SetNx doesn't overwrites an existing value.
func (t *Tree[Key, Value]) SetNx(key Key, value Value) (added bool) {
	return t.insertNode(key, value, false)
}

// Del deletes value by key. O(logn). It returns false,
// if key doesn't exits.
func (t *Tree[Key, Value]) Del(key Key) (deleted bool) {
	var node = t.findNode(key)
	deleted = (node != t.sentinel)
	t.deleteNode(node)
	return
}

// Get O(logn). It returns zero value, if key doesn't exist.
func (t *Tree[Key, Value]) Get(key Key) Value {
	return t.findNode(key).value
}

// GetEx O(logn). It returns false, if key doesn't exist.
func (t *Tree[Key, Value]) GetEx(key Key) (val Value, ok bool) {
	var node = t.findNode(key)
	return node.value, node != t.sentinel
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
// It just changes index of value O(2logn).
func (t *Tree[Key, Value]) Move(oldKey, newKey Key) (moved bool) {
	if n := t.findNode(oldKey); n != t.sentinel {
		t.insertNode(newKey, n.value, true)
		t.deleteNode(n)
		return true
	}
	return // false
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

// Min returns minimum indexed and its value O(logn)
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
	walkFunc WalkFunc[Key, Value]) (err error) {

	if n.key > from {
		if n.left != sentinel {
			if err = n.left.walkLeft(sentinel, from, to, walkFunc); err != nil {
				return
			}
		}
	}

	if n.key >= from && n.key <= to {
		if err = walkFunc(n.key, n.value); err != nil {
			return
		}
	}

	if n.key < to {
		if n.right != sentinel {
			err = n.right.walkLeft(sentinel, from, to, walkFunc)
			if err != nil {
				return
			}
		}
	}

	return // nil
}

func (n *node[Key, Value]) walkRight(sentinel *node[Key, Value], from, to Key,
	walkFunc WalkFunc[Key, Value]) (err error) {

	if n.key < from {
		if n.right != sentinel {
			err = n.right.walkRight(sentinel, from, to, walkFunc)
			if err != nil {
				return
			}
		}
	}

	if n.key <= from && n.key >= to {
		if err = walkFunc(n.key, n.value); err != nil {
			return
		}
	}

	if n.key > to {
		if n.left != sentinel {
			err = n.left.walkRight(sentinel, from, to, walkFunc)
			if err != nil {
				return
			}
		}
	}

	return // nil
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
// The Tree shouldn't be modified inside the WalkFunc.
func (t *Tree[Key, Value]) Walk(from, to Key,
	walkFunc WalkFunc[Key, Value]) (err error) {

	switch {
	case from == to:
		var node = t.findNode(from)
		if node != t.sentinel {
			return walkFunc(node.key, node.value)
		}
		return
	case from < to:
		return t.root.walkLeft(t.sentinel, from, to, walkFunc)
	default: // to < from
	}

	return t.root.walkRight(t.sentinel, from, to, walkFunc)
}

// Slice returns all values at given range if any.
func (t *Tree[Key, Value]) Slice(from, to Key) (vals []Value) {
	var walkFunc = func(_ Key, value Value) error {
		vals = append(vals, value)
		return nil
	}
	t.Walk(from, to, walkFunc)
	return
}

// SliceKeys returns all keys at given range if any.
func (t *Tree[Key, Value]) SliceKeys(from, to Key) (keys []Key) {
	var walkFunc = func(key Key, _ Value) (err error) {
		keys = append(keys, key)
		return
	}
	t.Walk(from, to, walkFunc)
	return
}
