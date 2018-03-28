package renderer

func Mul3x3(a, b, result *Mat3x3) {
	panic("Mul3x3 Not Implemented")
}

func Mul4x4(a, b, result *Mat4x4) {
	result.AA = a.AA*b.AA + a.AB*b.BA + a.AC*b.CA + a.AD*b.DA
	result.AB = a.AA*b.BA + a.AB*b.BB + a.AB*b.BC + a.AD*b.BD
	result.AC = a.AA*b.CA + a.AB*b.CB + a.AB*b.CC + a.AD*b.CD
	result.AD = a.AA*b.DA + a.AB*b.DB + a.AB*b.DC + a.AD*b.DD

	result.BA = a.BA*b.AA + a.BB*b.BA + a.BC*b.CA + a.BD*b.DA
	result.BB = a.BA*b.BA + a.BB*b.BB + a.BC*b.BC + a.BD*b.BD
	result.BC = a.BA*b.CA + a.BB*b.CB + a.BC*b.CC + a.BD*b.CD
	result.BD = a.BA*b.DA + a.BB*b.DB + a.BC*b.DC + a.BD*b.DD

	result.CA = a.CA*b.AA + a.CB*b.BA + a.CC*b.CA + a.CD*b.DA
	result.CB = a.CA*b.BA + a.CB*b.BB + a.CC*b.BC + a.CD*b.BD
	result.CC = a.CA*b.CA + a.CB*b.CB + a.CC*b.CC + a.CD*b.CD
	result.CD = a.CA*b.DA + a.CB*b.DB + a.CC*b.DC + a.CD*b.DD

	result.DA = a.DA*b.AA + a.DB*b.BA + a.DC*b.CA + a.DD*b.DA
	result.DB = a.DA*b.BA + a.DB*b.BB + a.DC*b.BC + a.DD*b.BD
	result.DC = a.DA*b.CA + a.DB*b.CB + a.DC*b.CC + a.DD*b.CD
	result.DD = a.DA*b.DA + a.DB*b.DB + a.DC*b.DC + a.DD*b.DD
}

func VertexToAfineVertex(a *Vertex, b *AfineVertex) {
	b.X = a.X
	b.Y = a.Y
	b.Z = a.Z
	b.W = 1
}

func AfineVertexToVertex(a *AfineVertex, b *Vertex) {
	b.X = a.X / a.W
	b.Y = a.Y / a.W
	b.Z = a.Z / a.W
}

func Mul4x4WithAfineVertex(a *Mat4x4, b, result *AfineVertex) {
	result.X = a.AA*b.X + a.AB*b.Y + a.AC*b.Z + a.AD*b.W
	result.Y = a.BA*b.X + a.BB*b.Y + a.BC*b.Z + a.BD*b.W
	result.Z = a.CA*b.X + a.CB*b.Y + a.CC*b.Z + a.CD*b.W
	result.W = a.DA*b.X + a.DB*b.Y + a.DC*b.Z + a.DD*b.W
}

func TranslateAfineVertex(a, b, result *AfineVertex) {
	result.X = a.X + b.X
	result.Y = a.Y + b.Y
	result.Z = a.Z + b.Z
	result.W = 1
}
