package vec

import "math"

type Color = Vec3
type Point = Vec3

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{
		X: u.X + v.X,
		Y: u.Y + v.Y,
		Z: u.Z + v.Z,
	}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}

func (v Vec3) Mul(u Vec3) Vec3 {
	return Vec3{
		X: u.X * v.X,
		Y: u.Y * v.Y,
		Z: u.Z * v.Z,
	}
}

func (v Vec3) Times(t float64) Vec3 {
	return Vec3{
		X: v.X * t,
		Y: v.X * t,
		Z: v.X * t,
	}
}

func (v Vec3) Div(t float64) Vec3 {
	return Vec3{
		X: v.X / t,
		Y: v.Y / t,
		Z: v.Z / t,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) UnitVector() Vec3 {
	return v.Div(v.Length())
}