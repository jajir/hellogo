#!/bin/sh
#
# Compile and run first example
#

# set correct working directory to project root
cd `dirname $0`
cd ..

./build/build.sh
./out/print ./example/example1.svg