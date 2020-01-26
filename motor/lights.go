package main

import (
	ev3 "github.com/ev3go/ev3"
	log "github.com/sirupsen/logrus"
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

func (l *Leds) GreenTurnOn() {
	max, err := ev3.GreenLeft.MaxBrightness()
	if err != nil {
		log.Fatalf("Unable to read max brightness of led: %v. There is err: %v", ev3.GreenLeft, err)
	}

	ev3.GreenLeft.SetBrightness(max / 2)
	ev3.GreenRight.SetBrightness(max / 2)
	ev3.RedLeft.SetBrightness(0)
	ev3.RedRight.SetBrightness(0)
}
