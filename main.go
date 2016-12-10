package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/DrTeePot/game/loader"
	"github.com/DrTeePot/game/render"
	"github.com/DrTeePot/game/shaders"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

	vertexShader := "shaders/vertexShader.glsl"
	fragmentShader := "shaders/fragmentShader.glsl"

	shader, err := shaders.NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	// REAL STUFFS
	vertecies := []float32{
		-0.5, 0.5, 0, // 0
		-0.5, -0.5, 0, // 1
		0.5, -0.5, 0, // 2
		0.5, 0.5, 0, // 3
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	model := loader.LoadToModel(vertecies, indices)

	// Configure global settings
	gl.DepthFunc(gl.LESS)

	gl.ClearColor(0.11, 0.545, 0.765, 0.0)

	for !window.ShouldClose() {
		render.Prepare()
		shader.Start()
		render.Render(model)
		shader.Stop()

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	shader.CleanUp()
	loader.CleanUp()
}
