# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files to the working directory
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port that your application will run on
EXPOSE 8080

# Define the command to start your application
CMD ["./main"]