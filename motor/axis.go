package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
	"sync"
)

type Axis struct {
	min, max int
	motor    *Ev3lmotor
}

func NewAxis(motor *Ev3lmotor) Axis {
	axis := Axis{0, 0, motor}
	return axis
}

func (axis *Axis) Init(wg *sync.WaitGroup, sensitivity float64) {
	defer wg.Done()
	axis.motor.SetPolarity(ev3dev.Normal)
	axis.runUntilObstacle(sensitivity)
	axis.motor.ReversePolarity()
	axis.min, axis.max = axis.runUntilObstacle(sensitivity)
	axis.min = -axis.min
	axis.max = -axis.max
	//	log.Printf("min & max [%d, %d], [%d, %d]", min, max, axis.min, axis.max)
	axis.motor.SetPolarity(ev3dev.Normal)
}

func (axis *Axis) runUntilObstacle(sensitivity float64) (startPos, finalPos int) {
	motor := axis.motor
	startPos = motor.Position()

	speed := motor.MaxSpeed() / 6

	motor.StartMotor(speed)

	time.Sleep(500 * time.Millisecond)

	currentSpeed := motor.CurrentSpeed()

	for count(speed, currentSpeed, sensitivity) {
		currentSpeed = motor.CurrentSpeed()
		//		log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
		time.Sleep(10 * time.Millisecond)
	}

	finalPos = motor.Position()

	motor.StopMotor()

	return startPos, finalPos
}

func count(speed, currentSpeed int, sensitivity float64) bool {
	//	log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
	return math.Abs(float64(speed-currentSpeed)) < float64(speed)/100*sensitivity
}

func (axis *Axis) Range() int {
	return axis.max - axis.min
}

func (axis *Axis) PrintInfo() {
	log.Printf("Motor %v {min=%d, max=%d, range=%d}",axis.motor.name, axis.min, axis.max, axis.Range())
}
