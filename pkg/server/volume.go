package server

type Volume struct {
	Id   uint64
	Size uint64
	Name string

	ChunkMap map[uint64]int
}

type VolumeCtl struct {
}

func (volctl *VolumeCtl) VolumeCreate(name string) *Volume {
	return nil
}

func (volctl *VolumeCtl) VolumeDelete(vol *Volume) error {
	return nil
}

func (volctl *VolumeCtl) VolumeList() {

}

func (volctl *VolumeCtl) Read() {

}

func (volctl *VolumeCtl) Write() {

}
