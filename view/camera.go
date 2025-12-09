package view

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position [3]float32
	Rotation [3]float32
	UpVector [3]float32
	View     mgl32.Mat4

	viewLocation int32
	view         mgl32.Mat4
}

func NewCamera(shaderProgram uint32) *Camera {
	var c Camera

	c.Position = [3]float32{0, -3, 0}
	c.UpVector = [3]float32{0, 1, 0}

	c.view = mgl32.LookAtV(
		c.Position, // Position
		c.Rotation, // LookAt
		c.UpVector, // Up Vector
	)

	translateMatrix := mgl32.Translate3D(c.Position[0], c.Position[1], c.Position[2])
	c.view = translateMatrix

	c.View = c.view
	c.viewLocation = gl.GetUniformLocation(shaderProgram, gl.Str("view\x00"))

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.View[0])

	return &c
}

func (c *Camera) MoveCamera(newPosition [3]float32) {
	c.Position = newPosition

	translateMatrix := mgl32.Translate3D(newPosition[0], newPosition[1], newPosition[2])
	c.view = translateMatrix

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])
}

func (c *Camera) RotateCamera(newRotation [3]float32) {
	c.Rotation = newRotation

	c.view = mgl32.LookAtV(
		c.Position,          // Position
		c.Rotation,          // LookAt
		mgl32.Vec3{0, 1, 0}, // Up Vector
	)

	gl.UniformMatrix4fv(c.viewLocation, 1, false, &c.view[0])
}
