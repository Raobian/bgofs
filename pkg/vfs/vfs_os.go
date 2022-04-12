package vfs

import (
	"os"
	"path/filepath"
)

type OsFS struct {
	basePath string
}

func CreateOsFS(name string) *OsFS {
	return &OsFS{
		basePath: name,
	}
}

func (fs *OsFS) joinPath(name string) string {
	return filepath.Join(fs.basePath, name)
}

func (fs *OsFS) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(fs.joinPath(name), perm)
}

func (fs *OsFS) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(fs.joinPath(name), flag, perm)
}

func (fs *OsFS) Open(name string) (File, error) {
	return os.Open(fs.joinPath(name))
}

func (fs *OsFS) Remove(name string) error {
	return os.Remove(fs.joinPath(name))
}

func (fs *OsFS) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs *OsFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
