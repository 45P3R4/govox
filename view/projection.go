package view

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var AspectRatio float32 = 800.0 / 600.0

var perspective mgl32.Mat4 = mgl32.Perspective(
	45.0, // FOV
	AspectRatio,
	0.1,    // Near plane
	1000.0, // Far plane
)

var ortho mgl32.Mat4 = mgl32.Ortho(-1, 1, -1, 1, 0.1, 100)

func InitPerspectiveProjetion(shaderProgram uint32) {
	projLocation := gl.GetUniformLocation(shaderProgram, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projLocation, 1, false, &perspective[0])
}

func InitOrthoProjetion(shaderProgram uint32) {
	projLocation := gl.GetUniformLocation(shaderProgram, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projLocation, 1, false, &ortho[0])
}
