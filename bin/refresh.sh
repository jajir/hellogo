#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

IP=192.168.1.46

cd `dirname $0`
cd ..

GOOS=linux GOARCH=arm GOARM=5  go build

scp hellogo robot@${IP}:/home/robot/

#scp ../out/waitkeys robot@${IP}:/home/robot/
#scp ../out/paint robot@${IP}:/home/robot/
#scp ../out/speaker robot@${IP}:/home/robot/
#scp ../out/motor robot@${IP}:/home/robot/
