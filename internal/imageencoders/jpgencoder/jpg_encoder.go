package jpgencoder

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

type JpgEncoder struct{}

func (encoder *JpgEncoder) Save(fileName string, img *image.RGBA) {
	f, _ := os.Create(fmt.Sprintf("%s.jpg", fileName))

	jpeg.Encode(f, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
}
