# Use the official Golang image to create a build artifact.
# This is the build stage.
FROM golang:1.23.0-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o bin/api .

# Start a new stage from scratch
FROM alpine:latest  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/api /usr/local/bin/api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["api"]