package trace

import (
	"math"
)

type Sphere struct {
	o2w	M44
	w2o	M44
	r	float64
	m	Material
}

func NewSphere(o2w *M44, m Material) *Sphere {
	return &Sphere{*o2w, *(o2w.inverse()), 1.0, m}
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

func (s *Sphere) Intersect(ray *Ray) (hit bool, t, u, v float64) {
	// Compute A, B and C coefficients
	r := NewRay(s.w2o.transformPt(&ray.o), s.w2o.rotateV3(&ray.d))
	r.t0 = ray.t0
	r.t1 = ray.t1

	a := Dot(&r.d, &r.d)
	l := PtDelta(&r.o, Origin)
	b := 2 * Dot(&r.d, l)
	c := Dot(l, l) - (s.r*s.r)

	sol, t0, t1 := solveQuadratic(a, b, c)
	hit = sol && t1 >= 0
	t = t0
	u = 0
	v = 0

	if !hit {
		return
	}

	// if t0 is less than zero, the intersection point is at t1
	if t < 0 {
		t = t1
	}

	iPt := PtAdd(&r.o, V3Mul(&r.d, t))
	theta := math.Acos(-iPt.y)
	phi := math.Atan2(iPt.z, -iPt.x)

	for phi < 0 {
		phi = phi + 2.0 * math.Pi;
	}

	u = phi / (2.0 * math.Pi)
	v = (math.Pi - theta) / math.Pi
	return
}

func (s *Sphere) material() Material {
	return s.m
}

func (s *Sphere) randomPt() *Pt {
	// TODO: pick a point on the surface
	return s.w2o.transformPt(Origin)
}
