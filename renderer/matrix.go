package renderer

func Mul3x3(a, b, result *Mat3x3) {
	result.AA = a.AA*b.AA + a.AB*b.BA + a.AC*b.CA
	result.AB = a.AA*b.AB + a.AB*b.BB + a.AC*b.CB
	result.AC = a.AA*b.AC + a.AB*b.BC + a.AC*b.CC

	result.BA = a.BA*b.AA + a.BB*b.BA + a.BC*b.CA
	result.BB = a.BA*b.AB + a.BB*b.BB + a.BC*b.CB
	result.BC = a.BA*b.AC + a.BB*b.BC + a.BC*b.CC

	result.CA = a.CA*b.AA + a.CB*b.BA + a.CC*b.CA
	result.CB = a.CA*b.AB + a.CB*b.BB + a.CC*b.CB
	result.CC = a.CA*b.AC + a.CB*b.BC + a.CC*b.CC
}

func Mul4x4(a, b, result *Mat4x4) {
	result.AA = a.AA*b.AA + a.AB*b.BA + a.AC*b.CA + a.AD*b.DA
	result.AB = a.AA*b.AB + a.AB*b.BB + a.AC*b.CB + a.AD*b.DB
	result.AC = a.AA*b.AC + a.AB*b.BC + a.AC*b.CC + a.AD*b.DC
	result.AD = a.AA*b.AD + a.AB*b.BD + a.AC*b.CD + a.AD*b.DD

	result.BA = a.BA*b.AA + a.BB*b.BA + a.BC*b.CA + a.BD*b.DA
	result.BB = a.BA*b.AB + a.BB*b.BB + a.BC*b.CB + a.BD*b.DB
	result.BC = a.BA*b.AC + a.BB*b.BC + a.BC*b.CC + a.BD*b.DC
	result.BD = a.BA*b.AD + a.BB*b.BD + a.BC*b.CD + a.BD*b.DD

	result.CA = a.CA*b.AA + a.CB*b.BA + a.CC*b.CA + a.CD*b.DA
	result.CB = a.CA*b.AB + a.CB*b.BB + a.CC*b.CB + a.CD*b.DB
	result.CC = a.CA*b.AC + a.CB*b.BC + a.CC*b.CC + a.CD*b.DC
	result.CD = a.CA*b.AD + a.CB*b.BD + a.CC*b.CD + a.CD*b.DD

	result.DA = a.DA*b.AA + a.DB*b.BA + a.DC*b.CA + a.DD*b.DA
	result.DB = a.DA*b.AB + a.DB*b.BB + a.DC*b.CB + a.DD*b.DB
	result.DC = a.DA*b.AC + a.DB*b.BC + a.DC*b.CC + a.DD*b.DC
	result.DD = a.DA*b.AD + a.DB*b.BD + a.DC*b.CD + a.DD*b.DD
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

func Mul3x3WithVector(a *Mat3x3, b, result *Vector) {
	result.X = a.AA*b.X + a.AB*b.Y + a.AC*b.Z
	result.Y = a.BA*b.X + a.BB*b.Y + a.BC*b.Z
	result.Z = a.CA*b.X + a.CB*b.Y + a.CC*b.Z
}
