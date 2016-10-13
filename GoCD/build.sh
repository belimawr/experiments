#!/bin/bash
export GOPATH=$(pwd)

mkdir -p bin
mkdir -p src/GoCD/

cp *.go src/GoCD/

cd src/GoCD/

go build -o ../../bin/GoCD

cd ../../

rm -rf ./src

