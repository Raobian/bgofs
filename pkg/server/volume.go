package server

import (
	"io"

	"github.com/Raobian/bgofs/pkg/common/log"
	pb "github.com/Raobian/bgofs/pkg/pb"
)

type VolumeService struct {
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
		log.DINFOf("server recv:%x off:%d len:%d Data:%s\n", res.Chkid, res.Offset, res.Length, string(res.Data))
	}
}
