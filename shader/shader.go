package shader

import (
	"fmt"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var fragmentShader uint32
var vertexShader uint32

func createShaders(vertex_shader_path string, fragment_shader_path string) (uint32, uint32) {
	vertexShaderSource, err := os.ReadFile(vertex_shader_path)
	if err != nil {
		fmt.Println(err)
	}
	fragmentShaderSource, err := os.ReadFile(fragment_shader_path)
	if err != nil {
		fmt.Println(err)
	}

	vertexShader, err := compileShader(string(vertexShaderSource), gl.VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
	}
	fragmentShader, err := compileShader(string(fragmentShaderSource), gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
	}

	return vertexShader, fragmentShader
}

func CreateShaderProgram(vertex_shader_path string, fragment_shader_path string) uint32 {
	program := gl.CreateProgram()
	vertexShader, fragmentShader = createShaders("resources/shaders/vert.glsl", "resources/shaders/frag.glsl")
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	return program
}

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
		logStr := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &logStr[0])
		return 0, &glfw.Error{}
	}

	return shader, nil
}

func ClearShaders() {
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(vertexShader)
}
