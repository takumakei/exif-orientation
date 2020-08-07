package orientation

import (
	"fmt"
	"io"

	"github.com/rwcarlsen/goexif/exif"
)

// Read returns an orientation, TopLeft in case of err != nil.
func Read(r io.Reader) (Orientation, error) {
	e, err := exif.Decode(r)
	if err != nil {
		return TopLeft, err
	}

	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return TopLeft, err
	}

	v, err := tag.Int(0)
	if err != nil {
		return TopLeft, err
	}

	o := Orientation(v)

	if !(1 <= v && v <= 8) {
		return TopLeft, fmt.Errorf("unknown rotation value: %d(%v)", v, o)
	}

	return o, nil
}
