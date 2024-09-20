package materials

import (
	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/tracer"
)

type Material interface {
	Scatter(r tracer.Ray, hitPoint geometry.Point3, hitPointNormal geometry.Vec3, isFrontFace bool) (tracer.Ray, colors.Color3)
}
