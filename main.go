package main

import (
	"fmt"
)

func main() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%d %d \n255\n", image_width, image_height)

	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {
			var r = float64(i) / (image_width - 1)
			var g = float64(j) / (image_height - 1)
			var b = 0.25

			fmt.Printf("%d %d %d\n", int(255.999*r), int(255.999*g), int(255.999*b))
		}
	}
}
