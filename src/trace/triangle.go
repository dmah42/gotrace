package trace

type Triangle struct {
	o2w M44
	w2o M44
	m   Material
	v   [3]V3
}

var (
	verts = [3]V3{{-1, -1, 0}, {1, -1, 0}, {0, 1, 0}}
)

func NewTriangle(o2w *M44, m Material) *Triangle {
	return &Triangle{*o2w, *(o2w.inverse()), m, verts}
}

func (tri *Triangle) Intersect(ray *Ray) (hit bool, t, u, v float64) {
	r := NewRay(tri.w2o.transformPt(&ray.o), tri.w2o.rotateV3(&ray.d))
	r.t0 = ray.t0
	r.t1 = ray.t1

	e1 := V3Sub(&tri.v[1], &tri.v[0])
	e2 := V3Sub(&tri.v[2], &tri.v[0])
	vp := Cross(&r.d, e2)

	det := Dot(e1, vp)

	hit = det != 0
	t = 0
	u = 0
	v = 0

	if !hit {
		return
	}

	vt := V3Sub(PtDelta(&r.o, &Pt{}), &tri.v[0])
	u = Dot(vt, vp) / det

	hit = u >= 0 && u <= 1
	if !hit {
		return
	}

	vq := Cross(vt, e1)
	v = Dot(&r.d, vq) / det

	hit = v >= 0 && (u+v) <= 1
	if !hit {
		return
	}

	t = Dot(e2, vq) / det
	hit = t > r.t0 && t < r.t1
	return
}

func (t *Triangle) material() Material {
	return t.m
}

func (t *Triangle) randomPt() *Pt {
	return &Pt{}
}
