#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

cd ..

GOOS=linux GOARCH=arm GOARM=5 go build .
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/speaker first/speaker.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/waitkeys buttons/waitkeys.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/paint screen/paint.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/find find/find.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/print print/print.go print/screen.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/motor motor/axis.go motor/ev3lmotor.go motor/screen.go motor/motor.go motor/touchSensor.go motor/lights.go
