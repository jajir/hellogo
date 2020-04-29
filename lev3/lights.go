package lev3

import (
	ev3 "github.com/ev3go/ev3"
	"log"
	"time"
	"math"
)

var Lights = NewLeds()

type Leds struct {
	running bool
}

func NewLeds() Leds {
	var out = *new(Leds)
	out.running = false
	return *new(Leds)
}

func (l *Leds) TurnOff() {
	l.running = false
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

func (l *Leds) HeartBeat() {
	l.running = true
	var max = getMaxBrightness() / 2
	var step float64 = 0
	var increment float64 = 0.05
	for l.running == true {
		val := int(math.Abs(math.Sin(step)) * float64(max))
		step+=increment
		ev3.RedLeft.SetBrightness(val)
		ev3.RedRight.SetBrightness(val)
		time.Sleep(30 * time.Millisecond)
	}
}

func getMaxBrightness() int {
	max, err := ev3.GreenLeft.MaxBrightness()
	if err != nil {
		log.Fatalf("Unable to read max brightness of led: %v. There is err: %v", ev3.GreenLeft, err)
	}
	return max
}

