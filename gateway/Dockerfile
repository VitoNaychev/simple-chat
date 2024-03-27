# Use an official Golang runtime as a parent image
FROM golang:1.21.4

# Set the working directory to /app
WORKDIR /app

# Copy the gateway service at /app/gateway
COPY ../gateway /app/gateway

# Copy go.mod and go.sum in /app
COPY ../go.mod /app

# Set the working directiry to /app/cmd
WORKDIR /app/gateway/cmd

# Build the Go application
RUN go build -o main

# Run the Go application
CMD ["./main"]