package trace

import "math"

type V3 struct {
	x, y, z float64
}

func NewV3(x, y, z float64) *V3 {
	return &V3{x, y, z}
}

func Dot(lhs, rhs *V3) float64 {
	return lhs.x*rhs.x + lhs.y*rhs.y + lhs.z*rhs.z
}

func Cross(lhs, rhs *V3) *V3 {
	x := lhs.y*rhs.z - lhs.z*rhs.y
	y := lhs.z*rhs.x - lhs.x*rhs.z
	z := rhs.x*rhs.y - lhs.y*rhs.x
	return NewV3(x, y, z)
}

func V3Add(lhs, rhs *V3) *V3 {
	return NewV3(lhs.x+rhs.x, lhs.y+rhs.y, lhs.z+rhs.z)
}

func V3Sub(lhs, rhs *V3) *V3 {
	return NewV3(lhs.x-rhs.x, lhs.y-rhs.y, lhs.z-rhs.z)
}

func V3Mul(v *V3, s float64) *V3 {
	return NewV3(v.x*s, v.y*s, v.z*s)
}

func (v *V3) LenSqr() float64 {
	return Dot(v, v)
}

func (v *V3) Len() float64 {
	return math.Sqrt(v.LenSqr())
}

func (v *V3) Norm() *V3 {
	l := v.Len()
	if l > 0.0 {
		return NewV3(v.x/l, v.y/l, v.z/l)
	}
	return NewV3(v.x, v.y, v.z)
}

