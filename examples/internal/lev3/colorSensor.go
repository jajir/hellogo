package lev3

import (
	ev3dev "github.com/ev3go/ev3dev"
	"log"
	"strconv"
	"time"
)

type ColorSensor struct {
	sensor *ev3dev.Sensor
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
	s := ts.sensor
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
	v, _ := s.Mode()
	log.Printf("Current Mode   : %s", v)
}

func (ts *ColorSensor) TestReflected() {
	ts.sensor.SetMode("COL-AMBIENT")
	for {
		v := ts.ReadValue()
		log.Printf("Reading color sensor value '%d'", v)
		time.Sleep(10 * time.Millisecond)
	}
}

func (ts *ColorSensor) ReadValue() int {
	v, err := ts.sensor.Value(0)
	if err != nil {
		log.Fatalf("Unable to read value from color sensor, bacause of: %v", err)
	}
	out, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("Unable to converts value '%s' to int", v)
	}
	return out
}

func (ts *ColorSensor) IsCovered() bool {
	i1 := ts.ReadValue()
	time.Sleep(10 * time.Millisecond)
	i2 := ts.ReadValue()
	time.Sleep(10 * time.Millisecond)
	i3 := ts.ReadValue()
	f := (float32(i1) + float32(i2) + float32(i3)) / 3.0
//	log.Printf("Reading color sensor value %v, %v, %v, avg: %f", i1, i2, i3, f)
	return f > 0.3
}
