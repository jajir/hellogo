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

# Make command executable file.
make ${command}

# Optionally deploy executable file to EV3.
 if [ "$?" -eq "0" ]
then
    # Execute just when program was compiled.
    deploy ${command}
fi

