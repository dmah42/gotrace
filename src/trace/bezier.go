package trace

const (
	divs int = 16
)

type Bezier struct {
	o2w M44
	w2o M44
	m   Material
	p   [](*PolyMesh)
}

func NewBezier(o2w *M44, m Material, verts []Pt, patches [][16]uint32) *Bezier {
	b := &Bezier{*o2w, *(o2w.inverse()), m, [](*PolyMesh){}}
	b.generatePolyMeshes(o2w, m, verts, patches)
	return b
}

func (b *Bezier) Intersect(ray *Ray) (hit bool, t, u, v float64) {
	hit = false
	t = ray.t1
	u = 0.0
	v = 0.0
	for i := range b.p {
		phit, pt, pu, pv := b.p[i].Intersect(ray)
		if phit && pt < t {
			hit = phit
			t = pt
			u = pu
			v = pv
		}
	}
	// NOTE: this will skip the entire object if any patch is behind the ray origin
	return hit && t > ray.t0, t, u, v
}

func evalBezierCurve(p []Pt, t float64) *Pt {
	iT := 1.0 - t
	b0 := iT * iT * iT
	b1 := 3 * t * iT * iT
	b2 := 3 * t * t * iT
	b3 := t * t * t

	p0 := PtScale(&p[0], b0)
	p1 := PtScale(&p[1], b1)
	p2 := PtScale(&p[2], b2)
	p3 := PtScale(&p[3], b3)
	return PtAdd(p0, PtAdd(p1, PtAdd(p2, p3.ToV3()).ToV3()).ToV3())
}

func evalBezierPatch(ctrlPts []Pt, u, v float64) *Pt {
	var uCurve [4]Pt
	for i := range uCurve {
		uCurve[i] = *(evalBezierCurve(ctrlPts[4*i:4*(i+1)], u))
	}
	return evalBezierCurve(uCurve[:], v)
}

func (b *Bezier) generatePolyMeshes(o2w *M44, m Material, points []Pt, patches [][16]uint32) {
	P := make([]Pt, (divs+1)*(divs+1))
	nverts := make([]uint32, divs*divs)
	verts := make([]uint32, divs*divs*4)

	controlPoints := make([]Pt, 16)

	for np := range patches {
		for i := 0; i < len(controlPoints); i++ {
			controlPoints[i] = points[patches[np][i]-1]
		}

		var k uint32 = 0
		for j := 0; j <= divs; j++ {
			for i := 0; i <= divs; i++ {
				u := float64(i) / float64(divs)
				v := float64(j) / float64(divs)
				P[k] = *evalBezierPatch(controlPoints, u, v)
				k++
			}
		}

		k = 0
		for j := 0; j < divs; j++ {
			for i := 0; i < divs; i++ {
				nverts[k] = 4
				verts[k*4] = uint32((divs+1)*j + i)
				verts[k*4+1] = uint32((divs+1)*(j+1) + i)
				verts[k*4+2] = uint32((divs+1)*(j+1) + (i + 1))
				verts[k*4+3] = uint32((divs+1)*j + (i + 1))
				k++
			}
		}

		b.p = append(b.p, NewPolyMesh(o2w, m, uint32(len(nverts)), nverts, verts, P, []V3{}))
	}
}

func (b *Bezier) material() Material {
	return b.m
}

func (b *Bezier) randomPt() *Pt {
	// TODO: random point on random polymesh
	return &Pt{}
}
