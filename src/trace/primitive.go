package trace

type Primitive interface {
	Intersect(r *Ray) (bool, float64)
	randomPt() *Pt

	material() Material
}
