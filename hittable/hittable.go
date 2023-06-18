package hittable

import (
	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

type HitRecord struct {
	p         vec.Point
	normal    vec.Vec3
	t         float64
	frontFace bool
}

func NewHitRecord(p vec.Point, normal vec.Vec3, t float64, frontFace bool) HitRecord {
	return HitRecord{p, normal, t, frontFace}
}

func isFrontFace(ray *raytrace.Ray, outwardNormal vec.Vec3) bool {
	frontFace := ray.Direction.Dot(outwardNormal) < 0
	if frontFace {
		return true
	} else {
		return false
	}
}

type Hittable interface {
	Hit(ray *raytrace.Ray, tMin float64, tMax float64) *HitRecord
}
