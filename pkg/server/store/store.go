package store

type Object struct {
	Volid    uint32
	NodeID   uint32
	DiskID   uint32
	ObjectID uint32
}

type ObjectIO struct {
	Obj    *Object
	Offset uint32
	Length uint32
	Data   []byte
}

type Store interface {
	Create(*Object) error
	Remove(*Object) error
	List() (*Object, error)
	Read(*ObjectIO) (int, error)
	Write(*ObjectIO) (int, error)
}

func NewStore(base string) Store {
	return NewPlainStore(base)
}
