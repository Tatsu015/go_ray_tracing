package raytrace

import "github.com/Tatsu015/go_ray_tracing.git/vec"

type Ray struct {
	Origin    vec.Vec3
	Direction vec.Vec3
}

func NewRay(origin vec.Vec3, direction vec.Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r *Ray) At(t float64) vec.Point {
	return vec.Add(r.Origin, vec.Times(r.Direction, t))
}
