package loader

import (
	"github.com/DrTeePot/game/model"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// list of vaos and vbos
var vaos []uint32
var vbos []uint32

func init() {
	vaos = make([]uint32, 10)
	vbos = make([]uint32, 10)
}

func CleanUp() {
	// needs the length and a pointer to the first element
	// c bindings are a pain
	gl.DeleteVertexArrays(int32(len(vaos)), &vaos[0])
	gl.DeleteBuffers(int32(len(vbos)), &vbos[0])
}

func LoadToModel(vertecies []float32, indices []uint32) model.RawModel {
	// Create Vertex array object
	var vertexArrayID uint32
	gl.GenVertexArrays(1, &vertexArrayID)
	gl.BindVertexArray(vertexArrayID)

	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertecies)*4, gl.Ptr(vertecies), gl.STATIC_DRAW)

	var indicesBufferID uint32
	gl.GenBuffers(1, &indicesBufferID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indicesBufferID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4, // uint32 is 4 bytes
		gl.Ptr(indices), gl.STATIC_DRAW)

	// allows for cleanup
	vaos = append(vaos, vertexArrayID)
	vbos = append(vbos, vertexBuffer)
	vbos = append(vbos, indicesBufferID)

	return model.NewRawModel(vertexBuffer, indicesBufferID, len(vertecies))
}

func bindIndicesBuffer(indices []uint32) uint32 {
	var vboID uint32
	gl.GenBuffers(2, &vboID)
	vbos = append(vbos, vboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*4, // float32 is 4 bytes
		gl.Ptr(indices), gl.STATIC_DRAW)
	return vboID
}
