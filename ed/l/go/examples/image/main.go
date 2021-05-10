package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	GopherImg = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`
)

func main() {
	imgFilePath := ""
	if len(os.Args) > 1 {
		imgFilePath = os.Args[1]
	}
	if imgFilePath != "" {
		fmt.Printf("imgFilePath: %+v \n", imgFilePath)
	}

	downloadAndConvertIfGIF()
	//img := decodeFile(imgFilePath)
	//toGIF(img)
	//toJPEG(img)
	//toPNG(img)
	//strImgToFile(GopherImg)
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}

func downloadAndConvertIfGIF() {
	//u := "https://www.clipartmax.com/png/small/21-219198_balloon-free-png-transparent-background-images-free-party-balloons-png.png"
	//u := "https://c0.klipartz.com/pngpicture/460/619/gratis-png-dibujo-corazon-stock-photography-corazon.png"
	//u := "https://blog.craigjoneswildlifephotography.co.uk/wp-content/uploads/2010/07/CMJ4893Jpeg-FB.jpg"
	//u := "https://raw.githubusercontent.com/Rapol/M101JS/master/project/mongomart/static/img/logo.jpg"
	//u := "https://i.pinimg.com/originals/65/ba/48/65ba488626025cff82f091336fbf94bb.gif"
	u := "https://cdn1.nuorder.com/product/fe2670dd16943a102cd52a68f6e67433.jpg"
	req, err := http.NewRequest("GET", u, nil)
	he(err)

	// GET image.
	client := &http.Client{}
	res, err := client.Do(req)
	he(err)
	defer res.Body.Close()

	// Read image twice.
	var body bytes.Buffer
	bodyReader := io.TeeReader(res.Body, &body)

	// Result reader.
	var r io.Reader
	// Write to file (Going to be 2nd read).
	r = &body

	// Detect format (1st read).
	img, fmtName, err := image.Decode(bodyReader)
	he(err)
	fmt.Printf("Format is: %v \n", fmtName)
	if fmtName == "gif" {
		buffer := bytes.NewBuffer(make([]byte, 0))
		err := png.Encode(buffer, img)
		he(err)
		// Overwrite result reader.
		r = buffer
		fmt.Printf("GIF converted to PNG \n")
	}

	// Write file (2nd read or read overwritten data).
	out, err := os.Create("/tmp/x.png")
	he(err)
	defer out.Close()
	io.Copy(out, r)
}

func strImgToFile(strImg string) {
	i := base64.NewDecoder(base64.StdEncoding, strings.NewReader(GopherImg))

	//_, fmtName, err := image.Decode(i)
	//he(err)
	//fmt.Printf("🎾 fmtName=%+v \n", fmtName)

	img, err := png.Decode(i)
	he(err)
	toPNG(img)
}

func decodeFile(imgFilePath string) image.Image {
	f, err := os.Open(imgFilePath)
	he(err)
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	he(err)
	fmt.Printf("🎾 ColorModel=%+v, fmtName=%+v \n", img.ColorModel(), fmtName)

	return img
}

func toGIF(img image.Image) {
	f, err := os.Create("/tmp/out.gif")
	he(err)
	defer f.Close()

	opt := gif.Options{
		NumColors: 256,
	}
	err = gif.Encode(f, img, &opt)
	he(err)
}

func toJPEG(img image.Image) {
	f, err := os.Create("/tmp/out.jpeg")
	he(err)
	defer f.Close()

	opt := jpeg.Options{
		Quality: 99,
	}
	err = jpeg.Encode(f, img, &opt)
	he(err)
}

func toPNG(img image.Image) {
	f, err := os.Create("/tmp/out.png")
	he(err)
	defer f.Close()

	enc := png.Encoder{
		CompressionLevel: png.DefaultCompression,
	}
	err = enc.Encode(f, img)
	he(err)
}
