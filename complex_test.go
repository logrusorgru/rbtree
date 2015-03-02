package ebony

import (
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
