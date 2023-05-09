package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	ev3control "github.com/jajir/hellogo/example/ev3control"
	readsvg "github.com/jajir/hellogo/example/readsvg"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Startint of printing")
	p := readsvg.Point{0, 0}
	var l readsvg.Line
	l = readsvg.NewLine{p, []readsvg.Point{{10, 0}, {10, 10}, {0, 10}, {0, 0}}}
	log.Info("Line " + l.String() + "")
	ev3control.Display.Write(0, 10, "Caligration started")
	ev3control.Display.Clean()
	var konec = make(chan string)
	go waitingForBack(konec)

	axis := ev3control.NewAxis("axisX", "outA", "in1")
	go initializeAxis(&axis)

	k := <-konec
	log.Info("Shutting down, bacause of " + k + "\n")
	ev3control.Speaker.Beep()
	ev3control.Display.Clean()
	ev3control.Display.Write(0, 10, "Back button was pressed.")
	ev3control.Display.Write(0, 25, "Program is terminated.")
	ev3control.Display.Close()
}

func initializeAxis(axis *ev3control.Axis) {
	err := axis.Calibrate()
	if err != nil {
		log.Fatalf("Unable to initialize: %v", err)
	}
	axis.PrintInfo()
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
