package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	ev3control "github.com/jajir/hellogo/example/ev3control"
	log "github.com/sirupsen/logrus"
)

func main() {

	var konec = make(chan string)
	ev3control.Display.Clean()
	ev3control.Display.Write(0, 10, "Caligration started")
	go waitingForBack(konec)

	motor := ev3control.NewEv3lmotor("outA", "axisX")
	motor.PrintInfo()
	go calibrate(motor)

	k := <-konec
	log.Info("Shutting down, bacause of " + k + "\n")
	ev3control.Speaker.Beep()
	motor.StopMotor()
	ev3control.Display.Clean()
	ev3control.Display.Write(0, 10, "Back button was pressed.")
	ev3control.Display.Write(0, 25, "Program is terminated.")
	ev3control.Display.Close()
}

func waitingForBack(konec chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		if e.Button == 1 {
			konec <- "Button Back was pressed."
		}
	}
}

func calibrate(motor ev3control.Ev3lmotor) {
	maxSpeed := motor.MaxSpeed()
	moveSpeed := maxSpeed / 5
	rollBackSpeed := maxSpeed / 10
	touchSensor := ev3control.NewTouchSensor("in1")
	touchSensor.PrintInfo()

	// start motor in one direction
	motor.SetNormalPolarity()

	motor.StartMotor(moveSpeed)

	//wait for touch sensor signal and mark position
	touchSensor.WaitUntilPressed()
	motor.StopMotor()
	motor.WaitUntilMotorStop()
	motor.ReversePolarity()
	motor.StartMotor(rollBackSpeed)
	touchSensor.WaitUntilReleased()
	positionStart := motor.Position()

	//measure oposit side of axis
	motor.StartMotor(moveSpeed)

	//wait for touch sensor signal and mark position
	touchSensor.WaitUntilPressed()
	motor.StopMotor()
	motor.WaitUntilMotorStop()
	motor.ReversePolarity()
	motor.StartMotor(rollBackSpeed)
	touchSensor.WaitUntilReleased()
	positionEnd := motor.Position()

	motor.StopMotor()

	//positions and difference between them are definition of robot axis
	log.WithFields(log.Fields{
		"Start position": positionStart,
		"End position":   positionEnd,
	}).Info("Calibratio is done.")
}
