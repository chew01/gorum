############################
# STEP 1 build executable binary
############################
FROM golang:1.20rc3 AS builder

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/goginapp/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /go/main

############################
# STEP 2 build a small image
############################
FROM ubuntu:22.04

WORKDIR /

# Copy our static executable.
COPY --from=builder /go/main /go/main

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

WORKDIR /go

# Run the Go Gin binary.
ENTRYPOINT ["/go/main"]