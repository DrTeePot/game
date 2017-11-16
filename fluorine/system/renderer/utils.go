package renderer

import (
	"math"

	"github.com/DrTeePot/game/components"
	"github.com/DrTeePot/game/maths"
	"github.com/DrTeePot/game/shaders"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func createProjectionMatrix() mgl32.Mat4 {
	fov := float32(70)
	nearPlane := float32(0.1)
	farPlane := float32(1000)

	// TODO make a display class that handles the display
	aspectRatio := float32(1260) / 720
	y_scale := float32(1 / float32(math.Tan(float64(mgl32.DegToRad(fov/2)))) * aspectRatio)
	x_scale := y_scale / aspectRatio
	frustrumLength := farPlane - nearPlane

	matrix := mgl32.Mat4{}
	matrix.Set(0, 0, x_scale)
	matrix.Set(1, 1, y_scale)
	matrix.Set(2, 2, -((farPlane + nearPlane) / frustrumLength))
	matrix.Set(3, 2, -1)
	matrix.Set(2, 3, -((2 * nearPlane * farPlane) / frustrumLength))
	matrix.Set(3, 3, 0)
	return matrix
}

func prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
