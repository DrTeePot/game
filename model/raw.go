package model

type RawModel struct {
	vao         uint32
	vertexCount int32
}

func NewRawModel(arrayObject uint32, vCount int) RawModel {
	return RawModel{
		vao:         arrayObject,
		vertexCount: int32(vCount),
	}
}

func (r RawModel) ID() uint32         { return r.vao }
func (r RawModel) VertexCount() int32 { return r.vertexCount }

type TexturedModel struct {
	model   RawModel
	texture Texture
}

func NewTexturedModel(model RawModel, texture Texture) TexturedModel {
	return TexturedModel{
		model:   model,
		texture: texture,
	}
}

func (t TexturedModel) Texture() Texture   { return t.texture }
func (t TexturedModel) RawModel() RawModel { return t.model }
