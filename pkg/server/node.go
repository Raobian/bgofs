package server

import "google.golang.org/grpc"

type Node struct {
	nid  uint32
	conn *grpc.ClientConn
}
