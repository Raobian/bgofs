
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null)
REVISIONDATE := $(shell git log HEAD --pretty=format:'%ad' --date short 2>/dev/null)

PKG=bgofs/pkg/version
LDFLAGS = -s -w
LDFLAGS += -X ${PKG}.revision=$(REVISION) \
		   -X ${PKG}.revisionDate=$(REVISIONDATE)

TARGET=bgofs

${TARGET}: main.go pb
	go build -o $@ 
# go build -ldflags="$(LDFLAGS)" -o $@ 

pb:
	protoc --gofast_out=plugins=grpc:. ./pkg/pb/volume.proto


clean:
	rm -rf ${TARGET} ./pkg/pb/*.pb.go