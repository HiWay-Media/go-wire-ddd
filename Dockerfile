FROM golang:1.21-alpine AS build_base
LABEL org.opencontainers.image.authors="allan.nava@hiway.media"
#
RUN mkdir -p /etc/ssl/certs/ && update-ca-certificates && apk add --no-cache git
# Set the Current Working Directory inside the container
WORKDIR /go/src/app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
# Build the Go app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./microservices/main.go
#
# Start fresh from a smaller image
#FROM scratch
FROM phusion/baseimage:focal-1.2.0
#
WORKDIR /app
RUN apt-get install -y ca-certificates
COPY --from=build_base /go/src/app/main /app/main
#
EXPOSE 8080
#
CMD ["./app/main"]
