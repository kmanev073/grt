package colors

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/kmanev073/grt/internal/utils"
)

type Color3 struct {
	R float64
	G float64
	B float64
}

func New(r float64, g float64, b float64) Color3 {
	return Color3{R: r, G: g, B: b}
}

func Black() Color3 {
	return Color3{}
}

func White() Color3 {
	return Color3{1, 1, 1}
}

func Random() Color3 {
	return Color3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandomInterval(min float64, max float64) Color3 {
	return Color3{
		R: utils.RandomFloat64Interval(min, max),
		G: utils.RandomFloat64Interval(min, max),
		B: utils.RandomFloat64Interval(min, max),
	}
}

func (c Color3) ToRGBA() color.RGBA {

	intensity := utils.NewInterval(0, 0.999)

	rGamma := linearToGamma(c.R)
	gGamma := linearToGamma(c.G)
	bGamma := linearToGamma(c.B)

	r := byte(256 * intensity.Clamp(rGamma))
	g := byte(256 * intensity.Clamp(gGamma))
	b := byte(256 * intensity.Clamp(bGamma))

	return color.RGBA{r, g, b, 0xFF}
}

func (c Color3) Add(u Color3) Color3 {
	return Color3{
		R: c.R + u.R,
		G: c.G + u.G,
		B: c.B + u.B,
	}
}

func (c Color3) Scale(factor float64) Color3 {
	return Color3{
		R: c.R * factor,
		G: c.G * factor,
		B: c.B * factor,
	}
}

func (c Color3) Merge(u Color3) Color3 {
	return Color3{
		R: c.R * u.R,
		G: c.G * u.G,
		B: c.B * u.B,
	}
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}
	return 0
}
