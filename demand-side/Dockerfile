# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /go/src/demand-side

# Download Go modules
COPY go.mod ./

COPY cmd/ ./cmd

COPY pkg/ ./pkg

RUN go mod download

CMD [ "go","run" ,"cmd/main.go"]

# TODO(isham) add multistage build