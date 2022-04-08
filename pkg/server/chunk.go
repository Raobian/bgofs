package server

type Chkid struct {
	volid uint32
	idx   uint32
}

type Chunk struct {
	// id Chkid

	rep  int
	reps []Object

	ref int
}

func OpenChunk(volid, idx uint32, rep int) *Chunk {
	return nil
}

func CreatChunk(volid, idx uint32, rep int) *Chunk {
	chk := &Chunk{
		// id: Chkid{
		// 	volid: volid,
		// 	idx:   idx,
		// },
		rep: rep,
		ref: 0,
	}

	return chk
}

func ListChunk() {

}

func (chk *Chunk) Delete() {

}

func (chk *Chunk) Write() {

}

func (chk *Chunk) Read() {

}

func (chk *Chunk) Ref() {
	chk.ref++
}

func (chk *Chunk) UnRef() {
	chk.ref--
	if chk.ref == 0 {

	}
}
