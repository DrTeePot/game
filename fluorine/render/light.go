package render

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Light struct {
	position mgl32.Vec3
	colour   mgl32.Vec3
}

func NewLight(p, c mgl32.Vec3) Light {
	return Light{p, c}
}

func (l Light) Position() mgl32.Vec3 { return l.position }
func (l Light) Colour() mgl32.Vec3   { return l.colour }
