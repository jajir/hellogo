#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

IP=192.168.1.46

#scp ../out/waitkeys robot@${IP}:/home/robot/
#scp ../out/paint robot@${IP}:/home/robot/
scp ../out/speaker robot@${IP}:/home/robot/
#scp ../out/find robot@${IP}:/home/robot/
#scp ../out/motor robot@${IP}:/home/robot/
#scp ../out/print robot@${IP}:/home/robot/
