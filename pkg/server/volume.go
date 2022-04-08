package server

type VolumeInfo struct {
	Id   uint64
	Size uint64
	Name string
}
type Volume struct {
	Info     VolumeInfo
	ChunkMap map[uint64]Chunk
}
