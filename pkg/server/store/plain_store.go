package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"
)

type PlainStore struct {
	base string
}

func PathValidate(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, 0x777)
		}
	}
}

func NewPlainStore(base string) *PlainStore {
	abs, _ := filepath.Abs(base)
	PathValidate(abs)
	return &PlainStore{
		base: abs,
	}
}

func (s *PlainStore) object2Path(obj *Object) string {
	return fmt.Sprintf("%s/%08x/%08x_%08x_%08x", s.base, obj.Volid, obj.NodeID, obj.DiskID, obj.ObjectID)
}

func (s *PlainStore) Create(obj *Object) error {
	path := s.object2Path(obj)
	log.DINFO("creae path:%s", path)
	d, _ := filepath.Split(path)
	PathValidate(d)
	file, err := os.Create(path)
	if err != nil {
		log.DERROR("create %s failed %v", path, err)
		return err
	}
	defer file.Close()

	if err := file.Truncate(int64(common.OBJSIZE)); err != nil {
		log.DERROR("truncate %s failed %v", path, err)
		return err
	}
	return nil
}

func (s *PlainStore) Remove(obj *Object) error {
	return os.Remove(s.object2Path(obj))
}

func (s *PlainStore) List() (*Object, error) {
	return nil, nil
}

func (s *PlainStore) Read(oio *ObjectIO) (int, error) {
	file, err := os.OpenFile(s.object2Path(oio.Obj), os.O_RDONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.ReadAt(oio.Data, int64(oio.Offset))
}

func (s *PlainStore) Write(oio *ObjectIO) (int, error) {
	file, err := os.OpenFile(s.object2Path(oio.Obj), os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.WriteAt(oio.Data, int64(oio.Offset))
}
