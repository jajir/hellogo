package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		maxX    = 1000.0
		maxY    = 1000.0
		cyclesX = 3.0 // kolil vrcholu to bude mit do stran
		cyclesY = 5.0  //kolil vrcholu to bude mit nahoru a dolu
	)

	var img = image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	var red = color.RGBA{255, 0, 0, 255} // Red

	step := 0.0001

	for c := 0.0; c < math.Pi*cyclesX*cyclesY; c += step {
		x := math.Cos(cyclesX*c)*maxX/2 + maxX/2
		y := math.Sin(cyclesY*c)*maxY/2 + maxY/2
		img.Set(int(x), int(y), red)
	}

	png.Encode(out, img) // NOTE: ignoring encoding errors
}
