/*

@example:
docker build -t cn007b/go:1.14-opencv -f ed/sh/docker/examples.Dockerfile/examples.go.blur.Dockerfile .
GOPATH=/gh/ed/l/go/examples/blur/app
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH -w $GOPATH cn007b/go:1.14-opencv sh -c '
	go run blur.go /gh/ed/l/python/examples/blur/z.png
'

*/

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

func gr_z(h int, w int) image.Rectangle {
	top := ti(0.14, h)
	bottom := ti(0.25, h)
	left := ti(0.60, w)
	right := ti(0.71, w)

	return image.Rectangle{Min: image.Point{X: left, Y: top}, Max: image.Point{X: right, Y: bottom}}
}

func gr_n(h int, w int) image.Rectangle {
	top := ti(0.065, h)
	bottom := ti(0.16, h)
	left := ti(0.49, w)
	right := ti(0.528, w)

	return image.Rectangle{Min: image.Point{X: left, Y: top}, Max: image.Point{X: right, Y: bottom}}
}

func gr(h int, w int) image.Rectangle {
	return gr_n(h, w)
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
	gocv.GaussianBlur(r, &r, image.Pt(99, 99), 10, 10, gocv.BorderDefault)
	r.Close()

	gocv.IMWrite(f+".r.png", img)
}

func main() {
	f := os.Args[1]
	b(f)
}
