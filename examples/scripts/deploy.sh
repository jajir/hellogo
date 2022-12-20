#!/bin/sh
#
# Deploy binaries to EV3 brick.
#

# set correct working directory
cd `dirname $0`
cd ..

# set configuraton
source ./scripts/conf.sh

#
# Find command name.
#
readCommandName $1
#echo "Command is '$command'"

# Deploy executable file to EV3.
deploy ${command}
