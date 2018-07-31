package render

import (
	"math"
	// "time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/components/mesh"
	"github.com/DrTeePot/game/fluorine/components/transform"
	"github.com/DrTeePot/game/fluorine/render/shaders"
	"github.com/DrTeePot/game/fluorine/store"
)

// TODO make a display class that handles the display
const (
	FIELD_OF_VIEW = float32(70)
	NEAR_PLANE    = float32(0.1)
	FAR_PLANE     = float32(1000)
	ASPECT_RATIO  = float32(1260) / 720
)

/*
Might also need to keep track of shaders here, since they may need
to be deleted at some point, like when the program ends?

We also need to remove VAO's from memory when things end?
*/
type RenderEngine struct {
	// this could be an array of renderer objects
	window *glfw.Window
	models []Model
	camera Camera
	// TODO probably the display, or else those constants should be on camera
}

func NewEngine(
	camera Camera,
	models []Model,
	shader shaders.BasicShader,
	w *glfw.Window,
) RenderEngine {

	// TODO user should get to specify shader and how it initalizes
	initializeOpenGL(shader)
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour

	return RenderEngine{
		window: w,
		camera: camera,
		models: models,
	}
}

func (r RenderEngine) Update(s store.Store) {

	// Update function
	transformEntities := s.Component(transform.TransformName).State()
	meshEntities := s.Component(mesh.MeshComponent).State()

	renderEntities := make(map[uint32][]transformMatrix)
	for entityID, meshData := range meshEntities {
		transformData, ok := transformEntities[entityID]
		if !ok {
			transformData = []float32{0, 0, 0, 0, 0, 0}
		}
		meshID := uint32(meshData[0])
		currentEntities := renderEntities[meshID]

		matrix := CreateTransformationMatrix(
			mgl32.Vec3{
				transformData[0], // position x
				transformData[1], // position y
				transformData[2], // position z
			},
			transformData[3], // rotation x
			transformData[4], // rotation y
			transformData[5], // rotation z
			1,                //scale
		)

		renderEntities[meshID] = append(currentEntities, matrix)
	}

	// TODO light component
	lights := []Light{NewLight(
		mgl32.Vec3{5, 5, -15},
		mgl32.Vec3{1, 1, 1},
	)}

	// **** RENDER LOOP **** //
	// TODO eventually this will need to take into account time offsets
	r.render(renderEntities, lights)

	r.window.SwapBuffers()
}

// The bulk of the renderer
func (r RenderEngine) render(entities map[uint32][]transformMatrix, lights []Light) {
	prepare()
	viewMatrix := CreateViewMatrix(r.camera)

	for meshID, transforms := range entities {
		r.models[meshID].Render(viewMatrix, lights, transforms)
	}
}

// Called when we first create the render engine
func initializeOpenGL(shader shaders.BasicShader) {
	projectionMatrix := createProjectionMatrix()

	shader.Start()
	shader.LoadProjectionMatrix(projectionMatrix)
	shader.Stop()
}

func createProjectionMatrix() mgl32.Mat4 {
	y_scale := float32(1 / float32(math.Tan(float64(mgl32.DegToRad(FIELD_OF_VIEW/2)))) * ASPECT_RATIO)
	x_scale := y_scale / ASPECT_RATIO
	frustrumLength := FAR_PLANE - NEAR_PLANE

	matrix := mgl32.Mat4{}
	matrix.Set(0, 0, x_scale)
	matrix.Set(1, 1, y_scale)
	matrix.Set(2, 2, -((FAR_PLANE + NEAR_PLANE) / frustrumLength))
	matrix.Set(3, 2, -1)
	matrix.Set(2, 3, -((2 * NEAR_PLANE * FAR_PLANE) / frustrumLength))
	matrix.Set(3, 3, 0)
	return matrix
}

func prepare() {
	// TODO this just needs to be enabled for shaders that use these features
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.DEPTH_TEST)

	// this needs to run each update
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
