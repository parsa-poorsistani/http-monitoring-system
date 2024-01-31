# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the Go application
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./main"]
