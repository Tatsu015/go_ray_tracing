package hittable

import "github.com/Tatsu015/go_ray_tracing.git/raytrace"

type HittableList struct {
	hittables []Hittable
}

func NewHittableList() HittableList {
	return HittableList{}
}

func (h *HittableList) Add(hittable Hittable) {
	h.hittables = append(h.hittables, hittable)
}

func (h *HittableList) Hit(ray *raytrace.Ray, tMin float64, tMax float64) *HitRecord {
	closeSoFar := tMax
	var hitRecord *HitRecord = nil

	for _, v := range h.hittables {
		rec := v.Hit(ray, tMin, closeSoFar)
		if rec != nil {
			closeSoFar = rec.t
			hitRecord = rec
		}

	}
	return hitRecord
}
