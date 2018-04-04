package renderer

import "log"

func NewCamera() *Camera {
	return &Camera{}
}

func Ortographic(left, right, bot, top, near, far float64) *Mat4x4 {
	w := right - left
	h := top - bot
	d := far - near

	orto := Mat4x4Identity
	orto.AA = 2 / w
	orto.BB = 2 / h
	orto.CC = -2 / d
	orto.AD = -(right + left) / w
	orto.BD = -(top + bot) / h
	orto.CD = -(far + near) / d

	return &orto
}

//http://rextester.com/DGL86880
func (c *Camera) LookAt(eye *Vertex, center *Vertex, up *Vector) {
	z := VectorFromVertex(center, eye)
	Normalize(z)
	x := CrossProduct(up, z)
	Normalize(x)
	y := CrossProduct(z, x)
	Normalize(y)

	c.ModelView = Mat4x4Identity
	c.ModelView.AA = x.X
	c.ModelView.AB = x.Y
	c.ModelView.AC = x.Z
	c.ModelView.AD = -dotVectorVertex(eye, x)
	c.ModelView.BA = y.X
	c.ModelView.BB = y.Y
	c.ModelView.BC = y.Z
	c.ModelView.BD = -dotVectorVertex(eye, y)
	c.ModelView.CA = z.X
	c.ModelView.CB = z.Y
	c.ModelView.CC = z.Z
	c.ModelView.CD = -dotVectorVertex(eye, z)

	log.Println("ViewModel", c.ModelView)
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

func dotVectorVertex(v1 *Vertex, v2 *Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}
