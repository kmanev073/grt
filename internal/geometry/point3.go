package geometry

type Point3 = Vec3

func NewPoint3(x float64, y float64, z float64) Point3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}
