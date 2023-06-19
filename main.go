package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Tatsu015/go_ray_tracing.git/hittable"
	"github.com/Tatsu015/go_ray_tracing.git/ppm"
	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

var WHITE = vec.NewVec3(1, 1, 1)
var BLUE = vec.NewVec3(0.5, 0.7, 1)

func rayColor(ray *raytrace.Ray, world *hittable.HittableList) vec.Vec3 {
	rec := world.Hit(ray, 0, math.Inf(0))
	if rec != nil {
		offset := vec.NewVec3(1, 1, 1)
		return rec.GetNormal().Add(offset).Times(0.5)
	}

	ud := &ray.Direction
	t := 0.5 * (ud.Y + 1)
	return WHITE.Times(1 - t).Add(BLUE.Times(t))
}

func main() {
	const ASPECT_RATIO = 16.0 / 9.0
	var image_width = 100
	var image_height = int(float64(image_width) / ASPECT_RATIO)
	viewport_height := 2.0
	viewport_width := ASPECT_RATIO * viewport_height
	focal_length := 1.0

	origin := vec.NewVec3(0, 0, 0)
	horizontal := vec.NewVec3(viewport_width, 0, 0)
	vertival := vec.NewVec3(0, viewport_height, 0)
	focal_pos := vec.NewVec3(0, 0, focal_length)
	lower_left_corner := origin.Sub(horizontal.Times(0.5)).Sub(vertival.Times(0.5)).Sub(focal_pos)

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
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)

			d := lower_left_corner.Add(horizontal.Times(u)).Add(vertival.Times(v)).Sub(origin)
			r := raytrace.NewRay(origin, d)

			c := rayColor(&r, &world)
			buf += ppm.WriteColor(c)
		}
	}
	fmt.Println(buf)
	fmt.Fprintf(os.Stderr, "\nDone!\n")
}
