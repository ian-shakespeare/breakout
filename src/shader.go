package breakout

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	Id uint32
}

func NewShader(vertexSource string, fragmentSource string) (Shader, error) {
	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return Shader{}, err
	}
	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return Shader{}, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var linkSuccess int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &linkSuccess)

	if linkSuccess == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(infoLog))

		return Shader{}, errors.New(infoLog)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return Shader{Id: program}, nil
}

func (s Shader) Delete() {
	gl.DeleteProgram(s.Id)
}

func (s Shader) Use() {
	gl.UseProgram(s.Id)
}

func (s Shader) SetMatrix4(name string, mat mgl32.Mat4) {
	uniform := gl.GetUniformLocation(s.Id, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(uniform, 1, false, &mat[0])
}

func (s Shader) SetVector3f(name string, vec mgl32.Vec3) {
	uniform := gl.GetUniformLocation(s.Id, gl.Str(name+"\x00"))
	gl.Uniform3f(uniform, vec.X(), vec.Y(), vec.Z())
}

func (s Shader) SetInteger(name string, value int32) {
	uniform := gl.GetUniformLocation(s.Id, gl.Str(name+"\x00"))
	gl.Uniform1i(uniform, value)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var compilationSuccess int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &compilationSuccess)

	if compilationSuccess == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(infoLog))

		return shader, errors.New(infoLog)
	}

	return shader, nil
}
