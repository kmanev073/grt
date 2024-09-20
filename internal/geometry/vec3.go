package geometry

import (
	"math"
	"math/rand/v2"

	"github.com/kmanev073/grt/internal/utils"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vec3) Opposite() Vec3 {
	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}

func (v Vec3) Subtract(u Vec3) Vec3 {
	return Vec3{
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}

func (v Vec3) Scale(factor float64) Vec3 {
	return Vec3{
		X: v.X * factor,
		Y: v.Y * factor,
		Z: v.Z * factor,
	}
}

func (v Vec3) Downscale(factor float64) Vec3 {
	return v.Scale(1 / factor)
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) IsNearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func (v Vec3) UnitVector() Vec3 {
	return v.Downscale(v.Length())
}

func (v Vec3) DotProduct(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) CrossProduct(u Vec3) Vec3 {
	return Vec3{
		X: v.Y*u.Z - v.Z*u.Y,
		Y: v.Z*u.X - v.X*u.Z,
		Z: v.X*u.Y - v.Y*u.X,
	}
}

func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Subtract(normal.Scale(2 * v.DotProduct(normal)))
}

func (v Vec3) Refract(n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(v.Opposite().DotProduct(n), 1.0)

	rOutPerp := v.Add(n.Scale(cosTheta)).Scale(etaiOverEtat)
	rOutParallel := n.Scale(-math.Sqrt(math.Abs(1 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func Random() Vec3 {
	return Vec3{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64(),
	}
}

func RandomInterval(min float64, max float64) Vec3 {
	return Vec3{
		X: utils.RandomFloat64Interval(min, max),
		Y: utils.RandomFloat64Interval(min, max),
		Z: utils.RandomFloat64Interval(min, max),
	}
}

func RandomUnitVector() Vec3 {
	for {
		p := RandomInterval(-1, 1)
		lensq := p.LengthSquared()

		if 1e-160 < lensq && lensq <= 1 {
			return p.Downscale(math.Sqrt(lensq))
		}
	}
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	onUnitSphere := RandomUnitVector()

	if onUnitSphere.DotProduct(normal) >= 0 {
		return onUnitSphere
	}

	return onUnitSphere.Opposite()
}

func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{2*rand.Float64() - 1, 2*rand.Float64() - 1, 0}
		if p.LengthSquared() < 1 {
			return p
		}
	}
}
