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

func Add(u Vec3, v Vec3) Vec3 {
	return Vec3{
		X: u.X + v.X,
		Y: u.Y + v.Y,
		Z: u.Z + v.Z,
	}
}

func Sub(u Vec3, v Vec3) Vec3 {
	return Vec3{
		X: u.X - v.X,
		Y: u.Y - v.Y,
		Z: u.Z - v.Z,
	}
}

func Mul(u Vec3, v Vec3) Vec3 {
	return Vec3{
		X: u.X * v.X,
		Y: u.Y * v.Y,
		Z: u.Z * v.Z,
	}
}

func Times(u Vec3, t float64) Vec3 {
	return Vec3{
		X: u.X * t,
		Y: u.X * t,
		Z: u.X * t,
	}
}

func Div(u Vec3, t float64) Vec3 {
	return Vec3{
		X: u.X / t,
		Y: u.Y / t,
		Z: u.Z / t,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) UnitVector() Vec3 {
	return Div(v, v.Length())
}
