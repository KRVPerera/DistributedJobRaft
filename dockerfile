#
#  Created by Rukshan Perera (rukshan.perera@student.oulu.fi)
#

# Use the official Go image as the base image
FROM golang:alpine3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

ARG CONFIG_FILE_PATH
ENV CGO_ENABLED=1

# Download and install the Go dependencies
RUN go mod tidy
RUN apk add --no-cache gcc musl-dev
RUN go get github.com/mattn/go-sqlite3
RUN go mod download

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
EXPOSE 8083
EXPOSE 2461
EXPOSE 2462
EXPOSE 2463

# Copy the rest of the project files
COPY . .
COPY ./config/$CONFIG_FILE_PATH config/config.xml
RUN go build -o bin

# Set the entry point f-ldflags='-extldflags "-static"'or the container
ENTRYPOINT ["/app/bin"]