package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Tatsu015/go_ray_tracing.git/hittable"
	"github.com/Tatsu015/go_ray_tracing.git/ppm"
	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/rtmath"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

var WHITE = vec.NewVec3(1, 1, 1)
var BLUE = vec.NewVec3(0.5, 0.7, 1)

func rayColor(ray *raytrace.Ray, world *hittable.HittableList, depth int) vec.Vec3 {
	if depth <= 0 {
		return vec.NewVec3(0, 0, 0)
	}
	rec := world.Hit(ray, 0.001, math.Inf(0))
	if rec != nil {
		t := rec.GetPoint().Add(vec.RandomInHemiSphere(rec.GetNormal()))
		refRay := raytrace.NewRay(rec.GetPoint(), t.Sub(rec.GetPoint()))
		return rayColor(&refRay, world, depth-1).Times(0.5)
	}

	ud := &ray.Direction
	t := 0.5 * (ud.Y + 1)
	return WHITE.Times(1 - t).Add(BLUE.Times(t))
}

func main() {
	const SAMPLE_PER_PIXCEL = 50
	const ASPECT_RATIO = 16.0 / 9.0
	const MAX_DEPTH = 50
	var image_width = 256
	var image_height = int(float64(image_width) / ASPECT_RATIO)
	camera := raytrace.NewCamera(ASPECT_RATIO)

	fmt.Printf("P3\n%d %d \n255\n", image_width, image_height)

	buf := ""

	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %3d", j)

		world := hittable.NewHittableList()
		h1 := hittable.NewSphere(vec.NewVec3(0, 0, -1), 0.5)
		world.Add(&h1)
		h2 := hittable.NewSphere(vec.NewVec3(0, -100.5, -1), 100)
		world.Add(&h2)

		for i := 0; i < image_width; i++ {
			sum := vec.NewVec3(0, 0, 0)
			for k := 0; k < SAMPLE_PER_PIXCEL; k++ {
				u := (float64(i) + rtmath.RandomDouble()) / float64(image_width-1)
				v := (float64(j) + rtmath.RandomDouble()) / float64(image_height-1)

				r := camera.GetRay(u, v)

				c := rayColor(&r, &world, MAX_DEPTH)
				sum = sum.Add(c)
			}
			buf += ppm.WriteColor(sum, SAMPLE_PER_PIXCEL)
		}
	}
	fmt.Println(buf)
	fmt.Fprintf(os.Stderr, "\nDone!\n")
}
