// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// find demonstrates finding the first available sensor for a driver name.
//
// Program provide information about all connected motors and sensors.
// With informatin about setting.
// Currently it supports just LEGO original peripherials.
//
package main

import (
	"fmt"

	ev3dev "github.com/ev3go/ev3dev"
)

func printMotorInfo(name string, motor *ev3dev.TachoMotor) {
	for _, command := range motor.Commands() {
		fmt.Printf("Motor '%s' support command    : %s\n", name, command)
	}
	for _, command := range motor.StopActions() {
		fmt.Printf("Motor '%s' support stop action: %s\n", name, command)
	}
	fmt.Printf("Motor '%s' has type           : %s\n", name, motor.Type())
	fmt.Printf("Motor '%s' have countPerRot   : %d\n", name, motor.CountPerRot())
	p, _ := motor.Polarity()
	fmt.Printf("Motor '%s' have polarity      : %s\n", name, p)
	pos, _ := motor.Position()
	fmt.Printf("Motor '%s' is in position     : %d\n", name, pos)
	fmt.Printf("Motor '%s' have max speed     : %d\n", name, motor.MaxSpeed())
}

func printMotors() {
	var motors = [...]string{"lego-43362",
		"lego-ev3-l-motor",
		"lego-ev3-m-motor",
		"lego-47154",
		"lego-70823",
		"lego-71427",
		"lego-74569",
		"lego-88002",
		"lego-88003",
		"lego-88004",
		"lego-8882",
		"lego-8883",
		"lego-9670",
		"lego-nxt-motor"}
	var outPorts = [...]string{"outA", "outB", "outC", "outD"}
	for _, port := range outPorts {
		for _, driver := range motors {
			motor, err := ev3dev.TachoMotorFor("ev3-ports:"+port, driver)
			if err != nil {
				//There no such motor connect at given port.
			} else {
				fmt.Printf("Found motor %s at port %s\n", driver, port)
				printMotorInfo(driver, motor)
			}
		}
	}
}

func printSensorInfo(name string, sensor *ev3dev.Sensor) {
	for _, command := range sensor.Commands() {
		fmt.Printf("Sensor '%s' support command      : %s\n", name, command)
	}
	for _, command := range sensor.Modes() {
		fmt.Printf("Sensor '%s' support mode         : %s\n", name, command)
	}
	fmt.Printf("Sensor '%s' has type             : %s\n", name, sensor.Type())
	fmt.Printf("Sensor '%s' have firmware version: %s\n", name, sensor.FirmwareVersion())
}

func printSensors() {
	var sensors = [...]string{"lego-ev3-us",
		"lego-ev3-gyro",
		"lego-ev3-color",
		"lego-ev3-touch",
		"lego-ev3-ir",
		"wedo-hub",
		"wedo-motion",
		"wedo-tilt",
		"lego-power-storage",
		"lego-nxt-temp",
		"lego-nxt-touch",
		"lego-nxt-light",
		"lego-nxt-sound",
		"lego-nxt-us"}
	var inPorts = [...]string{"in1", "in2", "in3", "in4"}
	for _, port := range inPorts {
		for _, driver := range sensors {
			sensor, err := ev3dev.SensorFor("ev3-ports:"+port, driver)
			if err != nil {
				//There no such motor connect at given port.
			} else {
				fmt.Printf("Found sensor %s at port %s\n", driver, port)
				printSensorInfo(driver, sensor)
			}
		}
	}
}

func main() {
	fmt.Printf("Starting find info about EV3 brick.\n")
	printMotors()
	printSensors()
}
