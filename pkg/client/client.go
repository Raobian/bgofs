package client

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/Raobian/bgofs/pkg/common"
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
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	vc = pb.NewVolumeServiceClient(conn)
	Upload()
}

func Upload() {
	f := "test_file"
	file, err := os.Open(f)
	defer file.Close()

	fst, err := file.Stat()
	fsize := fst.Size()
	log.Printf("-- bian -- fsize:%d", fsize)

	stream, err := vc.Upload(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
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
			log.Fatalf("failed to read")
			return
		}

		chkid := (uint64(1) << 32) & uint64(cid)
		stream.Send(&pb.Chunk{
			Chkid:  chkid,
			Offset: 0,
			Length: uint32(n),
			Data:   []byte(buf),
		})

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Upload get response err: %v", err)
	}
	log.Printf("--- bian -- recv: code:%d msg:%s", res.Code, res.Msg)
}
