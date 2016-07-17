package trace

import "math"

type V3 struct {
	X, Y, Z float64
}

func Dot(lhs, rhs *V3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func Cross(lhs, rhs *V3) *V3 {
	x := lhs.Y*rhs.Z - lhs.Z*rhs.Y
	y := lhs.Z*rhs.X - lhs.X*rhs.Z
	z := lhs.X*rhs.Y - lhs.Y*rhs.X
	return &V3{x, y, z}
}

func V3Add(lhs, rhs *V3) *V3 {
	return &V3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

func V3Sub(lhs, rhs *V3) *V3 {
	return &V3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func V3Mul(v *V3, s float64) *V3 {
	return &V3{v.X * s, v.Y * s, v.Z * s}
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
		return &V3{v.X / l, v.Y / l, v.Z / l}
	}
	return &V3{v.X, v.Y, v.Z}
}
