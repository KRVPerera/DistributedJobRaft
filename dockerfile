# Use the official Go image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o bin

# Set the entry point for the container
ENTRYPOINT ["./app/bin"]