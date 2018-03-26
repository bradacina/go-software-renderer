package renderer

//ColorIntensity determines the intensity of the color of a triangle given a light direction
func ColorIntensity(a, b, c *Vertex, lightVector *Vector, color *RGBA) {
	v0 := VectorFromVertex(a, b)
	v1 := VectorFromVertex(a, c)

	// calculate the normal of the triangle face
	normal := CrossProduct(v0, v1)
	Normalize(normal)

	intensity := DotProduct(normal, lightVector)

	if intensity < 0 {
		color.Alpha = 0
	} else {
		color.Blue = byte(float64(color.Blue) * intensity)
		color.Green = byte(float64(color.Green) * intensity)
		color.Red = byte(float64(color.Red) * intensity)
	}
}
