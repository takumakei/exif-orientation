package orientation

import (
	"image"
	"image/draw"
)

// NormalizeCopy returns new copied image that of orientation is TopLeft.
//
// The image returned would be created by creator.
func NormalizeCopy(src image.Image, orientation Orientation, creator func(image.Rectangle) draw.Image) draw.Image {
	return normalize[int(orientation)-1](src, creator)
}

// Normalize returns an image that of orientation is TopLeft.
// In case of orientation is TopLeft, Normalize returns src itself.
// Otherwise the image returned would be created by NewRGBA.
func Normalize(src image.Image, orientation Orientation) image.Image {
	if orientation == TopLeft {
		return src
	}
	return normalize[int(orientation)-1](src, NewRGBA)
}

// NewRGBA returns image.NewRGBA(bounds).
//
// Normalize uses NewRGBA.
func NewRGBA(bounds image.Rectangle) draw.Image { return image.NewRGBA(bounds) }

var normalize = []func(image.Image, func(image.Rectangle) draw.Image) draw.Image{
	topLeft,
	topRight,
	bottomRight,
	bottomLeft,
	leftTop,
	rightTop,
	rightBottom,
	leftBottom,
}

func topLeft(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(bounds)
	for y := t; y < b; y++ {
		for x := l; x < r; x++ {
			i.Set(x, y, src.At(x, y))
		}
	}
	return i
}

func topRight(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(bounds)
	u := r - 1
	for y := t; y < b; y++ {
		for x := l; x < r; x++ {
			i.Set(x, y, src.At(u-x, y))
		}
	}
	return i
}

func bottomRight(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(bounds)
	u := r - 1
	v := b - 1
	for y := t; y < b; y++ {
		for x := l; x < r; x++ {
			i.Set(x, y, src.At(u-x, v-y))
		}
	}
	return i
}

func bottomLeft(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(bounds)
	v := b - 1
	for y := t; y < b; y++ {
		for x := l; x < r; x++ {
			i.Set(x, y, src.At(x, v-y))
		}
	}
	return i
}

func leftTop(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(image.Rect(0, 0, b-t, r-l))
	dx := b - t - 1
	for y := b - 1; y >= t; y-- {
		dy := 0
		for x := l; x < r; x++ {
			i.Set(dx, dy, src.At(x, y))
			dy++
		}
		dx--
	}
	return i
}

func rightTop(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(image.Rect(0, 0, b-t, r-l))
	dx := 0
	for y := b - 1; y >= t; y-- {
		dy := 0
		for x := l; x < r; x++ {
			i.Set(dx, dy, src.At(x, y))
			dy++
		}
		dx++
	}
	return i
}

func rightBottom(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(image.Rect(0, 0, b-t, r-l))
	dx := b - t - 1
	for y := t; y < b; y++ {
		dy := 0
		for x := r - 1; x >= l; x-- {
			i.Set(dx, dy, src.At(x, y))
			dy++
		}
		dx--
	}
	return i
}

func leftBottom(src image.Image, f func(image.Rectangle) draw.Image) draw.Image {
	bounds := src.Bounds()
	l, r := bounds.Min.X, bounds.Max.X // left, right
	t, b := bounds.Min.Y, bounds.Max.Y // top, bottom
	i := image.NewRGBA(image.Rect(0, 0, b-t, r-l))
	dx := 0
	for y := t; y < b; y++ {
		dy := 0
		for x := r - 1; x >= l; x-- {
			i.Set(dx, dy, src.At(x, y))
			dy++
		}
		dx++
	}
	return i
}
