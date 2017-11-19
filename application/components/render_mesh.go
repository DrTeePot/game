package components

import (
	"bufio"
	"os"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

// TODO simplify render_mesh + mesh_utils
type Mesh struct {
	File string

	// TODO should this data be here?
	loaded bool
	mesh   meshData
}

// TODO maybe we don't want to cache this? then only renderer can see it
type meshData struct {
	vao         uint32
	vertexCount int32
}

// fulfills system.MeshLoader interface
// TODO if meshData is removed, this doesn't need to be a pointer
func (r *Mesh) LoadMesh() (vao uint32, vertexCount int32, err error) {
	if r.loaded {
		return r.mesh.vao, r.mesh.vertexCount, nil
	}

	// we haven't loaded, load our mesh
	vao, vertexCount, err = loadMeshFile(r.File)
	r.mesh = meshData{
		vao:         vao,
		vertexCount: vertexCount,
	}

	// if an error didn't occur, then we loaded correctly :)
	if err != nil {
		r.loaded = true
	}

	return vao, vertexCount, err
}

func loadMeshFile(filename string) (vao uint32, vc int32, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	vertecies := []mgl32.Vec3{}
	textureCoords := []mgl32.Vec2{}
	normals := []mgl32.Vec3{}
	faces := []face{}

	var textureArray []float32
	var normalsArray []float32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// we still have a newline
		tokens := strings.Split(scanner.Text(), " ")

		switch string(tokens[0]) {
		case "v":
			vertex := mgl32.Vec3{
				parse(tokens[1]),
				parse(tokens[2]),
				parse(tokens[3]),
			}
			vertecies = append(vertecies, vertex)

		case "vt":
			texture := mgl32.Vec2{
				parse(tokens[1]),
				parse(tokens[2]),
			}
			textureCoords = append(textureCoords, texture)

		case "vn":
			vertex := mgl32.Vec3{
				parse(tokens[1]),
				parse(tokens[2]),
				parse(tokens[3]),
			}
			normals = append(normals, vertex)

		case "f":
			vertex1 := strings.Split(tokens[1], "/")
			vertex2 := strings.Split(tokens[2], "/")
			vertex3 := strings.Split(tokens[3], "/")

			faces = append(faces, createFace(vertex1, vertex2, vertex3))
		}
	}
	if err = scanner.Err(); err != nil {
		return // error
	}

	// make the model
	indices := []uint32{}
	textureArray = make([]float32, len(vertecies)*2)
	normalsArray = make([]float32, len(vertecies)*3)
	verteciesArray := make([]float32, len(vertecies)*3)
	for _, f := range faces {
		indices = processVertex(f[0],
			indices, textureCoords, normals,
			textureArray, normalsArray)
		indices = processVertex(f[1],
			indices, textureCoords, normals,
			textureArray, normalsArray)
		indices = processVertex(f[2],
			indices, textureCoords, normals,
			textureArray, normalsArray)
	}

	// convert vertecies
	for i, v := range vertecies {
		verteciesArray[i*3] = v.X()
		verteciesArray[i*3+1] = v.Y()
		verteciesArray[i*3+2] = v.Z()
	}

	vao, vcount := loadMeshToOpenGL(
		verteciesArray,
		indices,
		textureArray,
		normalsArray)

	return vao, vcount, nil
}
