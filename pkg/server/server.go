package server

import (
	"context"
	"io"
	"net"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"

	pb "github.com/Raobian/bgofs/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (vs *VolumeService) Create(ctx context.Context, req *pb.VolumeInfo) (*pb.VolumeCtlResponse, error) {

	return &pb.VolumeCtlResponse{
		Code:  0,
		Msg:   "ok",
		Volid: 1,
	}, nil
}

func (vs *VolumeService) Remove(ctx context.Context, req *pb.VolumeInfo) (*pb.VolumeCtlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func (vs *VolumeService) Write(srv pb.VolumeService_WriteServer) error {
	log.DINFO("write start")
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			log.DINFO("write end")
			return srv.SendAndClose(&pb.VolumeResponse{
				Code: 0,
				Msg:  "ok",
			})
		}
		if err != nil {
			return err
		}
		log.DINFO("server recv:%x off:%d len:%d Data:%s\n", res.Volid, res.Offset, res.Length, string(res.Data))
	}
}

func (vs *VolumeService) Read(srv pb.VolumeService_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
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
