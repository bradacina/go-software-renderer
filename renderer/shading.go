package renderer

func GrayShade(a, b, c *Vertex, lightVector *Vector, color *ARGB) {
	v0 := VectorFromVertex(a, b)
	v1 := VectorFromVertex(a, c)

	normal := CrossProduct(v0, v1)
	Normalize(normal)

	intensity := DotProduct(normal, lightVector)

	if intensity < 0 {
		color.Alpha = 0
	} else {
		color.Blue = byte(255 * intensity)
		color.Green = byte(255 * intensity)
		color.Red = byte(255 * intensity)
	}
}
