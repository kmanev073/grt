package metal

import (
	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/tracer"
)

type metal struct {
	albedo colors.Color3
	fuzz   float64
}

func New(albedo colors.Color3, fuzz float64) metal {
	return metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

func (m *metal) Scatter(r tracer.Ray, hitPoint geometry.Point3, hitPointNormal geometry.Vec3, isFrontFace bool) (tracer.Ray, colors.Color3) {

	reflected := r.Direction.Reflect(hitPointNormal)

	reflectedDirection := reflected.UnitVector().Add(geometry.RandomUnitVector().Scale(m.fuzz))

	for reflectedDirection.DotProduct(hitPointNormal) <= 0 {
		reflectedDirection = reflected.UnitVector().Add(geometry.RandomUnitVector().Scale(m.fuzz))
	}

	scattered := tracer.Ray{
		Origin:    hitPoint,
		Direction: reflectedDirection,
	}

	return scattered, m.albedo
}
