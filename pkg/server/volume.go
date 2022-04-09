package server

type VolumeInfo struct {
	Id   uint64
	Size uint64
	Name string
}

type Volume struct {
	Info     VolumeInfo
	ChunkMap map[uint64]int
}

type VolumeService struct {
}

func (vs *VolumeService) VolumeCreate(name string) *Volume {
	return nil
}

func (vs *VolumeService) VolumeDelete(vol *Volume) error {
	return nil
}

func (vs *VolumeService) VolumeList() {

}

func (vol *Volume) Read() {

}

func (vol *Volume) Write() {

}
