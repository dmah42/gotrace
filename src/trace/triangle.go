package trace

import (
	"math"
)

type Triangle struct {
	o2w	M44
	w2o	M44
	m	Material
	v	[3]V3
}

func NewTriangle(o2w *M44) *Triangle {
	return &Triangle{*o2w, *(o2w.inverse()), NewSolidColor(Color{0.0, 1.0, 0.8}),
			 [3]V3{{-1, -1, 0}, {1, -1, 0}, {0, 1, 0},}}
}

func (tri *Triangle) Intersect(ray *Ray) (bool, float64) {
	r := NewRay(tri.w2o.transformPt(&ray.o), tri.w2o.rotateV3(&ray.d))
	r.t0 = ray.t0
	r.t1 = ray.t1

	e1 := V3Sub(&tri.v[1], &tri.v[0])
	e2 := V3Sub(&tri.v[2], &tri.v[0])
	vp := Cross(&r.d, e2)

	det := Dot(e1, vp)
	if det == 0 {
		return false, 0
	}

	vt := V3Sub(PtDelta(&r.o, Origin), &tri.v[0])
	u := Dot(vt, vp) / det
	if u < 0 || u > 1 {
		return false, 0
	}

	vq := Cross(vt, e1)
	v := Dot(&r.d, vq) / det

	if v < 0 || v > 1 {
		return false, 0
	}

	t := Dot(e2, vq) / det
	return t > r.t0 && t < r.t1, t
}

func (t *Triangle) isLight() bool {
	return false
}

func (t *Triangle) material() Material {
	return t.m
}

func (t *Triangle) randomPt() *Pt {
	return Origin
}
