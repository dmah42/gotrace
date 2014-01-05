package main

import "trace"

const (
	imageW = 640
	imageH = 480
)

func main() {
	c := trace.NewContext(imageW, imageH)

	o2w := trace.NewM44()
	o2w.Translate(trace.NewV3(-2, 0, -5))/*.scale(trace.NewV3(1.0, 2.0, 1.0))*/
	c.AddPrimitive(trace.NewSphere(o2w))

	o2w = trace.NewM44()
	o2w.Translate(trace.NewV3(0, 0, -5)).Scale(trace.NewV3(1, 2, 1))
	c.AddPrimitive(trace.NewSphere(o2w))

	o2w = trace.NewM44()
	o2w.Translate(trace.NewV3(2, 0, -5))
	c.AddPrimitive(trace.NewTriangle(o2w))

	image := trace.Render(c)
	trace.WriteImageToPPM(image, "render")
}
