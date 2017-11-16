package renderer

import (
	"math"

	"github.com/DrTeePot/game/components"
	"github.com/DrTeePot/game/maths"
	"github.com/DrTeePot/game/system"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type renderEntity struct {
	*ecs.Basic
	// TODO need these components
	*Transform  // location, rotation, scale
	*Renderable // mesh and texture
}

type lightEntity struct {
	*ecs.Basic
	// TODO create
	*Transform
	*Colour
}

type Renderer struct {
	shader   BasicShader // TODO generalize
	entities []renderEntity
	lights   []lightEntity
	camera   Camera
}

func NewRenderer(shader BasicShader, camera Camera) Renderer {
	return Renderer{
		camera:   camera,
		shader:   shader,
		entities: []renderEntity{},
		lights:   []lightEntity{},
	}
}

func (r *Renderer) Init() {
	// setup openGL
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.DepthFunc(gl.LESS)

	// utils
	projectionMatrix := createProjectionMatrix()

	// shader.go
	r.shader.Start()
	r.shader.LoadProjectionMatrix(projectionMatrix)
	r.shader.Stop()
}

func (s *Renderer) AddEntity(id ecs.Basic, transform Transform, render Renderable) error {
	// TODO add entities to system
}

func (s *Renderer) AddLight(id ecs.Basic, transform Transform, colour Colour) error {
	// TODO add light to system

}

func (s *Renderer) Remove(system.Basic) {
	// TODO implement removal of things from system
}

/*
* Called each time we need to update our components
* dt - time in milliseconds since last update
 */
func (s *Renderer) Update(dt uint32) {
	// our draw calls go here
	prepare() // utils.go
	for index, e := range entities {
		shader = s.shader // TODO get rid of this, make all s.shader

		mesh := e.Mesh
		texture := e.Texture

		shader.Start()
		shader.LoadLightPosition(s.lights[0].Position()) // TODO implement Transform.Position
		shader.LoadLightColour(s.lights[0].Colour())     // TODO implement Colour.Colour()

		viewMatrix := maths.CreateViewMatrix(camera)
		shader.LoadViewMatrix(viewMatrix)

		// bind our VAO and the buffers we're using
		gl.BindVertexArray(mesh.ID())
		gl.EnableVertexAttribArray(0) // enable vertecies
		gl.EnableVertexAttribArray(1) // enable textures
		gl.EnableVertexAttribArray(2) // enable normals

		r := e.Rotation()

		transformationMatrix := maths.CreateTransformationMatrix(
			e.Position(),
			r.X(), r.Y(), r.Z(),
			e.Scale())
		shader.LoadTransformationMatrix(transformationMatrix)
		shader.LoadSpecular(texture.Shine(), texture.Reflectivity())

		// setup texture
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture.ID())

		// draw the model
		gl.DrawElements(gl.TRIANGLES, mesh.VertexCount(),
			gl.UNSIGNED_INT, nil) // draw using elements array

		// cleanup our VAO
		gl.DisableVertexAttribArray(0) // disable vertecies
		gl.DisableVertexAttribArray(1) // disable textures
		gl.DisableVertexAttribArray(2) // disable normals
		gl.BindVertexArray(0)          // unbind model VAO

		shader.Stop()
	}
}

func (s *Renderer) Start() {
	// this is where we run our renderer loop
	// Does Update need to run in main thread?

}

// Cleanup our stuff
func (s *Renderer) Delete() {
	s.shader.Delete()
	// loader.Cleanup (remove entity VAO's from memory)
}
