package client

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/Raobian/bgofs/pkg/config"
	pb "github.com/Raobian/bgofs/pkg/pb"

	"google.golang.org/grpc"
)

var vc pb.VolumeServiceClient
var conn *grpc.ClientConn

func connectToServer() pb.VolumeServiceClient {
	var err error
	conn, err = grpc.Dial(config.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.DFATAL("net.Connect err: %v", err)
	}
	vc = pb.NewVolumeServiceClient(conn)
	return vc
}

func closeClient() {
	conn.Close()
}

func Upload(fname string) {
	connectToServer()
	defer closeClient()

	log.DINFO("fname:%s opening...", fname)
	file, err := os.Open(fname)
	if err != nil {
		log.DFATAL("Upload file:%s open failed err: %v", fname, err)
	}
	defer file.Close()

	finfo, err := file.Stat()
	if err != nil {
		log.DFATAL("Upload file:%s stat failed err: %v", fname, err)
	}
	size := finfo.Size()

	sname := strings.Split(fname, "/")
	filename := sname[len(sname)-1]

	res, err := vc.Create(context.Background(), &pb.VolumeInfo{
		Name:  filename,
		Size_: uint64(size),
	})
	if err != nil {
		log.DFATAL("Upload file:%s rpc create failed err: %v", fname, err)
	}

	stream, err := vc.Write(context.Background())
	if err != nil {
		log.DFATAL("Upload file:%s rpc write stream failed err: %v", fname, err)
	}
	defer stream.CloseSend()

	buf := make([]byte, config.CHKSIZE)
	var off uint64
	off = 0
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return
		}

		vio := &pb.VolumeIO{
			Volid:  res.Volid,
			Offset: off,
			Length: uint32(n),
			Data:   buf,
		}
		log.DINFO("write to server vol:%v offset:%v len:%v", vio.Volid, vio.Offset, vio.Length)

		err = stream.Send(vio)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.DFATAL("stream send volid:%v off:%v len:%v failed err:%v", res.Volid, off, n, err)
		}
		off += uint64(n)
	}

	res1, err := stream.CloseAndRecv()
	if err != nil {
		log.DFATAL("Upload get response err: %v", err)
	}
	log.DINFO("--- bian -- recv: code:%d msg:%s", res1.Code, res1.Msg)
}
