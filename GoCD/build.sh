#!/bin/bash
export GOPATH=$(pwd)

mkdir -p src/GoCD/

cp *.go src/GoCD/

cd src/GoCD/

mkdir -p bin
mkdir -p reports

go build -o bin/GoCD
go test -coverprofile=reports/coverage.out
go tool cover -html=reports/coverage.out -o reports/coverage.html
go test -v | go-junit-report > reports/report.xml
