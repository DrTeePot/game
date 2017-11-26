package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render/shaders"
)

type transformMatrix = mgl32.Mat4

type Model struct {
	id           uint32
	shader       shaders.BasicShader
	meshID       uint32
	vertexCount  int32
	textureID    uint32
	shine        float32
	reflectivity float32
}

// TODO NewModel

// Create a new entry in the transform array
func (m Model) Instantiate() {

}

// TODO somehow improve how we associate shaders and entities... maybe an int?
func (m Model) Render(
	viewMatrix mgl32.Mat4,
	lights []Light,
	transforms []transformMatrix,
) {
	// prepare shader
	m.shader.Start()

	// loop over this based on lights in scene
	m.shader.LoadLight(lights[0].position, lights[0].colour)

	m.shader.LoadViewMatrix(viewMatrix)

	m.shader.LoadSpecular(m.shine, m.reflectivity)

	// bind our VAO and the buffers we're using
	gl.BindVertexArray(m.meshID)
	gl.EnableVertexAttribArray(0) // enable vertecies
	gl.EnableVertexAttribArray(1) // enable textures
	gl.EnableVertexAttribArray(2) // enable normals

	// setup texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, m.textureID)

	// draw the model
	// draw using elements array
	for _, transform := range transforms {
		m.shader.LoadTransformationMatrix(transform)
		gl.DrawElements(gl.TRIANGLES, m.vertexCount, gl.UNSIGNED_INT, nil)
	}

	// cleanup our VAO
	gl.DisableVertexAttribArray(0) // disable vertecies
	gl.DisableVertexAttribArray(1) // disable textures
	gl.DisableVertexAttribArray(2) // disable normals
	gl.BindVertexArray(0)          // unbind model VAO

	m.shader.Stop()
}
