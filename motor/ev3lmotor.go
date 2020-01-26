package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
)

type Ev3lmotor struct {
	m *ev3dev.TachoMotor
}

func NewEv3lmotor(port string) Ev3lmotor {
	a, err := ev3dev.TachoMotorFor("ev3-ports:"+port, "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left motor in outA: %v", err)
	}
	return Ev3lmotor{a}
}

func (m *Ev3lmotor) PrintInfo() {
	a := m.m
	for _, command := range a.Commands() {
		log.Printf("Command    : %s", command)
	}
	for _, command := range a.StopActions() {
		log.Printf("Stop action: %s", command)
	}
	log.Printf("Driver     : %s", a.Driver())
	log.Printf("Type       : %s", a.Type())
	log.Printf("countPerRot: %d", a.CountPerRot())
	p, _ := a.Polarity()
	log.Printf("Polarity   : %s", p)
	pos, _ := a.Position()
	log.Printf("Position   : %d", pos)
	log.Printf("Max speed  : %d", a.MaxSpeed())
}

func (m *Ev3lmotor) MaxSpeed() int {
	return m.m.MaxSpeed()
}

func (m *Ev3lmotor) CurrentSpeed() (currentSpeed int) {
	currentSpeed, err := m.m.Speed()
	if err != nil {
		log.Fatalf("Unable to read motor speed: %v", err)
	}
	return
}

func (m *Ev3lmotor) StartMotor(speed int) {
	err := m.m.
		SetSpeedSetpoint(speed).
		Command("run-forever").
		Err()
	if err != nil {
		log.Fatalf("Motor doesn't started: %v", err)
	}
}

func (m *Ev3lmotor) StopMotor() {
	err := m.m.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor stopping failed: %v", err)
	}
}

func (m *Ev3lmotor) ReversePolarity() {
	a := m.m
	polarity, err := a.Polarity()
	if err != nil {
		log.Fatalf("Unable to read polarity: %v", err)
	}
	if polarity == ev3dev.Normal {
		a.SetPolarity(ev3dev.Inversed)
	}
	if polarity == ev3dev.Inversed {
		a.SetPolarity(ev3dev.Normal)
	}
}

func (m *Ev3lmotor) SetPolarity(polarity ev3dev.Polarity) {
	err := m.m.SetPolarity(polarity).Err()
	if err != nil {
		log.Fatalf("Set polarity failed: %v", err)
	}
}

func (m *Ev3lmotor) Position() int {
	pos, err := m.m.Position()
	if err != nil {
		log.Fatalf("Motor start posion reading failed: %v", err)
	}
	return pos
}
