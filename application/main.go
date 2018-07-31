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
	window := render.NewWindow(windowWidth, windowHeight, "Fluorine")
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
		window,
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

	// Hard coded ID's for the mesh and entity I added
	entityID := uint32(0)
	dragonMeshID := float32(0)

	// Actions, which get dispatched into the store
	addDragonMesh := mesh.SetMesh(entityID, dragonMeshID)
	moveDragonMesh := transform.SetPosition(entityID, 0, -5, -20)
	rotateAction := transform.IncreaseRotation(entityID, 1, 0, 0)

	// spawns goroutine
	myStore.DispatchFloat(addDragonMesh)
	myStore.DispatchFloat(moveDragonMesh)
	myStore.DispatchFloat(rotateAction)

	// **** MAIN LOOP **** //
	/*
		- start render engine (runs in this thread, this has to happen
		last)
	*/
	fluorine.Start()

	for !window.ShouldClose() {
		// TODO this is where systems that need to run in main thread should
		// happen, other systems can be started in goroutines
		// Should call 'Update' on each system
		// the thing added to the store should just be a buffered channel
		// that will send the store to the relevent update function

		continue
	}

	fluorine.Done()
}
