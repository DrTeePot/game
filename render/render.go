package render

import (
	"math"

	"github.com/DrTeePot/game/entity"
	"github.com/DrTeePot/game/maths"
	"github.com/DrTeePot/game/shaders"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func Initialize(shader shaders.BasicShader) {
	projectionMatrix := createProjectionMatrix()
	shader.Start()
	shader.LoadProjectionMatrix(projectionMatrix)
	shader.Stop()
}

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

func Prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// TODO RawModel as an interface?
func Render(entity entity.Entity, shader shaders.BasicShader) {
	tModel := entity.Model
	model := tModel.RawModel()

	// bind our VAO and the buffers we're using
	gl.BindVertexArray(model.ID())
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)

	transformationMatrix := maths.CreateTransformationMatrix(
		entity.Position,
		entity.RotX, entity.RotY, entity.RotZ,
		entity.Scale)
	shader.LoadTransformationMatrix(transformationMatrix)

	// setup texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tModel.Texture().ID())

	// draw the model
	gl.DrawElements(gl.TRIANGLES, model.VertexCount(),
		gl.UNSIGNED_INT, nil)

	// cleanup our VAO
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.BindVertexArray(0)
}
