package trace

import (
	"fmt"
	"log"
)

const maxDepth = 3
var background = NewColor(1.0, 0.0, 1.0)

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
	var hitPrim *Primitive = nil
	for _, p := range prims {
		// TODO: recurse and return color
		hit, t := p.Intersect(r)
		//fmt.Printf("  %v %.2f (%.2f, %.2f]\n", hit, t, r.t0, r.t1)
		if hit && t < minT && t > r.t0 {
			minT = t
			hitPrim = &p
		}
	}

	if hitPrim == nil {
		return *background
	}

	// TODO: reflection/refraction

	// TODO: lights
	/*
	hitPt := PtAdd(&r.o, V3Mul(&r.d, mindist))

	c := Color{0.0, 0.0, 0.0}
	for _, l := range prims {
		if !l.isLight() {
			continue
		}
		lightPt := l.randomPt()
		shadowRay := NewRay(hitPt, PtDelta(lightPt, hitPt))
		for i, p := range prims {
			if i == hitInx {
				continue
			}
			if hit, _ := p.Intersect(shadowRay); !hit {
				c.Add(p.material().color().Mul(l.material().color()))
			}
		}
	}
	return &c
	*/
	return (*hitPrim).material().color()
}

func Render(ctx *Context) Image {
	log.Printf("Rendering...\n")
	image := makeImage(ctx.imgW, ctx.imgH)

	rayO := ctx.camera.c2w.transformPt(Origin)
	// TODO: multithread this - tiles?
	for y := range image {
		for x := range image[y] {
			// convert to screen space
			xx := (2.0 * (float64(x) + 0.5) / float64(ctx.imgW) - 1.0) * ctx.camera.angle * ctx.aspectRatio
			yy := (1.0 - 2.0 * (float64(y) + 0.5) / float64(ctx.imgH)) * ctx.camera.angle

			//fmt.Printf("  %.3f %.3f\n", xx, yy)
			camPos := ctx.camera.c2w.transformPt(NewPt(xx, yy, -1))
//			fmt.Printf("  %+v\n", camPos)
			//rayD := c2w.rotateV3(NewV3(xx, yy, -1))
			rayD := PtDelta(camPos, rayO)
			ray := NewRay(rayO, rayD)

			c := trace(ray, ctx.primitives, 0)
			image[y][x].Add(&c)
		}
	}

	return image
}
