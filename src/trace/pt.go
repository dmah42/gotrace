package trace

var Origin = &Pt{0, 0, 0}

type Pt struct {
	x, y, z float64
}

func NewPt(x, y, z float64) *Pt {
	return &Pt{x, y, z}
}

func PtAdd(p *Pt, v *V3) *Pt {
	return NewPt(p.x+v.x, p.y+v.y, p.z+v.z)
}

func PtSub(p *Pt, v *V3) *Pt {
	return NewPt(p.x-v.x, p.y-v.y, p.z-v.z)
}

func PtDelta(lhs, rhs *Pt) *V3 {
	return NewV3(lhs.x-rhs.x, lhs.y-rhs.y, lhs.z-rhs.z)
}
