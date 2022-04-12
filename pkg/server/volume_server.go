package server

import (
	"github.com/Raobian/bgofs/pkg/server/store"
)

type VolumeServer struct {
	maxId uint64
	vmap  map[uint64]*Volume
	store store.Store
}

var GvolumeServer *VolumeServer

func init() {
	basedir := "./"
	GvolumeServer = &VolumeServer{
		maxId: 0,
		store: store.NewStore(basedir),
	}
}

func CreateVolume(name string, size uint64) {
	vid := GvolumeServer.maxId
	v := &Volume{
		Id:   vid,
		Name: name,
		Size: size,
	}

	GvolumeServer.vmap[vid] = v
	GvolumeServer.maxId++
}
