package renderer

type TextureShader struct {
	texture        *Buffer
	light          *Vector
	a, b, c        *Vertex // triangle vertices
	at, bt, ct     *Point  // texture coordinates for the triangle
	lightIntensity float64
	shouldIgnore   bool
}

func NewTextureShader(a, b, c *Vertex, at, bt, ct *Point, texture *Buffer, light *Vector) *TextureShader {
	ts := &TextureShader{
		texture: texture,
		light:   light,
		a:       a,
		b:       b,
		c:       c,
		at:      at,
		bt:      bt,
		ct:      ct,
	}

	ts.resolveLightIntensity()

	return ts
}

// ShouldIgnore returns true if this triangle face should not be drawn
func (ts *TextureShader) ShouldIgnore() bool {
	return ts.shouldIgnore
}

// ShadeFragment returns a color that corresponds to the barycentric coordinates given
func (ts *TextureShader) ShadeFragment(u, v, w float64, color *RGBA) {
	if ts.shouldIgnore {
		color.Alpha = 0
		color.Blue = 0
		color.Green = 0
		color.Red = 0
		return
	}

	ts.resolveColor(u, v, w, color)
	color.Blue = byte(float64(color.Blue) * ts.lightIntensity)
	color.Green = byte(float64(color.Green) * ts.lightIntensity)
	color.Red = byte(float64(color.Red) * ts.lightIntensity)
}

func (ts *TextureShader) resolveLightIntensity() {
	v0 := VectorFromVertex(ts.a, ts.b)
	v1 := VectorFromVertex(ts.a, ts.c)

	// calculate the normal of the triangle face
	normal := CrossProduct(v0, v1)
	Normalize(normal)

	ts.lightIntensity = DotProduct(normal, ts.light)

	if ts.lightIntensity < 0 {
		ts.shouldIgnore = true
	}
}

func (ts *TextureShader) resolveColor(u, v, w float64, color *RGBA) {

	p := Vector{Vertex{
		X: float64(ts.at.X)*u + float64(ts.bt.X)*v + float64(ts.ct.X)*w,
		Y: float64(ts.at.Y)*u + float64(ts.bt.Y)*v + float64(ts.ct.Y)*w}}

	ts.texture.Read(int(p.X), int(p.Y), color)
}
