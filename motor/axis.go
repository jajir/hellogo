package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type Axis struct {
	min, max int
	motor    *Ev3lmotor
}

func NewAxis(motor *Ev3lmotor) Axis {
	axis := Axis{0, 0, motor}
	return axis
}

func (axis *Axis) Init() {
	axis.motor.SetPolarity(ev3dev.Normal)
	max, min := axis.runUntilObstacle()
	axis.motor.ReversePolarity()
	axis.min, axis.max = axis.runUntilObstacle()
	axis.min = -axis.min
	axis.max = -axis.max
	//	log.Printf("min & max [%d, %d], [%d, %d]", min, max, axis.min, axis.max)
	axis.motor.SetPolarity(ev3dev.Normal)
}

func (axis *Axis) runUntilObstacle() (startPos, finalPos int) {
	motor := axis.motor
	startPos = motor.Position()

	speed := motor.MaxSpeed() / 6

	motor.StartMotor(speed)

	time.Sleep(500 * time.Millisecond)

	currentSpeed := motor.CurrentSpeed()

	for count(speed, currentSpeed) {
		currentSpeed = motor.CurrentSpeed()
		//		log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
		time.Sleep(10 * time.Millisecond)
	}

	finalPos = motor.Position()

	motor.StopMotor()

	return startPos, finalPos
}

func count(speed, currentSpeed int) bool {
	//	log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
	return math.Abs(float64(speed-currentSpeed)) < float64(speed)/100*10
}

func (axis *Axis) Range() int {
	return axis.max - axis.min
}

func (axis *Axis) PrintInfo() {
	log.Printf("Motor {min=%d, max=%d, range=%d}", axis.min, axis.max, axis.Range())
}
