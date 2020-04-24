package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	filename := "/Users/kovpakvolodymyr/web/kovpak/gh/ed/l/python/examples/blur/z.png"

	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}
	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
