package sphere

import (
	"math"

	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/hittables"
	"github.com/kmanev073/grt/internal/materials"
	"github.com/kmanev073/grt/internal/tracer"
	"github.com/kmanev073/grt/internal/utils"
)

type sphere struct {
	material materials.Material

	Center geometry.Point3
	Radius float64
}

func New(center geometry.Point3, radius float64, material materials.Material) *sphere {
	return &sphere{
		material: material,
		Center:   center,
		Radius:   radius,
	}
}

func (s *sphere) Hit(r tracer.Ray, rayT utils.Interval) *hittables.HitRecord {

	oc := s.Center.Subtract(r.Origin)

	a := r.Direction.LengthSquared()
	h := r.Direction.DotProduct(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return nil
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtd) / a
		if !rayT.Surrounds(root) {
			return nil
		}
	}

	hitRecord := hittables.HitRecord{
		T:     root,
		Point: r.At(root),
	}

	outwardNormal := hitRecord.Point.Subtract(s.Center).Downscale(s.Radius)
	hitRecord.SetFaceNormal(r, outwardNormal)

	hitRecord.Material = s.material

	return &hitRecord
}
