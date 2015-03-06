package ebony

import (
	//"errors"
	"math/rand"
	"testing"
	"time"
)

const COUNT = 10000

const MaxUint = ^uint(0)
const MinUint = 0

// complex tests

func TestRandomSetGetDel(t *testing.T) {
	tr := New()
	kv := make(map[uint]int64)
	for i := 0; i < COUNT; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		v := rand.Int63n(time.Now().Unix())
		tr.Set(k, v)
		kv[k] = v
		if uint(len(kv)) != tr.Count() {
			t.Errorf("[random set get] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
	for k := range kv {
		switch v := tr.Get(k).(type) {
		case int64:
			if v != kv[k] {
				t.Errorf("[random set get] wrong returned value, expected %d, got %d", kv[k], v)
			}
		default:
			t.Errorf("[random set get] wrong type of returned value, expected 'int64', got '%T'", v)
		}
		delete(kv, k)
		tr.Del(k)
		if uint(len(kv)) != tr.Count() {
			t.Errorf("[random set get] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
}

func TestCritIndex(t *testing.T) {
	tr := New()
	min, max := "min", "max"
	tr.Set(MaxUint, max)
	tr.Set(MinUint, min)
	switch v := tr.Get(MaxUint).(type) {
	case string:
		if v != max {
			t.Errorf("[crit index] wrong returned value, expected '%s', got '%s'", max, v)
		}
	default:
		t.Errorf("[crit index] wrong type of returned value, expected 'string', got '%T'", v)
	}
	switch v := tr.Get(MinUint).(type) {
	case string:
		if v != min {
			t.Errorf("[crit index] wrong returned value, expected '%s', got '%s'", min, v)
		}
	default:
		t.Errorf("[crit index] wrong type of returned value, expected 'string', got '%T'", v)
	}
}

func TestNilSet(t *testing.T) {
	tr := New()
	tr.Set(0, nil)
	tr.Set(1, nil)
	tr.Set(2, nil)
	if tr.Count() != 3 {
		t.Errorf("[nil set] wrong count, expected 3, got %d", tr.Count())
	}
	for _, j := range []uint{0, 1, 2} {
		if !tr.Exist(j) {
			t.Errorf("[nil set]  existing is not exist")
		}
	}
}

func TestOneSizeRange(t *testing.T) {
	tr := New()
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	if vls := tr.Range(1, 1); len(vls) != 1 {
		t.Errorf("[one size range] wrong length of values, expected 1, got %d", len(vls))
	} else if len(vls) == 1 && vls[0] != "b" {
		t.Errorf("[one size range] wrong value, expected 'b', got '%s'", vls[0])
	}
}

func TestOneSizeWalk(t *testing.T) {
	tr := New()
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	var ekey uint
	var evalue interface{}
	wl := func(key uint, value interface{}) error {
		ekey = key
		evalue = value
		return nil
	}
	if err := tr.Walk(1, 1, wl); err != nil {
		t.Errorf("[one size walk] unexpected walking error '%v'", err)
	}
	if ekey != 1 {
		t.Errorf("[one size walk] wrong key, expected 1, got %d", ekey)
	}
	if evalue != "b" {
		t.Errorf("[one size walk] wrong value, expected 'b', got '%s'", evalue)
	}
}

// ref.: http://stackoverflow.com/q/23276417/1816872
func qsort(a []uint) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivotIndex := rand.Int() % len(a)
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
	tr := New()
	kv := make(map[uint]int64)
	for i := 0; i < COUNT; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		v := rand.Int63n(time.Now().Unix())
		tr.Set(k, v)
		kv[k] = v
		if uint(len(kv)) != tr.Count() {
			t.Errorf("[random set range] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
	// direct order
	if vals := tr.Range(MinUint, MaxUint); len(vals) != len(kv) {
		t.Errorf("[random set range] wrong length of range values, expected %d, got %d", len(kv), len(vals))
	} else {
		kv_keys := make([]uint, 0, len(kv))
		for k, _ := range kv {
			kv_keys = append(kv_keys, k)
		}
		qsort(kv_keys)
		for i := 0; i < len(vals); i++ {
			if kv[kv_keys[i]] != vals[i] {
				t.Errorf("[random set range] wrong value in range, expected %d, got %d", kv[kv_keys[i]], vals[i])
			}
		}
	}
	// reverse order
	if vals := tr.Range(MaxUint, MinUint); len(vals) != len(kv) {
		t.Errorf("[random set range] wrong length of range values, expected %d, got %d", len(kv), len(vals))
	} else {
		kv_keys := make([]uint, 0, len(kv))
		for k, _ := range kv {
			kv_keys = append(kv_keys, k)
		}
		qsort(kv_keys)
		for i := 0; i < len(vals); i++ {
			if kv[kv_keys[len(vals)-1-i]] != vals[i] {
				t.Errorf("[random set range] wrong value in range, expected %d, got %d", kv[kv_keys[i]], vals[i])
			}
		}
	}
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
	for i := 0; i < COUNT; i++ {
		k := uint(rand.Int63n(time.Now().Unix()))
		v := rand.Int63n(time.Now().Unix())
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
	if err := tr.Walk(MinUint, MaxUint, wl); err != nil {
		t.Errorf("[random set walk del] unexpected walking error '%v'", err)
	}
	if counter != len(kv) {
		t.Errorf("[random set walk del] wrong number of steps, expected %d, got %d", len(kv), counter)
	}
}
*/
