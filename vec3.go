package main

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

func Add(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x + v.x,
		y: u.y + v.y,
		z: u.z + v.z,
	}
}

func Sub(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x - v.x,
		y: u.y - v.y,
		z: u.z - v.z,
	}
}

func Mul(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x * v.x,
		y: u.y * v.y,
		z: u.z * v.z,
	}
}

func Times(u Vec3, t float64) Vec3 {
	return Vec3{
		x: u.x * t,
		y: u.x * t,
		z: u.x * t,
	}
}

func Div(u Vec3, t float64) Vec3 {
	return Vec3{
		x: u.x / t,
		y: u.y / t,
		z: u.z / t,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) UnitVector() Vec3 {
	return Div(v, v.Length())
}
