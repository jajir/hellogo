#!/bin/sh
#
# Compile givec code lacaly and run it.
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
makeLocally ${command}

# Optionally execute program.
 if [ "$?" -eq "0" ]
then
    # Execute just when program was compiled.
    ./bin/${command} $2 $3 $4 $5
fi

