package renderer

type Shader interface {
	ShadeFragment(u, v, w float64, color *RGBA)
	ShouldIgnore() bool
}
