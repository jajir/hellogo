package ev3control

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Axis struct {
	name                                 string
	motor                                Ev3lmotor
	touchSensor                          TouchSensor
	position, positionStart, positionEnd int
}

func NewAxis(axisName string, motorPortName string, touchSensorPortName string) Axis {
	var axis = *new(Axis)
	axis.name = axisName
	axis.motor = NewEv3lmotor(motorPortName, axis.name)
	axis.touchSensor = NewTouchSensor(touchSensorPortName)
	axis.position = -1
	axis.positionStart = -1
	axis.positionEnd = -1
	return axis
}

func (axis *Axis) PrintInfo() {
	log.WithFields(log.Fields{
		"Name":             axis.name,
		"Start position":   axis.positionStart,
		"End position":     axis.positionEnd,
		"Current position": axis.position,
	}).Info("Information about Axis.")
}

func (axis *Axis) IsCalibrated() bool {
	return axis.positionStart != -1 && axis.positionEnd != -1
}

func (axis *Axis) AxisLength() (int, error) {
	if axis.positionStart == -1 && axis.positionEnd == -1 {
		return -1, fmt.Errorf("length of axis '%s' can't be determined because axis is not calibrated", axis.name)
	}
	return axis.positionEnd - axis.positionStart, nil
}

func (axis *Axis) moveSpeed() int {
	maxSpeed := axis.motor.MaxSpeed()
	moveSpeed := maxSpeed / 5
	return moveSpeed
}

func (axis *Axis) rollBackSpeed() int {
	maxSpeed := axis.motor.MaxSpeed()
	rollBackSpeed := maxSpeed / 10
	return rollBackSpeed
}

func (axis *Axis) Calibrate() error {
	if axis.touchSensor.IsPressed() {
		return fmt.Errorf("can't calibrate Axis '%s' because touch sensor is presed", axis.name)
	}
	// start motor in one direction
	axis.motor.SetNormalPolarity()
	moveUntilSensorIsTouched(*axis)
	axis.positionStart = axis.motor.Position()

	moveUntilSensorIsTouched(*axis)
	axis.positionEnd = axis.motor.Position()
	axis.position = axis.motor.Position()
	return nil
}

func moveUntilSensorIsTouched(axis Axis) {
	//measure oposit side of axis
	axis.motor.StartMotor(axis.moveSpeed())

	//wait for touch sensor signal and mark position
	axis.touchSensor.WaitUntilPressed()
	axis.motor.StopMotor()
	axis.motor.WaitUntilMotorStop()
	axis.motor.ReversePolarity() //change direction to release touch sensor
	axis.motor.StartMotor(axis.rollBackSpeed())
	axis.touchSensor.WaitUntilReleased()
	axis.motor.StopMotor()
}
