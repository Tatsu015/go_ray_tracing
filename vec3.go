package main

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

func add(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x + v.x,
		y: u.y + v.y,
		z: u.z + v.z,
	}
}

func sub(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x - v.x,
		y: u.y - v.y,
		z: u.z - v.z,
	}
}

func mul(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x * v.x,
		y: u.y * v.y,
		z: u.z * v.z,
	}
}

func times(u Vec3, t float64) Vec3 {
	return Vec3{
		x: u.x * t,
		y: u.x * t,
		z: u.x * t,
	}
}

func div(u Vec3, t float64) Vec3 {
	return Vec3{
		x: u.x / t,
		y: u.y / t,
		z: u.z / t,
	}
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.length_squared())
}

func (v Vec3) length_squared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}
