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

var (
	globalString    string
	globalBool      bool
	globalErr       error
	globalSlice     []string
	globalSliceKeys []int
)

func init() {
	rand.Seed(1050)
}

func Benchmark(b *testing.B) {
	for _, nt := range []struct {
		name string
		tr   TreeInterface[int, string]
	}{
		{"tree", New[int, string]()},
		{"thread-safe", NewThreadSafe[int, string]()},
	} {
		b.Run(nt.name, func(b *testing.B) {
			b.Run("sequential", func(b *testing.B) {
				b.Run("set", sequentialSet(nt.tr))
				b.Run("set-nx", sequentialSetNx(nt.tr))
				b.Run("get", sequentialGet(nt.tr))
				b.Run("get-ex", sequentialGetEx(nt.tr))
				b.Run("del", sequentialDel(nt.tr))
				b.Run("is-exist", sequentialIsExists(nt.tr))
				b.Run("move", sequentialMove(nt.tr))
				b.Run("min", sequentialMin(nt.tr))
				b.Run("max", sequentialMax(nt.tr))
				b.Run("walk", sequentialWlk(nt.tr))
				b.Run("slice", sequentialSlice(nt.tr))
				b.Run("slice-keys", sequentialSliceKeys(nt.tr))
			})
			b.Run("random", func(b *testing.B) {
				b.Run("set", randomSet(nt.tr))
				b.Run("set-nx", randomSetNx(nt.tr))
				b.Run("get", randomGet(nt.tr))
				b.Run("get-ex", randomGetEx(nt.tr))
				b.Run("del", randomDel(nt.tr))
				b.Run("is-exists", randomIsExists(nt.tr))
				b.Run("move", randomMove(nt.tr))
				b.Run("min", randomMin(nt.tr))
				b.Run("max", randomMax(nt.tr))
				b.Run("walk", randomWalk(nt.tr))
				b.Run("slice", randomSlice(nt.tr))
				b.Run("slice-keys", randomSliceKeys(nt.tr))
			})
		})
	}
}

func sequentialSet(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.ReportAllocs()
	}
}

func sequentialSetNx(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.SetNx(i, "")
		}
		b.ReportAllocs()
	}
}

func sequentialGet(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalString = tr.Get(i)
		}
		b.ReportAllocs()
	}
}

func sequentialGetEx(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalString, globalBool = tr.GetEx(i)
		}
		b.ReportAllocs()
	}
}

func sequentialDel(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalBool = tr.Del(i)
		}
		b.ReportAllocs()
	}
}

func sequentialIsExists(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalBool = tr.IsExist(i)
		}
		b.ReportAllocs()
	}
}

func sequentialMove(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		var ln = tr.Len()
		for i := 0; i < ln; i++ {
			globalBool = tr.Move(i, ln-i)
		}
		b.ReportAllocs()
	}
}

func sequentialMin(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, globalString = tr.Min()
		}
		b.ReportAllocs()
	}
}

func sequentialMax(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, globalString = tr.Max()
		}
		b.ReportAllocs()
	}
}

func sequentialWlk(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		var walkFunc = func(int, string) error {
			return nil
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalErr = tr.Walk(math.MinInt, math.MaxInt, walkFunc)
		}
		b.ReportAllocs()
	}
}

func sequentialSlice(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalSlice = tr.Slice(math.MinInt, math.MaxInt)
		}
		b.ReportAllocs()
	}
}

func sequentialSliceKeys(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(i, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalSliceKeys = tr.SliceKeys(math.MinInt, math.MaxInt)
		}
		b.ReportAllocs()
	}
}

// random

func shuffle(ary []int) {
	for i := range ary {
		j := rand.Intn(i + 1)
		ary[i], ary[j] = ary[j], ary[i]
	}
}

func randomSet(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			ks = append(ks, int(rand.Int63n(math.MaxInt)))
		}
		tr.Empty()
		b.StartTimer()
		for _, t := range ks {
			tr.Set(t, "")
		}
		b.ReportAllocs()
	}
}

func randomSetNx(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			ks = append(ks, int(rand.Int63n(math.MaxInt)))
		}
		tr.Empty()
		b.StartTimer()
		for _, t := range ks {
			tr.SetNx(t, "")
		}
		b.ReportAllocs()
	}
}

func randomGet(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			k := int(rand.Int63n(math.MaxInt))
			ks = append(ks, k)
			tr.Set(k, "")
		}
		shuffle(ks)
		b.StartTimer()
		for _, k := range ks {
			globalString = tr.Get(k)
		}
		b.ReportAllocs()
	}
}

func randomGetEx(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			k := int(rand.Int63n(math.MaxInt))
			ks = append(ks, k)
			tr.Set(k, "")
		}
		shuffle(ks)
		b.StartTimer()
		for _, k := range ks {
			globalString, globalBool = tr.GetEx(k)
		}
		b.ReportAllocs()
	}
}

func randomDel(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			k := int(rand.Int63n(math.MaxInt))
			ks = append(ks, k)
			tr.Set(k, "")
		}
		shuffle(ks)
		b.StartTimer()
		for _, t := range ks {
			globalBool = tr.Del(t)
		}
		b.ReportAllocs()
	}
}

func randomIsExists(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		ks := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			k := int(rand.Int63n(math.MaxInt))
			ks = append(ks, k)
			tr.Set(k, "")
		}
		shuffle(ks)
		b.StartTimer()
		for _, t := range ks {
			globalBool = tr.IsExist(t)
		}
		b.ReportAllocs()
	}
}

func randomMove(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			tr.Set(int(rand.Int63n(math.MaxInt)), "")
		}
		var ln = tr.Len()
		b.StartTimer()
		for i := 0; i < ln; i++ {
			globalBool = tr.Move(i, ln-i)
		}
		b.ReportAllocs()
	}
}

func randomMin(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			var k = int(rand.Int63n(math.MaxInt))
			tr.Set(k, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, globalString = tr.Min()
		}
		b.ReportAllocs()
	}
}

func randomMax(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			var k = int(rand.Int63n(math.MaxInt))
			tr.Set(k, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, globalString = tr.Max()
		}
		b.ReportAllocs()
	}
}

func randomWalk(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			var k = int(rand.Int63n(math.MaxInt))
			tr.Set(k, "")
		}
		var walkFunc = func(_ int, _ string) error {
			return nil
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalErr = tr.Walk(math.MinInt, math.MaxInt, walkFunc)
		}
		b.ReportAllocs()
	}
}

func randomSlice(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			var k = int(rand.Int63n(math.MaxInt))
			tr.Set(k, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalSlice = tr.Slice(math.MinInt, math.MaxInt)
		}
		b.ReportAllocs()
	}
}

func randomSliceKeys(tr TreeInterface[int, string]) func(b *testing.B) {
	return func(b *testing.B) {
		b.StopTimer()
		tr.Empty()
		for i := 0; i < b.N; i++ {
			var k = int(rand.Int63n(math.MaxInt))
			tr.Set(k, "")
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			globalSliceKeys = tr.SliceKeys(math.MinInt, math.MaxInt)
		}
		b.ReportAllocs()
	}
}
