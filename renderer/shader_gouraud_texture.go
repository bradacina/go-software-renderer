package renderer

import (
	"math"
)

type GouraudTextureShader struct {
	light        *Vector
	a, b, c      *Vertex // triangle vertices
	ai, bi, ci   float64 // light intensity at the 3 vertices
	at, bt, ct   *Point  // texture coordinates at the 3 vertice
	texture      *Buffer
	normalMap    *Buffer // texture containing normal map (if present)
	shouldIgnore bool
}

func NewGouraudTextureShader(
	a, b, c *Vertex,
	an, bn, cn, light *Vector,
	at, bt, ct *Point,
	texture *Buffer,
	normalMap *Buffer) *GouraudTextureShader {

	gs := &GouraudTextureShader{
		light: light,
		a:     a, b: b, c: c,
		at: at, bt: bt, ct: ct,
		texture:   texture,
		normalMap: normalMap}

	gs.resolveLightIntensity(an, bn, cn)

	return gs
}

// ShouldIgnore returns true if this triangle face should not be drawn
func (gs *GouraudTextureShader) ShouldIgnore() bool {
	return gs.shouldIgnore
}

// ShadeFragment returns a color that corresponds to the barycentric coordinates given
func (gs *GouraudTextureShader) ShadeFragment(u, v, w float64, color *RGBA) {
	if gs.shouldIgnore {
		color.Alpha = 255
		color.Blue = 0
		color.Green = 0
		color.Red = 0
		return
	}

	// texture coords
	col := int(u*float64(gs.at.X) + v*float64(gs.bt.X) + w*float64(gs.ct.X))
	row := int(u*float64(gs.at.Y) + v*float64(gs.bt.Y) + w*float64(gs.ct.Y))

	var lightIntensity float64

	// light
	if gs.normalMap != nil {
		gs.normalMap.Read(col, row, color)
		normal := Vector{Vertex{
			X: gs.colorChannelToCoord(color.Red),
			Y: gs.colorChannelToCoord(color.Green),
			Z: gs.colorChannelToCoord(color.Blue)}}

		Normalize(&normal)

		// todo: apply transform matrix to normal

		lightIntensity = DotProduct(&normal, gs.light)

	} else {

		lightIntensity = u*gs.ai + v*gs.bi + w*gs.ci
	}

	if lightIntensity < 0 {
		color.Alpha = 255
		color.Blue = 0
		color.Green = 0
		color.Red = 0
		return
	}

	gs.texture.Read(col, row, color)

	color.Blue = byte(float64(color.Blue) * lightIntensity)
	color.Green = byte(float64(color.Green) * lightIntensity)
	color.Red = byte(float64(color.Red) * lightIntensity)
}

func (gs *GouraudTextureShader) resolveLightIntensity(an, bn, cn *Vector) {

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

func (gs *GouraudTextureShader) colorChannelToCoord(channel byte) float64 {
	return (float64(channel) - 128.0) / 128.0
}
