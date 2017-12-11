package render

import (
	//	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

// TODO pretty sure mgl has this functionality
func CreateTransformationMatrix(translation mgl32.Vec3, rx, ry, rz, scale float32) mgl32.Mat4 {
	// TODO create different methods for these, look at Java code from gr 12
	// new identity matrix
	matrix := mgl32.Ident4()

	// set tranlation
	// assuming matrix row and column start at 0
	matrix.Set(0, 3, translation.X())
	matrix.Set(1, 3, translation.Y())
	matrix.Set(2, 3, translation.Z())

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

	// scale the matrix
	matrix.Mul(scale)

	//	fmt.Println(matrix)

	return matrix
}
