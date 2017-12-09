package main

import (
	"runtime"

	"github.com/DrTeePot/game/fluorine"
	"github.com/DrTeePot/game/fluorine/components/transform"
	"github.com/DrTeePot/game/fluorine/render"
	"github.com/DrTeePot/game/fluorine/store"
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

	// load shaders

	// load models, using shaders

	/*
		fluorine:
		- create reducers, likely in a different file
		- create store with reducers
		- generate top level engine object with store and sub-engines
		- start
	*/
	somethingFloat := store.NewFloatReducer(1, store.FloatNoOp)

	testCom := store.NewUniversalComponent_float32("test", somethingFloat)
	transformCom := transform.CreateTransformComponent()

	registeredComponents := store.NewRegistry([]store.UniversalComponent_float32{
		testCom,
		transformCom,
	})

	store := store.CreateStore(registeredComponents)

	fluorine := fluorine.New(
		window,
		engine,
		store,
	)

	// **** MAIN LOOP **** //
	/*
		- start render engine (runs in this thread, this has to happen
		last)
	*/
	fluorine.Start()
}
