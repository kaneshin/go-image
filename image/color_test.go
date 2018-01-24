package image

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Average(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			c := color.RGBA{197, 255, 0, 255}
			img.Set(j, i, c)
		}
	}

	t.Run("AverageRGBA", func(t *testing.T) {
		ta := assert.New(t)
		r, g, b, a := AverageRGBA(nil)
		ta.EqualValues(0, r)
		ta.EqualValues(0, g)
		ta.EqualValues(0, b)
		ta.EqualValues(1, a)

		r, g, b, a = AverageRGBA(img)
		ta.EqualValues(197, uint8(r))
		ta.EqualValues(255, uint8(g))
		ta.EqualValues(0, uint8(b))
		ta.EqualValues(255, uint8(a))
	})

	t.Run("AverageHexString", func(t *testing.T) {
		ta := assert.New(t)
		ta.Equal("#000000", AverageHexString(nil))
		ta.Equal("#c5ff00", AverageHexString(img))
	})
}
