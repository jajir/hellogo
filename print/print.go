package main

import (
	ev3dev "github.com/ev3go/ev3dev"
	"github.com/jajir/hellogo/lev3"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

var touchSensor lev3.TouchSensor

func main() {
	defer lev3.Display.Close()
	defer lev3.Speaker.Close()
	defer lev3.Lights.TurnOff()
	var errorChannel = make(chan string)
	go lev3.Lights.HeartBeat()

	touchSensor = lev3.NewTouchSensor("in1")

	if touchSensor.IsPressed() {
		touchSensor.WaitUntilReleased()
		calibratePen()
		touchSensor.WaitUntilReleased()
	}

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

	//	calibratePortal()
	//	calibrateFeeding()
	//	calibratePen()

	//	col := lev3.NewColorSensor("ev3-ports:in4")
	//	col.PrintInfo()
	//	col.TestAmbient()

	time.Sleep(31 * time.Minute)
}

func calibrateFeeding() {
	//kdyz je 1, pak ej papir
}

func calibratePortal() {
	var motor3 lev3.Ev3lmotor = lev3.NewEv3lmotor("lego-ev3-l-motor", "outC", "motorPortal")
	axisZ := lev3.NewAxis(&motor3)

	var wg sync.WaitGroup
	wg.Add(1)

	go axisZ.Init(&wg, float64(5))

	wg.Wait()

	axisZ.PrintInfo()
}

func calibratePen() int {
	/* kalibrace funguje pokud je pero ve
	 vytazene pozici, jinak vraci zkreslene hodnoty, default je 1250. */
	var motor lev3.Ev3lmotor = lev3.NewEv3lmotor("lego-ev3-m-motor", "outB", "motorPortal")

	var startPos, endPos int

	speed := motor.MaxSpeed() / 6

	startPos = motor.Position()
	motor.StartMotor(-speed)
	touchSensor.WaitUntilPressed()
	endPos = motor.Position()
	motor.StopMotor()
	var measure1 = endPos - startPos
	log.Printf("Measure 1 %d, start: %d. stop: %d", measure1, startPos, endPos)

	touchSensor.WaitUntilReleased()

	startPos = motor.Position()
	motor.StartMotor(speed)
	touchSensor.WaitUntilPressed()
	endPos = motor.Position()
	motor.StopMotor()
	var measure2 = endPos - startPos
	log.Printf("Measure 2 %d, start: %d. stop: %d", measure2, startPos, endPos)

	touchSensor.WaitUntilReleased()

	measure := int((math.Abs(float64(measure1)) + math.Abs(float64(measure2))*3) / 4)
	log.Printf("Measure results %d", measure)

	return measure
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
