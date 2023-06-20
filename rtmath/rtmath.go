package rtmath

import "math"

func Degrees2Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
