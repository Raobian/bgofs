package server

import (
	"io"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"
	pb "github.com/Raobian/bgofs/pkg/pb"
)

type VolumeService struct {
	chkCache common.Cache
	meta     *Meta
}

func NewVolumeService() *VolumeService {
	return &VolumeService{
		chkCache: *common.NewCache(1024),
		meta:     NewMeta(),
	}
}

func (vs *VolumeService) Upload(srv pb.VolumeService_UploadServer) error {
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			log.DINFO("-- bian -- eof")
			return srv.SendAndClose(&pb.ChunkResponse{
				Code: 0,
				Msg:  "ok",
			})
		}
		if err != nil {
			return err
		}
		log.DINFO("server recv:%x off:%d len:%d Data:%s\n", res.Chkid, res.Offset, res.Length, string(res.Data))
		// chkid := Chkid{
		// 	volid: uint32(res.Chkid),
		// }
		v, ok := vs.chkCache.Get(res.Chkid)
		if !ok {
			vs.chkCache.Add(res.Chkid, res.Offset)
			v = res.Offset
		}
		log.DINFO("chkcache get val:%v", v)
	}
}
