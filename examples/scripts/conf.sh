#!/bin/sh
#
# Contains configurations and shared functions.
#

#
# Set IP address of EV3 brick.
#
IP=192.168.2.3

#
# Deploy executable to EV3.
#
deploy() {
    local name=$1
    local target=./bin/${name}

    # verify if file exists
    if [ -f $target ]
    then
        echo "Uploading '${target}' to EV3"
        scp -Cq $target robot@${IP}:/home/robot/
    else
        echo "File to upload '${target}' doesn't exists."
        exit
    fi
}

#
# Compile program executable on EV3.
#
make(){
    local name=$1
    local src=cmd/${name}/${name}.go
    local target=./bin/${name}
    echo "Compiling '${src}' to '${target}'"
    GOOS=linux GOARCH=arm GOARM=5 go build -o  ${target} ${src}    
}

#
# Compile program locally.
#
makeLocally(){
    local name=$1
    local src=cmd/${name}/${name}.go
    local target=./bin/${name}
    echo "Compiling '${src}' to '${target}'"
    go build -o  ${target} ${src}
}


#
# Define command from directory cmd which will be compiled and uploaded to EV3 brick.
# Command name will store to variable command.
readCommandName(){
    local param=$1
    if [ -z $param ]
    then
        echo "Missing parameter program name. Program name is dir name from './cmd/'"
        exit
    fi
    command=$param
}

