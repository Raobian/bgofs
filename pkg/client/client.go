package client

import (
	"context"
	"io"
	"os"

	"github.com/Raobian/bgofs/pkg/common"
	"github.com/Raobian/bgofs/pkg/common/log"
	pb "github.com/Raobian/bgofs/pkg/pb"

	"google.golang.org/grpc"
)

// Address 连接地址
const Address string = ":8000"

var vc pb.VolumeServiceClient

func Run(fname string) {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.DFATAL("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	vc = pb.NewVolumeServiceClient(conn)
	Upload(fname)
}

func Upload(fname string) {
	log.DINFO("fname:%s opening...", fname)
	file, err := os.Open(fname)
	defer file.Close()

	stream, err := vc.Upload(context.Background())
	if err != nil {
		log.DFATAL("Upload list err: %v", err)
	}
	defer stream.CloseSend()

	buf := make([]byte, common.CHKSIZE)
	cid := uint32(0)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			log.DFATAL("failed to read")
			return
		}

		stream.Send(&pb.Chunk{
			Chkid: &pb.Chkid{
				Volid: 1,
				Id:    cid,
			},
			Offset: 0,
			Length: uint32(n),
			Data:   []byte(buf),
		})

		if err == io.EOF {
			break
		}
		if err != nil {
			log.DFATAL("stream request err: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.DFATAL("Upload get response err: %v", err)
	}
	log.DINFO("--- bian -- recv: code:%d msg:%s", res.Code, res.Msg)
}
