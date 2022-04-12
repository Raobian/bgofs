package kvengine

import (
	"errors"
)

type KVEngine interface {
	Close() error
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	List(prefix string) ([]string, error)
}

type Iterator interface {
	Valid() bool
	Key() string
	Value() string
	Next() error
	Close()
}

var (
	ENotFound = errors.New("not found")
)

func IsENotFound(err error) bool {
	return errors.Is(err, ENotFound)
}
