package entity

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position         mgl32.Vec3
	Pitch, Yaw, Roll float32

	movementVector mgl32.Vec3
	rotationVector mgl32.Vec3
}

func (c *Camera) AddVelocity(v mgl32.Vec3) {
	c.movementVector = c.movementVector.Add(v)
}

func (c *Camera) Move(axis mgl32.Vec3, velocity float32) func() {
	return c.movementFunction(axis.Mul(velocity))
}

// should be on entity
func (c *Camera) Update() {
	// c.IncreasePosition is also a function
	c.Position = c.Position.Add(c.movementVector)
}

func (c *Camera) movementFunction(v mgl32.Vec3) func() {
	return func() {
		c.AddVelocity(v)
	}
}
