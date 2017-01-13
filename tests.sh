#!/bin/sh
# To get more detailed output, pass -v to this script.

exec go test "$@" ./...
