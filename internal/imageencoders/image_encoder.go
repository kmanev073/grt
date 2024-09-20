package imageencoders

import "image"

type ImageEncoder interface {
	Save(fileName string, img *image.RGBA)
}
