package trace

type PolyMesh struct {
	o2w	M44
	w2o	M44
	m	Material
	tri	[]uint32
	v	[]Pt
	n	[]V3
}

func NewPolyMesh(o2w *M44, m Material, nv, v []uint32, pts []Pt, norms []V3) *PolyMesh {
	mesh := &PolyMesh{*o2w, *(o2w.inverse()), m, []uint32{}, []Pt{}, []V3{}}
	var nt uint32 = 0
	var faceIndex uint32 = 0
	var maxVertIndex uint32 = 0
	for i := range nv {
		nt = nt + nv[i] - 2
		var j uint32 = 0
		for ; j < nv[i]; j++ {
			if v[faceIndex + j] > maxVertIndex {
				maxVertIndex = v[faceIndex + j]
			}
		}
		faceIndex = faceIndex + nv[i]
	}
	mesh.v = make([]Pt, maxVertIndex + 1)
	for i := range mesh.v {
		mesh.v[i] = *(o2w.transformPt(&(pts[i])))
	}
	mesh.tri = make([]uint32, nt * 3)
	triIndex := 0
	faceIndex = 0
	for i := range nv {
		var j uint32 = 0
		for ; j < nv[i] - 2; j++ {
			mesh.tri[triIndex] = v[faceIndex]
			mesh.tri[triIndex + 1] = v[faceIndex + j + 1]
			mesh.tri[triIndex + 2] = v[faceIndex + j + 2]
			triIndex = triIndex + 3
		}
		faceIndex = faceIndex + nv[i]
	}
	return mesh
}

func (m *PolyMesh) Intersect(ray *Ray) (hit bool, t, u, v float64) {
	// TODO
	return false, 0, 0, 0
}

func (m *PolyMesh) material() Material {
	return m.m
}

func (m *PolyMesh) randomPt() *Pt {
	return m.w2o.transformPt(Origin)
}
