package utils

import "math"

type Interval interface {
	Min() float64
	Max() float64
	Size() float64
	Contains(x float64) bool
	Surrounds(x float64) bool
	Clamp(x float64) float64
}

type interval struct {
	min float64
	max float64
}

func NewInterval(minValue float64, maxValue float64) *interval {
	return &interval{
		min: minValue,
		max: maxValue,
	}
}

func EmptyInterval(minValue float64, maxValue float64) *interval {
	return &interval{
		min: math.MaxFloat64,
		max: -math.MaxFloat64,
	}
}

func UniverseInterval(minValue float64, maxValue float64) *interval {
	return &interval{
		min: -math.MaxFloat64,
		max: math.MaxFloat64,
	}
}

func (i *interval) Min() float64 {
	return i.min
}

func (i *interval) Max() float64 {
	return i.max
}

func (i *interval) Size() float64 {
	return i.max - i.min
}

func (i *interval) Contains(x float64) bool {
	return i.min <= x && x <= i.max
}

func (i *interval) Surrounds(x float64) bool {
	return i.min < x && x < i.max
}

func (i *interval) Clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}

	if x > i.max {
		return i.max
	}

	return x
}
