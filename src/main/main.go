package main

import "trace"

const (
	imageW = 640
	imageH = 480
)

var (
	sphereColor = trace.Color{0, 1, 0.8}
	triangleColor = trace.Color{0.8, 1, 0}
	lightColor = trace.Color{1, 1, 1}
)

func main() {
	c := trace.NewContext(imageW, imageH)

	o2w := trace.NewM44()
	o2w.Translate(trace.NewV3(-2, 0, -5))
	c.AddPrimitive(trace.NewSphere(o2w, trace.NewSolidColor(sphereColor)))

	o2w = trace.NewM44()
	o2w.Translate(trace.NewV3(0, 0, -2)).Scale(trace.NewV3(0.2, 0.2, 0.2))
	c.AddPrimitive(trace.NewSphere(o2w, trace.NewLight(lightColor)))

	o2w = trace.NewM44()
	o2w.Translate(trace.NewV3(2, 0, -5))
	c.AddPrimitive(trace.NewTriangle(o2w, trace.NewSolidColor(triangleColor)))

	image := trace.Render(c)
	trace.WriteImageToPPM(image, "render")
}
