package model

type Texture struct {
	id uint32
}

func NewTexture(id uint32) Texture {
	return Texture{id: id}
}

func (t Texture) ID() uint32 {
	return t.id
}
