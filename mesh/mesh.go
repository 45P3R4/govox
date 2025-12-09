package mesh

import (
	"vox/actor"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	actor.Actor
	vertices []float32
	indices  []uint32

	location int32
	model    mgl32.Mat4
}

func NewMesh(shaderProgram uint32, vertices []float32, indices []uint32) Mesh {
	var s Mesh = Mesh{
		vertices: vertices,
		indices:  indices,
	}

	s.location = gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
	s.model = mgl32.Ident4()
	s.Scale = [3]float32{1, 1, 1}

	return s
}

func (m *Mesh) UpdateModelMatrix() {
	m.model = mgl32.Translate3D(m.Position[0], m.Position[1], m.Position[2])
}
