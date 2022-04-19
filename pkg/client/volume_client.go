package client

import (
	"context"

	"github.com/Raobian/bgofs/pkg/common/log"
	pb "github.com/Raobian/bgofs/pkg/pb"
)

func CreateVolume(name string, size uint64) error {
	res, err := Vclient.Create(context.Background(), &pb.VolumeInfo{
		Name:  name,
		Size_: uint64(size),
	})
	if err != nil {
		log.DFATAL("Upload file:%s rpc create failed err: %v", name, err)
	}
	log.DINFO("create volume %s id:%d", name, res.Volid)
	return nil
}
