# Dockerfile

# Use the official Golang image as a build stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Check if the binary was created
RUN ls -l ./main  # This will list the files in the current directory

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Check if the binary was copied
RUN ls -l ./main  # This will list the files in the current directory

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]