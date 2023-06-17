package hittable

import (
	"math"

	"github.com/Tatsu015/go_ray_tracing.git/raytrace"
	"github.com/Tatsu015/go_ray_tracing.git/vec"
)

type Sphere struct {
	center vec.Point
	radius float64
}

func NewSphere(center vec.Point, radius float64) Sphere {
	return Sphere{
		center,
		radius,
	}
}

func (s *Sphere) Hit(ray *raytrace.Ray, tMin float64, tMax float64) *HitRecord {
	oc := ray.Origin.Sub(s.center)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.radius*s.radius
	d := b*b - a*c

	if d > 0 {
		root := math.Sqrt(d)
		t := (-b - root) / a
		if t < tMax && tMin < t {
			p := ray.At(t)
			r := NewHitRecord(p, p.Sub(s.center).UnitVector(), t)
			return &r
		}
		t = (-b + root) / a
		if t < tMax && tMin < t {
			p := ray.At(t)
			r := NewHitRecord(p, p.Sub(s.center).UnitVector(), t)
			return &r
		}
	}
	return nil
}
