package render

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render/shaders"
)

type Renderer interface {
	Render()
}

// TODO rename?
type Renderable struct {
	meshID       uint32
	vertexCount  int32
	textureID    uint32
	shine        float32
	reflectivity float32
	transform    mgl32.Mat4 // TODO create nice methods to update this
}

type shadedEntities [][]Renderable

func (e shadedEntities) entitiesByShader(shaderID uint32) []Renderable {
	return e[shaderID]
}

/*
Might also need to keep track of shaders here, since they may need
to be deleted at some point, like when the program ends?

We also need to remove VAO's from memory when things end?
*/
type RenderEngine struct {
	entities []Renderable // TODO eventually segment by shader
	lights   []Light
	programs [1]shaders.BasicShader // TODO make more general
	camera   Camera
}

func NewEngine(
	entities []Renderable,
	lights []Light,
	camera Camera,
) RenderEngine {
	// return an object that does the things
	// has a shader

	shader, err := shaders.NewBasicShader()
	if err != nil {
		panic(err)
	}

	initializeOpenGL(shader)

	return RenderEngine{
		entities: entities,
		lights:   lights,
		camera:   camera,
		programs: [1]shaders.BasicShader{shader},
	}
}

func (r RenderEngine) Start(window *glfw.Window) {
	// run a loop that just calls renderAll
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour
	for !window.ShouldClose() {
		// **** RENDER LOOP **** //
		r.renderAll() // eventually pass in real time between frames

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	r.programs[0].Delete()
}

// The bulk of the renderer
func (r RenderEngine) renderAll() {
	prepare()
	for _, e := range r.entities {
		viewMatrix := CreateViewMatrix(r.camera)
		e.Render(viewMatrix, r.lights, r.programs[0])
	}
}

func (r RenderEngine) AddLight(light Light) {
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

// TODO move to individual entities specifying their render methods
// TODO somehow improve how we associate shaders and entities... maybe an int?
func (r Renderable) Render(viewMatrix mgl32.Mat4, lights []Light, shader shaders.BasicShader) {
	// prepare shader
	shader.Start()

	// loop over this based on lights in scene
	shader.LoadLight(lights[0].position, lights[0].colour)

	shader.LoadViewMatrix(viewMatrix)
	shader.LoadTransformationMatrix(r.transform)
	shader.LoadSpecular(r.shine, r.reflectivity)

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

	shader.Stop()
}
