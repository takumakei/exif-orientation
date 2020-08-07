package orientation_test

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	orientation "github.com/takumakei/exif-orientation"
)

func Example() {
	b, err := ioutil.ReadFile("./examples/Landscape_7.jpg")
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(b)

	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	r.Reset(b)
	o, _ := orientation.Read(r)
	img = orientation.Normalize(img, o)

	f, err := os.Create("normalized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
		log.Fatal(err)
	}
	// Output:
}
