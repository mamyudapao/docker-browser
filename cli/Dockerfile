FROM golang:1.19-alpine as dev
RUN apt update && apt add git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
