package dielectric

import (
	"math"
	"math/rand/v2"

	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/tracer"
)

type dielectric struct {
	albedo          colors.Color3
	refractionIndex float64
}

func New(albedo colors.Color3, refractionIndex float64) dielectric {
	return dielectric{
		albedo:          albedo,
		refractionIndex: refractionIndex,
	}
}

func (d *dielectric) Scatter(r tracer.Ray, hitPoint geometry.Point3, hitPointNormal geometry.Vec3, isFrontFace bool) (tracer.Ray, colors.Color3) {

	ri := d.refractionIndex

	if isFrontFace {
		ri = 1 / d.refractionIndex
	}

	unitDirection := r.Direction.UnitVector()

	cosTheta := math.Min(unitDirection.Opposite().DotProduct(hitPointNormal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	var refracted geometry.Vec3

	if ri*sinTheta > 1 || reflectance(cosTheta, ri) > rand.Float64() {
		refracted = unitDirection.Reflect(hitPointNormal)
	} else {
		refracted = unitDirection.Refract(hitPointNormal, ri)
	}

	scattered := tracer.Ray{Origin: hitPoint, Direction: refracted}

	return scattered, d.albedo
}

func reflectance(cosine float64, refractionIndex float64) float64 {
	r0 := math.Pow((1-refractionIndex)/(1+refractionIndex), 2)
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
