package main

/*
TODO:
Pri startu nacit stav led a pri konci ho obnovit.
Tlacitka na kostce spusti kalibraci.
NOTES:
Color sensor doesn't work correctly on dayligh.
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sync"
	"time"

	ev3dev "github.com/ev3go/ev3dev"
	"github.com/jajir/hellogo/internal/lev3"
)

var touchSensor lev3.TouchSensor
var motorPortal, motorFeeder lev3.Ev3lmotor
var penMotor lev3.PenMotor

func main() {
	defer shutdownAll()
	var errorChannel = make(chan string)
	go lev3.Lights.HeartBeat()
	lev3.Speaker.Beep()
	lev3.Display.Clean()
	printInfo()

	touchSensor = lev3.NewTouchSensor("in1")
	penMotor = lev3.NewPenMotor("lego-ev3-m-motor", "outB", "motorPortal")
	motorFeeder = lev3.NewEv3lmotor("lego-ev3-l-motor", "outD", "motorPortal")
	motorPortal = lev3.NewEv3lmotor("lego-ev3-l-motor", "outC", "motorPortal")
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
	penMotor.StopMotor()
	motorPortal.StopMotor()
	motorFeeder.StopMotor()
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
				printRectangle()
			}
			if e.Button == 4 {
				fmt.Printf("It's Central\n")
			}
			if e.Button == 8 {
				fmt.Printf("It's Right\n")
				calibratePen()
			}
			if e.Button == 16 {
				fmt.Printf("It's Up\n")
				calibrePaper()
			}
			if e.Button == 32 {
				fmt.Printf("It's Down\n")
				calibratePortal()
			}
		}
	}
}

func drawFile(errorChannel chan string, file string) {
	lev3.Display.Clean()
	lev3.Display.Write(0, 10, "Budu malovat podle souboru ...")
	log.Println("Budu malovat podle souboru")

	dat, err := ioutil.ReadFile("pok.txt")
	if err != nil {
		log.Fatalf("Unable to open file. %v\n", err)
	}
	var pic lev3.Picture
	err = json.Unmarshal(dat, &pic)
	if err != nil {
		log.Fatalf("Unable to unmarshal data. %v\n", err)
	}
	draw(pic)
}

func draw(pic lev3.Drawable) {
	maxSpeed := motorPortal.MaxSpeed()
	log.Printf("step 0\n")

	log.Printf("step 1\n")

	//pen is at [500,0] move it to start position.
	startPoint := pic.GetStartPoint()
	motorPortal.Turn(maxSpeed, startPoint.GetX()-500)
	motorFeeder.Turn(maxSpeed, -startPoint.GetY())
	penMotor.Down()
	log.Printf("step 2\n")

	for i := 0; i < pic.GetStepsCount(); i++ {
		diff := pic.GetStepDiff(i)
		motorX := motorPortal.GetMotor()
		motorY := motorFeeder.GetMotor()
		if diff.GetX() > diff.GetY() {
			//x is maxSpeed, y speed will be counted
			speedX := maxSpeed
			speedY := int(math.Round(float64(diff.GetY()) / float64(diff.GetX()) * float64(maxSpeed)))

			motorX.SetSpeedSetpoint(speedX)
			motorX.SetPositionSetpoint(diff.GetX())
			motorX.SetStopAction("hold")
			motorX.Command("run-to-rel-pos")

			motorY.SetSpeedSetpoint(speedY)
			motorY.SetPositionSetpoint(diff.GetY())
			motorY.SetStopAction("hold")
			motorY.Command("run-to-rel-pos")
		} else {
			//y is maxSpeed, x speed will be counted
			speedX := int(math.Round(float64(diff.GetX()) / float64(diff.GetY()) * float64(maxSpeed)))
			speedY := maxSpeed

			motorX.SetSpeedSetpoint(speedX)
			motorX.SetPositionSetpoint(diff.GetX())
			motorX.SetStopAction("hold")
			motorX.Command("run-to-rel-pos")

			motorY.SetSpeedSetpoint(speedY)
			motorY.SetPositionSetpoint(diff.GetY())
			motorY.SetStopAction("hold")
			motorY.Command("run-to-rel-pos")
		}
		//waint until motors runs to selected place
		ev3dev.Wait(motorX, ev3dev.Running, 0, 0, false, 20*time.Second)
		ev3dev.Wait(motorY, ev3dev.Running, 0, 0, false, 20*time.Second)
	}

	penMotor.Up()
}

func printRectangle() {
	speed := motorPortal.MaxSpeed() / 6
	penMotor.Down()

	pageHeight := 1000
	pageWidth := 1000

	//page width 1074, it's portal
	motorPortal.Turn(speed, pageWidth/2)

	motorFeeder.Turn(speed, -pageHeight)

	motorPortal.Turn(speed, -pageWidth)

	motorFeeder.Turn(speed, pageHeight)

	motorPortal.Turn(speed, pageWidth/2)

	penMotor.Up()
}

func calibrePaper() {
	//kdyz je 1, pak ej papir
	col := lev3.NewColorSensor("ev3-ports:in4")
	col.PrintInfo()

	//	for {
	//		log.Printf("Is covered: %v", col.IsCovered())
	//	}

	if col.IsCovered() {
		log.Printf("Is covered")
		motorFeeder.StartMotor(100)
		for col.IsCovered() {
		}
		motorFeeder.StopMotor()
	} else {
		log.Printf("Is not covered")
		motorFeeder.StartMotor(-100)
		for !col.IsCovered() {
		}
		motorFeeder.StopMotor()
	}
	//	col.TestReflected()
	log.Printf("Paper is ready")
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

func calibratePen() {
	penMotor.Up()
	time.Sleep(15 * time.Second)
	penMotor.Down()
}
