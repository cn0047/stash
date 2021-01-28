package main

import (
	"image"

	"gocv.io/x/gocv"
)

func b(f string) {
	img := gocv.IMRead(f, gocv.IMReadUnchanged)
	defer img.Close()

	top, bottom, left, right := 286, 308, 111, 144
	r := image.Rectangle{Min: image.Point{X: left, Y: top}, Max: image.Point{X: right, Y: bottom}}
	rec := img.Region(r)
	gocv.GaussianBlur(rec, &rec, image.Pt(3, 9), 5, 5, gocv.BorderDefault)
	rec.Close()

	gocv.IMWrite("img.after.go.png", img)
}

func main() {
	b("/Users/kovpakvolodymyr/web/kovpak/gh/ed/l/python/examples/blur/img.before.png")
}
