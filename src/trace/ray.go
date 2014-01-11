package trace

const infinity = 10000.0

type Ray struct {
	o Pt
	d V3

	t0, t1 float64
}

func NewRay(o *Pt, d *V3) *Ray {
	r := &Ray{*o, *(d.Norm()), 0, infinity}
	return r
}
