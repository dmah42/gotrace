package trace

import "log"

type PolyMesh struct {
	o2w  M44
	w2o  M44
	m    Material
	ntri uint32
	tri  []uint32
	p    []Pt
	n    []V3
}

func NewPolyMesh(o2w *M44, m Material, np uint32, nv, v []uint32, pts []Pt, norms []V3) *PolyMesh {
	mesh := &PolyMesh{*o2w, *(o2w.inverse()), m, 0, []uint32{}, []Pt{}, []V3{}}
	var faceIndex uint32 = 0
	var maxVertIndex uint32 = 0
	var i uint32 = 0
	for ; i < np; i++ {
		mesh.ntri = mesh.ntri + nv[i] - 2
		var j uint32 = 0
		for ; j < nv[i]; j++ {
			if v[faceIndex+j] > maxVertIndex {
				maxVertIndex = v[faceIndex+j]
			}
		}
		faceIndex = faceIndex + nv[i]
	}
	mesh.p = make([]Pt, maxVertIndex+1)
	for i := range mesh.p {
		mesh.p[i] = *(o2w.transformPt(&(pts[i])))
		// TODO: bounding box
	}
	log.Printf("Creating polymesh: %d polygons, %d triangles\n", np, mesh.ntri)
	mesh.tri = make([]uint32, mesh.ntri*3)
	faceIndex = 0
	var triIndex uint32 = 0
	for i = 0; i < np; i++ {
		var j uint32 = 0
		for ; j < nv[i]-2; j++ {
			mesh.tri[triIndex] = v[faceIndex]
			mesh.tri[triIndex+1] = v[faceIndex+j+1]
			mesh.tri[triIndex+2] = v[faceIndex+j+2]
			triIndex += 3
		}
		faceIndex += nv[i]
	}
	return mesh
}

func intersectTriangle(r *Ray, v0, v1, v2 *Pt) (hit bool, t, u, v float64) {
	e1 := PtDelta(v1, v0)
	e2 := PtDelta(v2, v0)

	pv := Cross(&r.d, e2)
	det := Dot(e1, pv)
	hit = det != 0
	if !hit {
		return
	}

	tv := PtDelta(&r.o, v0)
	u = Dot(tv, pv) / det
	hit = u >= 0 && u <= 1
	if !hit {
		return
	}

	qv := Cross(tv, e1)
	v = Dot(&r.d, qv) / det
	hit = v >= 0 && u+v <= 1
	if !hit {
		return
	}

	t = Dot(e2, qv) / det
	hit = true
	return
}

func (m *PolyMesh) Intersect(ray *Ray) (hit bool, t, u, v float64) {
	r := NewRay(m.w2o.transformPt(&ray.o), m.w2o.rotateV3(&ray.d))
	r.t0 = ray.t0
	r.t1 = ray.t1

	minT := r.t1
	minU := 0.0
	minV := 0.0

	var i uint32 = 0
	for ; i < m.ntri; i++ {
		triIndex := i * 3
		v0 := m.p[m.tri[triIndex]]
		v1 := m.p[m.tri[triIndex+1]]
		v2 := m.p[m.tri[triIndex+2]]

		hit, t, u, v := intersectTriangle(r, &v0, &v1, &v2)
		// NOTE: this will get triangles behind the start of the ray. Should check t vs r.t0
		if hit && t < minT {
			minT = t
			minU = u
			minV = v
		}
	}

	// NOTE: this might discard a triangle that was behind the ray but not include triangles that should have been hit in front
	return minT < r.t1 && minT > r.t0, minT, minU, minV
}

func (m *PolyMesh) material() Material {
	return m.m
}

func (m *PolyMesh) randomPt() *Pt {
	return m.w2o.transformPt(Origin)
}
