package pngencoder

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type PngEncoder struct{}

func (encoder *PngEncoder) Save(fileName string, img *image.RGBA) {
	f, _ := os.Create(fmt.Sprintf("%s.png", fileName))
	png.Encode(f, img)
}
