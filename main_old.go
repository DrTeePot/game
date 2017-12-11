package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/DrTeePot/game/fluorine/display"
	"github.com/DrTeePot/game/fluorine/input"
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

type Dragon struct {
	// TODO not used because entity does this kinda?
	Mesh      components.Mesh
	Texture   components.Texture
	Transform components.Transform3D
}

func NewDragon(mesh string, tex string, pos mgl32.Vec3, rot mgl32.Vec3) {
	return entity.NewEntity("assets/dragon.obj", "assets.blank.png",
		mgl32.Vec3{0, -5, -20}, mgl32.Vec3{0, 0, 0})
}

func main() {
	window := display.New()
	defer display.Exit()

	// **** SETUP **** //
	// shaders
	vertexShader := "shaders/vertexShader.glsl"
	fragmentShader := "shaders/fragmentShader.glsl"
	shader, err := renderer.NewBasicShader(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	renderSystem := renderer.NewRenderer(shader)
	renderSystem.Init()

	e := NewDragon("assets/dragon.obj", "assets/blank.png",
		mgl32.Vec3{0, -5, -20}, mgl32.Vec3{0, 0, 0})

	// camera
	var camera entity.Camera // new 0'd camera

	// light
	coolLight := light.Create(
		mgl32.Vec3{5, 5, -15},
		mgl32.Vec3{1, 1, 1},
	)

	// unit axis
	x := mgl32.Vec3{1, 0, 0}
	y := mgl32.Vec3{0, 1, 0}
	z := mgl32.Vec3{0, 0, 1}

	// **** KEYBORD INPUT **** //
	keyboard := input.NewKeyboardListener(window)
	// TODO rename this binding function
	keyboard.OnMovementKey(glfw.KeyW, camera.Move(z, -0.2), camera.Move(z, 0.2))
	keyboard.OnMovementKey(glfw.KeyS, camera.Move(z, 0.2), camera.Move(z, -0.2))
	keyboard.OnMovementKey(glfw.KeyD, camera.Move(x, 0.2), camera.Move(x, -0.2))
	keyboard.OnMovementKey(glfw.KeyA, camera.Move(x, -0.2), camera.Move(x, 0.2))

	// jump?
	keyboard.OnMovementKey(glfw.KeyLeftShift, camera.Move(y, -0.2), camera.Move(y, 0.2))
	keyboard.OnMovementKey(glfw.KeySpace, camera.Move(y, 0.2), camera.Move(y, -0.2))

	// **** MAIN LOOP **** //
	gl.ClearColor(0.11, 0.545, 0.765, 0.0) // set background colour
	for !window.ShouldClose() {
		// rotate our entity
		e.RotateBy(mgl32.Vec3{0, 1, 0})
		camera.Update() // movement

		// **** RENDER LOOP **** //
		renderSystem.Update(0) // eventually pass in real time between frames

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	shader.Delete()
	model.CleanUp()
}
