# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install the application dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /app/main

# Expose port 8080 (or any other port your application listens on)
EXPOSE 8080

# Run the executable when the container starts
CMD ["/app/main"]