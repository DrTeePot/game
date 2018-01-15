package fluorine

import (
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/store"
)

type transformMatrix = mgl32.Mat4
type empty struct{} // uses 0 space

/*
TODO fluorine is an event loop, needs to have a list of
systems to run that fulfil some update interface.

systems will have dependencies on certain components.

namespace components using component name, check if requisite
component exists on init.
*/

type fluorine struct {
	window     *glfw.Window
	systems    []store.System
	store      store.Store
	renderTick *time.Ticker
}

func New(
	w *glfw.Window,
	y []store.System,
	s store.Store,
) fluorine {
	dependencies := make(map[string]empty)

	for _, componentName := range s.RegisteredComponents() {
		dependencies[componentName] = empty{}
	}

	for _, system := range y {
		for _, componentName := range system.Dependencies() {
			_, ok := dependencies[componentName]
			if !ok {
				panic("Dependencies of systems are not met")
			}
		}
	}

	return fluorine{
		window:     w,
		systems:    y,
		store:      s,
		renderTick: time.NewTicker(time.Duration(time.Second / 60)),
	}
}

func (f fluorine) Start() {
	// setup OpenGL
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour

	// run a loop that just calls renderAll
	for !f.window.ShouldClose() {
		// update state according to dispatched actions
		f.store.Update()

		select {
		case <-f.renderTick.C:
			for _, system := range f.systems {
				// TODO this takes parameters
				system.Update(f.store)
			}

			// TODO window, input adn swapping buffer
			//   should not be here
			// Show the things we drew on the buffer
			f.window.SwapBuffers()
		default:
		}

		// grab input events
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	f.store.Close()
	// TODO delete all the stuff from openGL
}
