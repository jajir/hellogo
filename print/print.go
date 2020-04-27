package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	"github.com/jajir/hellogo/lev3"
	"log"
	"strconv"
	"time"
)

func main() {
	defer lev3.Display.Close()
	defer lev3.Speaker.Close()
	defer lev3.Lights.TurnOff()
	var errorChannel = make(chan string)
	go lev3.Lights.HeartBeat()

	var touchSensor lev3.TouchSensor = lev3.NewTouchSensor("in1")
	go touchSensor.Watch(errorChannel)
	go run(errorChannel)

	k := <-errorChannel
	log.Printf("Shutting down, bacause of %s.\n", k)
	StopAllMotors()
	lev3.Speaker.Beep()
}

func run(errorChannel chan string) {
	lev3.Display.Clean()
	lev3.Display.Write(0, 10, "Starting ...")
	log.Println("Ahoj")
	
	col := lev3.NewColorSensor("ev3-ports:in4")
	col.PrintInfo()
	col.TestAmbient()
	
	time.Sleep(31 * time.Minute)
}

func StopAllMotors() {
	StopMotors("lego-ev3-m-motor")
	StopMotors("lego-ev3-l-motor")
}

func detectSensor() {
	var name string = "lego-ev3-color"
	for i := 1; i < 5; i++ {
		var port string = "ev3-ports:in" + strconv.Itoa(i)
		s, err := ev3dev.SensorFor(port, name)
		if err == nil {
			log.Printf("Na %s je senzor. %v", port, s.Type())
		} else {
			log.Printf("Na %s nic neni. %v", port, err)
		}
	}
}

func StopMotors(driver string) {
	var last *ev3dev.TachoMotor
	for {
		var tmp ev3dev.TachoMotor
		err := ev3dev.FindAfter(last, &tmp, driver)
		if err != nil {
			break
		}
		last = &tmp
		log.Printf("Stopping motor: %v", last)
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
