package kvengine

import (
	"testing"
)

func TestRdb(t *testing.T) {
	rdb := NewRedisKV()
	key := "rdb_key_1"
	val := "rdb_val_1"

	err := rdb.Set(key, []byte(val))
	if err != nil {
		t.Fatalf("set failed %v", err)
	}

	gv, err := rdb.Get(key)
	if err != nil {
		t.Fatalf("Get failed %v", err)
	}

	if val != string(gv) {
		t.Fatalf("expect %v got %v", val, string(gv))
	}
	t.Logf("expect %v got %v", val, string(gv))

	err = rdb.Delete(key)
	if err != nil {
		t.Fatalf("Del %v failed %v", key, err)
	}

	gv, err = rdb.Get(key)
	if !IsENotFound(err) {
		t.Fatalf("Get failed %v", err)
	}
}
