package renderer

import "math"

type GouraudShader struct {
	light        *Vector
	a, b, c      *Vertex // triangle vertices
	ai, bi, ci   float64 // light intensity at the 3 vertices
	shouldIgnore bool
}

func NewGouraudShader(a, b, c *Vertex, an, bn, cn, light *Vector) *GouraudShader {
	gs := &GouraudShader{light: light, a: a, b: b, c: c}

	gs.resolveLightIntensity(an, bn, cn)

	return gs
}

// ShouldIgnore returns true if this triangle face should not be drawn
func (gs *GouraudShader) ShouldIgnore() bool {
	return gs.shouldIgnore
}

// ShadeFragment returns a color that corresponds to the barycentric coordinates given
func (gs *GouraudShader) ShadeFragment(u, v, w float64, color *RGBA) {
	if gs.shouldIgnore {
		color.Alpha = 255
		color.Blue = 0
		color.Green = 0
		color.Red = 0
		return
	}

	lightIntensity := u*gs.ai + v*gs.bi + w*gs.ci

	color.Alpha = 255
	color.Blue = byte(255 * lightIntensity)
	color.Green = byte(255 * lightIntensity)
	color.Red = byte(255 * lightIntensity)
}

func (gs *GouraudShader) resolveLightIntensity(an, bn, cn *Vector) {

	Normalize(an)
	Normalize(bn)
	Normalize(cn)

	gs.ai = DotProduct(an, gs.light)
	gs.bi = DotProduct(bn, gs.light)
	gs.ci = DotProduct(cn, gs.light)

	if gs.ai < 0 && gs.bi < 0 && gs.ci < 0 {
		gs.shouldIgnore = true
	}

	gs.ai = math.Max(gs.ai, 0)
	gs.bi = math.Max(gs.bi, 0)
	gs.ci = math.Max(gs.ci, 0)
}
