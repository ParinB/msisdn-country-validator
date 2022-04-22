FROM golang:1.16-apline3.13 AS builder
RUN mkdir /backend
WORKDIR /backend
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .

RUN go install github.com/githubnemo/CompileDaemon@latest
EXPOSE 3001
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main




FROM golang:1.16-alpine3.13 as builder

RUN mkdir / jpay
WORKDIR /backend

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get \
    && go get github.com/gin-contrib/cors \
    && go get github.com/gin-gonic/gin \
    && go get github.com/mattn/go-sqlite3 \
    && go get github.com/stretchr/testify

RUN go build -o example main.go

FROM alpine:latest
EXPOSE 5555


