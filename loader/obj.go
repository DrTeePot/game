package loader

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/DrTeePot/game/model"

	"github.com/go-gl/mathgl/mgl32"
)

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

// TODO create type loader that can be added to models?
func LoadObjModel(filename string) (model model.RawModel, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return model, err
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
		return model, err
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

	return LoadToModel(verteciesArray, indices, textureArray, normalsArray), nil

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
