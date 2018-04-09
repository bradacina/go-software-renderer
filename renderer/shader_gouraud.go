package renderer

type GouraudShader struct {
	light        *Vector
	a, b, c      *Vertex // triangle vertices
	an, bn, cn   *Vertex // normals of the 3 vertices
	shouldIgnore bool
}

func NewGouraudShader(a, b, c, an, bn, cn *Vertex, light *Vector) *GouraudShader {
	gs := &GouraudShader{
		light: light, a: a, b: b, c: c,
		an: an, bn: bn, cn: cn}

	gs.resolveShouldIgnore()

	return gs
}

// ShouldIgnore returns true if this triangle face should not be drawn
func (gs *GouraudShader) ShouldIgnore() bool {
	return gs.shouldIgnore
}

// ShadeFragment returns a color that corresponds to the barycentric coordinates given
func (gs *GouraudShader) ShadeFragment(u, v, w float64, color *RGBA) {
	if gs.shouldIgnore {
		color.Alpha = 0
		color.Blue = 0
		color.Green = 0
		color.Red = 0
		return
	}

	normal := Vector{Vertex{
		X: float64(gs.an.X)*u + float64(gs.bn.X)*v + float64(gs.cn.X)*w,
		Y: float64(gs.an.Y)*u + float64(gs.bn.Y)*v + float64(gs.cn.Y)*w,
		Z: float64(gs.an.Z)*u + float64(gs.bn.Z)*v + float64(gs.cn.Z)*w}}

	lightIntensity := DotProduct(&normal, gs.light)

	color.Alpha = 255
	color.Blue = byte(255 * lightIntensity)
	color.Green = byte(255 * lightIntensity)
	color.Red = byte(255 * lightIntensity)
}

func (gs *GouraudShader) resolveShouldIgnore() {
	v0 := VectorFromVertex(gs.a, gs.b)
	v1 := VectorFromVertex(gs.a, gs.c)

	// calculate the normal of the triangle face
	normal := CrossProduct(v0, v1)
	Normalize(normal)

	lightIntensity := DotProduct(normal, gs.light)

	if lightIntensity < 0 {
		gs.shouldIgnore = true
	}
}
