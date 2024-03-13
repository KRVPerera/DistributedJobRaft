# Use the official Go image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

ARG CONFIG_FILE_PATH

# Download and install the Go dependencies
RUN go get -u github.com/mattn/go-sqlite3 v1.14.6
RUN go mod download

# Copy the rest of the project files
COPY . .
COPY ./config/$CONFIG_FILE_PATH config/config.xml
RUN go build -o bin

# Set the entry point for the container
ENTRYPOINT ["/app/bin"]