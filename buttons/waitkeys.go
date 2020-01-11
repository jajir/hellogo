// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// waitkeys demonstrates key waiting. It should be run from the command line.
// It requires ^C to terminate.
package main

import (
	"fmt"
	"time"
	"github.com/ev3go/ev3dev"
	"github.com/ev3go/ev3"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	var konec = make(chan string)
	
	go waitingForBack(konec)
	go run(konec)
	
	k := <- konec
	log.Info("Shutting down, bacause of " + k + "\n")
	zhasni()
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


func run(konec chan string) {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}
	
	var maxBrightness,_ = ev3.GreenLeft.MaxBrightness()
	maxBrightness=100
	
	fmt.Printf("Max brightness is %v", maxBrightness)
	
	ev3.GreenLeft.SetBrightness(maxBrightness)
	ev3.GreenRight.SetBrightness(maxBrightness)
	
	time.Sleep(2 * time.Second)
	
	ev3.RedLeft.SetBrightness(maxBrightness)
	ev3.RedRight.SetBrightness(maxBrightness)
	
	time.Sleep(30 * time.Second)
	
	for e := range w.Events {
		fmt.Printf("%+v\n", e)
		if 1 == e.Button {
			fmt.Printf("Zpet")
		}
	}
}
	
func zhasni(){
	ev3.GreenLeft.SetBrightness(0)
	ev3.GreenRight.SetBrightness(0)
	ev3.RedLeft.SetBrightness(0)
	ev3.RedRight.SetBrightness(0)		
}
