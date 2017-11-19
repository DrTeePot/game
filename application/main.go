package main

import (
	"runtime"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/render"
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
	window := render.NewWindow(windowWidth, windowHeight, "game")
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
	entities := []render.Renderable{{}}
	lights := []render.Light{render.NewLight(
		mgl32.Vec3{5, 5, -15},
		mgl32.Vec3{1, 1, 1},
	)}
	camera := render.Camera{}

	renderEngine := render.NewEngine(entities, lights, camera)

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
	// TODO this won't work properly right now since it will only
	// render entities that were passed into NewEngine, since it's pass
	// by value
	renderEngine.Start(window)

}
