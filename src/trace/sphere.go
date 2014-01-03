package trace

import (
	"math"
)

type Sphere struct {
	o2w	M44
	w2o	M44
	r	float64
}

func NewSphere(o2w *M44) *Sphere {
	return &Sphere{*o2w, *(o2w.inverse()), 1.0}
}

func solveQuadratic(a, b, c float64) (bool, float64, float64) {
	disc := b * b - 4 * a * c
	if disc < 0.0 {
		return false, 0.0, 0.0
	}
	var q float64
	if b < 0 {
		q = -0.5 * (b - math.Sqrt(disc))
	} else {
		q = -0.5 * (b + math.Sqrt(disc))
	}
	t0 := q / a
	t1 := c / q
	if t0 > t1 {
		temp := t0
		t0 = t1
		t1 = temp
	}
	return true, t0, t1
}

func (s *Sphere) Intersect(ray *Ray) (bool, float64) {
	// Compute A, B and C coefficients
	r := NewRay(s.w2o.transformPt(&ray.o), s.w2o.rotateV3(&ray.d))
	r.t0 = ray.t0
	r.t1 = ray.t1

	a := Dot(&r.d, &r.d)
	l := PtDelta(&r.o, Origin)
	b := 2 * Dot(&r.d, l)
	c := Dot(l, l) - (s.r*s.r)

	sol, t0, t1 := solveQuadratic(a, b, c)
	if !sol || t1 < 0 {
		return false, 0.0
	}

	// if t0 is less than zero, the intersection point is at t1
	if t0 < 0 {
		return true, t1
	}
	return true, t0
}

func (s *Sphere) isLight() bool {
	return false
}

func (s *Sphere) color() *Color {
	return &Color{0.0, 1.0, 0.8}
}

func (s *Sphere) randomPt() *Pt {
	return Origin
}
