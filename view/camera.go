package view

import (
	"math"
	"vox/actor"
	"vox/input"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var speed float32 = 0.1

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

func (c *Camera) UpdateMatrix() {
	posX := c.Position[0]
	posY := c.Position[1]
	posZ := c.Position[2]

	c.Rotation[0] = mgl32.Clamp(c.Rotation[0], -90, 90)

	translateMatrix := mgl32.Translate3D(posX, posY, posZ)

	pitch := mgl32.DegToRad(c.Rotation[0])
	yaw := mgl32.DegToRad(c.Rotation[1])
	roll := mgl32.DegToRad(c.Rotation[2])

	rotationQuat := mgl32.AnglesToQuat(yaw, pitch, roll, mgl32.YXZ)
	rotationMatrix := rotationQuat.Mat4()

	modelMatrix := translateMatrix.Mul4(rotationMatrix)

	c.view = modelMatrix.Inv()

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])
}

func (c *Camera) GetForwardYaw() (forward mgl32.Vec3) {
	yaw := float64(mgl32.DegToRad(c.Rotation[1]))

	forwardX := -math.Sin(yaw)
	forwardY := 0
	forwardZ := -math.Cos(yaw)

	forward = mgl32.Vec3{
		float32(forwardX),
		float32(forwardY),
		float32(forwardZ),
	}.Normalize()

	return
}

func (c *Camera) GetRightYaw() (forward mgl32.Vec3) {
	return c.GetForwardYaw().Cross(c.UpVector)
}

func (c *Camera) Update() {
	if input.InputMap[glfw.KeyW] {
		c.Position = c.Position.Add(c.GetForwardYaw().Mul(speed))
	}
	if input.InputMap[glfw.KeyS] {
		c.Position = c.Position.Sub(c.GetForwardYaw().Mul(speed))
	}
	if input.InputMap[glfw.KeyA] {
		c.Position = c.Position.Sub(c.GetRightYaw().Mul(speed))
	}
	if input.InputMap[glfw.KeyD] {
		c.Position = c.Position.Add(c.GetRightYaw().Mul(speed))
	}
	if input.InputMap[glfw.KeySpace] {
		c.Position = c.Position.Add(c.UpVector.Mul(speed))
	}
	if input.InputMap[glfw.KeyLeftControl] {
		c.Position = c.Position.Sub(c.UpVector.Mul(speed))
	}

	sensetive := float32(3)
	c.Rotation[1] = float32(-input.MouseX) / sensetive
	c.Rotation[0] = float32(-input.MouseY) / sensetive

	c.UpdateMatrix()

}
