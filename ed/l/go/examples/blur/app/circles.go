package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	filename := "/Users/kovpakvolodymyr/web/kovpak/gh/ed/l/python/examples/blur/z.png"
	resfilename := "z.res.png"

	window := gocv.NewWindow("detected circles")
	defer window.Close()

	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	defer img.Close()

	gocv.MedianBlur(img, &img, 5)

	cimg := gocv.NewMat()
	defer cimg.Close()

	gocv.CvtColor(img, &cimg, gocv.ColorGrayToBGR)

	circles := gocv.NewMat()
	defer circles.Close()

	gocv.HoughCirclesWithParams(
		img,
		&circles,
		gocv.HoughGradient,
		1,                     // dp
		float64(img.Rows()/8), // minDist
		75,                    // param1
		20,                    // param2
		10,                    // minRadius
		0,                     // maxRadius
	)

	blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}

	for i := 0; i < circles.Cols(); i++ {
		v := circles.GetVecfAt(0, i)
		// if circles are found
		if len(v) > 2 {
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])

			gocv.Circle(&cimg, image.Pt(x, y), r, blue, 2)
			gocv.Circle(&cimg, image.Pt(x, y), 2, red, 3)
		}
	}

	gocv.IMWrite(resfilename, cimg)
	return
	for {
		window.IMShow(cimg)

		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
