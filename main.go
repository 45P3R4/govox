package main

import (
	"math/rand"
	"runtime"
	"vox/input"
	"vox/mesh"
	"vox/shader"
	"vox/view"
	"vox/window"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var shaderProgram uint32

func main() {
	runtime.LockOSThread()

	glfwInit()
	defer glfw.Terminate()

	window := window.CreateWindow(800, 600, "govox")
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

	for i := 0; i < 1000; i++ {
		mesh := mesh.NewMesh(shaderProgram, mesh.GetCubeVertices(), mesh.GetCubeIndices())

		randX := float32(rand.Intn(101) - 50)
		randY := float32(rand.Intn(101) - 50)
		randZ := float32(rand.Intn(101) - 50)
		meshPool.AppendMesh(mesh, [3]float32{randX, randY, randZ})
	}

	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_TEST)

	// MAIN LOOP
	for !window.ShouldClose() {

		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)

		camera.Update()

		meshPool.UpdateMeshes()
		meshPool.Draw(shaderProgram)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	meshPool.ClearBuffers()
	gl.DeleteProgram(shaderProgram)
	shader.ClearShaders()
}
