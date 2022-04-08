package server

import (
	"github.com/Raobian/bgofs/pkg/kvengine"
)

type Meta struct {
	kv kvengine.KVEngine
}

func NewMeta() *Meta {
	kv := kvengine.NewRedisKV()
	return &Meta{
		kv: kv,
	}
}
