package shaders

import (
	"path/filepath"

	math "github.com/go-gl/mathgl/mgl32"
)

const (
	vertexShader   = "../fluorine/render/shaders/vertexShader.glsl"
	fragmentShader = "../fluorine/render/shaders/fragmentShader.glsl"

	transformationMatrix = "transformationMatrix"
	projectionMatrix     = "projectionMatrix"
	viewMatrix           = "viewMatrix"
	lightPosition        = "lightPosition"
	lightColour          = "lightColour"
	shineDamper          = "shineDamper"
	reflectivity         = "reflectivity"
)

// for documentation
type basicShader interface {
	Start()
	Stop()
	Delete()
	LoadTransformationMatrix(math.Mat4)
	LoadProjectionMatrix(math.Mat4)
	LoadViewMatrix(math.Mat4)
	LoadLight(position math.Vec3, colour math.Vec3)
	LoadSpecular(float32, float32)
}

type BasicShader struct {
	program program

	// uniform variable locations
	transformationMatrix int32
	projectionMatrix     int32
	viewMatrix           int32
	lightPosition        int32
	lightColour          int32
	shineDamper          int32
	reflectivity         int32
}

// NewBasicShader creates a Shader using the shaders from file specified
func NewBasicShader() (BasicShader, error) {
	vertexAbsPath, err := filepath.Abs(vertexShader)
	fragmentAbsPath, err := filepath.Abs(fragmentShader)
	program, err := NewShaderProgram(vertexAbsPath, fragmentAbsPath)
	if err != nil {
		return BasicShader{}, err
	}

	// bind attributes
	program.BindAttribute(0, "position")
	program.BindAttribute(1, "textureCoords")
	program.BindAttribute(2, "normal")

	// attach and link shaders
	err = program.LinkProgram()
	if err != nil {
		return BasicShader{}, err
	}

	// get shader uniform locations
	t := program.GetUniformLocation(transformationMatrix)
	p := program.GetUniformLocation(projectionMatrix)
	v := program.GetUniformLocation(viewMatrix)
	lp := program.GetUniformLocation(lightPosition)
	lc := program.GetUniformLocation(lightColour)
	s := program.GetUniformLocation(shineDamper)
	r := program.GetUniformLocation(reflectivity)

	return BasicShader{
		program:              program,
		transformationMatrix: t,
		projectionMatrix:     p,
		viewMatrix:           v,
		lightPosition:        lp,
		lightColour:          lc,
		shineDamper:          s,
		reflectivity:         r,
	}, nil
}

func (s BasicShader) Start()  { s.program.Start() }
func (s BasicShader) Stop()   { s.program.Stop() }
func (s BasicShader) Delete() { s.program.Delete() }

// Load to uniform variables
func (s BasicShader) LoadTransformationMatrix(matrix math.Mat4) {
	LoadMatrix(s.transformationMatrix, matrix)
}
func (s BasicShader) LoadProjectionMatrix(matrix math.Mat4) {
	LoadMatrix(s.projectionMatrix, matrix)
}
func (s BasicShader) LoadViewMatrix(matrix math.Mat4) {
	LoadMatrix(s.viewMatrix, matrix)
}
func (s BasicShader) LoadLight(position, colour math.Vec3) {
	LoadVector(s.lightPosition, position)
	LoadVector(s.lightColour, colour)
}
func (s BasicShader) LoadSpecular(shine float32, ref float32) {
	LoadFloat(s.shineDamper, shine)
	LoadFloat(s.reflectivity, ref)
}
