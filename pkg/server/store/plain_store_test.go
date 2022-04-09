package store

import (
	"bytes"
	"testing"
)

func TestPlainStore(t *testing.T) {
	ps := NewPlainStore("./")

	obj := &Object{
		Volid:    10,
		NodeID:   1,
		DiskID:   2,
		ObjectID: 28,
	}
	err := ps.Create(obj)
	if err != nil {
		t.Fatalf("create failed %v", err)
	}

	buf := []byte("-- test -- data --")
	bbuf := bytes.NewBuffer(buf)

	oio := &ObjectIO{
		Obj:    obj,
		Offset: 512,
		Length: uint32(bbuf.Len()),
		Data:   buf,
	}

	n, err := ps.Write(oio)
	if err != nil {
		t.Fatalf("write failed %v", err)
	}
	t.Logf("write n:%d bbuf.len:%d", n, bbuf.Len())

	buf1 := make([]byte, bbuf.Len())
	oio.Data = buf1
	n, err = ps.Read(oio)
	if err != nil {
		t.Fatalf("read failed %v", err)
	}

	t.Logf("get n:%d, %s", n, string(buf1))
	if string(buf) != string(buf1) {
		t.Fatalf("compare failed, buf:%s buf1:%s", string(buf), string(buf1))
	}
}

func TestPathValidate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PathValidate(tt.args.path)
		})
	}
}
