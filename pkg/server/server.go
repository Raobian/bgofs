package server

import (
	"io"
	"net"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"

	pb "github.com/Raobian/bgofs/pkg/pb"

	"google.golang.org/grpc"
)

func Init() {
	log.SetLevel(log.INFO)
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
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
		log.DINFO("server recv:%x off:%d len:%d Data:%s\n", res.Chkid, res.Offset, res.Length, string(res.Data))
		// chkid := Chkid{
		// 	volid: uint32(res.Chkid),
		// }
	}
}

func Server() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.DFATAL("net.Listen err: %v", err)
	}
	log.DINFO(Address + " net.Listing...")

	var grpcOpts = []grpc.ServerOption{
		grpc.MaxMsgSize(int(common.MaxMsgSize)),
	}
	grpcServer := grpc.NewServer(grpcOpts...)

	// 在gRPC服务器注册我们的服务
	pb.RegisterVolumeServiceServer(grpcServer, &VolumeService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.DFATAL("grpcServer.Serve err: %v", err)
	}
}
