package trace

type Pt struct {
	X, Y, Z float64
}

func PtAdd(p *Pt, v *V3) *Pt {
	return &Pt{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func PtSub(p *Pt, v *V3) *Pt {
	return &Pt{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func PtScale(p *Pt, f float64) *Pt {
	return &Pt{p.X * f, p.Y * f, p.Z * f}
}

func PtDelta(lhs, rhs *Pt) *V3 {
	return &V3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func (p *Pt) ToV3() *V3 {
	return PtDelta(p, &Pt{})
}
