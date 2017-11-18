package shaders

import (
	"github.com/DrTeePot/game/light"

	math "github.com/go-gl/mathgl/mgl32"
)

type BasicShader interface {
	Shader
	LoadTransformationMatrix(math.Mat4)
	LoadProjectionMatrix(math.Mat4)
	LoadViewMatrix(math.Mat4)
	// next three -> 1
	LoadLightPosition(math.Vec3)
	LoadLightColour(math.Vec3)
	LoadLight(light.Light) // light should be an entity, or a component
	LoadSpecular(float32, float32)
}

type basicShader struct {
	// TODO use better composition and helpers?
	// s.program all the time is kinda weird
	program program // struct that adds useful shader methods

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
	program.BindAttribute(2, "normal")

	// attach and link shaders
	err = program.LinkProgram()
	if err != nil {
		return nil, err
	}

	// get shader uniform locations
	t := program.GetUniformLocation("transformationMatrix")
	p := program.GetUniformLocation("projectionMatrix")
	v := program.GetUniformLocation("viewMatrix")
	lp := program.GetUniformLocation("lightPosition")
	lc := program.GetUniformLocation("lightColour")
	s := program.GetUniformLocation("shineDamper")
	r := program.GetUniformLocation("reflectivity")

	return basicShader{
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

func (s basicShader) Start()  { s.program.Start() }
func (s basicShader) Stop()   { s.program.Stop() }
func (s basicShader) Delete() { s.program.Delete() }

// Load to uniform variables
func (s basicShader) LoadTransformationMatrix(matrix math.Mat4) {
	s.program.LoadMatrix(s.transformationMatrix, matrix)
}
func (s basicShader) LoadProjectionMatrix(matrix math.Mat4) {
	s.program.LoadMatrix(s.projectionMatrix, matrix)
}
func (s basicShader) LoadViewMatrix(matrix math.Mat4) {
	s.program.LoadMatrix(s.viewMatrix, matrix)
}
func (s basicShader) LoadLightPosition(vec math.Vec3) {
	s.program.LoadVector(s.lightPosition, vec)
}
func (s basicShader) LoadLightColour(vec math.Vec3) {
	s.program.LoadVector(s.lightColour, vec)
}
func (s basicShader) LoadLight(l light.Light) {
	s.LoadLightPosition(l.Position())
	s.LoadLightColour(l.Colour())
}
func (s basicShader) LoadSpecular(shine float32, ref float32) {
	s.program.LoadFloat(s.shineDamper, shine)
	s.program.LoadFloat(s.reflectivity, ref)
}
