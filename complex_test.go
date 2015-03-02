package ebony

import (
	"github.com/logrusorgru/cry"
	"testing"
)

const COUNT = 10000

// complex tests

func TestRandomSetGet(t *testing.T) {
	tr := New()
	kv := make(map[uint]int64)
	for i := 0; i < COUNT; i++ {
		k64, _ := cry.Uint64()
		k := uint(k64)
		v, _ := cry.Int64()
		tr.Set(k, v)
		kv[k] = v
	}
	if uint(len(kv)) != tr.Count() {
		t.Errorf("[random set get] wrong count, expected %d, got %d", len(kv), tr.Count())
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
	}
}
