package lev3

import (
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"time"
)

type TouchSensor struct {
	//FIXME rename to sensor
	s      *ev3dev.Sensor
	status int
}

func NewTouchSensor(port string) TouchSensor {
	s, err := ev3dev.SensorFor("ev3-ports:"+port, "lego-ev3-touch")
	if err != nil {
		log.Fatalf("failed to find touch sensor: %v", err)
	}
	if s.NumValues() != 1 {
		log.Fatalf("Touch sensor return unexpected number of values (%v)", s.NumValues())
	}
	return TouchSensor{s, 0}
}

func (ts *TouchSensor) PrintInfo() {
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

func (ts *TouchSensor) Watch(eventChanel chan string) {
	for {
		time.Sleep(10 * time.Millisecond)
		s := ts.s
		v, err := s.Value(0)
		if err != nil {
			log.Fatalf("failed to get of value from touch sensor. Err: %v", err)
		}
		if v == "0" {
			if ts.status == 1 {
				ts.status = 0
				eventChanel <- "Touch sensor was released."
			}
		}
		if v == "1" {
			if ts.status == 0 {
				ts.status = 1
			}
		}
	}
}

func (ts *TouchSensor) IsPressed() bool {
	sensor := ts.s
	val, err := sensor.Value(0)
	if err != nil {
		log.Fatalf("failed to get of value from touch sensor. Err: %v", err)
	}
	return val == "1"
}

func (ts *TouchSensor) WaitUntilPressed() {
	sensor := ts.s
	var val string
	for val != "1" {
		time.Sleep(10 * time.Millisecond)
		var err error
		val, err = sensor.Value(0)
		if err != nil {
			log.Fatalf("failed to get of value from touch sensor. Err: %v", err)
		}
	}
}

func (ts *TouchSensor) WaitUntilReleased() {
	sensor := ts.s
	var val string
	for val != "0" {
		time.Sleep(10 * time.Millisecond)
		var err error
		val, err = sensor.Value(0)
		if err != nil {
			log.Fatalf("failed to get of value from touch sensor. Err: %v", err)
		}
	}
}
