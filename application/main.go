package main

import (
	"path"
	"runtime"

	"github.com/DrTeePot/game/fluorine"
	"github.com/DrTeePot/game/fluorine/components/mesh"
	"github.com/DrTeePot/game/fluorine/components/transform"
	"github.com/DrTeePot/game/fluorine/render"
	"github.com/DrTeePot/game/fluorine/render/shaders"
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
	_, mainFileLocation, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	applicationDirectory := path.Dir(mainFileLocation)

	/*
		Input handler:
		- create new input handler
		- bind things to keys
		- start engine
	*/

	camera := render.Camera{}

	shader, err := shaders.NewBasicShader()
	if err != nil {
		panic(err)
	}

	dragon, err := render.NewModel(
		"dragon",
		shader,
		applicationDirectory+"/assets/dragon.obj",
		applicationDirectory+"/assets/blank.png",
		0.5,
		0.5,
	)
	if err != nil {
		panic(err)
	}

	models := []render.Model{
		dragon, // 0
	}

	engine := render.NewEngine(
		camera,
		models,
		shader,
	)

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
	meshCom := mesh.CreateComponent()

	registeredComponents := store.NewRegistry([]store.UniversalComponent_float32{
		testCom,
		transformCom,
		meshCom,
	})

	myStore := store.CreateStore(registeredComponents)

	fluorine := fluorine.New(
		window,
		[]store.System{
			store.NewSystem(
				[]string{
					transformCom.Name(),
					meshCom.Name(),
				},
				engine.Update,
			),
		},
		myStore,
	)

	entityID := uint32(0)
	dragonMeshID := float32(0)
	addDragonMesh := mesh.SetMesh(entityID, dragonMeshID)
	moveDragonMesh := transform.SetPosition(entityID, 0, -5, -20)

	// spawns goroutine
	myStore.DispatchFloat(addDragonMesh)
	myStore.DispatchFloat(moveDragonMesh)

	// **** MAIN LOOP **** //
	/*
		- start render engine (runs in this thread, this has to happen
		last)
	*/
	fluorine.Start()
}
