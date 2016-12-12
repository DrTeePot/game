package entity

import (
	"github.com/DrTeePot/game/model"
	"github.com/go-gl/mathgl/mgl32"
)

// TODO revisit this, almost certainly not the best way of structuring this
type Entity struct {
	Model            model.TexturedModel
	Position         mgl32.Vec3
	RotX, RotY, RotZ float32
	Scale            float32
}

func (e *Entity) IncreasePosition(dx, dy, dz float32) {
	e.Position = e.Position.Add(mgl32.Vec3{dx, dy, dz})
}

func (e *Entity) IncreaseRotation(dx, dy, dz float32) {
	// TODO this seems sketchy...
	e.RotX += dx
	e.RotY += dy
	e.RotZ += dz
}
