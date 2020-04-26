package main

import (
	"fmt"
	"time"
)

func main() {
	Display.Clean()
	Display.Write(0, 10, "Starting ...")

	fmt.Println("Ahoj")

	time.Sleep(30 * time.Millisecond)

	Display.Close()
}
