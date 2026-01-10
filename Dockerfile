#Build image

FROM golang:tip-alpine3.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src/ ./src
RUN go build -o record-catalog ./src
