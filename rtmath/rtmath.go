package rtmath

import (
	"math"
	"math/rand"
	"time"
)

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

func RandomDouble() float64 {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return rand.Float64()
}
