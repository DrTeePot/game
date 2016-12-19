package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/DrTeePot/game/entity"
	"github.com/DrTeePot/game/input"
	"github.com/DrTeePot/game/light"
	"github.com/DrTeePot/game/loader"
	"github.com/DrTeePot/game/maths"
	"github.com/DrTeePot/game/model"
	"github.com/DrTeePot/game/render"
	"github.com/DrTeePot/game/shaders"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
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
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// TODO manage all this in a display thing
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "GAME", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// **** SETUP **** //
	// shaders
	vertexShader := "shaders/vertexShader.glsl"
	fragmentShader := "shaders/fragmentShader.glsl"

	shader, err := shaders.NewBasicShader(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	render.Initialize(shader)

	// model and texture
	rawModel, err := loader.LoadObjModel("assets/dragon.obj")
	if err != nil {
		fmt.Println("problem loading model")
		panic(err)
	}
	textureID, err := loader.LoadTexture("assets/blank.png")
	if err != nil {
		panic(err)
	}
	texture := model.NewTexture(textureID)
	texture.SetShine(10)
	texture.SetReflectivity(1)
	model := model.NewTexturedModel(rawModel, texture)

	e := entity.Entity{
		Model:    model,
		Position: mgl32.Vec3{0, -5, -20},
		RotX:     0,
		RotY:     0,
		RotZ:     0,
		Scale:    1,
	}

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
	// TODO rename this keybinding
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
		e.IncreaseRotation(0, 1, 0)
		camera.Update()

		// **** RENDER LOOP **** //
		render.Prepare()
		shader.Start()
		shader.LoadLightPosition(coolLight.Position())
		shader.LoadLightColour(coolLight.Colour())

		// render the piece of the world that the camera
		//    is looking at
		viewMatrix := maths.CreateViewMatrix(camera)
		shader.LoadViewMatrix(viewMatrix)

		render.Render(e, shader)

		shader.Stop()

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// **** CLEANUP **** //
	shader.Delete()
	loader.CleanUp()
}
