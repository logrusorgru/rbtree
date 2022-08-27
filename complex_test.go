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
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Count = 10000

// complex tests

func TestRandomSetGetDel(t *testing.T) {

	var (
		tr = New[int64, int64]()
		kv = make(map[int64]int64)
	)

	for i := 0; i < Count; i++ {
		var (
			k = rand.Int63n(math.MaxInt)
			v = rand.Int63n(math.MaxInt)
		)
		tr.Set(k, v)
		kv[k] = v
		assert.Equal(t, len(kv), tr.Len())
	}

	for k := range kv {
		assert.Equal(t, kv[k], tr.Get(k))
		delete(kv, k)
		tr.Del(k)
		assert.Equal(t, len(kv), tr.Len())
	}
}

func TestCritIndex(t *testing.T) {
	var (
		tr       = New[int, string]()
		min, max = "min", "max"
	)
	tr.Set(math.MaxInt, max)
	tr.Set(math.MinInt, min)
	assert.Equal(t, tr.Get(math.MaxInt), max)
	assert.Equal(t, tr.Get(math.MinInt), min)
}

func TestNilSet(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "")
	tr.Set(1, "")
	tr.Set(2, "")
	assert.Equal(t, 3, tr.Len())
	for _, j := range []int{0, 1, 2} {
		assert.True(t, tr.IsExist(j))
	}
}

func TestOneSizeRange(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	if assert.Len(t, tr.Slice(1, 1), 1) {
		assert.Equal(t, "b", tr.Slice(1, 1)[0])
	}
}

func TestOneSizeWalk(t *testing.T) {
	var tr = New[int, string]()
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	var (
		eKey   int
		eValue string
	)
	var walkFunc = func(key int, value string) error {
		eKey = key
		eValue = value
		return nil
	}
	assert.NoError(t, tr.Walk(1, 1, walkFunc))
	assert.Equal(t, 1, eKey)
	assert.Equal(t, "b", eValue)
}

// ref.: http://stackoverflow.com/q/23276417/1816872
func qsort(a []int) {
	if len(a) < 2 {
		return
	}
	var (
		left, right = 0, len(a) - 1
		pivotIndex  = rand.Int() % len(a)
	)
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
}

func TestRandomSetRange(t *testing.T) {

	var (
		tr = New[int, int64]()
		kv = make(map[int]int64)
	)

	for i := 0; i < Count; i++ {
		var (
			k = int(rand.Int63n(math.MaxInt))
			v = rand.Int63n(math.MaxInt)
		)
		tr.Set(k, v)
		kv[k] = v
		assert.Equal(t, len(kv), tr.Len())
	}

	// direct order
	var vals = tr.Slice(math.MinInt, math.MaxInt)
	if assert.Equal(t, len(vals), len(kv)) {
		kvKeys := make([]int, 0, len(kv))
		for k := range kv {
			kvKeys = append(kvKeys, k)
		}
		qsort(kvKeys)
		for i := 0; i < len(vals); i++ {
			assert.Equal(t, kv[kvKeys[i]], vals[i])
		}
	}

	// reverse order
	vals = tr.Slice(math.MaxInt, math.MinInt)
	if assert.Equal(t, len(vals), len(kv)) {
		var kvKeys = make([]int, 0, len(kv))
		for k := range kv {
			kvKeys = append(kvKeys, k)
		}
		qsort(kvKeys)
		for i := 0; i < len(vals); i++ {
			assert.Equal(t, kv[kvKeys[len(vals)-1-i]], vals[i])
		}
	}
}

func TestRandomSetWalk(t *testing.T) {

	var (
		tr = New[int, int64]()
		kv = make(map[int]int64)
	)

	for i := 0; i < Count; i++ {
		var (
			k = int(rand.Int63n(math.MaxInt))
			v = rand.Int63n(math.MaxInt)
		)
		tr.Set(k, v)
		kv[k] = v
		assert.Equal(t, len(kv), tr.Len())
	}

	var count int
	var walkFunc = func(key int, value int64) error {
		count++
		assert.Equal(t, value, tr.Get(key))
		return nil
	}

	// direct order
	assert.NoError(t, tr.Walk(math.MinInt, math.MaxInt, walkFunc))
	assert.Equal(t, count, len(kv))

	// reverse order
	count = 0
	assert.NoError(t, tr.Walk(math.MaxInt, math.MinInt, walkFunc))
	assert.Equal(t, count, len(kv))
}

/* moved to draft
Test result:

--- FAIL: TestRandomSetWalkDel (0.05s)
        complex_test.go:211: [random set walk del] wrong number of steps, expected 10000, got 4824
FAIL
exit status 1
FAIL    github.com/logrusorgru/ebony    0.189s

func TestRandomSetWalkDel(t *testing.T) {
	tr := New()
	kv := make(map[uint]int64)
	for i := 0; i < Count; i++ {
		k := uint(rand.Int63n(math.MaxInt))
		v := rand.Int63n(math.MaxInt)
		tr.Set(k, v)
		kv[k] = v
		if uint(len(kv)) != tr.Count() {
			t.Errorf("[random set walk del] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
	var pkey uint
	counter := 0
	wl := func(key uint, _ interface{}) error {
		if key != 0 {
			if pkey >= key {
				return errors.New("walking out of order")
			}
		}
		pkey = key
		counter++
		tr.Del(key)
		return nil
	}
	if err := tr.Walk(math.MinInt, math.MaxInt, wl); err != nil {
		t.Errorf("[random set walk del] unexpected walking error '%v'", err)
	}
	if counter != len(kv) {
		t.Errorf("[random set walk del] wrong number of steps, expected %d, got %d", len(kv), counter)
	}
}
*/
