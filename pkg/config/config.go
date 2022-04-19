package config

import "github.com/Raobian/bgofs/pkg/common/log"

var (
	CHKSIZE     = 1 << 22
	OBJSIZE     = CHKSIZE
	MAX_IO_SIZE = 4096
	IOV_MAX     = CHKSIZE / MAX_IO_SIZE

	GRPCAddr       = ":8000"
	GRPCMaxMsgSize = CHKSIZE + 4096

	DefaultLogLevel = log.INFO

	EtcdEndPointer = "127.0.0.1:2379"
	RedisAddr      = "127.0.0.1:6379"
)
