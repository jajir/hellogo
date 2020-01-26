package main

import (
	"fmt"
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

func main() {

	var konec = make(chan string)

	go waitingForBack(konec)
	go run(konec)

	k := <-konec
	log.Info("Shutting down, bacause of " + k + "\n")
	Lights.TurnOff()
	StopAllMotors()
	Display.Close()
}

func waitingForBack(konec chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		fmt.Printf("%+v\n", e)
		if 1 == e.Button {
			fmt.Printf("Zpet")
			konec <- "Button Back was pressed."
		}
	}
}

func StopAllMotors() {
	var last *ev3dev.TachoMotor
	for {
		var tmp ev3dev.TachoMotor
		err := ev3dev.FindAfter(last, &tmp, "lego-ev3-l-motor")
		if err != nil {
			break
		}
		last = &tmp
		fmt.Printf("Stopping motor: %v", last)
		StopMotor(last)
	}
}

func run(konec chan string) {
	Lights.GreenTurnOn()
	Display.Clean()
	Display.Write(0, 10, "Ahoj lidi")

	a, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left motor in outA: %v", err)
	}

	PrintInfo(a)

	//Time to write previous function to console
	time.Sleep(30 * time.Millisecond)

	//	go PrintStats(a)

	//	runForAWhile(a)

	max := runUntilObstacle(a)
	ReversePolarity(a)
	min := runUntilObstacle(a)

	log.Printf("Motor rande is: (%d, %d)", min, max)

	time.Sleep(30 * time.Millisecond)
}

func PrintStats(a *ev3dev.TachoMotor) {
	Display.Clean()
	for {
		astat, _ := a.State()
		aspeed, _ := a.Speed()

		msg := fmt.Sprintf("outA: %s %d", astat, aspeed)
		log.Printf(msg)

		Display.Write(10, 10, msg)
		time.Sleep(10 * time.Millisecond)
	}
}

func PrintInfo(a *ev3dev.TachoMotor) {
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

func runUntilObstacle(a *ev3dev.TachoMotor) int {
	startPos, err := a.Position()
	if err != nil {
		log.Fatalf("Motor start posion reading failed: %v", err)
	}

	speed := a.MaxSpeed() / 6

	StartMotor(a, speed)
	time.Sleep(500 * time.Millisecond)

	currentSpeed := CurrentSpeed(a)

	log.Printf("1) Target speed %d, current speed %d", speed, currentSpeed)

	for count(speed, currentSpeed) {
		currentSpeed = CurrentSpeed(a)
		log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
		time.Sleep(10 * time.Millisecond)
	}

	finalPos, err := a.Position()
	if err != nil {
		log.Fatalf("Motor final posion reading failed: %v", err)
	}

	StopMotor(a)

	return finalPos - startPos
}

func count(speed, currentSpeed int) bool {
	log.Printf("Target speed %d, current speed %d", speed, currentSpeed)
	return math.Abs(float64(speed-currentSpeed)) < float64(speed)/100*10
}

func CurrentSpeed(a *ev3dev.TachoMotor) (currentSpeed int) {
	currentSpeed, err := a.Speed()
	if err != nil {
		log.Fatalf("Unable to read motor speed: %v", err)
	}
	return
}

func StartMotor(a *ev3dev.TachoMotor, speed int) {
	err := a.
		SetSpeedSetpoint(speed).
		Command("run-forever").
		Err()
	if err != nil {
		log.Fatalf("Motor doesn't started: %v", err)
	}
}

func StopMotor(a *ev3dev.TachoMotor) {
	err := a.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor stopping failed: %v", err)
	}
}

func ReversePolarity(a *ev3dev.TachoMotor) {
	polarity, err := a.Polarity()
	if err != nil {
		log.Fatalf("Unable to read polarity: %v", err)
	}
	if polarity == ev3dev.Normal {
		SetPolarity(a, ev3dev.Inversed)
	}
	if polarity == ev3dev.Inversed {
		SetPolarity(a, ev3dev.Normal)
	}
}

func SetPolarity(a *ev3dev.TachoMotor, polarity ev3dev.Polarity) {
	err := a.SetPolarity(polarity).Err()
	if err != nil {
		log.Fatalf("Set polarity failed: %v", err)
	}
}

func runForAWhile(a *ev3dev.TachoMotor) {
	err := a.
		SetRampUpSetpoint(200 * time.Millisecond).
		SetRampDownSetpoint(200 * time.Millisecond).
		SetPosition(0).
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}

	max := a.MaxSpeed() / 4

	err = a.
		SetSpeedSetpoint(max).
		Command("run-forever").
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}

	time.Sleep(time.Second * 2)

	err = a.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}
}
