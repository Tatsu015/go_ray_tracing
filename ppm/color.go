package ppm

import (
	"fmt"

	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

func WriteColor(pixcelColor vec.Color, samplePerPixcel int) string {
	c := pixcelColor.Div(float64(samplePerPixcel)).Sqrt().Clamp(0, 0.999).Times(256)
	buf := fmt.Sprintf("%d %d %d\n", int(c.X), int(c.Y), int(c.Z))
	return buf
}
