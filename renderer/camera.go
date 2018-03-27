package renderer

func ProjectionOnCenter(distance float64) *Mat4x4 {
	result := Mat4x4Identity
	result.DC = -1 / distance

	return &result
}
