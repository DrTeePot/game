package system

import (
	"github.com/DrTeePot/game/components"

	"github.com/EngoEngine/ecs" // just used for BasicEntity component
)

type MeshLoader interface {
	LoadMesh() (vao uint32, vertexCount uint32)
}
type TextureLoader interface {
	LoadTexture() // TODO add return signiture
}

type Renderable interface {
	MeshLoader
	TextureLoader
}

type renderer struct {
	camera nil
	shader

	lights
	entities
}

func Renderer(shader Shader, camera Camera) {
	// do something with the shader and camera, return a renderer

}

func (r *renderer) AddEntity(id ecs.BasicEntity, t Transform, r Renderable) {
	// renderable interface has a Load() function that returns a VAO and
	//  a vertex count so we can add it.
}

func (r *renderer) AddLight(id ecs.BasicEntity, t Transform, c Colour) {

}

func (r *renderer) Remove(id ecs.BasicEntity) {
	// remove from whichever
}

func (r *renderer) update(dt uint32) {
	// actually do the rendering
	// one frame
}

func (r *renderer) Start() {
	// run a loop that calls update
}

func (r *renderer) Stop() {
	// stop our loop
}

func (r *renderer) Delete() {
	// stop our loop if necessary and clean up
	// remove vaos, shader, etc.
}
