package trace

type Primitive interface {
	// TODO: uv coords
	Intersect(r *Ray) (hit bool, t, u, v float64)
	randomPt() *Pt

	material() Material
}
