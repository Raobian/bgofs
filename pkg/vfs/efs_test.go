package vfs

import "testing"

func TestEFs(t *testing.T) {
	efs := NewEFs()

	efs.Mkdir("test", 0755)
	efs.Mkdir("test/tt", 0755)
	efs.Create("test/ok")
	list_fs, err := efs.OpenDir("test")
	if err != nil {
		t.Fatalf("opend failed %v", err)
	}
	for i, f := range list_fs {
		t.Logf("fs[%v]=%v", i, f)
	}
}
