package vec

import (
	"math"

	"github.com/Tatsu015/go_ray_tracing.git/rtmath"
)

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

func RandomVec3() Vec3 {
	return NewVec3(rtmath.RandomDouble(), rtmath.RandomDouble(), rtmath.RandomDouble())
}

func RandomVec3Inrange(min float64, max float64) Vec3 {
	return NewVec3(rtmath.RandomDoubleInRange(min, max), rtmath.RandomDoubleInRange(min, max), rtmath.RandomDoubleInRange(min, max))
}

func RandomUnitVector() Vec3 {
	a := rtmath.RandomDoubleInRange(0, 2*math.Pi)
	z := rtmath.RandomDoubleInRange(-1, 1)
	r := math.Sqrt((1 - z*z))
	return NewVec3(r*math.Cos(a), r*math.Sin(a), z)
}

func RandomInUnitSphere() Vec3 {
	for {
		p := RandomVec3Inrange(-1, 1)
		if p.LengthSquared() < 1 {
			return p
		}
	}
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
		Y: v.Y * t,
		Z: v.Z * t,
	}
}

func (v Vec3) Div(t float64) Vec3 {
	return Vec3{
		X: v.X / t,
		Y: v.Y / t,
		Z: v.Z / t,
	}
}

func (v Vec3) Clamp(min float64, max float64) Vec3 {
	return Vec3{
		X: rtmath.Clamp(v.X, min, max),
		Y: rtmath.Clamp(v.Y, min, max),
		Z: rtmath.Clamp(v.Z, min, max),
	}
}

func (v Vec3) Sqrt() Vec3 {
	return Vec3{
		X: math.Sqrt(v.X),
		Y: math.Sqrt(v.Y),
		Z: math.Sqrt(v.Z),
	}
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
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
