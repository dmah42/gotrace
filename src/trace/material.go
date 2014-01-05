package trace

type Material interface {
	// TODO reflect, refract?
	emissive()		Color
	diffuse(u, v float64)	Color
}

// SolidColor
type SolidColor struct {
	c	Color
}

func NewSolidColor(c Color) *SolidColor {
	return &SolidColor{c}
}

func (m *SolidColor) emissive() Color {
	return colorBlack
}

func (m *SolidColor) diffuse(u, v float64) Color {
	return m.c
}

// Light
type Light struct {
	c	Color
}

func NewLight(c Color) *Light {
	return &Light{c}
}

func (m *Light) emissive() Color {
	return m.c
}

func (m *Light) diffuse(u, v float64) Color {
	return colorBlack
}
