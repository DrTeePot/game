package render

import (
	"math"

	"github.com/DrTeePot/game/fluorine/render/shaders"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer interface {
	Render()
}

type lightEntity struct {
	position math.Vec3
	colour   math.Vec3
}

// TODO rename?
type Renderable struct {
	meshID       uint32
	vertexCount  uint32
	textureID    uint32
	shine        uint32
	reflectivity uint32
	transform    mgl32.Mat4 // TODO create nice methods to update this
}

type shadedEntities [][]Renderable

func (e shadedEntities) entitiesByShader(shaderID uint32) {
	return e[shaderID]
}

/*
Might also need to keep track of shaders here, since they may need
to be deleted at some point, like when the program ends?

We also need to remove VAO's from memory when things end?
*/
type RenderEngine struct {
	entities shadedEntities
	lights   []lightEntity
	shaders  []shaders.BasicShader // actually not this
	camera   Camera
}

func NewEngine(entities []Renderable, lights lightEntity, camera Camera) {
	// return an object that does the things
	// has a shader
	initializeOpenGL()

	return RenderEngine{
		entities: entities,
		lights:   lights,
		camera:   camera,
	}
}

func (r RenderEngine) Start() {
	// run a loop that just calls renderAll
}

// The bulk of the renderer
func (r RenderEngine) renderAll() {
	prepare()
	for _, e := range r.entities {
		e.render(r.camera.viewMatrix, r.lights)
	}
}

func (r RenderEngine) AddLight(light lightEntity) {
	// TODO pretty obvious, will use a channel
}

func (r RenderEngine) AddRenderable(entity Renderable) {
	// TODO pretty obvious, will use a channel
}

// Called when we first create the render engine
func initializeOpenGL(shader shaders.BasicShader) {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.DepthFunc(gl.LESS)

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

// Gets run on each update?
func prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Should this be a method on a Renderable? probably this should be an interface, let each component take care of how it wants to be rendered?
func (r Renderable) render(viewMatrix mgl32.Mat4, lights []lightEntity) {
	// prepare shader
	r.shader.Start()

	// loop over this based on lights in scene
	r.shader.LoadLightPosition(lights[0].position)
	r.shader.LoadLightColour(lights[0].colour)

	r.shader.LoadViewMatrix(viewMatrix)
	r.shader.LoadTransformationMatrix(r.transform)
	r.shader.LoadSpecular(r.shine, r.reflectivity)

	// bind our VAO and the buffers we're using
	gl.BindVertexArray(r.meshID)
	gl.EnableVertexAttribArray(0) // enable vertecies
	gl.EnableVertexAttribArray(1) // enable textures
	gl.EnableVertexAttribArray(2) // enable normals

	// setup texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, r.textureID)

	// draw the model
	// draw using elements array
	gl.DrawElements(gl.TRIANGLES, r.vertexCount, gl.UNSIGNED_INT, nil)

	// cleanup our VAO
	gl.DisableVertexAttribArray(0) // disable vertecies
	gl.DisableVertexAttribArray(1) // disable textures
	gl.DisableVertexAttribArray(2) // disable normals
	gl.BindVertexArray(0)          // unbind model VAO

	r.shader.Stop()
}
