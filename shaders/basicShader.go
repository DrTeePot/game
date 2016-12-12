package shaders

import (
	math "github.com/go-gl/mathgl/mgl32"
)

type BasicShader interface {
	Shader
	LoadTransformationMatrix(math.Mat4)
	LoadProjectionMatrix(math.Mat4)
}

type basicShader struct {
	program program // struct that adds useful shader methods

	// uniform variable locations
	transformationMatrix int32
	projectionMatrix     int32
}

// NewBasicShader creates a Shader using the shaders from file specified
// TODO should the shader files be specified here, or added to this file?
func NewBasicShader(vertexShader, fragmentShader string) (BasicShader, error) {
	// get a shader base so i can use util functions
	program, err := newShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		return nil, err
	}

	// bind attributes
	// TODO this should be in the specific shader created
	program.BindAttribute(0, "position")
	program.BindAttribute(1, "textureCoords")

	// attach and link shaders
	err = program.LinkProgram()
	if err != nil {
		return nil, err
	}

	// get shader uniform locations
	t := program.GetUniformLocation("transformationMatrix")
	p := program.GetUniformLocation("projectionMatrix")

	return basicShader{
		program:              program,
		transformationMatrix: t,
		projectionMatrix:     p,
	}, nil
}

func (s basicShader) Start()  { s.program.Start() }
func (s basicShader) Stop()   { s.program.Stop() }
func (s basicShader) Delete() { s.program.Delete() }

func (s basicShader) LoadTransformationMatrix(matrix math.Mat4) {
	s.program.LoadMatrix(s.transformationMatrix, matrix)
}

func (s basicShader) LoadProjectionMatrix(matrix math.Mat4) {
	s.program.LoadMatrix(s.projectionMatrix, matrix)
}
