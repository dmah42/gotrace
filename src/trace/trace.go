package trace

import (
	"fmt"
	"log"
	"time"
)

type stats struct {
	primaryRay, rayObjectTest, rayObjectIsect	uint64
}

const maxDepth = 3
var (
	background = NewColor(1, 0, 1)
	renderStats = stats{0, 0, 0}
	renderTime time.Duration
)

// x and y are coordinates on the image plane
func primaryRay(x, y float64) *Ray {
	imgPt := NewPt(x, y, 10)
	eyePt := Origin

	return NewRay(eyePt, PtDelta(imgPt, eyePt))
}

func trace(r *Ray, prims []Primitive, d uint32) Color {
	if d == maxDepth {
		fmt.Printf("  hit maxdepth\n")
		return Color{1.0, 1.0, 1.0}
	}

	minT := r.t1
	minU := 0.0
	minV := 0.0
	hitInx := -1
	for i := range prims {
		renderStats.rayObjectTest++
		hit, t, u, v := prims[i].Intersect(r)
		if hit && t < minT && t > r.t0 {
			renderStats.rayObjectIsect++
			minT = t
			minU = u
			minV = v
			hitInx = i
		}
	}

	if hitInx == -1 {
		return *background
	}

	// TODO: reflection/refraction
	// TODO: lights

	c := prims[hitInx].material().diffuse(minU, minV)
	// fmt.Printf("  %v\n", c)
	return c
}

func Render(ctx *Context) Image {
	log.Println("Rendering")
	image := makeImage(ctx.imgW, ctx.imgH)

	startTime := time.Now()
	rayO := ctx.camera.c2w.transformPt(Origin)
	// TODO: multithread this - tiles?
	for y := range image {
		for x := range image[y] {
			// convert to screen space
			xx := (2.0 * (float64(x) + 0.5) / float64(ctx.imgW) - 1.0) * ctx.camera.angle * ctx.aspectRatio
			yy := (1.0 - 2.0 * (float64(y) + 0.5) / float64(ctx.imgH)) * ctx.camera.angle

			//fmt.Printf("  %.3f %.3f\n", xx, yy)
			//camPos := ctx.camera.c2w.transformPt(NewPt(xx, yy, -1))
//			fmt.Printf("  %+v\n", camPos)
			//rayD := PtDelta(camPos, rayO)
			rayD := ctx.camera.c2w.rotateV3(NewV3(xx, yy, -1))
			ray := NewRay(rayO, rayD)
			renderStats.primaryRay++

			c := trace(ray, ctx.primitives, 0)
			image[y][x].Add(&c)
		}
	}
	renderTime = time.Since(startTime)
	log.Printf("Stats: %+v\n", renderStats)
	log.Printf("Time: %s\n", renderTime)
	return image
}
