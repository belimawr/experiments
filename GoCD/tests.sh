export GOPATH=$(pwd)

mkdir -p src/GoCD/
mkdir -p src/GoCD/reports

cp *.go src/GoCD/

cd src/GoCD/

go test -coverprofile=reports/coverage.out
go tool cover -html=reports/coverage.out -o reports/coverage.html
go test -v > ./test.tmp

ret_code=$?

cat test.tmp | go-junit-report > reports/report.xml

rm -rf ../../reports
mv reports ../../

cd ../../
rm -rf ./src/

exit $ret_code
