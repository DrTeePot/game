package loader

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"github.com/DrTeePot/game/model"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// TODO investigate generalizing this and making a Loader interface and struct
// can load from vertecies, load from obj, etc

// list of vaos and vbos
var vaos []uint32
var vbos []uint32
var textures []uint32

func init() {
	vaos = make([]uint32, 10)
	vbos = make([]uint32, 10)
	textures = make([]uint32, 10)
}

func CleanUp() {
	// needs the length and a pointer to the first element
	// c bindings are a pain
	gl.DeleteVertexArrays(int32(len(vaos)), &vaos[0])
	gl.DeleteBuffers(int32(len(vbos)), &vbos[0])
	gl.DeleteTextures(int32(len(textures)), &textures[0])
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

// Loads vertecies (v), indices (i), and texture coodinates (t) to OpenGL and
//  stores relevent information in a TexturedModel
func LoadToModel(
	v []float32,
	i []uint32,
	t []float32,
	n []float32,
) model.RawModel {

	vao := createVAO()            // create vertex array object
	_ = storeArrayBuffer(0, v, 3) // store vertices
	_ = storeArrayBuffer(1, t, 2) // store texture coordinates
	_ = storeArrayBuffer(2, n, 3) // store normal coordinates
	storeElementArrayBuffer(i)    // store element arraw, we can only have one
	unbindVAO()                   // let other vao's load

	return model.NewRawModel(vao, len(i))
}

// LoadTexture loads a png file into a texture, returns the ID
func LoadTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
