package vfs

import (
	"io"
	"os"
)

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	// Readdir(count int) ([]os.FileInfo, error)
	// Readdirnames(n int) ([]string, error)
	// Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	// WriteString(s string) (ret int, err error)
}

type VFS interface {
	Create(name string) (File, error)
	Mkdir(name string, perm os.FileMode) error
	// MkdirAll(path string, perm os.FileMode) error
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Open(name string) (File, error)
	Remove(name string) error
	// RemoveAll(path string) error
	Rename(oldname, newname string) error
	Stat(name string) (os.FileInfo, error)
	// Name() string
	// Chmod(name string, mode os.FileMode) error
	// Chown(name string, uid, gid int) error
	// Chtimes(name string, atime time.Time, mtime time.Time) error
}
