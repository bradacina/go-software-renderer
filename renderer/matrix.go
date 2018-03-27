package renderer

func Mul3x3(a, b, result *Mat3x3) {
	panic("Mul3x3 Not Implemented")
}

func Mul4x4(a, b, result *Mat4x4) {
	panic("Mul3x3 Not Implemented")

}

func VectorToAfineVector(a *Vector, b *AfineVector) {
	b.X = a.X
	b.Y = a.Y
	b.Z = a.Z
	b.W = 1
}

func AfineVectorToVector(a *AfineVector, b *Vector) {
	b.X = a.X / a.W
	b.Y = a.Y / a.W
	b.Z = a.Z / a.W
}

func Mul4x4WithAfineVector(a *Mat4x4, b, result *AfineVector) {
	result.X = a.AA*b.X + a.AB*b.Y + a.AC*b.Z + a.AD*b.W
	result.Y = a.BA*b.X + a.BB*b.Y + a.BC*b.Z + a.BD*b.W
	result.Z = a.CA*b.X + a.CB*b.Y + a.CC*b.Z + a.CD*b.W
	result.W = a.DA*b.X + a.DB*b.Y + a.DC*b.Z + a.DD*b.W
}
