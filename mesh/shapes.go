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

// Еще один метод
func GetSquareIndices() []uint32 {
	return []uint32{
		0, 1, 2,
		2, 3, 0,
	}
}
