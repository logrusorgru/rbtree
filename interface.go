package rbtree

import "golang.org/x/exp/constraints"

type TreeInterface[Key constraints.Ordered, Value any] interface {
	Set(Key, Value) bool
	SetNx(Key, Value) bool
	Del(Key) bool
	Get(Key) Value
	GetEx(Key) (Value, bool)
	IsExist(Key) bool
	Len() int
	Empty()
	Move(Key, Key) bool
	Max() (Key, Value)
	Min() (Key, Value)
	Walk(Key, Key, WalkFunc[Key, Value]) error
	Slice(Key, Key) []Value
	SliceKeys(Key, Key) []Key
}
