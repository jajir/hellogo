package main

import (
	"fmt"
	ev3dev "github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {

	var errorChannel = make(chan string)

	go waitingForBack(errorChannel)
	go run(errorChannel)

	k := <-errorChannel
	log.Info("Shutting down, bacause of " + k + "\n")
	Lights.TurnOff()
	StopAllMotors()
	Display.Close()
}

func waitingForBack(errorChannel chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		fmt.Printf("%+v\n", e)
		if 1 == e.Button {
			fmt.Printf("Zpet")
			errorChannel <- "Button Back was pressed."
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

func StopMotor(a *ev3dev.TachoMotor) {
	err := a.
		SetStopAction("coast").
		Command("stop").
		Err()
	if err != nil {
		log.Fatalf("Motor stopping failed: %v", err)
	}
}

func run(errorChannel chan string) {
	Lights.GreenTurnOn()
	Display.Clean()
	Display.Write(0, 10, "Starting ...")

	var touchSensor TouchSensor = NewTouchSensor("in1")
	touchSensor.PrintInfo()
	go touchSensor.watch(errorChannel)

	//	var motor1 Ev3lmotor = NewEv3lmotor("lego-ev3-l-motor", "outA")
	//	axisX := NewAxis(&motor1)
	//	axisX.Init(float64(10))
	//	axisX.PrintInfo()
	//
	//	var motor2 Ev3lmotor = NewEv3lmotor("lego-ev3-l-motor", "outC")
	//	axisY := NewAxis(&motor2)
	//	axisY.Init(float64(7))
	//	axisY.PrintInfo()
	//
	//	var motor3 Ev3lmotor = NewEv3lmotor("lego-ev3-m-motor", "outB")
	//	axisZ := NewAxis(&motor3)
	//	axisZ.Init(float64(11))
	//	axisZ.PrintInfo()

	//Time to write previous function to console
	time.Sleep(30 * time.Millisecond)
}

//func runForAWhile(a *ev3dev.TachoMotor) {
//	err := a.
//		SetRampUpSetpoint(200 * time.Millisecond).
//		SetRampDownSetpoint(200 * time.Millisecond).
//		SetPosition(0).
//		Err()
//	if err != nil {
//		log.Fatalf("Motor failed: %v", err)
//	}
//
//	max := a.MaxSpeed() / 4
//
//	err = a.
//		SetSpeedSetpoint(max).
//		Command("run-forever").
//		Err()
//	if err != nil {
//		log.Fatalf("Motor failed: %v", err)
//	}
//
//	time.Sleep(time.Second * 2)
//
//	err = a.
//		SetStopAction("coast").
//		Command("stop").
//		Err()
//	if err != nil {
//		log.Fatalf("Motor failed: %v", err)
//	}
//}

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
