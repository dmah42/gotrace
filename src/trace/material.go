package trace

type Material interface {
	// TODO reflect, refract, light?
	// TODO per-point color (ie, texture)
	color()	Color
}

type SolidColor struct {
	c	Color
}

func NewSolidColor(c Color) *SolidColor {
	return &SolidColor{c}
}

func (m *SolidColor) color() Color {
	return m.c
}
