package mesh

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

var VAO uint32
var VBO uint32
var EBO uint32

func BindArrays(vertices []float32, indices []uint32) {

	// Buffer creation
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.GenBuffers(1, &EBO)

	gl.BindVertexArray(VAO)

	// Vertex buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Vertex attributes
	// Position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, nil)
	gl.EnableVertexAttribArray(0)
	// Color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, nil)
	gl.EnableVertexAttribArray(1)
}

func ClearArrays() {
	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteBuffers(1, &EBO)
}
