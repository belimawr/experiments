GOPATH=$(shell pwd)

all:
	mkdir -p src/GoCD/
	mkdir -p reports
	mkdir -p bin

	cp *.go src/GoCD/

	go get -u github.com/jstemmer/go-junit-report

	cd src/GoCD/ && go build -o GoCD
	cd src/GoCD/ && go test -coverprofile=coverage.out
	cd src/GoCD/ && go tool cover -html=coverage.out -o coverage.html
	cd src/GoCD/ && go test -v | ../../bin/go-junit-report > report.xml

	rm -rf bin/* pkg

	cp src/GoCD/GoCD ./bin/
	cp src/GoCD/coverage.* ./reports
	cp src/GoCD/report.xml ./reports

	rm -rf src/
