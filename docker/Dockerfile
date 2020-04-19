# @TODO License INFORMATION.

# -------------
# Build stage

FROM golang:1.14 AS build-env

RUN mkdir -p /build/

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo \
        -o /build/erply_server \
        main.go

# -------------
# Image creation stage

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /build/erply_server /app/
COPY .env /app/.env
