package main

import (
	"fmt"
	"os"

	"github.com/Tatsu015/go_ray_tracing.git/ppm"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

func main() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%d %d \n255\n", image_width, image_height)

	buf := ""

	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %3d", j)
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			b := 0.25

			c := vec.NewVec3(r, g, b)
			buf += ppm.WriteColor(c)
		}
	}
	fmt.Println(buf)
	fmt.Fprintf(os.Stderr, "\033[2K\rDone!\n")
}
