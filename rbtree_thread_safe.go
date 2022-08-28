package rbtree

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type TreeThreadSafe[Key constraints.Ordered, Value any] struct {
	mx   sync.RWMutex
	tree *Tree[Key, Value]
}

// NewThreadSafe creates the new thread-safe RB-Tree.
func NewThreadSafe[Key constraints.Ordered, Value any]() (
	tts *TreeThreadSafe[Key, Value]) {

	tts = &TreeThreadSafe[Key, Value]{
		tree: New[Key, Value](),
	}
	return
}

// ToThreadSafe wraps a Tree.
func ToThreadSafe[Key constraints.Ordered, Value any](tree *Tree[Key, Value]) (
	tts *TreeThreadSafe[Key, Value]) {

	tts = &TreeThreadSafe[Key, Value]{
		tree: tree,
	}
	return
}

// Tree returns underlying Tree.
func (t *TreeThreadSafe[Key, Value]) Tree() (tree *Tree[Key, Value]) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree
}

// Set the value. O(logn). This will overwrite the existing value.
func (t *TreeThreadSafe[Key, Value]) Set(key Key, value Value) (added bool) {
	t.mx.Lock()
	defer t.mx.Unlock()

	return t.tree.Set(key, value)
}

// SetNx doesn't overwrites an existing value.
func (t *TreeThreadSafe[Key, Value]) SetNx(key Key, value Value) (added bool) {
	t.mx.Lock()
	defer t.mx.Unlock()

	return t.tree.SetNx(key, value)
}

// Del deletes value by key. O(logn). It returns false,
// if key doesn't exits.
func (t *TreeThreadSafe[Key, Value]) Del(key Key) (deleted bool) {
	t.mx.Lock()
	defer t.mx.Unlock()

	return t.tree.Del(key)
}

// Get O(logn). It returns zero value, if key doesn't exist.
func (t *TreeThreadSafe[Key, Value]) Get(key Key) Value {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Get(key)
}

// GetEx O(logn). It returns false, if key doesn't exist.
func (t *TreeThreadSafe[Key, Value]) GetEx(key Key) (val Value, ok bool) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.GetEx(key)
}

// IsExist O(logn)
func (t *TreeThreadSafe[Key, Value]) IsExist(key Key) bool {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.IsExist(key)
}

// Len O(1)
func (t *TreeThreadSafe[Key, Value]) Len() int {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Len()
}

// Move moves the value from one index to another. Silent.
// It just changes index of value O(2logn).
func (t *TreeThreadSafe[Key, Value]) Move(oldKey, newKey Key) (moved bool) {
	t.mx.Lock()
	defer t.mx.Unlock()

	return t.tree.Move(oldKey, newKey)
}

// Empty makes the tree empty O(1). It returns the Tree itself.
func (t *TreeThreadSafe[Key, Value]) Empty() {
	t.mx.Lock()
	defer t.mx.Unlock()

	t.tree.Empty()
}

// Max returns maximum index and its value O(logn)
func (t *TreeThreadSafe[Key, Value]) Max() (Key, Value) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Max()
}

// Min returns minimum indexed and its value O(logn)
func (t *TreeThreadSafe[Key, Value]) Min() (Key, Value) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Min()
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
func (t *TreeThreadSafe[Key, Value]) Walk(from, to Key,
	walkFunc WalkFunc[Key, Value]) (err error) {

	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Walk(from, to, walkFunc)
}

// Slice returns all values at given range if any.
func (t *TreeThreadSafe[Key, Value]) Slice(from, to Key) (vals []Value) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.Slice(from, to)
}

// SliceKeys returns all keys at given range if any.
func (t *TreeThreadSafe[Key, Value]) SliceKeys(from, to Key) (keys []Key) {
	t.mx.RLock()
	defer t.mx.RUnlock()

	return t.tree.SliceKeys(from, to)
}
