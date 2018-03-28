package renderer

import "log"

func NewCamera() *Camera {
	return &Camera{}
}

func ProjectionOnCenter(distance float64) *Mat4x4 {
	result := Mat4x4Identity
	result.DC = -1 / distance

	return &result
}

func (c *Camera) LookAt(eye *Vertex, center *Vertex, up *Vector) {
	z := VectorFromVertex(center, eye)
	Normalize(z)
	x := CrossProduct(up, z)
	Normalize(x)
	y := CrossProduct(z, x)
	Normalize(y)

	c.translation.X = -eye.X
	c.translation.Y = -eye.Y
	c.translation.Z = -eye.Z
	c.translation.W = 0

	c.mInv = Mat4x4Identity
	c.mInv.AA = x.X
	c.mInv.AB = x.Y
	c.mInv.AC = x.Z
	c.mInv.BA = y.X
	c.mInv.BB = y.Y
	c.mInv.BC = y.Z
	c.mInv.CA = z.X
	c.mInv.CB = z.Y
	c.mInv.CC = z.Z

	log.Println("Translation", c.translation)
	log.Println("M inv", c.mInv)
}

//ViewPort creates a ViewPort matrix that scales
// object coordinates from [-1,1],[-1,1] (object space)
// to [x,x+w][y,y+h] (viewport space)
func ViewPort(x, y, w, h float64) *Mat4x4 {
	viewPort := Mat4x4Identity
	depthResolution := 255.0
	viewPort.AD = float64(x + w/2.0)
	viewPort.BD = float64(y + w/2.0)
	viewPort.CD = depthResolution / 2.0
	viewPort.AA = float64(w / 2.0)
	viewPort.BB = float64(h / 2.0)
	viewPort.CC = depthResolution / 2.0

	return &viewPort
}

func (c *Camera) DebugVertex(v *AfineVertex) {
	debugTranslate, debugMInv := AfineVertex{}, AfineVertex{}
	TranslateAfineVertex(v, &c.translation, &debugTranslate)

	Mul4x4WithAfineVertex(&c.mInv, &debugTranslate, &debugMInv)

	log.Println(debugTranslate)
	log.Println(debugMInv)
}
