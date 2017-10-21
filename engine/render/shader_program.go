package render

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	math "github.com/go-gl/mathgl/mgl32"
)

type Shader interface {
	Start()
	Stop()
	Delete()
}

type shaderProgram struct {
	programID        uint32
	vertexShaderID   uint32
	fragmentShaderID uint32
}

func newShaderProgram(vertexShaderSource, fragmentShaderSource string) (program, error) {
	vertexShader, err := compileShaderFromFile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return program{}, err
	}

	fragmentShader, err := compileShaderFromFile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return program{}, err
	}

	programID := gl.CreateProgram()

	return program{
		programID:        programID,
		vertexShaderID:   vertexShader,
		fragmentShaderID: fragmentShader,
	}, nil
}

func (s shaderProgram) linkProgram() error {
	gl.AttachShader(s.programID, s.vertexShaderID)
	gl.AttachShader(s.programID, s.fragmentShaderID)
	gl.LinkProgram(s.programID)

	var status int32
	gl.GetProgramiv(s.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(s.programID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.programID, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	// Detach shaders
	gl.DetachShader(s.programID, s.vertexShaderID)
	gl.DetachShader(s.programID, s.fragmentShaderID)

	// delete the shaders
	gl.DeleteShader(s.vertexShaderID)
	gl.DeleteShader(s.fragmentShaderID)

	return nil
}

func (s shaderProgram) Start() {
	gl.UseProgram(s.programID)
}

func (s shaderProgram) Stop() {
	gl.UseProgram(0)
}

func (s shaderProgram) Delete() {
	s.Stop()
	// since we slated the shaders for deletion when we attached
	//  them to the program, this will delete them.
	gl.DetachShader(s.programID, s.vertexShaderID)
	gl.DetachShader(s.programID, s.fragmentShaderID)
	gl.DeleteProgram(s.programID)
}

func (p shaderProgram) GetUniformLocation(uniformName string) int32 {
	variableChar := &[]uint8(uniformName)[0]

	return gl.GetUniformLocation(p.programID, variableChar)
}

func (p shaderProgram) LoadFloat(location int32, value float32) {
	gl.Uniform1f(location, value)
}

func (p shaderProgram) LoadVector(location int32, vector math.Vec3) {
	gl.Uniform3f(location, vector.X(), vector.Y(), vector.Z())
}

func (p shaderProgram) LoadBoolean(location int32, value bool) {
	// go initializes to 0
	var toLoad float32
	if value {
		toLoad = 1
	}
	gl.Uniform1f(location, toLoad)
}

func (p shaderProgram) LoadMatrix(location int32, matrix math.Mat4) {
	// the second parameter is the number of matrices being passed
	// we convert matrix to a *float32
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (s shaderProgram) BindAttribute(attribute uint32, variableName string) {
	// variable name needs to be a *uint8
	variableChar := &[]uint8(variableName)[0]

	gl.BindAttribLocation(s.programID, attribute, variableChar)
}

// compileShaderFromFile takes a filename and a shader type and returns a shader
func compileShaderFromFile(filename string, shaderType uint32) (uint32, error) {
	shaderSource, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	shader, err := compileShader(string(shaderSource)+"\x00", shaderType)
	if err != nil {
		return shader, err
	}
	return shader, nil
}

// compileShader takes in a string of GLSL source code and returns a shader
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
