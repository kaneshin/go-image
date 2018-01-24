package image

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"
)

func AverageRGBA(img image.Image) (r, g, b, a uint32) {
	if img == nil {
		return 0, 0, 0, 1
	}

	m := resize.Resize(1, 1, img, resize.Lanczos3)
	if m == nil {
		return 0, 0, 0, 1
	}

	c := m.At(0, 0)
	if c == nil {
		return 0, 0, 0, 1
	}
	return c.RGBA()
}

func AverageHexString(img image.Image) string {
	if img == nil {
		return "#000000"
	}

	r, g, b, _ := AverageRGBA(img)
	return fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
}
