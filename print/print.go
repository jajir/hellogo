package main

import (
	"fmt"
	"github.com/jajir/hellogo/lev3"
	"time"
)

func main() {
	defer lev3.Display.Close()
	defer lev3.Speaker.Close()
	defer lev3.Lights.TurnOff()
	
	lev3.Lights.GreenTurnOn()

	lev3.Display.Clean()
	lev3.Display.Write(0, 10, "Starting ...")

	fmt.Println("Ahoj")

	lev3.Speaker.Beep()

	time.Sleep(30 * time.Millisecond)
}
