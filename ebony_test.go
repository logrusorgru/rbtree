package ebony

import "testing"

// foundation suit

/*
moved to: not necessary

func TestInit(t *testing.T) {
	if sentinel.left != sentinel {
		t.Error("[init] left child of sentinel is not reference to self")
	}
	if sentinel.right != sentinel {
		t.Error("[init] right child of sentinel is not reference to self")
	}
	if sentinel.color != black {
		t.Error("[init] sentinel color is not black")
	}
}
*/

// basic suit

func TestNew(t *testing.T) {
	tr := New()
	if tr.count != 0 {
		t.Error("[new] count != 0")
	}
	if tr.root != sentinel {
		t.Error("[new] root != sentinel")
	}
}

func TestSet(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	if tr.root.id != 0 {
		t.Errorf("[set] wrong id, expected 0, got %d", tr.root.id)
	}
	switch v := tr.root.value.(type) {
	case string:
		if v != x {
			t.Errorf("[set] wrong returned value, expected '%s', got '%s'", x, v)
		}
	default:
		t.Errorf("[set] wrong type of returned value, expected 'string', got '%T'", v)
	}
	if tr.count != 1 {
		t.Errorf("[set] wrong count, expected 1, got %d", tr.count)
	}
}

func TestDel(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	tr.Del(0)
	if tr.count != 0 {
		t.Errorf("[del] wrong count after del, expected 0, got %d", tr.count)
	}
	if tr.root != sentinel {
		t.Error("[del] wrong tree state after del")
	}
}

func TestGet(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	val := tr.Get(0)
	switch v := val.(type) {
	case string:
		if v != x {
			t.Errorf("[get] wrong returned value, expected 'x', got '%s'", v)
		}
	default:
		t.Errorf("[get] wrong type of returned value, expected 'string', got '%T'", v)
	}
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
	if v := tr.Get(579); v != nil {
		t.Errorf("[get] wrong returned value, expected nil, got '%v'", v)
	}
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
}

func TestExist(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	val := tr.Exist(0)
	if !val {
		t.Error("[exist] existing is not exist")
	}
	val = tr.Exist(12)
	if val {
		t.Error("[exist] not existing is exist")
	}
}

func TestCount(t *testing.T) {
	x := "x"
	tr := New()
	if tr.Count() != 0 {
		t.Errorf("[count] wrong count, expected 0, got %d", tr.Count())
	}
	tr.Set(0, x)
	if tr.Count() != 1 {
		t.Errorf("[count] wrong count, expected 1, got %d", tr.Count())
	}
	tr.Set(1, x)
	if tr.Count() != 2 {
		t.Errorf("[count] wrong count, expected 2, got %d", tr.Count())
	}
	tr.Del(1)
	if tr.Count() != 1 {
		t.Errorf("[count] wrong count, expected 1, got %d", tr.Count())
	}
	tr.Del(0)
	if tr.Count() != 0 {
		t.Errorf("[count] wrong count, expected 0, got %d", tr.Count())
	}
}

func TestMove(t *testing.T) {
	x := "x"
	tr := New()
	tr.Set(0, x)
	tr.Move(0, 1)
	val := tr.Get(1)
	switch v := val.(type) {
	case string:
		if v != x {
			t.Errorf("[move] wrong returned value, expected '%s', got '%s'", x, v)
		}
	default:
		t.Errorf("[move] wrong type of returned value, expected 'string', got '%T'", v)
	}
	if tr.count != 1 {
		t.Errorf("[move] wrong count, expected 0, got %d", tr.count)
	}
}

func TestFlush(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Flush()
	if tr.count != 0 {
		t.Error("[flush] count != 0")
	}
	if tr.root != sentinel {
		t.Error("[flush] root != sentinel")
	}
}

func TestMax(t *testing.T) {
	max := "max"
	maxi := uint(6)
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(maxi, max)
	tr.Set(3, "m")
	tr.Set(4, "n")
	tr.Set(5, "o")
	i, val := tr.Max()
	if i != maxi {
		t.Errorf("[max] wrong index of min, expected %d, got %d", maxi, i)
	}
	switch v := val.(type) {
	case string:
		if v != max {
			t.Errorf("[max] wrong returned value, expected '%s', got '%s'", max, v)
		}
	default:
		t.Errorf("[max] wrong type of returned value, expected 'string', got '%T'", v)
	}
}

func TestMin(t *testing.T) {
	min := "min"
	mini := uint(0)
	tr := New()
	tr.Set(1, "x")
	tr.Set(2, "y")
	tr.Set(3, "z")
	tr.Set(mini, min)
	tr.Set(4, "m")
	tr.Set(5, "n")
	tr.Set(6, "o")
	i, val := tr.Min()
	if i != mini {
		t.Errorf("[min] wrong index of min, expected %d, got %d", mini, i)
	}
	switch v := val.(type) {
	case string:
		if v != min {
			t.Errorf("[min] wrong returned value, expected '%s', got '%s'", min, v)
		}
	default:
		t.Errorf("[min] wrong type of returned value, expected 'string', got '%T'", v)
	}
}

func TestRange(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	{
		vls := tr.Range(1, 3)
		if len(vls) != 3 {
			t.Errorf("[range] wrong range length, expected 3, got %d", len(vls))
		}
		r13 := []string{"y", "z", "m"}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i] != r13[i] {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i], vls[i])
			}
		}
	}
	{
		vls := tr.Range(3, 1)
		if len(vls) != 3 {
			t.Errorf("[range] wrong range length, expected 3, got %d", len(vls))
		}
		r13 := []string{"m", "z", "y"}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i] != r13[i] {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i], vls[i])
			}
		}
	}
	{
		vls := tr.Range(1, 9)
		if len(vls) != 4 {
			t.Errorf("[range] wrong range length, expected 4, got %d", len(vls))
		}
		r13 := []string{"y", "z", "m", "n"}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i] != r13[i] {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i], vls[i])
			}
		}
	}
	tr.Del(0)
	{
		vls := tr.Range(4, 0)
		if len(vls) != 4 {
			t.Errorf("[range] wrong range length, expected 4, got %d", len(vls))
		}
		r13 := []string{"n", "m", "z", "y"}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i] != r13[i] {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i], vls[i])
			}
		}
	}
}

func TestWalk(t *testing.T) {
	tr := New()
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	type pair struct {
		Key   uint
		Value interface{}
	}
	vls := []pair{}
	wl := func(key uint, value interface{}) error {
		vls = append(vls, pair{key, value})
		return nil
	}
	{
		if err := tr.Walk(1, 3, wl); err != nil {
			t.Errorf("[range] unexpected walking error '%v'", err)
		}
		if len(vls) != 3 {
			t.Errorf("[range] wrong range length, expected 3, got %d", len(vls))
		}
		r13 := []pair{
			{1, "y"},
			{2, "z"},
			{3, "m"},
		}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i].Value != r13[i].Value {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i].Value, vls[i].Value)
			}
			if vls[i].Key != r13[i].Key {
				t.Errorf("[range] wrong key, expected '%d', got '%d'", r13[i].Key, vls[i].Key)
			}
		}
	}
	{
		vls = nil
		if err := tr.Walk(3, 1, wl); err != nil {
			t.Errorf("[range] unexpected walking error '%v'", err)
		}
		if len(vls) != 3 {
			t.Errorf("[range] wrong range length, expected 3, got %d", len(vls))
		}
		r13 := []pair{
			{3, "m"},
			{2, "z"},
			{1, "y"},
		}
		for i := 0; i < len(vls) && i < len(r13); i++ {
			if vls[i].Value != r13[i].Value {
				t.Errorf("[range] wrong value, expected '%s', got '%s'", r13[i].Value, vls[i].Value)
			}
			if vls[i].Key != r13[i].Key {
				t.Errorf("[range] wrong key, expected '%d', got '%d'", r13[i].Key, vls[i].Key)
			}
		}
	}
	{
		vls := []interface{}{}
		wl = func(_ uint, value interface{}) error {
			if value == "z" {
				return Stop
			}
			vls = append(vls, value)
			return nil
		}
		if err := tr.Walk(1, 3, wl); err != nil && err != Stop {
			t.Errorf("[range] unexpected walking error '%v'", err)
		}
		if len(vls) != 1 {
			t.Errorf("[range] wrong walking result length, expected 1, got %d", len(vls))
		}
		if vls[0] != "y" {
			t.Errorf("[range] wrong walking result, expected [y], got %v", vls)
		}
	}
	{
		vls := []interface{}{}
		wl = func(_ uint, value interface{}) error {
			if value == "z" {
				return Stop
			}
			vls = append(vls, value)
			return nil
		}
		if err := tr.Walk(3, 1, wl); err != nil && err != Stop {
			t.Errorf("[range] unexpected walking error '%v'", err)
		}
		if len(vls) != 1 {
			t.Errorf("[range] wrong walking result length, expected 1, got %d", len(vls))
		}
		if vls[0] != "m" {
			t.Errorf("[range] wrong walking result, expected [m], got %v", vls)
		}
	}
}
