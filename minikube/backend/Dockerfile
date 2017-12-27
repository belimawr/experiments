FROM golang:1.9-alpine as builder

RUN apk --no-cache upgrade && apk --no-cache add git
RUN go get github.com/golang/dep/cmd/dep
COPY ./ /go/src/github.com/belimawr/experiments/minikube/backend
WORKDIR /go/src/github.com/belimawr/experiments/minikube/backend
RUN dep ensure && go build

FROM alpine

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/belimawr/experiments/minikube/backend/backend /opt/

WORKDIR /opt

EXPOSE 3000

CMD ["./backend"]

