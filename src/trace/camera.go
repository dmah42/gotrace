package trace

import "math"

type Camera struct {
	c2w	M44
	w2c	M44
	focal	float64
	angle	float64
}

func NewCamera(focal float64, position V3) *Camera {
	c := &Camera{}
	c.c2w = *NewM44()
	c.translate(&position)
	c.setFocal(focal)
	return c
}

func (c *Camera) translate(v *V3) {
	c.c2w.Translate(v)
	c.w2c = *c.c2w.inverse()
}

func (c *Camera) setFocal(f float64) {
	c.focal = f
	c.angle = math.Tan(f * 0.5 * math.Pi / 180.0)
}
