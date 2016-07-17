package trace

import (
	"testing"
)

func TestPt(t *testing.T) {
	p := &Pt{1.0, 2.0, 3.0}
	if p.X != 1.0 && p.Y != 2.0 && p.Z != 3.0 {
		t.Errorf("want {1.0, 2.0, 3.0}, got %v", *p)
	}
}

func TestPtAdd(t *testing.T) {
	cases := []struct {
		lhs  *Pt
		rhs  *V3
		want Pt
	}{
		{lhs: &Pt{1.0, 2.0, 3.0}, rhs: &V3{4.0, 5.0, 6.0}, want: Pt{5.0, 7.0, 9.0}},
	}

	for _, tt := range cases {
		gotAdd := PtAdd(tt.lhs, tt.rhs)
		if tt.want != *gotAdd {
			t.Errorf("want %v, got %v\n", tt.want, *gotAdd)
		}
	}
}

func TestPtSub(t *testing.T) {
	cases := []struct {
		lhs  *Pt
		rhs  *V3
		want Pt
	}{
		{lhs: &Pt{1.0, 2.0, 3.0}, rhs: &V3{4.0, 5.0, 6.0}, want: Pt{-3.0, -3.0, -3.0}},
	}

	for _, tt := range cases {
		gotSub := PtSub(tt.lhs, tt.rhs)
		if tt.want != *gotSub {
			t.Errorf("want %v, got %v\n", tt.want, *gotSub)
		}
	}
}

func TestPtScale(t *testing.T) {
	cases := []struct {
		lhs  *Pt
		f    float64
		want Pt
	}{
		{lhs: &Pt{1.0, 2.0, 3.0}, f: 0.5, want: Pt{0.5, 1.0, 1.5}},
	}

	for _, tt := range cases {
		got := PtScale(tt.lhs, tt.f)
		if tt.want != *got {
			t.Errorf("want %v, got %v\n", tt.want, *got)
		}
	}
}

func TestPtDelta(t *testing.T) {
	cases := []struct {
		lhs, rhs *Pt
		want     V3
	}{
		{lhs: &Pt{1.0, 2.0, 3.0}, rhs: &Pt{4.0, 5.0, 6.0}, want: V3{-3.0, -3.0, -3.0}},
	}

	for _, tt := range cases {
		gotDelta := PtDelta(tt.lhs, tt.rhs)
		if tt.want != *gotDelta {
			t.Errorf("want %v, got %v\n", tt.want, *gotDelta)
		}
	}
}
