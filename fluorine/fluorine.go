package fluorine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render"
)

type transformMatrix = mgl32.Mat4

type fluorine struct {
	window       *glfw.Window
	renderEngine render.RenderEngine
	store        store
}

func New(
	window *glfw.Window,
	renderEngine render.RenderEngine,
	store store,
) fluorine {
	return fluorine{
		window:       window,
		renderEngine: renderEngine,
		store:        store,
	}
}

func (f fluorine) Start() {
	// setup OpenGL
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour

	// run a loop that just calls renderAll
	for !f.window.ShouldClose() {
		// update state according to dispatched actions
		f.store.update()

		// TODO transform data into a way that the render engine can process

		// **** RENDER LOOP **** //
		entities := [][]transformMatrix{
			[]transformMatrix{
				transformMatrix{0, 0, 0, 0},
				transformMatrix{0, 0, 0, 0},
			},
		}
		lights := []render.Light{render.NewLight(
			mgl32.Vec3{0, 0, 0},
			mgl32.Vec3{0, 0, 0},
		)}
		// eventually this will need to take into account time offsets
		f.renderEngine.Update(entities, lights)

		// Show the things we drew on the buffer
		f.window.SwapBuffers()

		// grab input events
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	f.store.Close()
	// TODO delete all the stuff from openGL
}
