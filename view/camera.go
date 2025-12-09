package view

import (
	"fmt"
	"vox/actor"
	"vox/input"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	actor.Actor
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	UpVector mgl32.Vec3

	shaderProgram uint32
	viewLocation  int32
	view          mgl32.Mat4
}

func NewCamera(shaderProgram uint32) *Camera {
	var c Camera

	c.shaderProgram = shaderProgram
	c.Position = [3]float32{0, 0, -3}
	c.UpVector = [3]float32{0, 1, 0}

	c.view = mgl32.LookAtV(
		c.Position, // Position
		c.Rotation, // LookAt
		c.UpVector, // Up Vector
	)

	c.viewLocation = gl.GetUniformLocation(c.shaderProgram, gl.Str("view\x00"))

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])

	return &c
}

func (c *Camera) MoveCamera(newPosition [3]float32) {
	c.Position = newPosition

	translateMatrix := mgl32.Translate3D(newPosition[0], newPosition[1], newPosition[2])
	c.view = translateMatrix

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])
}

func (c *Camera) UpdateMatrix() {
	translateMatrix := mgl32.Translate3D(c.Position[0], c.Position[1], c.Position[2])

	pitch := mgl32.DegToRad(c.Rotation[0])
	yaw := mgl32.DegToRad(c.Rotation[1])
	roll := mgl32.DegToRad(c.Rotation[2])

	rotationQuat := mgl32.AnglesToQuat(yaw, pitch, roll, mgl32.YXZ)
	rotationMatrix := rotationQuat.Mat4()

	modelMatrix := translateMatrix.Mul4(rotationMatrix)

	c.view = modelMatrix.Inv()

	fmt.Println(c.view)

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])
}

func (camera *Camera) Update() {
	if input.InputMap[glfw.KeyW] {
		camera.Position[2] -= 0.1
	}
	if input.InputMap[glfw.KeyS] {
		camera.Position[2] += 0.1
	}
	if input.InputMap[glfw.KeyA] {
		camera.Position[0] += 0.1
	}
	if input.InputMap[glfw.KeyD] {
		camera.Position[0] -= 0.1
	}
	if input.InputMap[glfw.KeySpace] {
		camera.Position[1] -= 0.1
	}
	if input.InputMap[glfw.KeyLeftControl] {
		camera.Position[1] += 0.1
	}

	sensetive := float32(3)
	camera.Rotation[1] = float32(input.MouseX) / sensetive
	camera.Rotation[0] = float32(input.MouseY) / sensetive

	camera.UpdateMatrix()

}
