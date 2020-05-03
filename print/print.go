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
	ev3dev "github.com/ev3go/ev3dev"
	"github.com/jajir/hellogo/lev3"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

var touchSensor lev3.TouchSensor
var motorPen, motorPortal, motorFeeder lev3.Ev3lmotor

func main() {
	defer lev3.Display.Close()
	defer lev3.Speaker.Close()
	defer lev3.Lights.TurnOff()
	var errorChannel = make(chan string)
	go lev3.Lights.HeartBeat()

	touchSensor = lev3.NewTouchSensor("in1")
	motorPen = lev3.NewEv3lmotor("lego-ev3-m-motor", "outB", "motorPortal")
	motorFeeder = lev3.NewEv3lmotor("lego-ev3-l-motor", "outD", "motorPortal")
	motorPortal = lev3.NewEv3lmotor("lego-ev3-l-motor", "outC", "motorPortal")

	if touchSensor.IsPressed() {
		touchSensor.WaitUntilReleased()
		calibratePen()
		touchSensor.WaitUntilReleased()
	}

	go touchSensor.Watch(errorChannel)

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		go drawFile(errorChannel, argsWithoutProg[0])
	} else {
		go run(errorChannel)
	}

	k := <-errorChannel
	log.Printf("Shutting down, bacause of %s.\n", k)
	StopAllMotors()
	lev3.Speaker.Beep()
}

func run(errorChannel chan string) {
	lev3.Display.Clean()
	lev3.Display.Write(0, 10, "Starting ...")
	log.Println("Ahoj")
	preparePen()

	drawRectangle()
	//	calibratePortal()
	//	drawRectangle()
	//	calibratePortal()
	//	preparePaper()
	//	calibratePen()

	time.Sleep(31 * time.Minute)
}

func drawFile(errorChannel chan string, file string) {
	lev3.Display.Clean()
	lev3.Display.Write(0, 10, "Budu malovat podle souboru ...")
	log.Println("Budu malovat podle souboru")
	preparePen()

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

func draw(pic lev3.Drawable){
		maxSpeed := motorPen.MaxSpeed()
	log.Printf("step 0\n")

	log.Printf("step 1\n")

	//pen is at [500,0] move it to start position.
	startPoint := pic.GetStartPoint()
	motorPortal.Turn(maxSpeed, startPoint.GetX()-500)
	motorFeeder.Turn(maxSpeed, -startPoint.GetY())
	motorPen.Turn(maxSpeed, -200) //pen down
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

	motorPen.Turn(maxSpeed, 200) //pen up
}

func drawRectangle() {
	//pen down
	speed := motorPen.MaxSpeed() / 6
	motorPen.Turn(speed, -100)

	pageHeight := 1000
	pageWidth := 1000

	//page width 1074, it's porla
	motorPortal.Turn(speed, pageWidth/2)

	motorFeeder.Turn(speed, -pageHeight)

	motorPortal.Turn(speed, -pageWidth)

	motorFeeder.Turn(speed, pageHeight)

	motorPortal.Turn(speed, pageWidth/2)

	//pen up
	motorPen.Turn(speed, 100)
}

func preparePen() {
	speed := motorPen.MaxSpeed() / 6
	motorPen.Turn(speed, -1200)
}

func preparePaper() {
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
	//turn pen back
	speed := motorPen.MaxSpeed() / 6
	motorPen.Turn(speed, 1200)

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
