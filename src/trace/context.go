package trace

type Context struct {
	imgW, imgH	uint32
	background	Color
	camera		Camera
	primitives	[]Primitive

	aspectRatio	float64
}

func NewContext(w, h uint32) *Context {
	c := &Context{imgW: w, imgH: h,
		      background: Color{1, 0, 1},
		      camera: *NewCamera(30.0, V3{0, 0, 5}),
		      primitives: []Primitive{},}
	c.aspectRatio = float64(w) / float64(h)
	return c
}

func (c *Context) AddPrimitive(p Primitive) {
	c.primitives = append(c.primitives, p)
}
