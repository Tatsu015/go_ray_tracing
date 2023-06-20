package ppm

import (
	"fmt"

	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

func WriteColor(pixcelColor vec.Color, samplePerPixcel int) string {
	c := pixcelColor.Div(float64(samplePerPixcel)).Times(256).Clamp(0, 0.999)
	buf := fmt.Sprintf("%d %d %d\n", int(c.X), int(c.Y), int(c.Z))
	return buf
}
