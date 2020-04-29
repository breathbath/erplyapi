# @TODO License INFORMATION.

# -------------
# Build stage

FROM golang:1.14 AS build-env

RUN mkdir -p /app/

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo \
        -o /app/erply_server \
        main.go
COPY .env /app/.env
COPY templates /app
COPY favicon.ico /app/favicon.ico
