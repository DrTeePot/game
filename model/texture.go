package model

type Texture struct {
	id uint32

	shine        float32 // specular power
	reflectivity float32 // specular intensity
}

func NewTexture(id uint32) Texture {
	return Texture{id: id}
}

func (t Texture) ID() uint32            { return t.id }
func (t Texture) Shine() float32        { return t.shine }
func (t Texture) Reflectivity() float32 { return t.reflectivity }

func (t *Texture) SetShine(s float32)        { t.shine = s }
func (t *Texture) SetReflectivity(r float32) { t.reflectivity = r }
