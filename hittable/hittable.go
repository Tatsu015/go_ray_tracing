package hittable

import (
	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

type HitRecord struct {
	p      vec.Point
	normal vec.Vec3
	t      float64
}

type Hittable interface {
	Hit(ray *raytrace.Ray, tMin float64, tMax float64) (bool, HitRecord)
}
