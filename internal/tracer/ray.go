package tracer

import "github.com/kmanev073/grt/internal/geometry"

type Ray struct {
	Origin    geometry.Point3
	Direction geometry.Vec3
}

func (r Ray) At(t float64) geometry.Point3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
