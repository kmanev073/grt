package lambertian

import (
	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/tracer"
)

type lambertian struct {
	albedo colors.Color3
}

func New(albedo colors.Color3) lambertian {
	return lambertian{
		albedo: albedo,
	}
}

func (l *lambertian) Scatter(r tracer.Ray, hitPoint geometry.Point3, hitPointNormal geometry.Vec3, isFrontFace bool) (tracer.Ray, colors.Color3) {

	scatterDirection := hitPointNormal.Add(geometry.RandomUnitVector())

	if scatterDirection.IsNearZero() {
		scatterDirection = hitPointNormal
	}

	scattered := tracer.Ray{
		Origin:    hitPoint,
		Direction: scatterDirection,
	}

	return scattered, l.albedo
}
