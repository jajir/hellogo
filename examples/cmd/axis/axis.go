package main

/*
TODO:
Pri startu nacit stav led a pri konci ho obnovit.
Tlacitka na kostce spusti kalibraci.
NOTES:
Color sensor doesn't work correctly on dayligh.
*/
import (
	"fmt"
	"log"
	"sync"

	ev3dev "github.com/ev3go/ev3dev"
	"github.com/jajir/hellogo/internal/lev3"
)

var touchSensor lev3.TouchSensor
var motor lev3.Ev3lmotor

func main() {
	defer shutdownAll()
	var errorChannel = make(chan string)
	go lev3.Lights.HeartBeat()
	lev3.Speaker.Beep()
	lev3.Display.Clean()
	printInfo()

	touchSensor = lev3.NewTouchSensor("in1")
	motor = lev3.NewEv3lmotor("lego-ev3-l-motor", "outA", "motorPortal")
	go touchSensor.Watch(errorChannel)

	go run(errorChannel)

	k := <-errorChannel
	log.Printf("Shutting down, because of %s.\n", k)
	lev3.Speaker.Beep()
}

func shutdownAll() {
	lev3.Display.Clean()
	lev3.Display.Close()
	lev3.Speaker.Close()
	lev3.Lights.TurnOff()
	motor.StopMotor()
}

const rowHeight int = 10

var row int = 0

func printInfo() {
	lev3.Display.Clean()
	row = 10
	line("How to use:")
	line("Touch sensor: immediately")
	line("    quit program")
	line("Left arrow: print")
	line("    rectangle.")
	line("Right arrow: calibre pen")
	line("    up and down.")
	line("Up arrow: calibre paper")
	line("    feeding.")
	line("Down arrow: calibre")
	line("    moving pen from left")
	line("    to right.")
}

func line(text string) {
	lev3.Display.Write(0, row, text)
	row += rowHeight
}

func run(errorChannel chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		// fmt.Printf("%+v\n", e)
		if e.Value == 0 {
			if e.Button == 1 {
				fmt.Printf("It's Back\n")
			}
			if e.Button == 2 {
				fmt.Printf("It's Left\n")
			}
			if e.Button == 4 {
				fmt.Printf("It's Central\n")
			}
			if e.Button == 8 {
				fmt.Printf("It's Right\n")
			}
			if e.Button == 16 {
				fmt.Printf("It's Up\n")
			}
			if e.Button == 32 {
				fmt.Printf("It's Down\n")
				calibratePortal()
			}
		}
	}
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
