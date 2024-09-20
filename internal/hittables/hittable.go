package hittables

import (
	"github.com/kmanev073/grt/internal/tracer"
	"github.com/kmanev073/grt/internal/utils"
)

type Hittable interface {
	Hit(r tracer.Ray, rayT utils.Interval) *HitRecord
}
