package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render/shaders"
)

type transformMatrix = mgl32.Mat4

type modeller interface {
	Instantiate()
	Render(
		viewMatrix mgl32.Mat4,
		lights []Light,
		transforms []transformMatrix)
}

type Model struct {
	// TODO eventually store a reference to the shader we want to use for this
	name        string
	shader      shaders.BasicShader
	meshID      uint32
	vertexCount int32

	// TODO these should all be their own components
	textureID    uint32
	shine        float32
	reflectivity float32
}

func NewModel(
	name string, // used in Model Library
	shader shaders.BasicShader,
	meshFile string,
	textureFile string,
	shine float32,
	reflectivity float32,
) (model Model, err error) {

	meshID, vertexCount, err := loadMeshFile(meshFile)
	if err != nil {
		return model, err
	}
	textureID, err := loadTexture(textureFile)
	if err != nil {
		return model, err
	}

	return Model{
		name,
		shader,
		meshID,
		vertexCount,
		textureID,
		shine,
		reflectivity,
	}, nil
}

// TODO somehow improve how we associate shaders and entities... maybe an int?
func (m Model) Render(
	viewMatrix mgl32.Mat4,
	lights []Light,
	transforms []transformMatrix,
) {
	m.setupShader(viewMatrix, lights)

	m.loadModelToGPU()

	// draw using elements array
	for _, transform := range transforms {
		m.shader.LoadTransformationMatrix(transform)
		gl.DrawElements(gl.TRIANGLES, m.vertexCount, gl.UNSIGNED_INT, nil)
	}

	m.cleanupModel()

	m.stopShader()
}

func (m Model) setupShader(
	viewMatrix mgl32.Mat4,
	lights []Light,
) {
	// prepare shader
	m.shader.Start()

	// loop over this based on lights in scene
	m.shader.LoadLight(lights[0].position, lights[0].colour)

	m.shader.LoadViewMatrix(viewMatrix)

	m.shader.LoadSpecular(m.shine, m.reflectivity)
}

func (m Model) loadModelToGPU() {
	// bind our VAO and the buffers we're using
	gl.BindVertexArray(m.meshID)
	gl.EnableVertexAttribArray(0) // enable vertecies
	gl.EnableVertexAttribArray(1) // enable textures
	gl.EnableVertexAttribArray(2) // enable normals

	// setup texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, m.textureID)
}

func (m Model) cleanupModel() {
	gl.DisableVertexAttribArray(0) // disable vertecies
	gl.DisableVertexAttribArray(1) // disable textures
	gl.DisableVertexAttribArray(2) // disable normals
	gl.BindVertexArray(0)          // unbind model VAO
}

func (m Model) stopShader() {
	m.shader.Stop()
}
