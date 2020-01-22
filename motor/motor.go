// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// waitkeys demonstrates key waiting. It should be run from the command line.
// It requires ^C to terminate.
package main

import (
	"fmt"
	"time"
	ev3dev "github.com/ev3go/ev3dev"
	ev3 "github.com/ev3go/ev3"
	log "github.com/sirupsen/logrus"
	"image/color"
	"image"
	"image/draw"
	"golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
)

func main() {
	
	var konec = make(chan string)
	
	ev3.LCD.Init(true)
	go waitingForBack(konec)
	go run(konec)
	
	k := <- konec
	log.Info("Shutting down, bacause of " + k + "\n")
	zhasni()
	ev3.LCD.Close()
}

func waitingForBack(konec chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}
	
	for e := range w.Events {
		fmt.Printf("%+v\n", e)
		if 1 == e.Button {
			fmt.Printf("Zpet")
			konec <- "Button Back was pressed." 
		}
	}
}

func addLabel(img draw.Image, x, y int, label string) {
    point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

    d := &font.Drawer{
        Dst:  img,
        Src:  image.NewUniform(color.Black),
        Face: basicfont.Face7x13,
        Dot:  point,
    }
    d.DrawString(label)
}


func clean() {
	draw.Draw(ev3.LCD,  ev3.LCD.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
}

func run(konec chan string) {
	clean()
	addLabel(ev3.LCD, 0, 10, "Ahoj lidi")
	
	a, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")	
	if err != nil {
		log.Fatalf("failed to find left motor in outA: %v", err)
	}
	
	a.SetPosition(10)
	
	astat, _ := a.State()
	aspeed, _ := a.Speed()
	log.Printf("outA: %s %d", astat, aspeed)
	
	time.Sleep(30 * time.Second)
}
	
func zhasni(){
	ev3.GreenLeft.SetBrightness(0)
	ev3.GreenRight.SetBrightness(0)
	ev3.RedLeft.SetBrightness(0)
	ev3.RedRight.SetBrightness(0)		
}
