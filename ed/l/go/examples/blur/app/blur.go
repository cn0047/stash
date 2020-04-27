// @example:
// go run ed/l/go/examples/blur/app/blur.go ed/l/python/examples/blur/z.png
package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func ti(f float32, i int) int {
	r := f * float32(i)
	return int(r)
}

func gr(h int, w int) image.Rectangle {
	top := ti(0.14, h)
	bottom := ti(0.25, h)
	left := ti(0.60, w)
	right := ti(0.71, w)

	return image.Rectangle{Min: image.Point{X: left, Y: top}, Max: image.Point{X: right, Y: bottom}}
}

func b(f string) {
	img := gocv.IMRead(f, gocv.IMReadUnchanged)
	defer img.Close()
	if img.Empty() {
		fmt.Printf("failed to read image: %s", f)
		return
	}

	s := img.Size()
	h := s[0]
	w := s[1]

	r := img.Region(gr(h, w))
	gocv.GaussianBlur(r, &r, image.Pt(99, 99), 10, 0, gocv.BorderDefault)
	r.Close()

	gocv.IMWrite(f+".r.png", img)
}

func main() {
	f := os.Args[1]
	b(f)
}
