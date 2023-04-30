#!/bin/sh
#
# Compile example project for EV3 brick.
#

# set correct working directory
cd `dirname $0`
cd ..

# set configuraton
source ./scripts/conf.sh

#
# Find command name.
#
readExampleProjectName $1
echo "Example project name is '$exampleProjectName'"

# Make command executable file.
make ${exampleProjectName}
