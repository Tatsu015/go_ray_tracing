package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Tatsu015/go_ray_tracing.git/ppm"
	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

var WHITE = vec.NewVec3(1, 1, 1)
var BLUE = vec.NewVec3(0.5, 0.7, 1)

func hitSphere(center *vec.Point, radius float64, ray *raytrace.Ray) float64 {
	oc := ray.Origin.Sub(*center)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - radius*radius
	d := b*b - a*c
	if d < 0 {
		return -1
	} else {
		return (-b - math.Sqrt(d)) / a
	}
}

func rayColor(r *raytrace.Ray) vec.Vec3 {
	center := vec.NewVec3(0, 0, -1)
	t := hitSphere(&center, 0.5, r)
	if t > 0 {
		n := r.At(t).Sub(center)
		return vec.NewVec3(n.X+1, n.Y+1, n.Z+1).Times(0.5)
	}
	ud := &r.Direction
	t = 0.5 * (ud.Y + 1)
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
		for i := 0; i < image_width; i++ {
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)

			d := lower_left_corner.Add(horizontal.Times(u)).Add(vertival.Times(v)).Sub(origin)
			r := raytrace.NewRay(origin, d)

			c := rayColor(&r)
			buf += ppm.WriteColor(c)
		}
	}
	fmt.Println(buf)
	fmt.Fprintf(os.Stderr, "\nDone!\n")
}
