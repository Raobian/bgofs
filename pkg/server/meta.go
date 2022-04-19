package server

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/Raobian/bgofs/pkg/config"
	"github.com/Raobian/bgofs/pkg/kvengine"
)

type Meta struct {
	kv kvengine.KVEngine
}

var meta *Meta

func init() {
	if !config.IsServer {
		return
	}

	kv := kvengine.NewRedisKV()
	meta = &Meta{
		kv: kv,
	}
}

func GetMeta() *Meta {
	return meta
}

const volume_prefix = "volumes/"

func id2Key(id uint64) string {
	return fmt.Sprintf("%s%08x", volume_prefix, id)
}

func GetVolume(id uint64) ([]byte, error) {
	return meta.kv.Get(id2Key(id))
}

func SetVolume(id uint64, info []byte) error {
	return meta.kv.Set(id2Key(id), info)
}

func ListVolume() {
	ks, err := meta.kv.List(volume_prefix)
	if err != nil {
		log.DFATAL("list failed %v", err)
	}
	for _, k := range ks {
		sk := strings.Split(k, volume_prefix)
		id, err := strconv.ParseUint(sk[len(sk)-1], 16, 64)
		if err != nil {
			log.DFATAL("bad volume: %s", sk)
		}

		log.DINFO("get id: %08x", id)
		_, err = meta.kv.Get(k)
		if err != nil {
			log.DFATAL("get %s info failed err:%v", k, err)
		}

	}
}
