package fluorine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/components/mesh"
	"github.com/DrTeePot/game/fluorine/components/transform"
	"github.com/DrTeePot/game/fluorine/render"
	"github.com/DrTeePot/game/fluorine/store"
)

type transformMatrix = mgl32.Mat4

type fluorine struct {
	window       *glfw.Window
	renderEngine render.RenderEngine
	store        store.Store
}

func New(
	w *glfw.Window,
	r render.RenderEngine,
	s store.Store,
) fluorine {
	return fluorine{
		window:       w,
		renderEngine: r,
		store:        s,
	}
}

func (f fluorine) Start() {
	// setup OpenGL
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour

	// run a loop that just calls renderAll
	for !f.window.ShouldClose() {
		// update state according to dispatched actions
		f.store.Update()

		transformEntities := f.store.State(transform.TransformName)
		meshEntities := f.store.State(mesh.MeshComponent)

		renderEntities := make(map[uint32][]transformMatrix)
		for entityID, meshData := range meshEntities {
			transformData, ok := transformEntities[entityID]
			if !ok {
				transformData = []float32{0, 0, 0, 0, 0, 0}
			}
			meshID := uint32(meshData[0])
			currentEntities := renderEntities[meshID]

			matrix := render.CreateTransformationMatrix(
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
		lights := []render.Light{render.NewLight(
			mgl32.Vec3{5, 5, -15},
			mgl32.Vec3{1, 1, 1},
		)}

		// **** RENDER LOOP **** //
		// TODO eventually this will need to take into account time offsets
		f.renderEngine.Update(renderEntities, lights)

		// Show the things we drew on the buffer
		f.window.SwapBuffers()

		// grab input events
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	f.store.Close()
	// TODO delete all the stuff from openGL
}
