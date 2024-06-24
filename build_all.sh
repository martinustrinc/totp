#!/bin/bash

# Windows Builds
#GOOS=windows GOARCH=386 go build main.go -o generator_time_win32.exe
#GOOS=windows GOARCH=amd64 go build main.go -o generator_time_win64.exe
$Env:GOOS = "windows"
$Env:GOARCH = "386"
go build generator_time_win32.exe

$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
go build generator_time_win64.exe

# Linux Builds
#GOOS=linux GOARCH=386 go build main.go -o generator_time_linux32
#GOOS=linux GOARCH=amd64 go build main.go -o generator_time_linux64
$Env:GOOS = "linux"
$Env:GOARCH = "386"
go build generator_time_linux32

$Env:GOOS = "linux"
$Env:GOARCH = "amd64"
go build generator_time_linux64
