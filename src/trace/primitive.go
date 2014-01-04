package trace

type Primitive interface {
	// TODO: uv coords
	Intersect(r *Ray) (bool, float64)
	randomPt() *Pt

	material() Material
}
