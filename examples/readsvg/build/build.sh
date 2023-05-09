#!/bin/sh
#
# Build final executable
#

# set correct working directory to project root
cd `dirname $0`
cd ..

go build -o ./out/print ./cmd/print/print.go
