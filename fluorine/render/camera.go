package render

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position         mgl32.Vec3
	Pitch, Yaw, Roll float32

	// TODO remove movement stuff
	movementVector mgl32.Vec3
	rotationVector mgl32.Vec3
	viewMatrix     mgl32.Mat4
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

func CreateViewMatrix(camera Camera) mgl32.Mat4 {
	matrix := mgl32.Ident4()

	rx := camera.Pitch
	ry := camera.Yaw
	rz := camera.Roll

	// set rotation
	// idk if this will do waht i want
	quat := mgl32.AnglesToQuat(
		mgl32.DegToRad(rx),
		mgl32.DegToRad(ry),
		mgl32.DegToRad(rz),
		mgl32.XYZ)
	rotation := quat.Mat4()

	// do we multiply these?
	matrix = matrix.Mul4(rotation)

	cameraPos := camera.Position
	negativeCamera := mgl32.Vec3{-cameraPos.X(), -cameraPos.Y(), -cameraPos.Z()}

	// tranlate by negativeCamera
	matrix.Set(0, 3, negativeCamera.X())
	matrix.Set(1, 3, negativeCamera.Y())
	matrix.Set(2, 3, negativeCamera.Z())

	return matrix
}
