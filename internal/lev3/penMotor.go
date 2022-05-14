package lev3

import (
	"time"

	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
)

const PEN_UP int = 0
const PEN_DOWN int = 1
const MOVE_LENGTH int = 1200

type PenMotor struct {
	penMotor *ev3dev.TachoMotor
	name     string
	status   int
	speed    int
}

func NewPenMotor(motorType, port, name string) PenMotor {
	motor, err := ev3dev.TachoMotorFor("ev3-ports:"+port, motorType)
	if err != nil {
		log.Fatalf("failed to find motor in port %s: %v", port, err)
	}
	p := new(PenMotor)
	p.status = PEN_DOWN
	p.penMotor = motor
	p.name = name
	p.speed = motor.MaxSpeed() / 6
	return *p
}

func (m *PenMotor) PrintInfo() {
	a := m.penMotor
	for _, command := range a.Commands() {
		log.Printf(m.name+" Command    : %s", command)
	}
	for _, command := range a.StopActions() {
		log.Printf(m.name+" Stop action: %s", command)
	}
	log.Printf(m.name+" Driver     : %s", a.Driver())
	log.Printf(m.name+" Type       : %s", a.Type())
	log.Printf(m.name+" countPerRot: %d", a.CountPerRot())
	p, _ := a.Polarity()
	log.Printf(m.name+" Polarity   : %s", p)
	pos, _ := a.Position()
	log.Printf(m.name+" Position   : %d", pos)
	log.Printf(m.name+" Max speed  : %d", a.MaxSpeed())
}

func (m *PenMotor) Up() {
	if m.status == PEN_DOWN {
		log.Printf("Moving Up")
		motor := m.penMotor
		motor.SetSpeedSetpoint(m.speed)
		motor.SetPositionSetpoint(MOVE_LENGTH)
		motor.SetStopAction("hold")
		motor.Command("run-to-rel-pos")
		ev3dev.Wait(motor, ev3dev.Running, 0, 0, false, 20*time.Second)
		m.status = PEN_UP
	}
}

func (m *PenMotor) Down() {
	if m.status == PEN_UP {
		log.Printf("Moving Down")
		motor := m.penMotor
		motor.SetSpeedSetpoint(m.speed)
		motor.SetPositionSetpoint(-MOVE_LENGTH)
		motor.SetStopAction("hold")
		motor.Command("run-to-rel-pos")
		ev3dev.Wait(motor, ev3dev.Running, 0, 0, false, 20*time.Second)
		m.status = PEN_DOWN
	}
}

func (m *PenMotor) StopMotor() {
	err := m.penMotor.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor stopping failed: %v", err)
	}
}
