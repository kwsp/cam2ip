package image

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/anthonynsimon/bild/transform"
	"github.com/pbnjay/pixfont"
)

func Rotate(img image.Image, angle int) image.Image {
	switch angle {
	case 90:
		img = transform.Rotate(img, 90, &transform.RotationOptions{ResizeBounds: true})
	case 180:
		img = transform.Rotate(img, 180, &transform.RotationOptions{ResizeBounds: true})
	case 270:
		img = transform.Rotate(img, 270, &transform.RotationOptions{ResizeBounds: true})
	}

	return img
}

func Flip(img image.Image, dir string) image.Image {
	switch dir {
	case "horizontal":
		img = transform.FlipH(img)
	case "vertical":
		img = transform.FlipV(img)
	}

	return img
}

func Timestamp(img image.Image, format string) image.Image {
	dimg, ok := img.(draw.Image)
	if !ok {
		b := img.Bounds()
		dimg = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(dimg, b, img, b.Min, draw.Src)
	}

	pixfont.DrawString(dimg, 10, 10, time.Now().Format(format), color.White)

	return dimg
}

func CropImageCenter(src image.Image, cropW, cropH int) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	if cropW > w || cropH > h {
		return src
	}

	// top-left corner of crop rectangle
	startX := bounds.Min.X + (w-cropW)/2
	startY := bounds.Min.Y + (h-cropH)/2

	cropRect := image.Rect(startX, startY, startX+cropW, startY+cropH)
	cropped := image.NewRGBA(image.Rect(0, 0, cropW, cropH))
	draw.Draw(cropped, cropped.Bounds(), src, cropRect.Min, draw.Src)

	return cropped
}
