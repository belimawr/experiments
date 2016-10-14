#!/bin/bash
export GOPATH=$(pwd)
GITHASH=`git rev-parse --short HEAD`

mkdir -p bin
mkdir -p src/GoCD/

cp *.go src/GoCD/

cd src/GoCD/

go build -o ../../bin/GoCD

cd ../../

rm -rf ./src

cp bin/GoCD bin/application

cd bin

zip -9 GoCD-$GITHASH-STG.zip application

rm GoCD application
