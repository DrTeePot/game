package components

import (
	"strconv"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// list of vaos and vbos
var vaos []uint32
var vbos []uint32
var textures []uint32

// CleanUp removes our VAO's and buffers from memory
func DeleteMesh() {
	// needs the length and a pointer to the first element
	// c bindings are a pain
	gl.DeleteVertexArrays(int32(len(vaos)), &vaos[0])
	gl.DeleteBuffers(int32(len(vbos)), &vbos[0])
	gl.DeleteTextures(int32(len(textures)), &textures[0])
}

func init() {
	vaos = []uint32{}
	vbos = []uint32{}
	textures = []uint32{}
}

func createVAO() uint32 {
	// Create Vertex array object
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)
	vaos = append(vaos, vertexArrayID)
	return vertexArrayID
}

func unbindVAO() {
	gl.BindVertexArray(0)
}

// stores a vertex array into a VBO at attributeNumber with coordinate size width
func storeArrayBuffer(attributeNumber uint32, vertecies []float32, width int32) uint32 {
	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertecies)*4, gl.Ptr(vertecies), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attributeNumber, width, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	vbos = append(vbos, vertexBuffer) // for cleanup
	return vertexBuffer
}

func storeElementArrayBuffer(indices []uint32) {
	var indicesBufferID uint32
	gl.GenBuffers(1, &indicesBufferID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indicesBufferID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4, // uint32 is 4 bytes
		gl.Ptr(indices), gl.STATIC_DRAW)
	vbos = append(vbos, indicesBufferID) // cleanup
}

func loadMeshToOpenGL(
	v []float32,
	i []uint32,
	t []float32,
	n []float32,
) (vao uint32, vertexCount int32) {
	vao = createVAO()             // create vertex array object
	_ = storeArrayBuffer(0, v, 3) // store vertices
	_ = storeArrayBuffer(1, t, 2) // store texture coordinates
	_ = storeArrayBuffer(2, n, 3) // store normal coordinates
	storeElementArrayBuffer(i)    // store element arraw, we can only have one
	unbindVAO()                   // let other vao's load

	return vao, int32(len(i))
}

// ****** LOADING HELPERS ****
type faceVertex struct {
	N uint32
	T uint32
	I uint32
}
type face [3]faceVertex

func createFace(vertecies ...[]string) face {
	f := face{}
	for i, v := range vertecies {
		f[i] = faceVertex{
			I: parseI(v[0]),
			T: parseI(v[1]),
			N: parseI(v[2]),
		}
	}
	return f

}

func parseI(b string) uint32 {
	p, _ := strconv.ParseUint(b, 10, 64)
	return uint32(p)
}
func parse(b string) float32 {
	p, _ := strconv.ParseFloat(b, 64)
	return float32(p)
}

func processVertex(
	vertexData faceVertex,
	indices []uint32,
	textures []mgl32.Vec2,
	normals []mgl32.Vec3,
	textureArray []float32,
	normalsArray []float32,
) []uint32 {
	// TODO this is a mess, heavy refactoring needed
	index := vertexData.I - 1
	indices = append(indices, index)

	// obj starts at 1 not 0
	currentTex := textures[vertexData.T-1]
	textureArray[index*2] = currentTex.X()
	// blender renders textures with reversed y axis
	textureArray[index*2+1] = 1 - currentTex.Y()

	currentNorm := normals[vertexData.N-1]
	normalsArray[index*3] = currentNorm.X()
	normalsArray[index*3+1] = currentNorm.Y()
	normalsArray[index*3+2] = currentNorm.Z()

	return indices
}
