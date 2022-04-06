package kvengine

import (
	"testing"
)

func TestKV(t *testing.T) {
	kv := NewKV()
	if kv == nil {
		t.Fatal("new kv failed")
	}

	k := "test_k"
	v := "test_v"

	kv.Set(k, []byte(v))

	gv, _ := kv.Get(k)
	if v != string(gv) {
		t.Fatalf("want %s but get %s", v, gv)
	}
	kv.Close()
}
