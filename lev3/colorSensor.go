package lev3

import (
	ev3dev "github.com/ev3go/ev3dev"
	"log"
	"time"
)

type ColorSensor struct {
	s      *ev3dev.Sensor
	status int
}

func NewColorSensor(port string) ColorSensor {
	s, err := ev3dev.SensorFor(port, "lego-ev3-color")
	if err != nil {
		log.Fatalf("failed to find touch sensor: %v", err)
	}
	return ColorSensor{s, 0}
}

func (ts *ColorSensor) PrintInfo() {
	s := ts.s
	log.Printf("Type           : %s", s.Type())
	log.Printf("Driver         : %s", s.Driver())
	log.Printf("BinDataFormat  : %s", s.BinDataFormat())
	for _, command := range s.Commands() {
		log.Printf("Command       : %s", command)
	}
	log.Printf("FirmwareVersion: %s", s.FirmwareVersion())
	for _, mode := range s.Modes() {
		log.Printf("Modes         : %s", mode)
	}
}

func (ts *ColorSensor) TestAmbient() {
	s := ts.s
	s.SetMode("COL-AMBIENT")
	for {
		n := s.NumValues()
		for i := 0; i < n; i++ {
			v, err := s.Value(i)
			if err != nil {
				log.Fatalf("failed to get of value %d: %v", i, err)
			}
			log.Printf("Number of values: %d, cx: %d, value '%s'", n, i, v)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
