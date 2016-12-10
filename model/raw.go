package model

type RawModel struct {
	indices     uint32
	vboID       uint32
	vertexCount int32
}

func NewRawModel(bufferID uint32, indices uint32, vCount int) RawModel {
	return RawModel{
		vboID:       bufferID,
		indices:     indices,
		vertexCount: int32(vCount),
	}
}

func (r RawModel) BufferID() uint32   { return r.vboID }
func (r RawModel) Indices() uint32    { return r.indices }
func (r RawModel) VertexCount() int32 { return r.vertexCount }
