package main

import (
	"runtime"

	"github.com/DrTeePot/game/fluorine"
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
	camera := render.Camera{}

	engine := render.NewEngine(
		camera,
		[][]render.Model{{render.Model{}}, {render.Model{}}},
	)

	somethingFloat := fluorine.NewFloatComponent(
		"test",
		1,
		fluorine.FloatNoOp,
	)
	somethingString := fluorine.NewStringComponent(
		"test",
		1,
		fluorine.StringNoOp,
	)

	arrayOfFloat := []fluorine.FloatComponent{
		somethingFloat,
	}

	arrayOfString := []fluorine.StringComponent{
		somethingString,
	}

	// TODO need a way to create array of components, or none
	store := fluorine.CreateStore(arrayOfFloat, arrayOfString)

	fluorine := fluorine.New(
		window,
		engine,
		store,
	)

	// load shaders

	// load entities, with shaders

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
	fluorine.Start()
	// TODO this won't work properly right now since it will only
	// render entities that were passed into NewEngine, since it's pass
	// by value

}
