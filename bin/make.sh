#!/bin/sh
#
# Deploy binaries to EV3 brick.
#
cd `dirname $0`
cd ..

#GOOS=linux GOARCH=arm GOARM=5 go build .
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/speaker examples/speaker/speaker.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/waitkeys examples/waitkeys/waitkeys.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/paint examples/paint/paint.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/find examples/find/find.go
#go build -o ./out/create_lissajous lissajous/create_lissajous.go
#go build -o ./out/test_lissajous print/test_lissajous.go GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/print print/print.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/motor motor/axis.go motor/ev3lmotor.go motor/screen.go motor/motor.go motor/touchSensor.go motor/lights.go
