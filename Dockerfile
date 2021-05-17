FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]