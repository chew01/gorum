FROM golang:1.20rc3 AS builder

ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/goginapp/
COPY .. .

RUN go get -d -v

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /go/main

FROM ubuntu:22.04

WORKDIR /

COPY --from=builder /go/main /go/main
COPY /internal/database/init.sql /go/internal/database/init.sql

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

WORKDIR /go

ENTRYPOINT ["/go/main"]
