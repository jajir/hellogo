package main

import (
	"fmt"

	ev3dev "github.com/ev3go/ev3dev"
	ev3control "github.com/jajir/hellogo/example/ev3control"
	log "github.com/sirupsen/logrus"
)

func main() {

	var konec = make(chan string)
	ev3control.Display.Clean()
	ev3control.Display.Write(0, 10, "Hello world!")
	ev3control.Display.Write(0, 25, "To terminate programm")
	ev3control.Display.Write(0, 40, " press Bact on EV3.")

	go waitingForBack(konec)

	k := <-konec
	log.Info("Shutting down, bacause of " + k + "\n")
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
		fmt.Printf("%+v\n", e)
		if e.Button == 1 {
			fmt.Printf("Zpet")
			konec <- "Button Back was pressed."
		}
	}
}
