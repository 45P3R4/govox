package mesh

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type MeshPool struct {
	meshes []Mesh
	VBO    uint32
	EBO    uint32
	VAO    uint32

	vertexData []float32
	indexData  []uint32

	totalIndices int32
	isUpdated    bool
}

func (mp *MeshPool) ClearBuffers() {
	gl.DeleteBuffers(1, &mp.VBO)
	gl.DeleteBuffers(1, &mp.EBO)
	gl.DeleteVertexArrays(1, &mp.VAO)
}

func NewMeshPool() *MeshPool {
	mp := &MeshPool{}
	gl.GenVertexArrays(1, &mp.VAO)
	gl.GenBuffers(1, &mp.VBO)
	gl.GenBuffers(1, &mp.EBO)

	return mp
}

func (mp *MeshPool) AppendMesh(mesh Mesh, position [3]float32) {
	mesh.Position = position
	mp.meshes = append(mp.meshes, mesh)
	mp.isUpdated = true
}

func (mp *MeshPool) UpdateBuffers() {
	if !mp.isUpdated && len(mp.vertexData) > 0 {
		return
	}

	mp.vertexData = nil
	mp.indexData = nil
	mp.totalIndices = 0

	var vertexOffset uint32 = 0

	for _, mesh := range mp.meshes {
		mp.vertexData = append(mp.vertexData, mesh.vertices...)

		for _, idx := range mesh.indices {
			mp.indexData = append(mp.indexData, idx+vertexOffset)
			mp.totalIndices++
		}

		vertexOffset += uint32(len(mesh.vertices) / 6)
	}

	mp.bindBuffers()

	mp.isUpdated = false
}

func (mp *MeshPool) Draw(shaderProgram uint32) {
	mp.UpdateBuffers()

	if mp.totalIndices == 0 {
		return
	}

	gl.BindVertexArray(mp.VAO)

	var currentIndex int32 = 0
	for i := range mp.meshes {
		mesh := &mp.meshes[i]

		if mesh.location == 0 {
			mesh.location = gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
		}

		// Устанавливаем матрицу модели для этого меша
		gl.UniformMatrix4fv(mesh.location, 1, false, &mesh.model[0])

		// Вычисляем количество индексов для этого меша
		indicesCount := int32(len(mesh.indices))

		// Рисуем только индексы этого меша
		gl.DrawElementsBaseVertex(
			gl.TRIANGLES,
			indicesCount,
			gl.UNSIGNED_INT,
			gl.PtrOffset(int(currentIndex*4)), // Смещение в байтах
			0,
		)

		currentIndex += indicesCount
	}

	gl.BindVertexArray(0)
}

func (mp *MeshPool) bindBuffers() {
	gl.BindVertexArray(mp.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, mp.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(mp.vertexData)*4, gl.Ptr(mp.vertexData), gl.STATIC_DRAW)

	// Position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, nil)
	gl.EnableVertexAttribArray(0)
	// Color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, mp.EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mp.indexData)*4, gl.Ptr(mp.indexData), gl.STATIC_DRAW)

	gl.BindVertexArray(0)
}

func (mp *MeshPool) UpdateMeshes() {
	for i := range mp.meshes {

		// r := float64(i + 1)
		// time := glfw.GetTime() / 100000
		// mp.meshes[i].Position[0] += float32(math.Sin(time * r))
		// mp.meshes[i].Position[1] += float32(math.Cos(time * r))
		// mp.meshes[i].Position[2] += float32(math.Cos(time * r))
		mp.meshes[i].UpdateModelMatrix()
	}
	mp.isUpdated = true
}
