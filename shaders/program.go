package shaders

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderBase struct {
	programID        uint32
	vertexShaderID   uint32
	fragmentShaderID uint32
}

func NewShaderProgram(vertexShaderSource, fragmentShaderSource string) (ShaderBase, error) {
	vertexShader, err := compileShaderFromFile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return ShaderBase{}, err
	}

	fragmentShader, err := compileShaderFromFile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return ShaderBase{}, err
	}

	program := gl.CreateProgram()

	shader := ShaderBase{
		programID:        program,
		vertexShaderID:   vertexShader,
		fragmentShaderID: fragmentShader,
	}

	// this should be in the specific shader created
	// shader.bindAttribute(0, "position")

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return ShaderBase{}, fmt.Errorf("failed to link program: %v", log)
	}

	// Detach shaders
	gl.DetachShader(program, vertexShader)
	gl.DetachShader(program, fragmentShader)

	// delete the shaders
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	// thinMatrix calls "bindAttributes" here, an anonymous class that does something

	return shader, nil
}

func (s ShaderBase) Start() {
	gl.UseProgram(s.programID)
}

func (s ShaderBase) Stop() {
	gl.UseProgram(0)
}

func (s ShaderBase) CleanUp() {
	s.Stop()
	// since we slated the shaders for deletion when we attached
	//  them to the program, this will delete them.
	gl.DetachShader(s.programID, s.vertexShaderID)
	gl.DetachShader(s.programID, s.fragmentShaderID)
	gl.DeleteProgram(s.programID)
}

func (s ShaderBase) bindAttribute(attribute uint32, variableName string) {
	// variable name needs to be a *uint8
	variableChar := &[]uint8(variableName)[0]

	gl.BindAttribLocation(s.programID, attribute, variableChar)
}

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
