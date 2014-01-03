package trace

type Primitive interface {
	Intersect(r *Ray) (bool, float64)
	randomPt() *Pt

	// TODO: material
	isLight() bool
	// TODO: pass point in
	color() *Color
}
