#!/bin/bash

mkdir -p bin
mkdir -p reports

go build -o bin/GoCD
go test -coverprofile=reports/coverage.out
go tool cover -html=reports/coverage.out -o reports/coverage.html
