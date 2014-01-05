package trace

type Primitive interface {
	Intersect(r *Ray) (hit bool, t, u, v float64)
	randomPt() *Pt

	material() Material
}
