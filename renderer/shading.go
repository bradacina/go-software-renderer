package renderer

//ColorIntensity determines the intensity of the color of a triangle given a light direction
func ColorIntensity(a, b, c *Vertex, lightVector *Vector) float64 {
	v0 := VectorFromVertex(a, b)
	v1 := VectorFromVertex(a, c)

	// calculate the normal of the triangle face
	normal := CrossProduct(v0, v1)
	Normalize(normal)

	intensity := DotProduct(normal, lightVector)
	return intensity
}
