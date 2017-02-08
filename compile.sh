#!/bin/bash

export GOPATH=$PWD 
echo ++++++++++++++++++++++
echo main.go building ...
echo ++++++++++++++++++++++
go build -o bin/main.bin src/main.go
echo main.bin running ...
echo ----------------------
#clear
./bin/main.bin
