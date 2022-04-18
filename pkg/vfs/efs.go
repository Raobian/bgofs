package vfs

import (
	"log"
	"os"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

type EFs struct {
	EndPointer string
	etcdClient *etcd.Client
}

func NewEFs() *EFs {
	efs := &EFs{
		EndPointer: "http://127.0.0.1:2379",
	}
	efs.etcdClient = etcd.NewClient([]string{efs.EndPointer})
	return efs
}

func (efs *EFs) Close() {
	efs.etcdClient.Close()
}

func (efs *EFs) Name() string {
	return "efs"
}

func (efs *EFs) Create(name string) (File, error) {
	_, err := efs.etcdClient.Set(name, "", 0)
	log.Printf("create %v err:%v", name, err)
	if err != nil {
		return nil, err
	}
	return efs.Open(name)
}

func (efs *EFs) Mkdir(name string, perm os.FileMode) error {
	_, err := efs.etcdClient.CreateDir(name, 0)
	log.Printf("mkdir %v err %v", name, err)
	// TODO: create md_info
	return err
}

func (efs *EFs) OpenDir(name string) ([]string, error) {
	res, err := efs.etcdClient.Get(name, false, false)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, e := range res.Node.Nodes {
		sp := strings.Split(e.Key, "/")
		f := sp[len(sp)-1]
		files = append(files, f)
		log.Printf("-- bian -- e:%v isdir:%v", e.Key, e.Dir)
	}
	return files, nil
}

func (efs *EFs) Remove(name string) error {
	_, err := efs.etcdClient.Get(name, false, false)
	if err != nil {
		return err
	}

	_, err = efs.etcdClient.DeleteDir(name)
	return err
}

func (efs *EFs) RemoveAll(path string) error {
	_, err := efs.etcdClient.Get(path, false, false)
	if err != nil {
		return err
	}
	_, err = efs.etcdClient.Delete(path, true)
	log.Printf("delete %v err:%v", path, err)
	return err
}

func (efs *EFs) Open(name string) (File, error) {
	_, err := efs.etcdClient.Get(name, false, false)
	if err != nil {
		return nil, err
	}
	efsf, err := GetFileByName(name)
	if err != nil {
		efsf = NewEFsFile(name)
	}
	return efsf, nil
}
