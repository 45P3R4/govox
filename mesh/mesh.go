package mesh

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vertices []float32
	indices  []uint32
	Scale    [3]float32
	Position [3]float32
	Angle    [3]float32

	location int32
	model    mgl32.Mat4
}

func NewMesh(shaderProgram uint32, vertices []float32, indices []uint32) *Mesh {
	var s Mesh = Mesh{
		vertices: vertices,
		indices:  indices,
	}

	s.location = gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
	s.model = mgl32.Ident4()
	s.Scale = [3]float32{1, 1, 1}

	return &s
}

func (s *Mesh) Update() {
	s.model = mgl32.Ident4()

	s.model = s.model.Mul4(mgl32.Translate3D(s.Position[0], s.Position[1], s.Position[2]))

	s.model = s.model.Mul4(mgl32.HomogRotate3DX(s.Angle[0]))
	s.model = s.model.Mul4(mgl32.HomogRotate3DY(s.Angle[1]))
	s.model = s.model.Mul4(mgl32.HomogRotate3DZ(s.Angle[2]))

	s.model = s.model.Mul4(mgl32.Scale3D(s.Scale[0], s.Scale[1], s.Scale[2]))

	gl.UniformMatrix4fv(s.location, 1, false, &s.model[0])

	BindArrays(s.vertices, s.indices)
}
