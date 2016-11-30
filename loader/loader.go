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

func LoadToVAO(positions []float32) model.RawModel {
	vaoID := createVAO()
	storeDataInAttributeList(0, positions)
	unbindVAO()
	return model.NewRawModel(vaoID, len(positions)/3)
}

func CleanUp() {
	// needs the length and a pointer to the first element
	// c bindings are a pain
	gl.DeleteVertexArrays(int32(len(vaos)), &vaos[0])
	gl.DeleteBuffers(int32(len(vbos)), &vbos[0])
}

func createVAO() uint32 {
	var vaoID uint32
	gl.GenVertexArrays(1, &vaoID)
	vaos = append(vaos, vaoID)
	gl.BindVertexArray(vaoID)
	return vaoID
}

func storeDataInAttributeList(attributeNumber uint32, data []float32) {
	// to use float64 data, use gl.DOUBLE instead of gl.FLOAT below
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	vbos = append(vbos, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	// multiply data by 4 because float32 is 4 bytes
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	// 3*4 because i have 3 points in a vertex, and each point is 4 bytes
	gl.VertexAttribPointer(attributeNumber, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func unbindVAO() {
	gl.BindVertexArray(0)
}
