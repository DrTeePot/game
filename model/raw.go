package model

type RawModel struct {
	vaoID       uint32
	vertexCount int32
}

func NewRawModel(id uint32, vCount int) RawModel {
	return RawModel{
		vaoID:       id,
		vertexCount: int32(vCount),
	}
}

func (r RawModel) VAOID() uint32      { return r.vaoID }
func (r RawModel) VertexCount() int32 { return r.vertexCount }
