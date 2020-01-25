package main

import (
	ev3 "github.com/ev3go/ev3"
)

var Lights = NewLeds()

type Leds struct {
}

func NewLeds() Leds {
	return *new(Leds)
}

func (l *Leds) TurnOff() {
	ev3.GreenLeft.SetBrightness(0)
	ev3.GreenRight.SetBrightness(0)
	ev3.RedLeft.SetBrightness(0)
	ev3.RedRight.SetBrightness(0)
}
