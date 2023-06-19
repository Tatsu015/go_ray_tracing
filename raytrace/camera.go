package raytrace

import "github.com/Tatsu015/go_ray_tracing.git/vec"

type Camera struct {
	origin          vec.Vec3
	lowerLeftCorner vec.Vec3
	horizontal      vec.Vec3
	vertical        vec.Vec3
}

func NewCamera(aspectRatio float64) Camera {
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := vec.NewVec3(0, 0, 0)
	horizontal := vec.NewVec3(viewportWidth, 0, 0)
	vertival := vec.NewVec3(0, viewportHeight, 0)
	focalPos := vec.NewVec3(0, 0, focalLength)
	lowerLeftCorner := origin.Sub(horizontal.Times(0.5)).Sub(vertival.Times(0.5)).Sub(focalPos)

	return Camera{
		origin,
		lowerLeftCorner,
		horizontal,
		vertival,
	}
}

func (c *Camera) GetRay(u float64, v float64) Ray {
	d := c.lowerLeftCorner.Add(c.horizontal.Times(u)).Add(c.vertical.Times(v)).Sub(c.origin)
	return NewRay(c.origin, d)
}
