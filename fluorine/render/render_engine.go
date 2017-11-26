package render

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render/shaders"
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
	models [][]Model // array of models by shader
	camera Camera
	// TODO probably the display, or else those constants should be on camera
}

func NewEngine(
	camera Camera,
	models [][]Model, // TODO make this an actual type
) RenderEngine {
	// return an object that does the things
	// has a shader

	shader, err := shaders.NewBasicShader()
	if err != nil {
		panic(err)
	}

	initializeOpenGL(shader)

	return RenderEngine{
		camera: camera,
		models: models,
	}
}

// The bulk of the renderer
func (r RenderEngine) Update(entities [][]transformMatrix, lights []Light) {
	prepare()
	// TODO this render loop needs to take into account [][][]TransformMatrix
	for _, shader := range r.models {
		for _, model := range shader {
			viewMatrix := CreateViewMatrix(r.camera)
			// TODO model.id might not work, more thought needed.
			model.Render(viewMatrix, lights, entities[model.id])
		}
	}
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

// Gets run on each update?
func prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
