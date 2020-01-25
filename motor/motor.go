package main

import (
	"fmt"
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {

	var konec = make(chan string)

	go waitingForBack(konec)
	go run(konec)

	k := <-konec
	log.Info("Shutting down, bacause of " + k + "\n")
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

func run(konec chan string) {
	Lights.TurnOff()
	Display.Clean()
	Display.Write(0, 10, "Ahoj lidi")

	a, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left motor in outA: %v", err)
	}

	for _, command := range a.Commands() {
		log.Printf("Command: %s", command)
	}

	err = a.
		SetRampUpSetpoint(200 * time.Millisecond).
		SetRampDownSetpoint(200 * time.Millisecond).
		SetPosition(0).
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}

	max := a.MaxSpeed() / 2

	err = a.
		SetSpeedSetpoint(max).
		Command("run-forever").
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}

	time.Sleep(time.Second / 4)

	err = a.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor failed: %v", err)
	}

	astat, _ := a.State()
	aspeed, _ := a.Speed()
	log.Printf("outA: %s %d", astat, aspeed)

	time.Sleep(30 * time.Second)
}
