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

package rbtree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewThreadSafe(t *testing.T) {
	var tr = NewThreadSafe[int, any]()
	assert.Zero(t, tr.tree.len, "len != 0")
	assert.Equal(t, tr.tree.root, tr.tree.sentinel, "root != sentinel")
}

func TestTreeThreadSafe(t *testing.T) {

	const (
		x = "x"
		y = "y"
		z = "z"
	)
	var tr = NewThreadSafe[int, string]()

	assert.Equal(t, "", tr.Get(579))
	var val, ok = tr.GetEx(579)
	assert.Zero(t, val)
	assert.False(t, ok)
	assert.False(t, tr.Del(579))

	tr.Set(0, x)
	assert.Equal(t, 1, tr.Len())
	assert.Equal(t, x, tr.Get(0))
	assert.False(t, tr.SetNx(0, x))
	assert.Equal(t, 1, tr.Len())
	assert.Equal(t, x, tr.Get(0))
	assert.True(t, tr.SetNx(1, y))
	assert.Equal(t, 2, tr.Len())
	assert.Equal(t, y, tr.Get(1))

	assert.True(t, tr.Del(1))
	assert.Equal(t, 1, tr.Len())
	assert.Equal(t, "", tr.Get(1))
	tr.Del(0)
	assert.Equal(t, 0, tr.Len())
	assert.Equal(t, "", tr.Get(0))

	assert.False(t, tr.IsExist(0))
	tr.Set(0, x)
	assert.True(t, tr.IsExist(0))

	assert.True(t, tr.Move(0, 1))
	assert.Equal(t, 1, tr.Len())
	assert.Equal(t, "", tr.Get(0))
	assert.Equal(t, x, tr.Get(1))

	assert.False(t, tr.Move(579, 975))
	assert.Equal(t, 1, tr.Len())
	assert.Equal(t, "", tr.Get(0))
	assert.Equal(t, x, tr.Get(1))

	tr.Empty()
	assert.Equal(t, 0, tr.Len())

	tr.Set(1, x)
	tr.Set(2, y)
	tr.Set(3, z)
	assert.Equal(t, 3, tr.Len())

	var maxKey, maxValue = tr.Max()
	assert.Equal(t, 3, maxKey)
	assert.Equal(t, z, maxValue)

	var minKey, minValue = tr.Min()
	assert.Equal(t, 1, minKey)
	assert.Equal(t, x, minValue)

	assert.EqualValues(t, []string{x, y, z}, tr.Slice(math.MinInt, math.MaxInt))
	assert.EqualValues(t, []string{x}, tr.Slice(1, 1))
	assert.EqualValues(t, []string{x, y}, tr.Slice(1, 2))
	assert.EqualValues(t, []string{x, y, z}, tr.Slice(1, 3))
	assert.EqualValues(t, []string{x, y, z}, tr.Slice(1, 4))
	assert.EqualValues(t, []string{z, y, x}, tr.Slice(5, 0))
	assert.Nil(t, tr.Slice(90, 210))

	assert.EqualValues(t, []int{1, 2, 3},
		tr.SliceKeys(math.MinInt, math.MaxInt))
	assert.EqualValues(t, []int{1}, tr.SliceKeys(1, 1))
	assert.EqualValues(t, []int{1, 2}, tr.SliceKeys(1, 2))
	assert.EqualValues(t, []int{1, 2, 3}, tr.SliceKeys(1, 3))
	assert.EqualValues(t, []int{1, 2, 3}, tr.SliceKeys(1, 4))
	assert.EqualValues(t, []int{3, 2, 1}, tr.SliceKeys(5, 0))
	assert.Nil(t, tr.SliceKeys(90, 210))

	type pair struct {
		key   int
		value string
	}

	var pairs []pair
	var err = tr.Walk(0, 100, func(key int, value string) (err error) {
		pairs = append(pairs, pair{key: key, value: value})
		return
	})
	assert.NoError(t, err)
	assert.EqualValues(t, []pair{
		{1, x},
		{2, y},
		{3, z},
	}, pairs)

	pairs = pairs[:0]
	err = tr.Walk(2, 3, func(key int, value string) (err error) {
		pairs = append(pairs, pair{key: key, value: value})
		return
	})
	assert.NoError(t, err)
	assert.EqualValues(t, []pair{
		{2, y},
		{3, z},
	}, pairs)

	err = tr.Walk(math.MinInt, math.MaxInt, func(int, string) error {
		return ErrStop
	})
	assert.Equal(t, ErrStop, err)

	tr.Del(1)
	tr.Del(2)
	tr.Del(3)
}
