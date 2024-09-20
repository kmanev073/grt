package hittableslist

import (
	"github.com/kmanev073/grt/internal/hittables"
	"github.com/kmanev073/grt/internal/tracer"
	"github.com/kmanev073/grt/internal/utils"
)

type HittablesList struct {
	Objects []hittables.Hittable
}

func (list *HittablesList) Hit(r tracer.Ray, rayT utils.Interval) *hittables.HitRecord {

	var closestHit *hittables.HitRecord = nil
	closestSoFar := rayT.Max()

	for _, object := range list.Objects {
		tempRec := object.Hit(r, utils.NewInterval(rayT.Min(), closestSoFar))

		if tempRec != nil {
			closestHit = tempRec
			closestSoFar = tempRec.T
		}
	}

	return closestHit
}
