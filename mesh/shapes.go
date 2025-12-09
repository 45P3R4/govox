package mesh

func GetTriangleVertices() []float32 {
	return []float32{
		// Pos         // Colors
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom left
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom right
		0.5, 0.5, 0.0, 0.0, 0.0, 1.0, // top right
	}
}

func GetTriangleIndices() []uint32 {
	return []uint32{
		0, 1, 2,
	}
}

func GetTriangleTextCoord() []float32 {
	return []float32{
		0.0, 0.0, // bottom left
		1.0, 0.0, // bottom right
		1.0, 1.0, // top right
	}
}

func GetSquareVertices() []float32 {
	return []float32{
		// Pos         // Colors
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom left
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom right
		0.5, 0.5, 0.0, 0.0, 0.0, 1.0, // top right
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, // top left
	}
}

func GetSquareIndices() []uint32 {
	return []uint32{
		0, 1, 2,
		2, 3, 0,
	}
}

func GetCubeVertices() []float32 {
	return []float32{
		// front
		-0.5, -0.5, 0.5, 1.0, 0.0, 0.0, // bottom left front
		0.5, -0.5, 0.5, 0.0, 1.0, 0.0, // bottom right front
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, // top left front
		-0.5, 0.5, 0.5, 1.0, 1.0, 0.0, // top right front

		// back
		-0.5, -0.5, -0.5, 1.0, 0.0, 1.0, // bottom left back
		0.5, -0.5, -0.5, 0.0, 1.0, 1.0, // bottom right back
		0.5, 0.5, -0.5, 0.5, 0.5, 0.5, // top left back
		-0.5, 0.5, -0.5, 1.0, 1.0, 1.0, // top right back
	}
}

func GetCubeIndices() []uint32 {
	return []uint32{
		// front
		0, 1, 2,
		2, 3, 0,

		// back
		4, 5, 6,
		6, 7, 4,

		// top
		3, 2, 6,
		6, 7, 3,

		// bottom
		0, 1, 5,
		5, 4, 0,

		// right
		1, 5, 6,
		6, 2, 1,

		// left
		0, 4, 7,
		7, 3, 0,
	}
}

func GetCubeTextureCoords() []float32 {
	return []float32{
		// front
		0.0, 0.0, // 0
		1.0, 0.0, // 1
		1.0, 1.0, // 2
		0.0, 1.0, // 3

		// back
		1.0, 0.0, // 4
		0.0, 0.0, // 5
		0.0, 1.0, // 6
		1.0, 1.0, // 7

		// top
		0.0, 1.0, // 3
		1.0, 1.0, // 2
		1.0, 0.0, // 6
		0.0, 0.0, // 7

		// bottom
		0.0, 0.0, // 0
		1.0, 0.0, // 1
		1.0, 1.0, // 5
		0.0, 1.0, // 4

		// right
		0.0, 0.0, // 1
		1.0, 0.0, // 5
		1.0, 1.0, // 6
		0.0, 1.0, // 2

		// left
		1.0, 0.0, // 0
		0.0, 0.0, // 4
		0.0, 1.0, // 7
		1.0, 1.0, // 3
	}
}

func GetCubeNormals() []float32 {
	return []float32{
		// front +Z
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,

		// back -Z
		0.0, 0.0, -1.0,
		0.0, 0.0, -1.0,
		0.0, 0.0, -1.0,
		0.0, 0.0, -1.0,

		// top +Y
		0.0, 1.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 1.0, 0.0,

		// bottom -Y
		0.0, -1.0, 0.0,
		0.0, -1.0, 0.0,
		0.0, -1.0, 0.0,
		0.0, -1.0, 0.0,

		// right +X
		1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,

		// left -X
		-1.0, 0.0, 0.0,
		-1.0, 0.0, 0.0,
		-1.0, 0.0, 0.0,
		-1.0, 0.0, 0.0,
	}
}
