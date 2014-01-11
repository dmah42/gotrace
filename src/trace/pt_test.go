package trace

import (
	"testing"
)

func TestNewPt(t *testing.T) {
	p := NewPt(1.0, 2.0, 3.0)
	if p.x != 1.0 && p.y != 2.0 && p.z != 3.0 {
		t.Errorf("want {1.0, 2.0, 3.0}, got %v", *p)
	}
}

func TestPtAdd(t *testing.T) {
	cases := []struct {
		lhs  *Pt
		rhs  *V3
		want Pt
	}{
		{lhs: NewPt(1.0, 2.0, 3.0), rhs: NewV3(4.0, 5.0, 6.0), want: Pt{5.0, 7.0, 9.0}},
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
		{lhs: NewPt(1.0, 2.0, 3.0), rhs: NewV3(4.0, 5.0, 6.0), want: Pt{-3.0, -3.0, -3.0}},
	}

	for _, tt := range cases {
		gotSub := PtSub(tt.lhs, tt.rhs)
		if tt.want != *gotSub {
			t.Errorf("want %v, got %v\n", tt.want, *gotSub)
		}
	}
}

func TestPtDelta(t *testing.T) {
	cases := []struct {
		lhs, rhs *Pt
		want     V3
	}{
		{lhs: NewPt(1.0, 2.0, 3.0), rhs: NewPt(4.0, 5.0, 6.0), want: V3{-3.0, -3.0, -3.0}},
	}

	for _, tt := range cases {
		gotDelta := PtDelta(tt.lhs, tt.rhs)
		if tt.want != *gotDelta {
			t.Errorf("want %v, got %v\n", tt.want, *gotDelta)
		}
	}
}
