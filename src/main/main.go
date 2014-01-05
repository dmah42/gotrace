package main

import "trace"

const (
	imageW = 640
	imageH = 480
)

func createRandomImage() [][]trace.Color {
	image := make([][]trace.Color, imageH)
	for i := range image {
		image[i] = make([]trace.Color, imageW)
		for j := range image[i] {
			image[i][j].R = float64(i) / float64(imageH)
			image[i][j].G = float64(j) / float64(imageW)
			image[i][j].B = 0.0

			//log.Printf("%d %d %+v | ", i, j, image[i][j])
		}
		//log.Printf("\n")
	}
	return image
}

func main() {
	c := trace.NewContext(imageW, imageH)

	o2w := trace.NewM44()
	o2w.Translate(trace.NewV3(0, 0, -5))/*.scale(trace.NewV3(1.0, 2.0, 1.0))*/
	//c.AddPrimitive(trace.NewSphere(o2w))
	c.AddPrimitive(trace.NewTriangle(o2w))

	image := trace.Render(c)
	//image := createRandomImage()
	trace.WriteImageToPPM(image, "render")
}
