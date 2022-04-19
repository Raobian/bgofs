package server

import (
	"context"
	"io"
	"net"

	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/Raobian/bgofs/pkg/config"

	pb "github.com/Raobian/bgofs/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Init() {
	// log.SetLevel(config.DefaultLogLevel)
}

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
		log.DINFO("server recv:%x off:%d len:%d\n", res.Volid, res.Offset, res.Length)
	}
}

func (vs *VolumeService) Read(srv pb.VolumeService_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func Server() {
	listener, err := net.Listen("tcp", config.GRPCAddr)
	if err != nil {
		log.DFATAL("net.Listen err: %v", err)
	}
	log.DINFO(" net.Listing on %s ...", config.GRPCAddr)

	var grpcOpts = []grpc.ServerOption{
		grpc.MaxMsgSize(int(config.GRPCMaxMsgSize)),
	}
	grpcServer := grpc.NewServer(grpcOpts...)

	// 在gRPC服务器注册我们的服务
	pb.RegisterVolumeServiceServer(grpcServer, &VolumeService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.DFATAL("grpcServer.Serve err: %v", err)
	}
}
