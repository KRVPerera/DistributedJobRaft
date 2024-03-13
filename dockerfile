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

# Copy the rest of the project files
COPY . .
COPY ./config/$CONFIG_FILE_PATH config/config.xml
RUN go build -ldflags='-s -w -extldflags "-static"' -o bin

# Set the entry point for the container
ENTRYPOINT ["/app/bin"]