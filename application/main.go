package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/input"
	"github.com/DrTeePot/game/fluorine/render"
	"github.com/DrTeePot/game/fluorine/render/light"
	"github.com/DrTeePot/game/fluorine/render/maths"
	"github.com/DrTeePot/game/fluorine/render/system/renderer"

	"./components"
	"./entity"
)

const (
	windowWidth  = 1260
	windowHeight = 720
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to
	//	be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	window := render.NewWindow()
	// probably also do cleanup?
	defer render.CloseWindow()

	// **** SETUP **** //

	/*
		Input handler:
		- create new input handler
		- bind things to keys
		- start engine
	*/

	/*
		Rules engine:
		- create rules, probably in a yaml
		- start engine
	*/

	/*
		Render engine:
		- Create new render engine
		- load shaders to render engine
		- load meshes and textures to render engine
		- figure out how lights are handled
		// TODO eventually make entites and lights using AddLight and AddEntity
	*/
	renderEngine = render.NewEngine(entities, lights)

	/*
		Terrain:
		- set parameters and generate
	*/

	/*
		Core:
		- create store
		- start state machine
		- generate top level entity with store
		- start
	*/

	// **** MAIN LOOP **** //
	/*
		- start render engine (runs in this thread, this has to happen
		last)
	*/
	renderEngine.Start(camera)

	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour
	for !window.ShouldClose() {
		// **** RENDER LOOP **** //
		render.Update(0) // eventually pass in real time between frames

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	shader.Delete()
}
