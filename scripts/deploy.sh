#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

#
# Set IP address of EV3 brick.
#
IP=192.168.1.46

cd `dirname $0`

#scp ../bin/waitkeys robot@${IP}:/home/robot/
#scp ../bin/paint robot@${IP}:/home/robot/
#scp ../bin/speaker robot@${IP}:/home/robot/
scp ../bin/find robot@${IP}:/home/robot/
#scp ../bin/motor robot@${IP}:/home/robot/
#scp ../bin/print robot@${IP}:/home/robot/
