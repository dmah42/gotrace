package trace

import (
	"testing"
)

func TestNewV3(t *testing.T) {
	v := NewV3(1.0, 2.0, 3.0)
	if v.x != 1.0 && v.y != 2.0 && v.z != 3.0 {
		t.Errorf("want {1.0, 2.0, 3.0}, got %v", *v)
	}
}

func TestDot(t *testing.T) {
	cases := []struct {
		lhs, rhs *V3
		want     float64
	}{
		{lhs: NewV3(1.0, 2.0, 3.0), rhs: NewV3(4.0, 5.0, 6.0), want: 32.0},
		{lhs: NewV3(2.0, 4.0, 6.0), rhs: NewV3(-1.0, -0.5, -0.3), want: -5.8},
	}

	for _, tt := range cases {
		got := Dot(tt.lhs, tt.rhs)
		if tt.want != got {
			t.Errorf("want %v, got %v\n", tt.want, got)
		}
	}
}

func TestCross(t *testing.T) {
	cases := []struct {
		lhs, rhs, want *V3
	}{
		{lhs: NewV3(1, 0, 0), rhs: NewV3(0, 1, 0), want: NewV3(0, 0, 1)},
	}

	for _, tt := range cases {
		got := Cross(tt.lhs, tt.rhs)
		if *tt.want != *got {
			t.Errorf("want %v, got %v\n", tt.want, got)
		}
	}
}

func TestV3Add(t *testing.T) {
	cases := []struct {
		lhs, rhs *V3
		want     V3
	}{
		{lhs: NewV3(1.0, 2.0, 3.0), rhs: NewV3(4.0, 5.0, 6.0), want: V3{5.0, 7.0, 9.0}},
	}

	for _, tt := range cases {
		gotAdd := V3Add(tt.lhs, tt.rhs)
		if tt.want != *gotAdd {
			t.Errorf("want %v, got %v\n", tt.want, *gotAdd)
		}
	}
}

func TestV3Sub(t *testing.T) {
	cases := []struct {
		lhs, rhs *V3
		want     V3
	}{
		{lhs: NewV3(1.0, 2.0, 3.0), rhs: NewV3(4.0, 5.0, 6.0), want: V3{-3.0, -3.0, -3.0}},
	}

	for _, tt := range cases {
		gotSub := V3Sub(tt.lhs, tt.rhs)
		if tt.want != *gotSub {
			t.Errorf("want %v, got %v\n", tt.want, *gotSub)
		}
	}
}

func TestV3Mul(t *testing.T) {
	cases := []struct {
		v    *V3
		s    float64
		want V3
	}{
		{v: NewV3(1.0, 2.0, 3.0), s: 0.5, want: V3{0.5, 1.0, 1.5}},
	}

	for _, tt := range cases {
		gotMul := V3Mul(tt.v, tt.s)
		if tt.want != *gotMul {
			t.Errorf("want %v, got %v\n", tt.want, *gotMul)
		}
	}
}

func TestLenSqr(t *testing.T) {
	cases := []struct {
		v    *V3
		want float64
	}{
		{v: NewV3(1.0, 2.0, 3.0), want: 14.0},
		{v: NewV3(-0.5, 1.0, 0.0), want: 1.25},
	}

	for _, tt := range cases {
		gotLenSqr := tt.v.LenSqr()
		if tt.want != gotLenSqr {
			t.Errorf("want %v, got %v\n", tt.want, gotLenSqr)
		}
	}
}

func TestNorm(t *testing.T) {
	cases := []struct {
		v    *V3
		want V3
	}{
		{v: NewV3(2.0, 3.0, 6.0), want: V3{2.0 / 7.0, 3.0 / 7.0, 6.0 / 7.0}},
	}

	for _, tt := range cases {
		gotNorm := tt.v.Norm()
		if tt.want != *gotNorm {
			t.Errorf("want %v, got %v\n", tt.want, *gotNorm)
		}
	}
}
