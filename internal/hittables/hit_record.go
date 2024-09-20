package hittables

import (
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/materials"
	"github.com/kmanev073/grt/internal/tracer"
)

type HitRecord struct {
	normal    geometry.Vec3
	frontFace bool

	Material materials.Material
	Point    geometry.Point3
	T        float64
}

func (hr *HitRecord) Normal() geometry.Vec3 {
	return hr.normal
}

func (hr *HitRecord) FrontFace() bool {
	return hr.frontFace
}

func (hr *HitRecord) SetFaceNormal(r tracer.Ray, outwardNormal geometry.Vec3) {

	hr.frontFace = r.Direction.DotProduct(outwardNormal) < 0

	if hr.frontFace {
		hr.normal = outwardNormal
	} else {
		hr.normal = outwardNormal.Opposite()
	}
}
