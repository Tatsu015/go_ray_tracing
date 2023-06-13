package ppm

import (
	"fmt"

	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

func WriteColor(pixcelColor vec.Color) string {
	buf := fmt.Sprintf("%d %d %d\n", int(255.999*pixcelColor.X), int(255.999*pixcelColor.Y), int(255.999*pixcelColor.Z))
	return buf
}
