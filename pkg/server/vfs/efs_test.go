package vfs

import (
	"fmt"
	"testing"
)

func TestDir(t *testing.T) {
	efs := NewEFs()
	if efs == nil {
		t.Fatalf("new efs failed")
	}

	t_dir := "test"
	efs.RemoveAll(t_dir)

	err := efs.Mkdir(t_dir, 0755)
	if err != nil {
		t.Fatalf("mkdir %s err:%v", t_dir, err)
	}

	err = efs.Mkdir(t_dir, 0755)
	if err == nil {
		t.Fatalf("mkdir %s err:%v should be exist", t_dir, err)
	}

	subdir_name := "sub_test"
	subdir := fmt.Sprintf("%s/%s", t_dir, subdir_name)
	err = efs.Mkdir(subdir, 0755)
	if err != nil {
		t.Fatalf("make subdir %s err:%v", subdir, err)
	}

	list_fs, err := efs.OpenDir(t_dir)
	if err != nil {
		t.Fatalf("opend %s failed %v", t_dir, err)
	}
	exist := false
	for _, f := range list_fs {
		if f == subdir_name {
			exist = true
		}
	}
	if !exist {
		t.Fatalf("subdir %s not exist", subdir)
	}

	err = efs.Remove(t_dir)
	if err == nil {
		t.Fatalf("remove %s err should be not empty", t_dir)
	}

	err = efs.RemoveAll(t_dir)
	if err != nil {
		t.Fatalf("remove %s err:%v", t_dir, err)
	}

	efs.Close()
}

func TestEFsFile(t *testing.T) {
	efs := NewEFs()

	dname := "test"
	efs.RemoveAll(dname)
	efs.Mkdir(dname, 0755)

	fname := "tfile"
	fpath := fmt.Sprintf("%s/%s", dname, fname)
	f, err := efs.Create(fpath)
	if err != nil {
		t.Fatalf("create failed %v", err)
	}
	if f.Name() != fpath {
		t.Fatalf("create fname:%s should be %s", f.Name(), fname)
	}

	f.Stat()

	f.Close()

	efs.Remove(fpath)
	efs.Remove(dname)
	efs.Close()
}
