package main

import (
	"fmt"
)

func main() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%d %d \n255\n", image_width, image_height)

	buf := ""

	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			b := 0.25

			c := Vec3{x: r, y: g, z: b}
			buf += WriteColor(c)
			fmt.Println(buf)
		}
	}
}
