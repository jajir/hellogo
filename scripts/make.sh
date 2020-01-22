#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

cd ..

GOOS=linux GOARCH=arm GOARM=5 go build .
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/speaker first/speaker.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/waitkeys buttons/waitkeys.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/paint screen/paint.go
GOOS=linux GOARCH=arm GOARM=5 go build -o ./out/motor motor/motor.go
