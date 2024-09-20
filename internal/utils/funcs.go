package utils

import (
	"math"
	"math/rand/v2"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func RandomFloat64Interval(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
