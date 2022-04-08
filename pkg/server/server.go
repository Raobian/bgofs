package server

import (
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
	pb.RegisterVolumeServiceServer(grpcServer, NewVolumeService())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.DFATAL("grpcServer.Serve err: %v", err)
	}
}
