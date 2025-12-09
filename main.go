package main

import (
	"runtime"
	"vox/input"
	"vox/mesh"
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

	vertexShader, fragmentShader = createShaders("shaders/vert.glsl", "shaders/frag.glsl")
	shaderProgram = createShaderProgram()
	gl.UseProgram(shaderProgram)

	view.InitPerspectiveProjetion(shaderProgram)

	camera := view.NewCamera(shaderProgram)
	println(camera)

	shape := mesh.NewMesh(shaderProgram, mesh.GetSquareVertices(), mesh.GetSquareIndices())

	input.InitInput()

	var newY float32 = camera.Position[0]
	var newZ float32 = camera.Position[1]
	var newX float32 = camera.Position[2]

	// MAIN LOOP
	for !window.ShouldClose() {

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		time := glfw.GetTime()

		if input.InputMap[glfw.KeyW] {
			newZ += 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}
		if input.InputMap[glfw.KeyS] {
			newZ -= 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}
		if input.InputMap[glfw.KeyA] {
			newX += 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}
		if input.InputMap[glfw.KeyD] {
			newX -= 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}
		if input.InputMap[glfw.KeySpace] {
			newY -= 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}
		if input.InputMap[glfw.KeyLeftControl] {
			newY += 0.1
			camera.MoveCamera([3]float32{newX, newY, newZ})
		}

		angle := float32(time * 50)
		shape.Angle = [3]float32{0, angle / 40, 0}

		shape.Update()

		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	mesh.ClearArrays()
	gl.DeleteProgram(shaderProgram)
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(vertexShader)
}
