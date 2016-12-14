package loader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DrTeePot/game/model"

	"github.com/go-gl/mathgl/mgl32"
)

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
	indices := []uint32{}

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
			// we have moved on
			// create the arrays for the next part
			textureArray = make([]float32, len(vertecies)*2)
			normalsArray = make([]float32, len(vertecies)*3)

			vertex1 := strings.Split(tokens[1], "/")
			vertex2 := strings.Split(tokens[2], "/")
			vertex3 := strings.Split(tokens[3], "/")

			indices = processVertex(vertex1,
				indices, textureCoords, normals,
				textureArray, normalsArray)
			indices = processVertex(vertex2,
				indices, textureCoords, normals,
				textureArray, normalsArray)
			indices = processVertex(vertex3,
				indices, textureCoords, normals,
				textureArray, normalsArray)

			for scanner.Scan() {
				tokens := strings.Split(scanner.Text(), " ")

				switch string(tokens[0]) {
				case "f":
					vertex1 := strings.Split(tokens[1], "/")
					vertex2 := strings.Split(tokens[2], "/")
					vertex3 := strings.Split(tokens[3], "/")

					indices = processVertex(vertex1,
						indices, textureCoords, normals,
						textureArray, normalsArray)
					indices = processVertex(vertex2,
						indices, textureCoords, normals,
						textureArray, normalsArray)
					indices = processVertex(vertex3,
						indices, textureCoords, normals,
						textureArray, normalsArray)
				}

			}
		}
	}
	if err = scanner.Err(); err != nil {
		return model, err
	}

	// convert
	verteciesArray := make([]float32, len(vertecies)*3)
	for i, v := range vertecies {
		verteciesArray[i*3] = v.X()
		verteciesArray[i*3+1] = v.Y()
		verteciesArray[i*3+2] = v.Z()
	}

	// fmt.Println(textureArray)
	return LoadToModel(verteciesArray, indices, textureArray), nil

}

func processVertex(
	vertexData []string,
	indices []uint32,
	textures []mgl32.Vec2,
	normals []mgl32.Vec3,
	textureArray []float32,
	normalsArray []float32,
) []uint32 {

	fmt.Println(vertexData)
	// TODO this is a mess, heavy refactoring needed
	index64, err := strconv.ParseUint(vertexData[0], 10, 64)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	index := uint32(index64) - 1
	indices = append(indices, index)
	fmt.Println(indices)

	index64, err = strconv.ParseUint(vertexData[1], 10, 64)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	// obj starts at 1 not 0
	currentTex := textures[index64-1]
	textureArray[index64*2] = currentTex.X()
	// blender renders textures with reversed y axis
	textureArray[(index64*2)+1] = 1 - currentTex.Y()

	index64, err = strconv.ParseUint(vertexData[2], 10, 64)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	currentNorm := normals[index64-1]
	normalsArray[index64*3] = currentNorm.X()
	normalsArray[index64*3+1] = currentNorm.Y()
	normalsArray[index64*3+2] = currentNorm.Z()

	return indices
}
