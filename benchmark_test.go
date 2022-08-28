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
)

func Benchmark(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.Run("set", sequentialSet)
		b.Run("set-nx", sequentialSetNx)
		b.Run("get", sequentialGet)
		b.Run("get-ex", sequentialGetEx)
		b.Run("del", sequentialDel)
		b.Run("is-exist", sequentialIsExists)
		b.Run("move", sequentialMove)
		b.Run("min", sequentialMin)
		b.Run("max", sequentialMax)
		b.Run("walk", sequentialWlk)
		b.Run("slice", sequentialRng)
	})
	b.Run("random", func(b *testing.B) {
		b.Run("set", randomSet)
		b.Run("set-nx", randomSetNx)
		b.Run("get", randomGet)
		b.Run("get-ex", randomGetEx)
		b.Run("del", randomDel)
		b.Run("is-exists", randomIsExists)
		b.Run("move", randomMove)
		b.Run("min", randomMin)
		b.Run("max", randomMax)
		b.Run("walk", randomWalk)
		b.Run("slice", randomSlice)
	})
}

func sequentialSet(b *testing.B) {
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.ReportAllocs()
}

func sequentialSetNx(b *testing.B) {
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.SetNx(i, "")
	}
	b.ReportAllocs()
}

func sequentialGet(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Get(i)
	}
	b.ReportAllocs()
}

func sequentialGetEx(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.GetEx(i)
	}
	b.ReportAllocs()
}

func sequentialDel(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Del(i)
	}
	b.ReportAllocs()
}

func sequentialIsExists(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.IsExist(i)
	}
	b.ReportAllocs()
}

func sequentialMove(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	var ln = tr.Len()
	for i := 0; i < ln; i++ {
		tr.Move(i, ln-i)
	}
	b.ReportAllocs()
}

func sequentialMin(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Min()
	}
	b.ReportAllocs()
}

func sequentialMax(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Max()
	}
	b.ReportAllocs()
}

func sequentialWlk(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	var walkFunc = func(int, string) error {
		return nil
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Walk(math.MinInt, math.MaxInt, walkFunc)
	}
	b.ReportAllocs()
}

func sequentialRng(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	tr.Slice(math.MinInt, math.MaxInt)
	b.ReportAllocs()
}

// random

func shuffle(ary []int) {
	for i := range ary {
		j := rand.Intn(i + 1)
		ary[i], ary[j] = ary[j], ary[i]
	}
}

func randomSet(b *testing.B) {
	b.StopTimer()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		ks = append(ks, int(rand.Int63n(math.MaxInt)))
	}
	var tr = New[int, string]()
	b.StartTimer()
	for _, t := range ks {
		tr.Set(t, "")
	}
	b.ReportAllocs()
}

func randomSetNx(b *testing.B) {
	b.StopTimer()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		ks = append(ks, int(rand.Int63n(math.MaxInt)))
	}
	var tr = New[int, string]()
	b.StartTimer()
	for _, t := range ks {
		tr.SetNx(t, "")
	}
	b.ReportAllocs()
}

func randomGet(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := int(rand.Int63n(math.MaxInt))
		ks = append(ks, k)
		tr.Set(k, "")
	}
	shuffle(ks)
	b.StartTimer()
	for _, k := range ks {
		tr.Get(k)
	}
	b.ReportAllocs()
}

func randomGetEx(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := int(rand.Int63n(math.MaxInt))
		ks = append(ks, k)
		tr.Set(k, "")
	}
	shuffle(ks)
	b.StartTimer()
	for _, k := range ks {
		tr.GetEx(k)
	}
	b.ReportAllocs()
}

func randomDel(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := int(rand.Int63n(math.MaxInt))
		ks = append(ks, k)
		tr.Set(k, "")
	}
	shuffle(ks)
	b.StartTimer()
	for _, t := range ks {
		tr.Del(t)
	}
	b.ReportAllocs()
}

func randomIsExists(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	ks := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := int(rand.Int63n(math.MaxInt))
		ks = append(ks, k)
		tr.Set(k, "")
	}
	shuffle(ks)
	b.StartTimer()
	for _, t := range ks {
		tr.IsExist(t)
	}
	b.ReportAllocs()
}

func randomMove(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		tr.Set(int(rand.Int63n(math.MaxInt)), "")
	}
	var ln = tr.Len()
	b.StartTimer()
	for i := 0; i < ln; i++ {
		tr.Move(i, ln-i)
	}
	b.ReportAllocs()
}

func randomMin(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		var k = int(rand.Int63n(math.MaxInt))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Min()
	}
	b.ReportAllocs()
}

func randomMax(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		var k = int(rand.Int63n(math.MaxInt))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Max()
	}
	b.ReportAllocs()
}

func randomWalk(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		var k = int(rand.Int63n(math.MaxInt))
		tr.Set(k, "")
	}
	var walkFunc = func(_ int, _ string) error {
		return nil
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Walk(math.MinInt, math.MaxInt, walkFunc)
	}
	b.ReportAllocs()
}

func randomSlice(b *testing.B) {
	b.StopTimer()
	var tr = New[int, string]()
	for i := 0; i < b.N; i++ {
		var k = int(rand.Int63n(math.MaxInt))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Slice(math.MinInt, math.MaxInt)
	}
	b.ReportAllocs()
}
