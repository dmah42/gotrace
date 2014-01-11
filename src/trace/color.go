package trace

type Color struct {
	R, G, B float64
}

var (
	colorBlack = Color{0, 0, 0}
	colorWhite = Color{1, 1, 1}
)

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func (c *Color) Add(rhs *Color) {
	c.R = c.R + rhs.R
	c.G = c.G + rhs.G
	c.B = c.B + rhs.B
}

func (c *Color) Mul(rhs *Color) *Color {
	return NewColor(c.R*rhs.R, c.G*rhs.G, c.B*rhs.B)
}

func (c *Color) Scale(s float64) {
	c.R = c.R * s
	c.G = c.G * s
	c.B = c.B * s
}
