package lev3

import (
	"image"
	"image/color"
	"image/draw"

	ev3 "github.com/ev3go/ev3"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

/**
User should work with this variable.
*/
var Display = NewScreen()

type Screen struct {
}

func NewScreen() Screen {
	ev3.LCD.Init(true)
	return *new(Screen)
}

func (s *Screen) Close() {
	ev3.LCD.Close()
}

func (s *Screen) Clean() {
	draw.Draw(ev3.LCD, ev3.LCD.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
}

func (s *Screen) Write(x, y int, label string) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  ev3.LCD,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
