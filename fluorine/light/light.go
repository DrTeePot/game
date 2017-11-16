package light

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Light struct {
	position mgl32.Vec3
	colour   mgl32.Vec3
}

func Create(p, c mgl32.Vec3) Light {
	return Light{p, c}
}

func (l Light) Position() mgl32.Vec3      { return l.position }
func (l Light) Colour() mgl32.Vec3        { return l.colour }
func (l *Light) SetPosition(p mgl32.Vec3) { l.position = p }
func (l *Light) SetColour(c mgl32.Vec3)   { l.colour = c }
