package renderer

//GrayShade determines the gray shade color of a triangle given a light direction
func GrayShade(a, b, c *Vertex, lightVector *Vector, color *RGBA) {
	v0 := VectorFromVertex(a, b)
	v1 := VectorFromVertex(a, c)

	// calculate the normal of the triangle face
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
