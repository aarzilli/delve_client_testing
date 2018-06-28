#!/bin/bash

ulimit -c unlimited
go build -gcflags='-N -l' panic.go
GOTRACEBACK=crash ./panic
mv core panic-core
echo ./panic ./panic-core
