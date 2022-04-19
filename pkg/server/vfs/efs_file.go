package vfs

import (
	"log"
	"os"
	"sync"
)

type EFsFile struct {
	FileName string
	Pos      int64
	ref      int32
}

type EFsRoot struct {
	MaxId uint64
	// fileMap map[string]*EFsFile
	// mu      sync.Mutex
	fileMap sync.Map
}

var Root EFsRoot

func init() {
	// Root.fileMap = make(map[string]*EFsFile)
}

func GetFileByName(name string) (*EFsFile, error) {
	// efsf := Root.fileMap[name]
	efsf, ok := Root.fileMap.Load(name)
	if !ok {
		return nil, os.ErrNotExist
	}
	return efsf.(*EFsFile), nil
}

func NewEFsFile(name string) *EFsFile {
	// sp := strings.Split(name, "/")
	// fname := sp[len(sp)-1]
	efsf := &EFsFile{
		FileName: name,
		Pos:      0,
		ref:      0,
	}
	// Root.fileMap[name] = efsf
	Root.fileMap.Store(name, efsf)
	efsf.Ref()
	return efsf
}

func (f *EFsFile) Ref() {
	f.ref++
}

func (f *EFsFile) UnRef() {
	f.ref--
	if f.ref == 0 {
		// delete(Root.fileMap, f.FileName)
		Root.fileMap.Delete(f.FileName)
	}
}

func (f *EFsFile) Name() string {
	return f.FileName
}

func (f *EFsFile) Stat() (os.FileInfo, error) {
	log.Printf("file :%s stat", f.FileName)
	return nil, nil
}

func (f *EFsFile) Close() error {
	log.Printf("file :%s close", f.FileName)
	f.UnRef()
	return nil
}

func (f *EFsFile) Read(p []byte) (n int, err error) {
	log.Printf("file :%s read p:%d", f.FileName, len(p))
	return 0, nil
}

/*
	SeekStart means relative to the start of the file,
	SeekCurrent means relative to the current offset, and
	SeekEnd means relative to the end.
*/
func (f *EFsFile) Seek(offset int64, whence int) (int64, error) {
	log.Printf("file :%s seek off:%d", f.FileName, offset)
	return 0, nil
}

func (f *EFsFile) Write(p []byte) (n int, err error) {
	log.Printf("file :%s write p:%d", f.FileName, len(p))
	return 0, nil
}
