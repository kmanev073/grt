package camera

import (
	"fmt"
	"image"
	"math"
	"math/rand/v2"
	"runtime"
	"time"

	"github.com/kmanev073/grt/internal/colors"
	"github.com/kmanev073/grt/internal/geometry"
	"github.com/kmanev073/grt/internal/hittables"
	"github.com/kmanev073/grt/internal/imageencoders"
	"github.com/kmanev073/grt/internal/tracer"
	"github.com/kmanev073/grt/internal/utils"
)

type Camera interface {
	Render(world hittables.Hittable, imageEncoder imageencoders.ImageEncoder)
}

type camera struct {
	imageWidth        int
	imageHeight       int
	center            geometry.Point3
	pixelDeltaU       geometry.Vec3
	pixelDeltaV       geometry.Vec3
	pixel00Location   geometry.Vec3
	samplesPerPixel   uint
	pixelSamplesScale float64
	maxDepth          uint
	defocusAngle      float64
	defocusDiskU      geometry.Vec3
	defocusDiskV      geometry.Vec3
}

func New(
	imageAspectRatio float32,
	imageWidth int,
	samplesPerPixel uint,
	maxDepth uint,
	vFov float64,
	lookFrom geometry.Point3,
	lookAt geometry.Point3,
	focusDist float64,
	defocusAngle float64) *camera {
	newCamera := camera{
		imageWidth:      imageWidth,
		imageHeight:     int(float32(imageWidth) / imageAspectRatio),
		center:          lookFrom,
		samplesPerPixel: samplesPerPixel,
		maxDepth:        maxDepth,
		defocusAngle:    defocusAngle,
	}

	theta := utils.DegreesToRadians(vFov)
	h := math.Tan(theta / 2)
	focusDistance := focusDist

	if focusDistance <= 0 {
		focusDistance = lookFrom.Subtract(lookAt).Length()
	}

	viewportHeight := 2 * h * focusDistance
	viewportAspectRatio := float64(newCamera.imageWidth) / float64(newCamera.imageHeight)
	viewportWidth := viewportHeight * viewportAspectRatio

	vup := geometry.NewVec3(0, 1, 0)
	w := lookFrom.Subtract(lookAt).UnitVector()
	u := vup.CrossProduct(w).UnitVector()
	v := w.CrossProduct(u)

	viewportU := u.Scale(viewportWidth)
	viewportV := v.Opposite().Scale(viewportHeight)

	newCamera.pixelDeltaU = viewportU.Downscale(float64(newCamera.imageWidth))
	newCamera.pixelDeltaV = viewportV.Downscale(float64(newCamera.imageHeight))

	viewportUpperLeft := newCamera.center.
		Subtract(w.Scale(focusDistance)).
		Subtract(viewportU.Downscale(2)).
		Subtract(viewportV.Downscale(2))

	newCamera.pixel00Location = viewportUpperLeft.Add(newCamera.pixelDeltaU.
		Add(newCamera.pixelDeltaV).
		Downscale(2))

	newCamera.pixelSamplesScale = 1 / float64(newCamera.samplesPerPixel)

	defocusRadius := focusDistance * math.Tan(utils.DegreesToRadians((newCamera.defocusAngle / 2)))
	newCamera.defocusDiskU = u.Scale(defocusRadius)
	newCamera.defocusDiskV = v.Scale(defocusRadius)

	return &newCamera
}

func (c *camera) Render(world hittables.Hittable, imageEncoder imageencoders.ImageEncoder) {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.imageWidth, c.imageHeight}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	startTime := time.Now()
	goroutinesGuard := make(chan struct{}, runtime.NumCPU())

	for y := 0; y < c.imageHeight; y++ {

		fmt.Println("Scanlines remaining:", c.imageHeight-y)

		for x := 0; x < c.imageWidth; x++ {

			goroutinesGuard <- struct{}{}
			go func() {
				pixelColor := colors.Black()

				for sample := 0; sample < int(c.samplesPerPixel); sample++ {
					r := c.getRay(x, y)
					pixelColor = pixelColor.Add(rayColor(r, c.maxDepth, world))
				}

				pixelColor = pixelColor.Scale(c.pixelSamplesScale)

				color := pixelColor.ToRGBA()

				img.Set(x, y, color)

				<-goroutinesGuard
			}()
		}

		imageEncoder.Save("image", img)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Done in", elapsedTime, "with", runtime.NumCPU(), "threads.")
}

func (c *camera) getRay(x int, y int) tracer.Ray {
	offset := sampleSquare()
	pixelSample := c.pixel00Location.
		Add(c.pixelDeltaU.Scale((float64(x) + offset.X))).
		Add(c.pixelDeltaV.Scale((float64(y) + offset.Y)))

	rayOrigin := c.center
	if c.defocusAngle > 0 {
		rayOrigin = c.sampleDefocusDisk()
	}
	rayDirection := pixelSample.Subtract(rayOrigin)

	return tracer.Ray{
		Origin:    rayOrigin,
		Direction: rayDirection,
	}
}

func sampleSquare() geometry.Vec3 {
	return geometry.Vec3{
		X: rand.Float64() - 0.5,
		Y: rand.Float64() - 0.5,
		Z: 0,
	}
}

func (c *camera) sampleDefocusDisk() geometry.Point3 {
	p := geometry.RandomInUnitDisk()
	return c.center.Add(c.defocusDiskU.Scale(p.X)).Add(c.defocusDiskV.Scale(p.Y))
}

func rayColor(r tracer.Ray, depth uint, world hittables.Hittable) colors.Color3 {

	if depth <= 0 {
		return colors.Black()
	}

	hitRecord := world.Hit(r, utils.NewInterval(0.001, math.MaxFloat64))

	if hitRecord == nil {
		return sky(r)
	}

	scattered, attenuation := hitRecord.Material.Scatter(r, hitRecord.Point, hitRecord.Normal(), hitRecord.FrontFace())

	return rayColor(scattered, depth-1, world).Merge(attenuation)
}

func sky(r tracer.Ray) colors.Color3 {
	y := r.Direction.UnitVector().Y
	a := 0.5 * (y + 1)

	return colors.White().Scale(1 - a).Add(colors.New(0.5, 0.7, 1).Scale(a))
}
