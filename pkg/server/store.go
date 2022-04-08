package server

import (
	"fmt"
	"os"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"
)

type Store interface {
	CreateObject(obj Object) error
	DeleteObject(obj Object) error
	ListObject() error
	WriteObject() error
	ReadObject() error
}

type PlainStore struct {
	base string
}

func PathValidate(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, 0x644)
		}
	}
}

func NewStore(base string) Store {
	PathValidate(base)
	return &PlainStore{
		base: base,
	}
}

func (store *PlainStore) CreateObject(obj Object) error {
	path := fmt.Sprintf("%s_%02d_%02d_%08d", store.base, obj.nid, obj.did, obj.oid)
	file, err := os.Create(path)
	if err != nil {
		log.DERROR("create %s failed", path)
		return err
	}
	defer file.Close()

	if err := file.Truncate(int64(common.OBJSIZE)); err != nil {
		log.DERROR("truncate %s failed", path)
		return err
	}

	return nil
}

func (store *PlainStore) DeleteObject(obj Object) error {

	return nil
}

func (store *PlainStore) ListObject() error {

	return nil
}

func (store *PlainStore) WriteObject() error {

	return nil
}

func (store *PlainStore) ReadObject() error {

	return nil
}
