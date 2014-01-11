package main

import (
	"math"
	"os"
	"strconv"
	"trace"
)

var (
	sphereColor   = trace.Color{0, 1, 0.8}
	triangleColor = trace.Color{0.8, 1, 0}
	lightColor    = trace.Color{1, 1, 1}
)

func polySphere(radius float64, divs uint32, o2w *trace.M44, material trace.Material) *trace.PolyMesh {
	numVertices := (divs-1)*divs + 2
	P := make([]trace.Pt, numVertices)
	N := make([]trace.V3, numVertices)

	// vertices
	u := -math.Pi / 2.0
	v := -math.Pi
	du := math.Pi / float64(divs)
	dv := 2 * math.Pi / float64(divs)

	P[0] = *trace.NewPt(0, -radius, 0)
	N[0] = *trace.NewV3(0, -radius, 0)

	var i uint32 = 0
	var k uint32 = 1
	for ; i < divs-1; i++ {
		u = u + du
		v = -math.Pi
		var j uint32 = 0
		for ; j < divs; j++ {
			x := radius * math.Cos(u) * math.Cos(v)
			y := radius * math.Sin(u)
			z := radius * math.Cos(u) * math.Sin(v)
			P[k] = *trace.NewPt(x, y, z)
			N[k] = *trace.NewV3(x, y, z)
			v = v + dv
			k++
		}
	}
	P[k] = *trace.NewPt(0, radius, 0)
	N[k] = *trace.NewV3(0, radius, 0)

	// polygons
	npoly := divs * divs
	nvertices := make([]uint32, npoly)
	vertices := make([]uint32, (6+(divs-1)*4)*divs)

	// connectivity lists
	var vid uint32 = 1
	var numV uint32 = 0
	var l uint32 = 0

	k = 0
	for i = 0; i < divs; i++ {
		var j uint32 = 0
		for ; j < divs; j++ {
			if i == 0 {
				nvertices[k] = 3
				vertices[l] = 0
				vertices[l+1] = j + vid
				if j == divs-1 {
					vertices[l+2] = vid
				} else {
					vertices[l+2] = j + vid + 1
				}
			} else if i == divs-1 {
				nvertices[k] = 3
				vertices[l] = j + vid + 1 - divs
				vertices[l+1] = vid + 1
				if j == divs-1 {
					vertices[l+2] = vid + 1 - divs
				} else {
					vertices[l+2] = j + vid + 2 - divs
				}
			} else {
				nvertices[k] = 4
				vertices[l] = j + vid + 1 - divs
				vertices[l+1] = j + vid + 1
				if j == divs-1 {
					vertices[l+2] = vid + 1
					vertices[l+3] = vid + 1 - divs
				} else {
					vertices[l+2] = j + vid + 2
					vertices[l+3] = j + vid + 2 - divs
				}
			}
			l += nvertices[k]
			k++
			numV++
		}
		vid = numV
	}

	return trace.NewPolyMesh(o2w, material, npoly, nvertices, vertices, P, N)
}

func main() {
	imageW, _ := strconv.ParseUint(os.Args[1], 0, 32)
	imageH, _ := strconv.ParseUint(os.Args[2], 0, 32)

	c := trace.NewContext(uint32(imageW), uint32(imageH))

	o2w := trace.NewM44()
	o2w.Translate(trace.NewV3(0, 0, -5))
	c.AddPrimitive(polySphere(2, 10, o2w, trace.NewSolidColor(sphereColor)))

	/*
		o2w := trace.NewM44()
		o2w.Translate(trace.NewV3(-2, 0, -5))
		c.AddPrimitive(trace.NewSphere(o2w, trace.NewSolidColor(sphereColor)))

		o2w = trace.NewM44()
		o2w.Translate(trace.NewV3(0, 0, -2)).Scale(trace.NewV3(0.2, 0.2, 0.2))
		c.AddPrimitive(trace.NewSphere(o2w, trace.NewLight(lightColor)))

		o2w = trace.NewM44()
		o2w.Translate(trace.NewV3(2, 0, -5))
		c.AddPrimitive(trace.NewTriangle(o2w, trace.NewSolidColor(triangleColor)))
	*/
	image := trace.Render(c)
	trace.WriteImageToPPM(image, "render")
}
