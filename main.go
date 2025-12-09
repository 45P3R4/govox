package main

import (
	"runtime"
	"vox/input"
	"vox/mesh"
	"vox/shader"
	"vox/view"
	"vox/window"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var vertexShader uint32
var fragmentShader uint32
var shaderProgram uint32

func main() {
	runtime.LockOSThread()

	glfwInit()
	defer glfw.Terminate()

	window := window.CreateWindow(800, 600, "VOX")
	defer window.Destroy()

	glInit()

	input.InitInput()

	shaderProgram = shader.CreateShaderProgram("resources/shaders/vert.glsl", "resources/shaders/frag.glsl")
	gl.UseProgram(shaderProgram)

	view.InitPerspectiveProjetion(shaderProgram)

	camera := view.NewCamera(shaderProgram)
	println(camera)

	m1 := mesh.NewMesh(shaderProgram, mesh.GetCubeVertices(), mesh.GetCubeIndices())
	m2 := mesh.NewMesh(shaderProgram, mesh.GetSquareVertices(), mesh.GetSquareIndices())

	meshPool := mesh.NewMeshPool()

	meshPool.AppendMesh(m1, [3]float32{0, 0, 0})
	meshPool.AppendMesh(m2, [3]float32{2, 0, 2})

	// gl.Enable(gl.CULL_FACE)

	// MAIN LOOP
	for !window.ShouldClose() {

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		// gl.Clear(gl.DEPTH_BUFFER_BIT)

		camera.Update()

		meshPool.Draw(shaderProgram)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	meshPool.ClearBuffers()
	gl.DeleteProgram(shaderProgram)
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(vertexShader)
}
