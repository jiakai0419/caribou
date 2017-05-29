#!/bin/sh

find . -type d | egrep -v "\.git|\.idea|\.$" | xargs go build
find . -type d | egrep -v "\.git|\.idea|\.$" | xargs go test
