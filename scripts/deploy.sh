#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

IP=192.168.1.46

#scp ../out/waitkeys robot@${IP}:/home/robot/
scp ../out/paint robot@${IP}:/home/robot/
