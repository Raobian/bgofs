package server

import (
	"fmt"

	"github.com/Raobian/bgofs/pkg/kvengine"
)

type Meta struct {
	kv kvengine.KVEngine
}

var meta *Meta

func init() {
	kv := kvengine.NewRedisKV()
	meta = &Meta{
		kv: kv,
	}
}

func GetMeta() *Meta {
	return meta
}

func id2Key(id uint64) string {
	return fmt.Sprintf("volume_%08x", id)
}

func GetVolume(id uint64) ([]byte, error) {
	return meta.kv.Get(id2Key(id))
}

func SetVolume(id uint64, info []byte) error {
	return meta.kv.Set(id2Key(id), info)
}
