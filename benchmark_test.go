//
// Copyright (c) 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package ebony

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSeqSet(b *testing.B) {
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.ReportAllocs()
}

func BenchmarkSeqGet(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Get(uint(i))
	}
	b.ReportAllocs()
}

func BenchmarkSeqDel(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Del(uint(i))
	}
	b.ReportAllocs()
}

func BenchmarkSeqExs(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Exist(uint(i))
	}
	b.ReportAllocs()
}

func BenchmarkSeqMov(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	ln := tr.Count()
	for i := uint(0); i < ln; i++ {
		tr.Move(uint(i), ln-uint(i))
	}
	b.ReportAllocs()
}

func BenchmarkSeqMin(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Min()
	}
	b.ReportAllocs()
}

func BenchmarkSeqMax(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Max()
	}
	b.ReportAllocs()
}

func BenchmarkSeqWlk(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	wl := func(_ uint, _ interface{}) error {
		return nil
	}
	b.StartTimer()
	tr.Walk(MinUint, MaxUint, wl)
	b.ReportAllocs()
}

func BenchmarkSeqRng(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(i), nil)
	}
	b.StartTimer()
	tr.Range(MinUint, MaxUint)
	b.ReportAllocs()
}

// random

func shuffle(ary []uint) {
	for i := range ary {
		j := rand.Intn(i + 1)
		ary[i], ary[j] = ary[j], ary[i]
	}
}

func BenchmarkRndSet(b *testing.B) {
	b.StopTimer()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		ks = append(ks, uint(rand.Int63n(time.Now().Unix())))
	}
	tr := New()
	b.StartTimer()
	for _, t := range ks {
		tr.Set(t, nil)
	}
	b.ReportAllocs()
}

func BenchmarkRndGet(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	for _, t := range ks {
		tr.Set(t, nil)
	}
	b.ReportAllocs()
}

func BenchmarkRndDel(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	for _, t := range ks {
		tr.Del(t)
	}
	b.ReportAllocs()
}

func BenchmarkRndExs(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	for _, t := range ks {
		tr.Exist(t)
	}
	b.ReportAllocs()
}

func BenchmarkRndMov(b *testing.B) {
	b.StopTimer()
	tr := New()
	for i := 0; i < b.N; i++ {
		tr.Set(uint(rand.Int63n(time.Now().Unix())), nil)
	}
	ln := tr.Count()
	b.StartTimer()
	for i := uint(0); i < ln; i++ {
		tr.Move(uint(i), ln-uint(i))
	}
	b.ReportAllocs()
}

func BenchmarkRndMin(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	for range ks {
		tr.Min()
	}
	b.ReportAllocs()
}

func BenchmarkRndMax(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	for range ks {
		tr.Max()
	}
	b.ReportAllocs()
}

func BenchmarkRndWlk(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	wl := func(_ uint, _ interface{}) error {
		return nil
	}
	b.StartTimer()
	tr.Walk(MinUint, MaxUint, wl)
	b.ReportAllocs()
}

func BenchmarkRndRng(b *testing.B) {
	b.StopTimer()
	tr := New()
	ks := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		ks = append(ks, k)
		tr.Set(k, nil)
	}
	shuffle(ks)
	b.StartTimer()
	tr.Range(MinUint, MaxUint)
	b.ReportAllocs()
}
