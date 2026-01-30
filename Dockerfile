FROM golang:1.21-bookworm
# This Dockerfile is used to bundle the source and all dependencies into an image for testing.

# Install gox for cross-compilation and golint for linting
RUN go install github.com/mitchellh/gox@latest && \
    go install golang.org/x/lint/golint@latest

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/p5k6/workflow-cli

ENV GO111MODULE=on

COPY ./_scripts /usr/local/bin

COPY go.mod go.sum ./
RUN go mod download

COPY . /go/src/github.com/p5k6/workflow-cli

RUN go build
