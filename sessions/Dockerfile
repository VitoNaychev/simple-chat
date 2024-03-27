# Use an official Golang runtime as a parent image
FROM golang:1.21.4

# Set the working directory to /app
WORKDIR /app

# Copy the sessions service at /app/sessions
COPY ../sessions /app/sessions

# Copy go.mod and go.sum in /app
COPY ../go.mod /app

# Set the working directiry to /app/cmd
WORKDIR /app/sessions/cmd

# Build the Go application
RUN go build -o main

# Run the Go application
CMD ["./main"]