#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

# set correct working directory
cd `dirname $0`
cd ..

#
# Set IP address of EV3 brick.
#
IP=192.168.1.46

#
# Define command from directory cmd which will be compiled and uploaded to EV3 brick.
#
command="find"

deploy() {
    local name=$1
    local target=./bin/${name}
    echo "Uploading '${target}' to EV3"
    scp ./bin/$name robot@${IP}:/home/robot/
}

make(){
    local name=$1
    local src=cmd/${name}/${name}.go
    local target=./bin/${name}
    echo "Compiling '${src}' to '${target}'"
    GOOS=linux GOARCH=arm GOARM=5 go build -o  ${target} ${src}    
}


#GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/speaker cmd/speaker/speaker.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/waitkeys cmd/waitkeys/waitkeys.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/paint cmd/paint/paint.go
#GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/find cmd/find/find.go
#go build -o ./bin/create_lissajous lissajous/create_lissajous.go

make ${command}
if [ "$?" -eq "0" ]
then
    # Do work when command exists on success
    deploy ${command}
fi

