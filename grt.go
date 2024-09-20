package main

import (
	"math/rand/v2"

	"github.com/kmanev073/grt/internal/camera"
	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/hittables"
	"github.com/kmanev073/grt/internal/hittables/hittableslist"
	"github.com/kmanev073/grt/internal/hittables/sphere"
	"github.com/kmanev073/grt/internal/imageencoders/pngencoder"
	"github.com/kmanev073/grt/internal/materials/dielectric"
	"github.com/kmanev073/grt/internal/materials/lambertian"
	"github.com/kmanev073/grt/internal/materials/metal"
)

func main() {
	const ImageAspectRatio = float32(16) / 9
	const ImageWidth = int(1200)
	const SamplesPerPixel = uint(1000)
	const MaxDepth = uint(50)
	const VFov = float64(20)
	const FocusDist = float64(10)
	const DefocusAngle = float64(0.6)

	var lookFrom = geometry.NewPoint3(13, 2, 3)
	var lookAt = geometry.NewPoint3(0, 0, 0)

	var groundMaterial = lambertian.New(colors.New(0.5, 0.5, 0.5))
	var world = &hittableslist.HittablesList{
		Objects: []hittables.Hittable{
			sphere.New(geometry.NewPoint3(0, -1000, 0), 1000, &groundMaterial),
		},
	}

	for i := -11; i < 11; i++ {
		for j := -11; j < 11; j++ {
			chooseMat := rand.Float64()
			center := geometry.NewPoint3(float64(i)+0.8*rand.Float64(), 0.2, float64(j)+0.8*rand.Float64())

			if center.Subtract(geometry.NewPoint3(4, 0.2, 0)).Length() > 0.9 &&
				center.Subtract(geometry.NewPoint3(-4, 0.2, 0)).Length() > 0.9 &&
				center.Subtract(geometry.NewPoint3(0, 0.2, 0)).Length() > 0.9 {

				if chooseMat < 0.34 {
					albedo := colors.Random().Merge(colors.Random())
					sphereMaterial := lambertian.New(albedo)
					sphere := sphere.New(center, 0.2, &sphereMaterial)
					world.Objects = append(world.Objects, sphere)
				} else if chooseMat < 0.67 {
					albedo := colors.RandomInterval(0.5, 1)
					fuzz := rand.Float64() / 2
					sphereMaterial := metal.New(albedo, fuzz)
					sphere := sphere.New(center, 0.2, &sphereMaterial)
					world.Objects = append(world.Objects, sphere)
				} else {
					albedo := colors.Random()
					sphereMaterial := dielectric.New(albedo, 1.5)
					sphere := sphere.New(center, 0.2, &sphereMaterial)
					world.Objects = append(world.Objects, sphere)
				}
			}
		}
	}

	material1 := dielectric.New(colors.White(), 1.5)
	sphere1 := sphere.New(geometry.NewPoint3(0, 1, 0), 1, &material1)
	world.Objects = append(world.Objects, sphere1)

	material2 := lambertian.New(colors.New(0.4, 0.2, 0.1))
	sphere2 := sphere.New(geometry.NewPoint3(-4, 1, 0), 1, &material2)
	world.Objects = append(world.Objects, sphere2)

	material3 := metal.New(colors.New(0.7, 0.6, 0.5), 0.0)
	sphere3 := sphere.New(geometry.NewPoint3(4, 1, 0), 1, &material3)
	world.Objects = append(world.Objects, sphere3)

	var imageEncoder = pngencoder.PngEncoder{}

	cam := camera.New(
		ImageAspectRatio,
		ImageWidth,
		SamplesPerPixel,
		MaxDepth,
		VFov,
		lookFrom,
		lookAt,
		FocusDist,
		DefocusAngle)

	cam.Render(world, &imageEncoder)
}
